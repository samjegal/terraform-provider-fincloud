package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/go-fincloud-helpers/fincloud"
)

var FincloudProvider *schema.Provider
var SupportedProviders map[string]terraform.ResourceProvider

func PreCheck(t *testing.T) {
	variables := []string{
		"FINCLOUD_CERTIFICATE_PATH",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for acceptance tests!", variable)
		}
	}
}

func Environment() (*fincloud.Environment, error) {
	return authentication.DetermineEnvironment("FINCLOUD")
}

func GetAuthConfig(t *testing.T) *authentication.Config {
	if os.Getenv(resource.TestEnvVar) == "" {
		t.Skip(fmt.Sprintf("Integration test skipped unless env '%s' set", resource.TestEnvVar))
		return nil
	}

	builder := authentication.Builder{
		CertTokenPath: os.Getenv("FINCLOUD_CERTIFICATE_PATH"),
	}
	config, err := builder.Build()
	if err != nil {
		t.Fatalf("Error building Fincloud Management Client: %+v", err)
		return nil
	}

	return config
}
