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
	c.Authorizer = authorizer
	c.Sender = sender.BuildSender("FinancialCloudResourceManagement")
}
