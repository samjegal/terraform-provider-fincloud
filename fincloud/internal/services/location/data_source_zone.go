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

var zoneDataSourceName = "fincloud_zone"

func dataSourceZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceZoneRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"code": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"region_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func dataSourceZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Location.ZoneClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	resp, err := client.Get(ctx)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("금융존 정보를 가져오는데 실패했습니다.: %+v", err)
		}
		return fmt.Errorf("금융존 정보를 가져오기 위한 요청을 만드는 중에 에러가 발생했습니다.: %+v", err)
	}

	for _, v := range *resp.Content {
		if err := d.Set("name", *v.ZoneName); err != nil {
			return fmt.Errorf("금융존 이름을 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("code", *v.ZoneCode); err != nil {
			return fmt.Errorf("금융존 코드를 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("description", *v.ZoneDescription); err != nil {
			return fmt.Errorf("금융존 설명 정보를 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("region_id", *v.RegionNo); err != nil {
			return fmt.Errorf("금융존의 리전 정보를 설정하는 중에 에러가 발생했습니다. err: %+v", err)
		}

		d.SetId(*v.ZoneNo)
		break
	}

	return nil
}
