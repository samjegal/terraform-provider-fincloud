package fincloud

import (
	"fmt"
	"strings"
)

const (
	EnvironmentFilepathName = "FINCLOUD_ENVIRONMENT_FILEPATH"

	NotAvailable = "N/A"
)

var environments = map[string]Environment{
	"FINCLOUD": FinancialCloud,
}

type ResourceIdentifier struct {
	Resource string `json:"resource"`
}

type Environment struct {
	Name                    string `json:"name"`
	ResourceManagerEndpoint string `json:"resourceManagerEndpoint"`
}

var (
	FinancialCloud = Environment{
		Name:                    "FinancialCloud",
		ResourceManagerEndpoint: "https://console.fin-ncloud.com",
	}
)

func EnvironmentFromName(name string) (Environment, error) {
	name = strings.ToUpper(name)
	env, ok := environments[name]
	if !ok {
		return env, fmt.Errorf("fincloud/helpers: There is no cloud environment matching the name %q", name)
	}

	return env, nil
}
