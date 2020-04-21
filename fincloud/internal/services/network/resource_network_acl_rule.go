package network

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/suppress"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var networkAclRuleResourceName = "fincloud_network_acl_rule"

func resourceNetworkAclRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkAclRuleCreateOrUpdate,
		Read:   resourceNetworkAclRuleRead,
		Update: resourceNetworkAclRuleCreateOrUpdate,
		Delete: resourceNetworkAclRuleDelete,

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
			"network_acl_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"network_acl_rules": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
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
								string(network.ICMP),
								string(network.TCP),
								string(network.UDP),
							}, true),
							DiffSuppressFunc: suppress.CaseDifference,
						},

						"allow": {
							Type:     schema.TypeBool,
							Required: true,
						},

						"cidr_block": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"port": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},

						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validate.NoEmptyStrings,
						},
					},
				},
				Set: networkAclRuleHashSet,
			},
		},
	}
}

func resourceNetworkAclRuleCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	networkAclId := d.Get("network_acl_id").(string)

	rules := expandNetworkAclTotalRule(d)
	resp, err := client.Network.AclRuleClient.CreateOrUpdate(ctx, networkAclId, rules)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Network ACL Rule을 생성/업데이트를 하는 중 문제가 발생했습니다. status: %d", resp.StatusCode)
	}

	d.SetId(networkAclId)

	return resourceNetworkAclRuleRead(d, meta)
}

func resourceNetworkAclRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	networkAclId := d.Id()

	content, err := flattenNetworkAclRuleList(ctx, client, networkAclId)
	if err != nil {
		d.SetId("")
		return err
	}

	if err = d.Set("rule", flattenNetworkAclRules(content.Content)); err != nil {
		d.SetId("")
		return err
	}

	return nil
}

func resourceNetworkAclRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()
	ruleClinet := client.Network.AclRuleClient

	networkAclId := d.Id()

	_, err := ruleClinet.CreateOrUpdate(ctx, networkAclId,
		network.ACLRuleListParameter{
			NetworkACLRules: &[]network.ACLRuleProperties{},
		})
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

// func substractContentSlice(ptr *[]network.ACLRuleParameter, index int) {
// 	slice := *ptr
// 	*ptr = append(slice[:index], slice[index+1:]...)
// }

func networkAclRuleHashSet(input interface{}) int {
	var buf bytes.Buffer

	if m, ok := input.(map[string]interface{}); ok {
		if v, ok := m["direction"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["priority"]; ok {
			buf.WriteString(fmt.Sprintf("%d-", v.(int)))
		}

		if v, ok := m["protocol"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["port"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["cidr_block"]; ok {
			buf.WriteString(fmt.Sprintf("%s-", v.(string)))
		}

		if v, ok := m["allow"]; ok {
			buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
		}

		if v, ok := m["description"]; ok {
			buf.WriteString(fmt.Sprintf("%s", v.(string)))
		}
	}

	return hashcode.String(buf.String())
}

func expandNetworkAclRule(r map[string]interface{}) *network.ACLRuleProperties {
	var dir bool
	if r["direction"].(string) == "inbound" {
		dir = true
	} else {
		dir = false
	}

	rule := network.ACLRuleProperties{
		IsInboundRule:    utils.Bool(dir),
		ProtocolTypeCode: network.ProtocolTypeCode(r["protocol"].(string)),
		IPBlock:          utils.String(r["cidr_block"].(string)),
		Priority:         utils.Int32(int32(r["priority"].(int))),
	}

	if v, ok := r["allow"].(bool); ok {
		rule.IsAllowRule = utils.Bool(v)
	}

	if v, ok := r["description"].(string); ok && v != "" {
		rule.Description = utils.String(v)
	}

	if v, ok := r["protocol"].(string); ok && v != "" {
		if v != "icmp" {
			rule.PortRange = utils.String(r["port"].(string))
		}
	}

	return &rule
}

func expandNetworkAclTotalRule(d *schema.ResourceData) network.ACLRuleListParameter {
	rules := make([]network.ACLRuleProperties, 0)

	ruleList := d.Get("rule").(*schema.Set).List()
	for _, rule := range ruleList {
		r := rule.(map[string]interface{})
		rules = append(rules, *expandNetworkAclRule(r))
	}

	return network.ACLRuleListParameter{
		NetworkACLRules: &rules,
	}
}

func flattenNetworkAclRuleList(ctx context.Context, client *clients.Client, id string) (*network.ACLRuleContentParameter, error) {
	inbound, err := client.Network.InboundClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	outbound, err := client.Network.OutboundClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	param := flattenNetworkAclRule(inbound.Content, outbound.Content)
	return &param, nil
}

func flattenNetworkAclRule(inbound *[]network.ACLRuleContentProperties, outbound *[]network.ACLRuleContentProperties) network.ACLRuleContentParameter {
	rules := make([]network.ACLRuleContentProperties, 0)

	rules = append(
		resourceRuleParameter(inbound),
		resourceRuleParameter(outbound)...,
	)

	return network.ACLRuleContentParameter{
		Content: &rules,
	}
}

func resourceRuleParameter(param *[]network.ACLRuleContentProperties) []network.ACLRuleContentProperties {
	rules := make([]network.ACLRuleContentProperties, 0)

	if param != nil {
		for _, v := range *param {
			rules = append(rules, v)
		}
	}

	return rules
}

func flattenNetworkAclRules(param *[]network.ACLRuleContentProperties) []interface{} {
	if param == nil {
		return []interface{}{}
	}

	output := make([]interface{}, 0)

	for _, v := range *param {
		rule := make(map[string]interface{})

		if v.RuleNo != nil {
			rule["id"] = int32(*v.RuleNo)
		}

		if v.IsInboundRule != nil {
			if *v.IsInboundRule == true {
				rule["direction"] = "inbound"
			} else {
				rule["direction"] = "outbound"
			}
		}

		protocol := string(network.ProtocolTypeCode(v.ProtocolTypeCode))

		switch protocol {
		case "icmp", "tcp", "udp":
			rule["protocol"] = protocol
		default:
			panic("Protocol 값이 존재하지 않습니다. protocol: " + protocol)
		}

		if v.IPBlock != nil {
			rule["cidr_block"] = string(*v.IPBlock)
		}

		if v.PortRange != nil {
			rule["port"] = string(*v.PortRange)
		}

		if v.IsAllowRule != nil {
			rule["allow"] = bool(*v.IsAllowRule)
		}

		if v.Priority != nil {
			rule["priority"] = int(*v.Priority)
		}

		if v.Description != nil {
			rule["description"] = string(*v.Description)
		}

		output = append(output, rule)
	}

	return output
}
