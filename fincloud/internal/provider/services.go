package provider

import (
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/common"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/location"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/vpn"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		compute.Registration{},
		location.Registration{},
		network.Registration{},
		vpn.Registration{},
	}
}
