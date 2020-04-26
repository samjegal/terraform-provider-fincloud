package network

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var netGatewayResourceName = "fincloud_nat_gateway"

func resourceNatGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceNatGatewayCreateOrUpdate,
		Read:   resourceNatGatewayRead,
		Update: resourceNatGatewayCreateOrUpdate,
		Delete: resourceNatGatewayDelete,

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
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"zone_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "110",
				ValidateFunc: validate.NoEmptyStrings,
			},

			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"address": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceNatGatewayCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	natGatewayClient := client.Network.NatGatewayClient

	name := d.Get("name").(string)
	resp, err := natGatewayClient.Create(ctx, expandNatGatewayParam(d))
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("NAT 게이트웨이 생성에 문제가 발생했습니다. NatGateway: %q", name)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "TERMTING"},
		Target:     []string{"RUN"},
		Refresh:    natGatewayStateRefreshFunc(client, name, expandNatGateaySearchParameter()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("NAT 게이트웨이 상태 코드를 가져오는 중 문제가 발생했습니다. NatGateway: %q: %+v", name, err)
	}

	return resourceNatGatewayRead(d, meta)
}

func resourceNatGatewayRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.NatGatewayClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resp, err := client.List(ctx, expandNatGateaySearchParameter())
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("전체 리스트 정보를 읽을 수 있는 요청 메시지를 만들 수가 없습니다. NatGateway: %q: %+v", name, err)
	}

	var props *network.NatGatewaySearchContentParameter
	for _, v := range *resp.Content {
		if name == *v.NatGatewayName {
			props = &v
			break
		}
	}

	if props != nil {
		if err := d.Set("address", *props.PublicIP); err != nil {
			return fmt.Errorf("NAT 게이트웨이 IP 주소를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		if err := d.Set("endpoint", fmt.Sprintf("%d", *props.EndpointNo)); err != nil {
			return fmt.Errorf("NAT 게이트웨이 엔드포인트를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		d.SetId(fmt.Sprintf("%d", *props.InstanceNo))
	}

	return nil
}

func resourceNatGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	natGatewayClient := client.Network.NatGatewayClient

	name := d.Get("name").(string)

	empty := new(network.NatGatewayParameter)
	_, err := natGatewayClient.Delete(ctx, d.Id(), empty)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "RUN", "TERMTING"},
		Target:     []string{"NOTFOUND"},
		Refresh:    natGatewayStateRefreshFunc(client, name, expandNatGateaySearchParameter()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("NAT 게이트웨이 상태 코드를 가져오는 중 문제가 발생했습니다. NatGateway: %q: %+v", name, err)
	}

	d.SetId("")

	return nil
}

func expandNatGatewayParam(d *schema.ResourceData) network.NatGatewayParameter {
	return network.NatGatewayParameter{
		ZoneNo:         utils.String(d.Get("zone_id").(string)),
		VpcNo:          utils.String(d.Get("vpc_id").(string)),
		NatGatewayName: utils.String(d.Get("name").(string)),
		Description:    utils.String(d.Get("description").(string)),
	}
}

func natGatewayStateRefreshFunc(client *clients.Client, name string, param network.NatGatewaySearchParameter) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		natGatewayClient := client.Network.NatGatewayClient
		ctx := client.StopContext

		resp, err := natGatewayClient.List(ctx, param)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. NatGateway: %q: %+v", name, err)
		}

		if resp.Content == nil {
			return nil, "NOTFOUND", fmt.Errorf("NAT 게이트웨이 정보가 존재하지 않습니다. NatGateway: %q: %+v", name, err)
		}

		for _, v := range *resp.Content {
			if name == *v.NatGatewayName {
				return v, string(network.NatGatewayStatusCode(v.StatusCode)), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}

func expandNatGateaySearchParameter() network.NatGatewaySearchParameter {
	var page int32 = 1
	var pageSize int32 = 100

	filter := make([]network.NatGatewaySearchFilterParameter, 0)

	return network.NatGatewaySearchParameter{
		PageNo:     utils.Int32(page),
		PageSizeNo: utils.Int32(pageSize),
		Filter:     &filter,
	}
}
