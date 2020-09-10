package vpc

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
)

var vpcResourceName = "fincloud_vpc"

func resourceVpc() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpcCreate,
		Read:   resourceVpcRead,
		Update: resourceVpcUpdate,
		Delete: resourceVpcDelete,

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
		},
	}
}

func resourceVpcCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.VpcClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	vpcName := d.Get("name").(string)
	cidrBlock := d.Get("cidr_block").(string)

	resp, err := client.Create(ctx, cidrBlock, vpcName)
	if err != nil {
		return fmt.Errorf("")
	}

	if *resp.CreateVpcResponse.ReturnCode != "0" {
		return nil // ERROR!
	}

	content := (*resp.CreateVpcResponse.VpcList)[0]
	id := *content.VpcNo

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "CREATING"},
		Target:     []string{"RUN"},
		Refresh:    vpcStateRefreshFunc(meta.(*clients.Client), id),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. VPC: %q", vpcName)
	}

	d.SetId(id)

	return resourceVpcRead(d, meta)
}

func resourceVpcRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.VpcClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := d.Id()
	vpcName := d.Get("name").(string)

	resp, err := client.GetDetail(ctx, id)
	if err != nil {
		return fmt.Errorf("VPC 정보의 요청에 문제가 발생했습니다. vpcName: %q: %+v", vpcName, err)
	}

	content := resp.GetVpcDetailResponse.VpcList
	if content != nil {
		if err := d.Set("name", (*content)[0].VpcName); err != nil {
			return fmt.Errorf("VPC 이름 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("cidr_block", (*content)[0].Ipv4CidrBlock); err != nil {
			return fmt.Errorf("VPC CIDR 주소 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		// if err := d.Set("status_code", string(*(*content)[0].VpcStatus.Code)); err != nil {
		// 	return fmt.Errorf("VPC 이름 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		// }
	}

	return nil
}

func resourceVpcUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVpcDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.VpcClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := d.Id()
	vpcName := d.Get("name").(string)

	_, err := client.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("VPC 삭제 요청에 대한 문제가 발생했습니다. err: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "CREATING", "TERMTING", "RUN"},
		Target:     []string{"NOTFOUND"},
		Refresh:    vpcStateRefreshFunc(meta.(*clients.Client), id),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("VPC 삭제 중 문제가 발생했습니다. vpcName: %s", vpcName)
	}

	d.SetId("")
	return nil
}

func vpcStateRefreshFunc(client *clients.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, err := client.Vpc.VpcClient.GetDetail(client.StopContext, id)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. VPC: %q: %+v", id, err)
		}

		if resp.GetVpcDetailResponse == nil {
			return nil, "EMPTY", fmt.Errorf("해당하는 정보가 존재하지 않습니다. VPC: %q: %+v", id, err)
		}

		for _, v := range *resp.GetVpcDetailResponse.VpcList {
			if id == *v.VpcNo {
				return v, string(*v.VpcStatus.Code), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}
