package compute

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var networkInterfaceResoureName = "fincloud_network_interface"

func resourceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfaceCreate,
		Read:   resourceNetworkInterfaceRead,
		Update: resourceNetworkInterfaceUpdate,
		Delete: resourceNetworkInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "normal",
				ValidateFunc: validation.StringInSlice([]string{
					"normal",
					"baremetal",
				}, true),
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

			"security_groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceNetworkInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	nicClient := client.Compute.NetworkInterfaceClient

	nic, err := expandNetworkInterfaceParameter(d)
	if err != nil {
		return err
	}

	resp, err := nicClient.Create(ctx, *nic)
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("네트워크 인터페이스를 생성 중 문제가 발생했습니다. %q: %+v", *nic.NetworkInterfaceName, err)
	}

	return resourceNetworkInterfaceRead(d, meta)
}

func resourceNetworkInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	nicClient := client.Compute.NetworkInterfaceClient

	name := d.Get("name")

	resp, err := nicClient.List(ctx, *expandNetworkInterfaceSearchParameter())
	if err != nil {
		return err
	}

	var props *compute.NetworkInterfaceContentProperties = nil

	for _, v := range *resp.Content {
		if name == *v.NetworkInterfaceName {
			props = &v
			break
		}
	}

	if props != nil {
		d.SetId(fmt.Sprintf("%d", *props.NetworkInterfaceNo))
	}

	return nil
}

func resourceNetworkInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.NetworkInterfaceClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	_, err := client.Delete(ctx, d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func expandNetworkInterfaceParameter(d *schema.ResourceData) (*compute.NetworkInterfaceParameter, error) {
	param := &compute.NetworkInterfaceParameter{
		VpcNo:                utils.String(d.Get("vpc_id").(string)),
		SubnetNo:             utils.String(d.Get("subnet_id").(string)),
		NetworkInterfaceName: utils.String(d.Get("name").(string)),
		Description:          utils.String(d.Get("description").(string)),
		BmYn:                 utils.String("N"),
	}

	if v := d.Get("address").(string); v != "" {
		param.OverlayIP = utils.String(v)
	}

	securityGroupList := make([]compute.NetworkInterfaceSecurityGroupsProperties, 0)
	securityGroupNumberList := make([]int32, 0)

	securityGroups := d.Get("security_groups").([]interface{})
	for _, group := range securityGroups {
		securityGroupId, _ := strconv.Atoi(group.(string))
		securityGroupList = append(securityGroupList,
			compute.NetworkInterfaceSecurityGroupsProperties{
				AccessControlGroupNo: utils.Int32(int32(securityGroupId)),
			})
		securityGroupNumberList = append(securityGroupNumberList, int32(securityGroupId))
	}

	param.AccessControlGroups = &securityGroupList
	param.AccessControlGroupNoList = &securityGroupNumberList

	if v := d.Get("type").(string); v == "baremetal" {
		param.BmYn = utils.String("Y")
	}

	return param, nil
}

func expandNetworkInterfaceSearchParameter() *compute.NetworkInterfaceSearchParameter {
	var page int32 = 1
	var pageSize int32 = 100

	filter := make([]compute.NetworkInterfaceSearchFilterProperties, 0)
	filter = append(filter, compute.NetworkInterfaceSearchFilterProperties{
		Field: utils.String("networkInterfaceName"),
	})

	return &compute.NetworkInterfaceSearchParameter{
		PageNo:     utils.Int32(page),
		PageSizeNo: utils.Int32(pageSize),
		Filter:     &filter,
	}
}
