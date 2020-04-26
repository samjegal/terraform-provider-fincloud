package vpn

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

func (r Registration) Name() string {
	return "Vpn"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		vpnConnectionDataSourceName: dataSourceVpnConnection(),
	}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		vpnConnectionResourceName: resourceVpnConnection(),
	}
}
