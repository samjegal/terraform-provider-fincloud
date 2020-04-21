package network

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var virtualPrivateCloudResourceName = "fincloud_virtual_private_cloud"

func resourceVirtualPrivateCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualPrivateCloudCreate,
		Read:   resourceVirtualPrivateCloudRead,
		Update: resourceVirtualPrivateCloudUpdate,
		Delete: resourceVirtualPrivateCloudDelete,

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

			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"status_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceVirtualPrivateCloudCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	vpcClient := client.Network.VirtualPrivateCloudClient

	vpcName := d.Get("name").(string)
	param := expandVirtualPrivateCloud(d)
	_, err := vpcClient.CreateOrUpdate(ctx, param)
	if err != nil {
		return fmt.Errorf("")
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "CREATING"},
		Target:     []string{"RUN"},
		Refresh:    virtualPrivateCloudStateRefreshFunc(client, vpcName),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. VPC: %q", vpcName)
	}

	return resourceVirtualPrivateCloudRead(d, meta)
}

func resourceVirtualPrivateCloudRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.VirtualPrivateCloudClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	vpcName := d.Get("name").(string)

	resp, err := client.List(ctx)
	if err != nil {
		return fmt.Errorf("VPC 정보의 요청에 문제가 발생했습니다. vpcName: %q: %+v", vpcName, err)
	}

	var props *network.VirtualPrivateCloudContentParameter
	for _, v := range *resp.Content {
		if vpcName == *v.VpcName {
			props = &v
			break
		}
	}

	if props != nil {
		if err := d.Set("name", *props.VpcName); err != nil {
			return fmt.Errorf("VPC 이름 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("cidr_block", *props.Ipv4Cidr); err != nil {
			return fmt.Errorf("VPC 이름 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("status_code", string(props.StatusCode)); err != nil {
			return fmt.Errorf("VPC 이름 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		d.SetId(*props.VpcNo)
	}

	return nil
}

func resourceVirtualPrivateCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	// FIXME: 금융 클라우드에서 구현되지 않은 상태이다.
	return nil
}

func resourceVirtualPrivateCloudDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()
	vpcClient := client.Network.VirtualPrivateCloudClient

	vpcNo := d.Id()
	vpcName := d.Get("name").(string)

	_, err := vpcClient.Delete(ctx, vpcNo, expandVirtualPrivateCloud(d))
	if err != nil {
		return fmt.Errorf("VPC 삭제 요청에 대한 문제가 발생했습니다. err: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "CREATING", "TERMTING"},
		Target:     []string{"NOTFOUND"},
		Refresh:    virtualPrivateCloudStateRefreshFunc(client, vpcName),
		Timeout:    30 * time.Minute,
		MinTimeout: 10 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("VPC 삭제 중 문제가 발생했습니다. vpcName: %s", vpcName)
	}

	d.SetId("")
	return nil
}

func expandVirtualPrivateCloud(d *schema.ResourceData) network.VirtualPrivateCloudParameter {
	vpcName := d.Get("name").(string)
	cidrBlock := d.Get("cidr_block").(string)

	return network.VirtualPrivateCloudParameter{
		VpcName:  utils.String(vpcName),
		Ipv4Cidr: utils.String(cidrBlock),
	}
}

func virtualPrivateCloudStateRefreshFunc(client *clients.Client, vpcName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := client.StopContext
		resp, err := client.Network.VirtualPrivateCloudClient.List(ctx)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. VPC: %q: %+v", vpcName, err)
		}

		if resp.Content == nil {
			return nil, "EMPTY", fmt.Errorf("VPC 정보가 존재하지 않습니다. VPC: %q: %+v", vpcName, err)
		}

		for _, v := range *resp.Content {
			if vpcName == *v.VpcName {
				return v, string(network.VirtualPrivateCloudStatusCode(v.StatusCode)), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}
