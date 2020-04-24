package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Registration struct{}

func (r Registration) Name() string {
	return "Network"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		virtualPrivateCloudDataSourceName: dataSourceVirtualPrivateCloud(),
	}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		virtualPrivateCloudResourceName: resourceVirtualPrivateCloud(),
		subnetResourceName:              resourceSubnet(),
		networkAclResourceName:          resourceNetworkAcl(),
		networkAclRuleResourceName:      resourceNetworkAclRule(),
		netGatewayResourceName:          resourceNatGateway(),
	}
}
