package compute

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var securityGroupResourceName = "fincloud_security_group"

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupCreateOrUpdate,
		Read:   resourceSecurityGroupRead,
		Update: resourceSecurityGroupCreateOrUpdate,
		Delete: resourceSecurityGroupDelete,

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

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"network_interface": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSecurityGroupCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	securityGroupClient := client.Compute.SecurityGroupClient

	name := d.Get("name").(string)

	resp, err := securityGroupClient.CreateOrUpdate(ctx, *expandSecurityGroupNameParameter(d, name))
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("Security Group 생성에 문제가 발생했습니다. SecurityGroup: %q: %+v", name, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "SET"},
		Target:     []string{"RUN"},
		Refresh:    securityGroupStateRefreshFunc(client, name),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. SecurityGroup: %q", name)
	}

	return resourceSecurityGroupRead(d, meta)
}

func resourceSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.SecurityGroupClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	resp, err := client.List(ctx)
	if err != nil {
		return fmt.Errorf("Security Group의 전체 리스트 정보를 가져오는 중 에러가 발생했습니다. SecurityGroup: %q: %+v", name, err)
	}

	var props *compute.SecurityGroupListContentParameter = nil

	for _, v := range *resp.Content {
		if name == *v.AccessControlGroupName {
			props = &v
			break
		}
	}

	if props != nil {
		if err := d.Set("network_interface", *props.NetworkInterfaceCount); err != nil {
			return fmt.Errorf("SeucrityGroup의 연동된 NIC 수를 설정하는 중 에러가 발생했습니다. err: %+v", err)
		}

		d.SetId(fmt.Sprintf("%v", *props.AccessControlGroupNo))
	}

	return nil
}

func resourceSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	securityGroupClient := client.Compute.SecurityGroupClient

	name := d.Get("name").(string)

	_, err := securityGroupClient.Delete(ctx, *expandSecurityGroupNumberParameter(d))
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"INIT", "SET", "RUN"},
		Target:     []string{"NOTFOUND"},
		Refresh:    securityGroupStateRefreshFunc(client, name),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("상세정보의 상태 코드를 가져오는 중 문제가 발생했다. SecurityGroup: %q, err: %+v", name, err)
	}

	d.SetId("")

	return nil
}

func securityGroupName(client *clients.Client, id string) (string, error) {
	ctx := client.StopContext
	resp, err := client.Compute.SecurityGroupClient.List((ctx))
	if err != nil {
		return "", err
	}

	if resp.Content == nil {
		return "", fmt.Errorf("SecurityGroup 정보가 존재하지 않습니다. SecurityGroup: %q: %+v", id, err)
	}

	securityGroupId, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	for _, v := range *resp.Content {
		if int32(securityGroupId) == *v.AccessControlGroupNo {
			return *v.AccessControlGroupName, nil
		}
	}

	return "", fmt.Errorf("SecurityGroup 정보가 존재하지 않습니다. SecurityGroup: %q", id)
}

func expandSecurityGroupNameParameter(d *schema.ResourceData, name string) *compute.SecurityGroupNameParameter {
	return &compute.SecurityGroupNameParameter{
		AccessControlGroupName: utils.String(name),
		VpcNo:                  utils.String(d.Get("vpc_id").(string)),
		Description:            utils.String(d.Get("description").(string)),
	}
}

func expandSecurityGroupNumberParameter(d *schema.ResourceData) *compute.SecurityGroupNumberParameter {
	return &compute.SecurityGroupNumberParameter{
		AccessControlGroupNo: utils.String(d.Id()),
		VpcNo:                utils.String(d.Get("vpc_id").(string)),
	}
}

func securityGroupStateRefreshFunc(client *clients.Client, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := client.StopContext
		resp, err := client.Compute.SecurityGroupClient.List((ctx))
		if err != nil {
			return nil, "NOTFOUND", fmt.Errorf("SecurityGroup 정보가 존재하지 않습니다. SecurityGroup: %q: %+v", name, err)
		}

		if resp.Content == nil {
			return nil, "NOTFOUND", fmt.Errorf("SecurityGroup 정보가 존재하지 않습니다. SecurityGroup: %q: %+v", name, err)
		}

		for _, v := range *resp.Content {
			if name == *v.AccessControlGroupName {
				return v, string(compute.SecurityRuleStatusCode(v.StatusCode)), nil
			}
		}

		return resp, "NOTFOUND", nil
	}
}
