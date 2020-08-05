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
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var sslvpnConnectionResourceName = "fincloud_sslvpn_connection"

func resourceSslVpnConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceSslVpnConnectionCreateOrUpdate,
		Read:   resourceSslVpnConnectionRead,
		Update: resourceSslVpnConnectionCreateOrUpdate,
		Delete: resourceSslVpnConnectionDelete,

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

func resourceSslVpnConnectionCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SslVpn.SslVpnConnectionClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	if d.Id() == "" {
		return fmt.Errorf("서버 SSL VPN은 수동으로 생성해서 인스턴스 번호를 이용해 terraform import 명령어로 작업 해야합니다.")
	}

	if d.HasChange("limit") {
		switch d.Get("limit").(int) {
		case 3, 5, 10:
			break
		default:
			return fmt.Errorf("서버 SSL VPN 스펙은 3, 5, 10의 사용자 제한만 지원합니다.")
		}

		_, err := client.UpdateSpec(ctx, expandSslVpnConnectionLimitUserParameter(d))
		if err != nil {
			return err
		}
	}

	return resourceSslVpnConnectionRead(d, meta)
}

func resourceSslVpnConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).SslVpn.SslVpnConnectionClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	if d.Id() == "" {
		return fmt.Errorf("서버 SSL VPN은 수동으로 생성해서 인스턴스 번호를 이용해 terraform import 명령어로 작업 해야합니다.")
	}

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

func resourceSslVpnConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func expandSslVpnConnectionLimitUserParameter(d *schema.ResourceData) sslvpn.LimitUserCountParameter {
	return sslvpn.LimitUserCountParameter{
		UserCountLimitation: utils.Int32(int32(d.Get("limit").(int))),
	}
}
