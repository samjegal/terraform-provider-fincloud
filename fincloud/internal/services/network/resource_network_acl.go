package network

import (
	"context"
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

var networkAclResourceName = "fincloud_network_acl"

func resourceNetworkAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkAclCreateOrUpdate,
		Read:   resourceNetworkAclRead,
		Update: resourceNetworkAclCreateOrUpdate,
		Delete: resourceNetworkAclDelete,

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

func resourceNetworkAclCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	aclClient := client.Network.AclClient

	name := d.Get("name").(string)
	param := expandNetworkAclParam(d, name)

	_, err := aclClient.CreateOrUpdate(ctx, param)
	if err != nil {
		return fmt.Errorf("생성 및 업데이트 중 에러가 발생했습니다. NetworkAcl: %q: %+v", name, err)
	}

	aclId, err := networkAclId(ctx, meta, name)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"USED"},
		Target:     []string{"RUN"},
		Refresh:    networkAclStateRefreshFunc(client, aclId),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. NetworkAcl %q", name)
	}

	return resourceNetworkAclRead(d, meta)
}

func resourceNetworkAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.AclClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	aclId, err := networkAclId(ctx, meta, name)
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, aclId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("상세정보를 읽을 수 있는 요청메시지를 만들 수가 없습니다. NetworkAcl: %q: %+v", aclId, err)
	}

	d.Set("name", name)

	if err := d.Set("vpc_id", resp.Content.VpcNo); err != nil {
		return fmt.Errorf("VPC Id 값을 설정하는 중 에러가 발생했습니다. err: %+v", err)
	}

	d.SetId(aclId)

	return nil
}

func resourceNetworkAclDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	aclClient := client.Network.AclClient

	aclId := d.Id()
	name := d.Get("name").(string)

	_, err := aclClient.Delete(ctx, aclId)
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"USED", "RUN"},
		Target:     []string{"NOTFOUND"},
		Refresh:    networkAclStateRefreshFunc(client, aclId),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. NetworkAcl %q", name)
	}

	d.SetId("")
	return nil
}

func expandNetworkAclParam(d *schema.ResourceData, name string) network.ACLRequestParameter {
	vpcId := d.Get("vpc_id").(string)
	description := d.Get("description").(string)

	return network.ACLRequestParameter{
		NetworkACLName: utils.String(name),
		VpcNo:          utils.String(vpcId),
		Description:    utils.String(description),
	}
}

func networkAclId(ctx context.Context, meta interface{}, name string) (string, error) {
	client := meta.(*clients.Client).Network.AclClient
	resp, err := client.List(ctx)
	if err != nil {
		return "", fmt.Errorf("전체 리스트 정보를 가져오는 중 에러가 발생했습니다. NetworkAcl: %+v", err)
	}

	for _, v := range *resp.Content {
		if name == *v.NetworkACLName {
			return fmt.Sprintf("%v", *v.NetworkACLNo), nil
		}
	}

	return "", fmt.Errorf("해당하는 정보가 존재하지 않습니다. NetworkAcl: %q: %+v", name, err)
}

func networkAclStateRefreshFunc(client *clients.Client, aclId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := client.StopContext
		resp, err := client.Network.AclClient.Get(ctx, aclId)
		if err != nil {
			if resp.StatusCode == http.StatusBadRequest {
				// error code 10001 확인 필요
				return resp, "NOTFOUND", nil
			}
			return nil, "NOTFOUND", fmt.Errorf("해당하는 정보가 존재하지 않습니다. NetworkAcl: %q: %+v", aclId, err)
		}

		return resp, *resp.Content.StatusCode, nil
	}
}
