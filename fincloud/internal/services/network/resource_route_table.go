package network

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var routeTableResourceName = "fincloud_route_table"

func resourceRouteTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouteTableCreate,
		Read:   resourceRouteTableRead,
		Update: resourceRouteTableUpdate,
		Delete: resourceRouteTableDelete,

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

			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"usage": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"private",
				}, true),
			},

			"subnet": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"route": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"cidr_block": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"endpoint": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				// Set: routeTableRuleHashSet,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	routeTableClient := client.Network.RouteTableClient

	name := d.Get("name").(string)

	param, err := expandRouteTableParameter(d)
	if err != nil {
		return err
	}

	// Route Table 생성
	resp, err := routeTableClient.Create(ctx, *param)
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("라우트 테이블 생성에 문제가 발생했습니다. RouteTable: %q: %+v", name, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"SET"},
		Target:     []string{"RUN"},
		Refresh:    routeTableStateRefreshFunc(client, name, expandRouteTableSearchParameter()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("라우트 테이블의 상태 코드를 가져오는 중 문제가 발생했다. Subnet: %q", name)
	}

	// 라우트 테이블 번호
	routeTableId, err := routeTableId(client, name)
	if err != nil {
		return err
	}

	d.SetId(routeTableId)

	// 연관 서브넷 설정
	if _, ok := d.GetOk("subnet"); ok {
		err = routeTableSubnetUpdate(client, d)
		if err != nil {
			return err
		}
	}

	// 라우트 룰 설정
	if _, ok := d.GetOk("route"); ok {
		resp, err = client.Network.RouteTableClient.Update(ctx, routeTableId,
			*expandRouteTableRuleParameter(client, d))
		if err != nil {
			return err
		}
	}

	return resourceRouteTableRead(d, meta)
}

func resourceRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.RouteTableClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	resp, err := client.List(ctx, expandRouteTableSearchParameter())
	if err != nil {
		return err
	}

	var props *network.RouteTableSearchContentParameter
	for _, v := range *resp.Content {
		if name == *v.RouteTableName {
			props = &v
			break
		}
	}

	if props != nil {
		d.SetId(fmt.Sprintf("%d", *props.RouteTableNo))
	}

	return nil
}

func resourceRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	_, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	// id := d.Id()

	if d.HasChange("subnet") {
		err := routeTableSubnetUpdate(client, d)
		if err != nil {
			return err
		}
	}

	// if d.HasChange("route") {
	// }

	return nil
}

func resourceRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()
	routeTableClient := client.Network.RouteTableClient

	id := d.Id()
	name := d.Get("name").(string)

	_, err := routeTableClient.Delete(ctx, id)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"SET", "RUN"},
		Target:     []string{"NOTFOUND"},
		Refresh:    routeTableStateRefreshFunc(client, name, expandRouteTableSearchParameter()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("라우트 테이블의 상태 코드를 가져오는 중 문제가 발생했다. Subnet: %q", name)
	}

	d.SetId("")

	return nil
}

func routeTableId(client *clients.Client, name string) (string, error) {
	ctx := client.StopContext
	content, err := client.Network.RouteTableClient.List(ctx, expandRouteTableSearchParameter())
	if err != nil {
		return "", err
	}

	var props *network.RouteTableSearchContentParameter
	for _, v := range *content.Content {
		if name == *v.RouteTableName {
			props = &v
			break
		}
	}

	return fmt.Sprintf("%d", *props.RouteTableNo), nil
}

func routeTableSubnetUpdate(client *clients.Client, d *schema.ResourceData) error {
	ctx := client.StopContext

	id := d.Id()
	name := d.Get("name").(string)

	subnetList := make([]string, 0)
	subnets := d.Get("subnet").([]interface{})
	for _, s := range subnets {
		subnetList = append(subnetList, s.(string))
	}

	routeTableNumber, _ := strconv.Atoi(id)
	_, err := client.Network.RouteTableSubnetClient.Update(ctx, id,
		network.RouteTableSubnetParameter{
			VpcNo:        utils.String(d.Get("vpc_id").(string)),
			RouteTableNo: utils.Int32(int32(routeTableNumber)),
			SubnetNos:    &subnetList,
		})
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"SET"},
		Target:     []string{"RUN"},
		Refresh:    routeTableStateRefreshFunc(client, name, expandRouteTableSearchParameter()),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("라우트 테이블의 상태 코드를 가져오는 중 문제가 발생했다. Subnet: %q", name)
	}

	return nil
}

// TODO: 라우트 룰 정보를 업데이트 하는 함수에 정리해야 한다.
func routeTableRuleUpdate() {

}

func expandRouteTableParameter(d *schema.ResourceData) (*network.RouteTableParameter, error) {
	param := &network.RouteTableParameter{
		VpcNo:          utils.String(d.Get("vpc_id").(string)),
		RouteTableName: utils.String(d.Get("name").(string)),
		Description:    utils.String(d.Get("description").(string)),
	}

	usage := d.Get("usage").(string)
	var isIGW string = "N"

	if usage == "public" {
		isIGW = "Y"
	}

	param.IgwYn = &isIGW

	return param, nil
}

func routeTableStateRefreshFunc(client *clients.Client, name string, param network.RouteTableSearchParameter) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		routeTableClient := client.Network.RouteTableClient

		ctx := client.StopContext
		resp, err := routeTableClient.List(ctx, param)
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. RouteTable: %q: %+v", name, err)
		}

		if resp.Content == nil {
			return nil, "NOTFOUND", fmt.Errorf("라우트 테이블 정보가 존재하지 않습니다. RouteTable: %q: %+v", name, err)
		}

		for _, v := range *resp.Content {
			if name == *v.RouteTableName {
				return v, string(network.RouteTableStatusCode(v.StatusCode)), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}

func expandRouteTableSearchParameter() network.RouteTableSearchParameter {
	var page int32 = 1
	var pageSize int32 = 100

	filter := make([]network.RouteTableSearchFilterParameter, 0)

	return network.RouteTableSearchParameter{
		PageNo:     utils.Int32(page),
		PageSizeNo: utils.Int32(pageSize),
		Filter:     &filter,
	}
}

func routeTableRuleHashSet(input interface{}) int {
	var buf bytes.Buffer

	if m, ok := input.(map[string]interface{}); ok {
		if v, ok := m["cidr_block"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["endpoint"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}
	}

	return hashcode.String(buf.String())
}

func expandRouteTableRuleParameter(client *clients.Client, d *schema.ResourceData) *network.RouteTableRuleParameter {
	id, _ := strconv.Atoi(d.Id())
	return &network.RouteTableRuleParameter{
		VpcNo:        utils.String(d.Get("vpc_id").(string)),
		RouteTableNo: utils.Int32(int32(id)),
		Route:        expandRouteTableRuleContentParameter(client, d),
	}
}

func expandRouteTableRuleContentParameter(client *clients.Client, d *schema.ResourceData) *[]network.RouteTableRuleContentParameter {
	output := make([]network.RouteTableRuleContentParameter, 0)

	routes := d.Get("route").(*schema.Set).List()
	for _, r := range routes {
		_ = r.(map[string]interface{})
	}

	return &output
}
