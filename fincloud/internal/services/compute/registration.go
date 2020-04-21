package compute

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Registration struct{}

func (r Registration) Name() string {
	return "Compute"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		securityGroupResourceName:     resourceSecurityGroup(),
		securityGroupRuleResourceName: resourceSecurityGroupRule(),

		networkInterfaceResoureName: resourceNetworkInterface(),
		serverResourceName:          resourceServer(),
		storageResourceName:         resourceStorage(),
		publicIpResourceName:        resourcePublicIp(),

		loginKeyResourceName:   resourceLoginKey(),
		initScriptResourceName: resourceInitScript(),
	}
}
