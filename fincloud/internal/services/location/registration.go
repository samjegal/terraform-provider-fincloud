package location

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Registration struct{}

func (r Registration) Name() string {
	return "Location"
}

func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fincloud_region": dataSourceRegion(),
		"fincloud_zone":   dataSourceZone(),
	}
}

func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}
