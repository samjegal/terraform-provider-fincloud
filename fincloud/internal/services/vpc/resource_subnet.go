package vpc

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
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
			"zone_code": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "FKR-1",
				ValidateFunc: validate.NoEmptyStrings,
			},

			// vpcNo
			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			// subnetName
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			// subnet
			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			// networkAclNo
			"network_acl_id": {
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
					"gen",   // general
					"loadb", // loadbalancer
					"bm",    // baremetal
				}, true),
			},
		},
	}
}

func resourceSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSubnetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
