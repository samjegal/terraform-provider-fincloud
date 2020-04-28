package authentication

import (
	"fmt"

	"github.com/samjegal/go-fincloud-helpers/fincloud"
)

func DetermineEnvironment(name string) (*fincloud.Environment, error) {
	env, err := fincloud.EnvironmentFromName(name)
	if err != nil {
		return nil, fmt.Errorf("A Financial Cloud Environment with name %q was not found: %+v", name, err)
	}

	return &env, nil
}
