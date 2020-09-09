package provider

import (
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/common"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/vpc"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		// compute.Registration{},
		// location.Registration{},
		// network.Registration{},
		// loadbalancer.Registration{},
		// sslvpn.Registration{},

		// new api version
		vpc.Registration{},
	}
}
