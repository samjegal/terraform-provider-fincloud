package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var serverResourceName = "fincloud_server"

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "11",
				ValidateFunc: validate.NoEmptyStrings,
			},

			"zone_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "110",
				ValidateFunc: validate.NoEmptyStrings,
			},

			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"spec": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"oscode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"login_key": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"path": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},
					},
				},
			},

			"network_interface": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"subnet_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"address": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"security_groups": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"init_script_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"root_password": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	serverClient := client.Compute.ServerClient

	name := d.Get("name").(string)

	param, err := expandServerParameter(d)
	if err != nil {
		return err
	}

	resp, err := serverClient.Create(ctx, *param)
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("서버 생성에 문제가 발생했다. Server: %q: %+v", name, err)
	}

	serverId, err := serverInstanceNumber(ctx, meta, name)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			"init", "booting", "creating", "setting up", "stopped",
			"shutting down", "terminating", "not found",
		},
		Target:     []string{"running"},
		Refresh:    serverStateRefreshFunc(client, serverId),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("서버 상세정보의 코드를 가져오는 중 문제가 발생했다. Server: %q", name)
	}

	d.SetId(serverId)

	return resourceServerRead(d, meta)
}

func resourceServerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	serverId := d.Id()

	serverClient := client.Compute.ServerClient

	_, err := serverClient.Get(ctx, d.Id())
	if err != nil {
		return err
	}

	if v, ok := d.GetOk("login_key"); ok {
		keys := v.([]interface{})
		var k map[string]interface{}
		for _, key := range keys {
			k = key.(map[string]interface{})
			break
		}

		name := k["name"].(string)
		path := k["path"].(string)

		pem, err := ioutil.ReadFile(path + name + ".pem")
		if err != nil {
			return fmt.Errorf("로그인 키 파일을 읽는 중 문제가 발생했습니다. LoginKey: %q: %+v", path+name+".pem", err)
		}

		keyFile := ioutil.NopCloser(bytes.NewReader(pem))
		resp, err := client.Compute.RootPasswordClient.Get(ctx, serverId, keyFile)
		if err != nil {
			return fmt.Errorf("서버 관리자 비밀번호 정보를 가져올 수가 없습니다. Server: %q: %+v", serverId, err)
		}

		d.Set("root_password", resp.Content.DecryptRootPassword)
	}

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()
	serverClient := client.Compute.ServerClient

	serverId := d.Id()

	_, err := serverClient.Shutdown(ctx, expandServerInstanceListParameter(d))
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			"init", "booting", "creating", "setting up", "running",
			"shutting down", "terminating", "not found",
		},
		Target:     []string{"stopped"},
		Refresh:    serverStateRefreshFunc(client, serverId),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("서버 상세정보의 상태 코드를 가져오는 중 문제가 발생했다. Server: %q", serverId)
	}

	_, err = serverClient.Delete(ctx, expandServerInstanceListParameter(d))
	if err != nil {
		return err
	}

	stateConf = &resource.StateChangeConf{
		Pending: []string{
			"init", "booting", "creating", "setting up", "running",
			"stopped", "shutting down", "terminating",
		},
		Target:     []string{"not found"},
		Refresh:    serverStateRefreshFunc(client, serverId),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("서버 상세정보의 상태 코드를 가져오는 중 문제가 발생했다. Server: %q", serverId)
	}

	d.SetId("")

	return nil
}

func expandServerParameter(d *schema.ResourceData) (*compute.ServerParameter, error) {
	param := &compute.ServerParameter{
		ReturnProtectionYn:   utils.String("N"),
		ServerCreateCount:    utils.String("1"),
		HostServerNameSameYn: utils.String("Y"),
		FeeSystemTypeCode:    utils.String("FXSUM"),
		MetaTypeCode:         utils.String("SYS"),

		ServerName:                utils.String(d.Get("name").(string)),
		ServerDesc:                utils.String(d.Get("description").(string)),
		ServerInstanceProductCode: utils.String(d.Get("spec").(string)),
		SoftwareProductCode:       utils.String(d.Get("oscode").(string)),

		RegionNo: utils.String(d.Get("region_id").(string)),
		VpcNo:    utils.String(d.Get("vpc_id").(string)),
		SubnetNo: utils.String(d.Get("subnet_id").(string)),
	}

	zoneId, _ := strconv.Atoi(d.Get("zone_id").(string))
	param.ZoneNo = utils.Int32(int32(zoneId))

	// Login Key
	if v, ok := d.GetOk("login_key"); ok {
		keys := v.([]interface{})
		for _, key := range keys {
			k := key.(map[string]interface{})
			param.LoginKeyName = utils.String(k["name"].(string))
			break
		}
	} else {
		return nil, fmt.Errorf("로그인 키가 존재하지 않습니다.")
	}

	// Network Interface
	param.NetworkInterfaces = expandServerNetworkInterfaceProperties(d)

	// Security Groups
	if v, ok := d.GetOk("security_groups"); ok {
		securityGroups := v.([]interface{})
		securityGroupNumberList := make([]int32, 0)
		for _, sec := range securityGroups {
			secId, _ := strconv.Atoi(sec.(string))
			securityGroupNumberList = append(securityGroupNumberList, int32(secId))
		}

		param.AccessControlGroupNoList = &securityGroupNumberList
	}

	// Init Script
	if v, ok := d.GetOk("init_script_id"); ok {
		param.InitConfigurationScriptNo = utils.String(v.(string))
	}

	return param, nil
}

func expandServerNetworkInterfaceProperties(d *schema.ResourceData) *[]compute.ServerNetworkInterfaceProperties {
	param := make([]compute.ServerNetworkInterfaceProperties, 0)
	networkInterface := d.Get("network_interface").([]interface{})

	var index int = 0

	for _, nic := range networkInterface {
		n := nic.(map[string]interface{})

		nicId, _ := strconv.Atoi(n["id"].(string))
		nicParam := compute.ServerNetworkInterfaceProperties{
			Order:              utils.Int32(int32(index)),
			SubnetNo:           utils.String(n["subnet_id"].(string)),
			NetworkInterfaceNo: utils.Int32(int32(nicId)),
			OverlayIP:          utils.String(n["address"].(string)),
		}

		if index == 0 {
			nicParam.DefaultYn = utils.String("Y")
		} else {
			nicParam.DefaultYn = utils.String("N")
		}

		securityGroupList := make([]compute.ServerNetworkInterfaceSecurityGroupProperties, 0)
		for _, sec := range n["security_groups"].([]interface{}) {
			securityGroupId, _ := strconv.Atoi(sec.(string))
			securityGroupList = append(securityGroupList,
				compute.ServerNetworkInterfaceSecurityGroupProperties{
					AccessControlGroupNo: utils.Int32(int32(securityGroupId)),
				})
		}

		nicParam.AccessControlGroups = &securityGroupList
		param = append(param, nicParam)

		index++
	}

	return &param
}

func serverInstanceNumber(ctx context.Context, meta interface{}, name string) (string, error) {
	client := meta.(*clients.Client).Compute.ServerClient

	resp, err := client.List(ctx, expandServerSearchParameter())
	if err != nil {
		return "", fmt.Errorf("서버 리스트 정보를 가져오는 중 에러가 발생했습니다. Server: %q: %+v", name, err)
	}

	for _, v := range *resp.Content {
		if name == *v.ServerName {
			return *v.InstanceNo, nil
		}
	}

	return "", fmt.Errorf("서버 정보가 존재하지 않습니다. Server: %q", name)
}

func serverStateRefreshFunc(client *clients.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := client.StopContext
		resp, err := client.Compute.ServerClient.List(ctx, expandServerSearchParameter())
		if err != nil || resp.Content == nil {
			return nil, "not found", fmt.Errorf("해당 서버 정보가 존재하지 않습니다. Server: %q: %+v", id, err)
		}

		for _, v := range *resp.Content {
			if id == *v.InstanceNo {
				return v, string(compute.ServerInstanceStatusName(v.InstanceStatusName)), nil
			}
		}

		return resp, "not found", nil
	}
}

func expandServerSearchParameter() compute.ServerSearchParameter {
	var page int32 = 1
	var pageSize int32 = 100

	return compute.ServerSearchParameter{
		PageNo:       utils.Int32(page),
		PageSizeNo:   utils.Int32(pageSize),
		SortedBy:     utils.String("instanceNo"),
		SortingOrder: utils.String("descending"),
		SubAccountNo: utils.String("0"),
		Owner:        utils.Bool(false),
		Filter:       nil,
	}
}

func expandServerInstanceListParameter(d *schema.ResourceData) compute.ServerInstanceListParameter {
	serverId := d.Id()
	serverList := make([]string, 0)
	serverList = append(serverList, serverId)

	return compute.ServerInstanceListParameter{
		InstanceNoList: &serverList,
	}
}
