package sslvpn

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/sslvpn"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
)

var sslvpnConnectionDataSourceName = "fincloud_sslvpn_connection"

func dataSourceSslVpnConnection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSslVpnConnectionRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"member": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"users": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"limit": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceSslVpnConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SslVpn.SslVpnConnectionClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.List(ctx)
	if err != nil {
		return err
	}

	var props *sslvpn.ContentParameter
	for _, v := range *resp.Content {
		if int32(id) == *v.InstanceNo {
			props = &v
			break
		}
	}

	if props != nil {
		if err := d.Set("name", *props.SslVpnName); err != nil {
			return fmt.Errorf("서버 SSL VPN의 이름을 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("member", *props.MemberNo); err != nil {
			return fmt.Errorf("서버 SSL VPN의 멤버 번호를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("users", *props.UserCount); err != nil {
			return fmt.Errorf("서버 SSL VPN의 현재 사용자 수를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("limit", *props.UserCountLimitation); err != nil {
			return fmt.Errorf("서버 SSL VPN의 최대 사용자 수를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}
	}

	return nil
}
