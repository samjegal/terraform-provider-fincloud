package provider

import (
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/common"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/location"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/network"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		compute.Registration{},
		network.Registration{},
		location.Registration{},
	}
}
