package fincloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.FincloudProvider()
}
