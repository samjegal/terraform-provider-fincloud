package network

import (
	"fmt"
	"time"

	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
)

var virtualPrivateCloudDataSourceName = "fincloud_virtual_private_cloud"

func dataSourceVirtualPrivateCloud() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVirtualPrivateRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVirtualPrivateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.VirtualPrivateCloudClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	vpcName := d.Get("name").(string)

	resp, err := client.List(ctx)
	if err != nil {
		return fmt.Errorf("VPC 정보의 요청에 문제가 발생했습니다. vpcName: %q: %+v", vpcName, err)
	}

	if resp.Content == nil {
		return fmt.Errorf("VPC 정보가 존재하지 않습니다.")
	}

	d.SetId("/vpcName/" + vpcName)

	for _, content := range *resp.Content {
		if vpcName == *content.VpcName {
			d.Set("name", *content.VpcName)
			d.Set("vpc_id", *content.VpcNo)
			d.Set("cidr_block", *content.Ipv4Cidr)
			d.Set("status_code", string(content.StatusCode))
			break
		}
	}

	return nil
}
