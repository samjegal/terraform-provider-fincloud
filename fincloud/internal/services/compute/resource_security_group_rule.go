package compute

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/suppress"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var securityGroupRuleResourceName = "fincloud_security_group_rule"

func resourceSecurityGroupRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupRuleCreateOrUpdate,
		Read:   resourceSecurityGroupRuleRead,
		Update: resourceSecurityGroupRuleCreateOrUpdate,
		Delete: resourceSecurityGroupRuleDelete,

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
			"vpc_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"security_group_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"direction": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"inbound",
								"outbound",
							}, true),
						},

						"protocol": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(compute.ICMP),
								string(compute.TCP),
								string(compute.UDP),
							}, true),
							DiffSuppressFunc: suppress.CaseDifference,
						},

						"cidr_block": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"security_group_id": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"port": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},
					},
				},
				Set: securityGroupRuleHashSet,
			},
		},
	}
}

func resourceSecurityGroupRuleCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	ruleClient := client.Compute.SecurityGroupRuleClient

	securityGroupId := d.Get("security_group_id").(string)
	name, err := securityGroupName(client, securityGroupId)
	if err != nil {
		return err
	}

	props := *expandSecurityGroupRulesProperties(client, d)

	err = resourceSecurityGroupRuleDelete(d, meta)
	if err != nil {
		return err
	}

	resp, err := ruleClient.CreateOrUpdate(ctx, securityGroupId, props)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("Security Group 룰 생성에 문제가 발생했습니다. SecurityGroup: %q: %+v", name, err)
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

	d.SetId(securityGroupId)

	return resourceSecurityGroupRuleRead(d, meta)
}

func resourceSecurityGroupRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	securityGroupId := d.Get("security_group_id").(string)

	inboundClient := client.Compute.InboundClient
	inbound, err := inboundClient.Get(ctx, securityGroupId)
	if err != nil {
		return err
	}

	outboundClient := client.Compute.OutboundClient
	outbound, err := outboundClient.Get(ctx, securityGroupId)
	if err != nil {
		return nil
	}

	d.Set("rule",
		append(
			flattenSecurityGroupRuleContentParameter(inbound.Content),
			flattenSecurityGroupRuleContentParameter(outbound.Content)...))

	return nil
}

func resourceSecurityGroupRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	ruleClient := client.Compute.SecurityGroupRuleClient

	securityGroupId := d.Get("security_group_id").(string)
	name, err := securityGroupName(client, securityGroupId)
	if err != nil {
		return err
	}

	_, err = ruleClient.CreateOrUpdate(ctx, securityGroupId,
		compute.SecurityGroupRulesProperties{
			VpcNo:                           utils.String(d.Get("vpc_id").(string)),
			AccessControlGroupInboundRules:  &[]compute.SecurityGroupRuleProperties{},
			AccessControlGroupOutboundRules: &[]compute.SecurityGroupRuleProperties{},
		})
	if err != nil {
		return err
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

	d.SetId("")

	return nil
}

func securityGroupRuleHashSet(input interface{}) int {
	var buf bytes.Buffer

	if m, ok := input.(map[string]interface{}); ok {
		if v, ok := m["direction"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["protocol"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["cidr_block"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["security_group_id"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["port"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["description"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}
	}

	return hashcode.String(buf.String())
}

func expandSecurityGroupRulesProperties(client *clients.Client, d *schema.ResourceData) *compute.SecurityGroupRulesProperties {
	return &compute.SecurityGroupRulesProperties{
		VpcNo:                           utils.String(d.Get("vpc_id").(string)),
		AccessControlGroupInboundRules:  expandSecurityGroupRuleParameter(client, d, "inbound"),
		AccessControlGroupOutboundRules: expandSecurityGroupRuleParameter(client, d, "outbound"),
	}
}

func expandSecurityGroupRuleParameter(client *clients.Client, d *schema.ResourceData, direction string) *[]compute.SecurityGroupRuleProperties {
	param := make([]compute.SecurityGroupRuleProperties, 0)

	rules := d.Get("rule").(*schema.Set).List()
	for _, rule := range rules {
		v := rule.(map[string]interface{})

		var inout bool = false

		if dir, ok := v["direction"].(string); ok {
			if dir != direction {
				continue
			}

			if direction == "inbound" {
				inout = true
			}
		}

		r := compute.SecurityGroupRuleProperties{
			IsInboundRule: &inout,
			Description:   utils.String(v["description"].(string)),
		}

		protocol := v["protocol"].(string)
		r.ProtocolTypeCode = compute.ProtocolTypeCode(protocol)

		sequence := v["security_group_id"].(string)
		if sequence != "" {
			seq, _ := strconv.Atoi(sequence)
			r.AccessControlGroupSequence = utils.Int32(int32(seq))

			name, err := securityGroupName(client, sequence)
			if err != nil {
				// TODO: 에러처리 필요
				panic(fmt.Sprintf("SecurityGroup의 이름이 존재하지 않습니다. sequence: %q", sequence))
			}
			r.AccessControlGroupName = &name
		}

		if sequence == "" {
			cidrBlock := v["cidr_block"].(string)
			r.IPBlock = &cidrBlock
		}

		if protocol != "icmp" {
			port := v["port"].(string)
			r.PortRange = &port
		}

		param = append(param, r)
	}

	return &param
}

func flattenSecurityGroupRuleContentParameter(rule *[]compute.SecurityGroupParameter) []interface{} {
	if rule == nil {
		return []interface{}{}
	}

	param := make([]interface{}, 0)

	for _, v := range *rule {
		r := make(map[string]interface{})

		if v.AccessControlGroupRuleNo != nil {
			r["id"] = fmt.Sprintf("%d", *v.AccessControlGroupRuleNo)
		}

		if v.IsInboundRule != nil {
			if *v.IsInboundRule == true {
				r["direction"] = "inbound"
			} else {
				r["direction"] = "outbound"
			}
		}

		protocol := strings.ToLower(string(compute.ProtocolTypeCode(v.ProtocolTypeCode)))

		switch protocol {
		case "icmp", "tcp", "udp":
			r["protocol"] = protocol
		default:
			panic("Protocol 값이 존재하지 않습니다. protocol: " + protocol)
		}

		if v.IPBlock != nil {
			r["cidr_block"] = string(*v.IPBlock)
		}

		if v.AccessControlGroupSequence != nil {
			r["security_group_id"] = fmt.Sprintf("%d", *v.AccessControlGroupSequence)
		}

		if v.PortRange != nil {
			r["port"] = string(*v.PortRange)
		}

		if v.Description != nil {
			r["description"] = string(*v.Description)
		}

		param = append(param, r)
	}

	return param
}
