package network

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var subnetResourceName = "fincloud_subnet"

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetCreate,
		Read:   resourceSubnetRead,
		Update: resourceSubnetUpdate,
		Delete: resourceSubnetDelete,

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

			"network_acl_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"private",
				}, true),
			},

			"purpose": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"normal",
					"loadbalancer",
					"baremetal",
				}, true),
			},
		},
	}
}

func resourceSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	subnetClient := client.Network.SubnetClient

	name := d.Get("name").(string)
	param, err := expandSubnetParam(d, name)
	if err != nil {
		return err
	}

	resp, err := subnetClient.Create(ctx, *param)
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("서브넷 생성에 문제가 발생했습니다. Subnet: %q, %q: %+v", name, *resp.Error.Message, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"CREATING"},
		Target:     []string{"RUN"},
		Refresh:    subnetStateRefreshFunc(client, name, expandSubnetSearchParameters()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("서브넷의 상태 코드를 가져오는 중 문제가 발생했다. Subnet: %q", name)
	}

	return resourceSubnetRead(d, meta)
}

func resourceSubnetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.SubnetClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	resp, err := client.List(ctx, expandSubnetSearchParameters())
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("전체 리스트 정보를 읽을 수 있는 요청 메시지를 만들 수가 없습니다. Subnet: %q: %+v", name, err)
	}

	var props *network.SubnetSearchContentParameter
	for _, v := range *resp.Content {
		if name == *v.SubnetName {
			props = &v
			break
		}
	}

	if props != nil {
		d.SetId(*props.SubnetNo)
	}

	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()
	subnetClient := client.Network.SubnetClient

	name := d.Get("name").(string)

	_, err := subnetClient.Delete(ctx, d.Id())
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"RUN", "TERMTING"},
		Target:     []string{"NOTFOUND"},
		Refresh:    subnetStateRefreshFunc(client, name, expandSubnetSearchParameters()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("서브넷 삭제중 문제가 발생했습니다. Subnet: %q", name)
	}

	d.SetId("")
	return nil
}

func expandSubnetParam(d *schema.ResourceData, name string) (*network.SubnetParameter, error) {
	param := &network.SubnetParameter{
		Subnet:     utils.String(d.Get("cidr_block").(string)),
		SubnetName: utils.String(d.Get("name").(string)),
		VpcNo:      utils.String(d.Get("vpc_id").(string)),
	}

	zoneId, _ := strconv.Atoi(d.Get("zone_id").(string))
	networkAclId, _ := strconv.Atoi(d.Get("network_acl_id").(string))

	param.ZoneNo = utils.Int32(int32(zoneId))
	param.NetworkACLNo = utils.Int32(int32(networkAclId))

	t := d.Get("type").(string)
	purpose := d.Get("purpose").(string)

	var isIGW string = "N"
	var isLB, isBM string = "N", "N"

	if t == "public" {
		if purpose == "loadbalancer" {
			return nil, fmt.Errorf("Public 서브넷에서는 로드밸런서를 사용할 수가 없습니다.")
		}

		isIGW = "Y"
	} else {
		switch purpose {
		case "normal":
			isLB = "N"
			isBM = "N"
		case "loadbalancer":
			isLB = "Y"
		case "baremetal":
			isBM = "Y"
		default:
			return nil, fmt.Errorf("서브넷 용도에 적합하지 않는 값이 입력되었습니다. purpose: %+v", purpose)
		}
	}

	param.IgwYn = &isIGW
	param.LbYn = &isLB
	param.BmYn = &isBM

	return param, nil
}

func subnetStateRefreshFunc(client *clients.Client, name string, param network.SubnetSearchParameter) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		subnetClient := client.Network.SubnetClient
		ctx := client.StopContext
		resp, err := subnetClient.List(ctx, param)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. Subnet: %q: %+v", name, err)
		}

		if resp.Content == nil {
			return nil, "NOTFOUND", fmt.Errorf("서브넷 정보가 존재하지 않습니다. Subnet: %q: %+v", name, err)
		}

		for _, v := range *resp.Content {
			if name == *v.SubnetName {
				return v, string(network.SubnetStatusCode(v.StatusCode)), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}

func expandSubnetSearchParameters() network.SubnetSearchParameter {
	var page int32 = 1
	var pageSize int32 = 100

	filter := make([]network.SubnetSearchFilterParameter, 0)
	filter = append(filter, network.SubnetSearchFilterParameter{
		Field: utils.String("subnetName"),
	})

	return network.SubnetSearchParameter{
		PageNo:     utils.Int32(page),
		PageSizeNo: utils.Int32(pageSize),
		Filter:     &filter,
	}
}
