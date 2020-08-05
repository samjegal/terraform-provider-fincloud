package provider

import (
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/common"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/loadbalancer"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/location"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/sslvpn"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		compute.Registration{},
		location.Registration{},
		network.Registration{},
		loadbalancer.Registration{},
		sslvpn.Registration{},
	}
}
