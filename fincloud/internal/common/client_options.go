package common

import (
	"github.com/Azure/go-autorest/autorest"
	fincloud "github.com/samjegal/go-fincloud-helpers/fincloud"
	"github.com/samjegal/go-fincloud-helpers/sender"
)

type ClientOptions struct {
	TerraformVersion string

	ResourceManagerAuthorizer autorest.Authorizer
	ResourceManagerEndpoint   string

	Environment fincloud.Environment
}

func (co ClientOptions) ConfigureClient(c *autorest.Client, authorizer autorest.Authorizer) {
	// c.Authorizer = authorizer
	c.Sender = sender.BuildSender("FinancialCloudResourceManagement")

	// Test
	c.AccessKey = "84E8670425073E2EF0B7"
	c.Secretkey = "9B48B2BE3E9DA67EAD2963A749A5ED0157F66C93"
}
