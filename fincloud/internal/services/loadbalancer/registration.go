package loadbalancer

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

func (r Registration) Name() string {
	return "LoadBalancer"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}
