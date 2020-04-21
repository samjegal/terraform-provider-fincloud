package clients

import (
	"context"

	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/go-fincloud-helpers/fincloud"
)

type ResourceManagerAccount struct {
	Environment fincloud.Environment
}

func NewResourceManagerAccount(ctx context.Context, config authentication.Config, env fincloud.Environment) (*ResourceManagerAccount, error) {
	account := ResourceManagerAccount{
		Environment: env,
	}

	return &account, nil
}
