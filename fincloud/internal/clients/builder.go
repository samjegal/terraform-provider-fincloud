package clients

import (
	"context"
	"fmt"

	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/go-fincloud-helpers/sender"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type ClientBuilder struct {
	AuthConfig       *authentication.Config
	TerraformVersion string
}

func Build(ctx context.Context, builder ClientBuilder) (*Client, error) {
	env, err := authentication.DetermineEnvironment(builder.AuthConfig.Environment)
	if err != nil {
		return nil, err
	}

	account, err := NewResourceManagerAccount(ctx, *builder.AuthConfig, *env)
	if err != nil {
		return nil, fmt.Errorf("Error building account: %+v", err)
	}

	client := Client{
		Account: account,
	}

	sender := sender.BuildSender("FinancialCloud")

	endpoint := env.ResourceManagerEndpoint
	auth, err := builder.AuthConfig.GetAuthorizationToken(sender, env.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	o := &common.ClientOptions{
		ResourceManagerAuthorizer: auth,
		ResourceManagerEndpoint:   endpoint,
		Environment:               *env,
	}

	if err := client.Build(ctx, o); err != nil {
		return nil, fmt.Errorf("Error building Client: %+v", err)
	}

	return &client, nil
}
