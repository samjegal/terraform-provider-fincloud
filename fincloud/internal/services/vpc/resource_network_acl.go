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

var networkACLResourceName = "fincloud_network_acl"

func resourceNetworkACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkACLCreateOrUpdate,
		Read:   resourceNetworkACLRead,
		Update: resourceNetworkACLCreateOrUpdate,
		Delete: resourceNetworkACLDelete,

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

			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceNetworkACLCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.NetworkACLClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	resp, err := client.Create(ctx,
		d.Get("vpc_id").(string),      // vpcNo
		name,                          // networkAclName
		d.Get("description").(string), // networkAclDescription
	)
	if err != nil {
		return fmt.Errorf("생성 및 업데이트 중 에러가 발생했습니다. NetworkACL: %q: %+v", name, err)
	}

	content := (*resp.CreateNetworkACLResponse.NetworkACLList)[0]
	id := *content.NetworkACLNo

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"USED"},
		Target:     []string{"RUN"},
		Refresh:    networkACLStateRefreshFunc(meta.(*clients.Client), id),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. NetworkACL: %q: %+v", name, err)
	}

	d.SetId(id)
	return nil
}
func resourceNetworkACLRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.NetworkACLClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := d.Id()
	name := d.Get("name").(string)

	resp, err := client.GetDetail(ctx, id)
	if err != nil {
		return fmt.Errorf("상세 정보의 요청에 문제가 발생했습니다. NetworkACL: %q: %+v", name, err)
	}

	content := resp.GetNetworkACLDetailResponse.NetworkACLList
	if content != nil {
		if err := d.Set("name", (*content)[0].NetworkACLName); err != nil {

		}
	}

	return nil
}

func resourceNetworkACLDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Vpc.NetworkACLClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := d.Id()
	name := d.Get("name").(string)

	_, err := client.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("삭제 요청에 문제가 발생했습니다. NetworkACL: %q: %+v", name, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"USED", "RUN"},
		Target:     []string{"NOTFOUND"},
		Refresh:    networkACLStateRefreshFunc(meta.(*clients.Client), id),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. NetworkACL: %q: %+v", name, err)
	}

	d.SetId("")
	return nil
}

func networkACLStateRefreshFunc(client *clients.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, err := client.Vpc.NetworkACLClient.GetDetail(client.StopContext, id)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. NetworkACL: %q: %+v", id, err)
		}

		if resp.GetNetworkACLDetailResponse == nil {
			return nil, "EMPTY", fmt.Errorf("해당하는 정보가 존재하지 않습니다. NetworkACL: %q: %+v", id, err)
		}

		for _, v := range *resp.GetNetworkACLDetailResponse.NetworkACLList {
			if id == *v.NetworkACLNo {
				return v, string(*v.NetworkACLStatus.Code), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}
