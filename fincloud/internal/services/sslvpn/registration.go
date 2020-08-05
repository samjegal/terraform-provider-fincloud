package sslvpn

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

func (r Registration) Name() string {
	return "SslVpn"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		sslvpnConnectionDataSourceName: dataSourceSslVpnConnection(),
	}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		sslvpnConnectionResourceName: resourceSslVpnConnection(),
	}
}
