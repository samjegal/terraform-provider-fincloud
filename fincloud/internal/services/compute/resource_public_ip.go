package compute

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var publicIpResourceName = "fincloud_public_ip"

func resourcePublicIp() *schema.Resource {
	return &schema.Resource{
		Create: resourcePublicIpCreate,
		Read:   resourcePublicIpRead,
		Update: resourcePublicIpUpdate,
		Delete: resourcePublicIpDelete,

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
			"server_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"address": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"assign": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePublicIpCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.PublicIpClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	serverId := d.Get("server_id").(string)

	resp, err := client.Create(ctx, extendPublicIPAddressServerInstanceParameter(d, serverId))
	if err != nil || resp.Response.StatusCode == http.StatusBadRequest {
		return err
	}

	var props *compute.PublicIPAddressSearchProperties = nil

	for _, v := range *resp.Content {
		props = &v
		break
	}

	if props != nil {
		d.SetId(*props.InstanceNo)
	}

	return resourcePublicIpRead(d, meta)
}

func resourcePublicIpRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.PublicIpClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	publicIpId := d.Id()

	resp, err := client.List(ctx, expandPublicIPAddressSearchFilterParameter())
	if err != nil {
		return err
	}

	var props *compute.PublicIPAddressSearchProperties = nil

	for _, v := range *resp.Content {
		if publicIpId == *v.InstanceNo {
			props = &v
			break
		}
	}

	if props != nil {
		if err := d.Set("address", *props.PublicIP); err != nil {
			return fmt.Errorf("공인 IP 주소를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if props.ServerInstanceNo != nil {
			if err := d.Set("assign", fmt.Sprint(*props.ServerInstanceNo)); err != nil {
				return fmt.Errorf("공인 IP 적용 서버값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
			}
		}
	}

	return nil
}

func resourcePublicIpUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.PublicIpClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	publicIpId := d.Id()
	serverId := d.Get("server_id").(string)
	assign := d.Get("assign").(string)

	if serverId != "" {
		_, err := client.Assign(ctx, publicIpId, extendPublicIPAddressServerInstanceParameter(d, serverId))
		if err != nil {
			return err
		}
	} else {
		_, err := client.Remove(ctx, publicIpId, extendPublicIPAddressServerInstanceParameter(d, assign))
		if err != nil {
			return err
		}
	}

	return resourcePublicIpRead(d, meta)
}

func resourcePublicIpDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.PublicIpClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	publicIpId := d.Id()
	var serverId string = ""

	if v, ok := d.GetOk("server_id"); ok {
		serverId = v.(string)
		_, err := client.Remove(ctx, publicIpId, extendPublicIPAddressServerInstanceParameter(d, serverId))
		if err != nil {
			return err
		}
	}

	_, err := client.Delete(ctx, publicIpId, extendPublicIPAddressServerInstanceParameter(d, serverId))
	if err != nil {
		return err
	}

	return nil
}

func extendPublicIPAddressServerInstanceParameter(d *schema.ResourceData, serverId string) compute.PublicIPAddressServerInstanceParameter {
	return compute.PublicIPAddressServerInstanceParameter{
		ServerInstanceNo: utils.String(serverId),
	}
}

func expandPublicIPAddressSearchFilterParameter() compute.PublicIPAddressSearchFilterParameter {
	var page int = 1
	var pageSize int = 100

	filter := make([]compute.PublicIPSearchFilterProperties, 0)
	filter = append(filter, compute.PublicIPSearchFilterProperties{
		Field: utils.String("publicIp"),
	})

	return compute.PublicIPAddressSearchFilterParameter{
		PageNo:     utils.Int32(int32(page)),
		PageSizeNo: utils.Int32(int32(pageSize)),
		Filter:     &filter,
	}
}
