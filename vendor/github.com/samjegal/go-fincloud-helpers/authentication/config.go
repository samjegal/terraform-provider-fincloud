package authentication

import (
	"github.com/Azure/go-autorest/autorest"
)

type Config struct {
	Environment string

	authMethod authMethod
}

func (c Config) GetAuthorizationToken(sender autorest.Sender, endpoint string) (autorest.Authorizer, error) {
	return c.authMethod.getAuthorizationToken(sender, endpoint)
}
