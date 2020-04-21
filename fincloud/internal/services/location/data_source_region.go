package location

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var regionDataSourceName = "fincloud_region"

func dataSourceRegion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRegionRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"code": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"usage": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func dataSourceRegionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Location.RegionClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	resp, err := client.Get(ctx)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Region 정보를 가져오는데 실패했습니다.: %+v", err)
		}
		return fmt.Errorf("Region 정보를 가져오기 위한 요청을 만드는 중에 에러가 발생했습니다.: %+v", err)
	}

	if resp.Content != nil {
		if err := d.Set("code", *resp.Content.RegionCode); err != nil {
			return fmt.Errorf("Region 코드값을 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("name", *resp.Content.RegionName); err != nil {
			return fmt.Errorf("Region 이름값을 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		usage := false
		if *resp.Content.UseYn == "Y" {
			usage = true
		}

		if err := d.Set("usage", usage); err != nil {
			return fmt.Errorf("Region 사용여부를 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}
	}

	d.SetId(*resp.Content.RegionNo)

	return nil
}
