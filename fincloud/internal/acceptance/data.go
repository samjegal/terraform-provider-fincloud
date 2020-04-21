package acceptance

import (
	"fmt"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/samjegal/go-fincloud-helpers/fincloud"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/provider"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var once sync.Once

type TestData struct {
	RandomInteger int
	RandomString  string

	Environment     fincloud.Environment
	EnvironmentName string

	ResourceName  string
	ResourceType  string
	resourceLabel string
}

func BuildTestData(t *testing.T, resourceType string, resourceLabel string) TestData {
	once.Do(func() {
		fincloudProvider := provider.TestFincloudProvider().(*schema.Provider)

		FincloudProvider = fincloudProvider
		SupportedProviders = map[string]terraform.ResourceProvider{
			"fincloud": fincloudProvider,
		}
	})

	env, err := Environment()
	if err != nil {
		t.Fatalf("Error retrieving Environment: %+v", err)
	}

	testData := TestData{
		RandomInteger: utils.AccRandTimeInt(),
		RandomString:  acctest.RandString(5),

		Environment:     *env,
		EnvironmentName: "FINCLOUD",

		ResourceName:  fmt.Sprintf("%s.%s", resourceType, resourceLabel),
		ResourceType:  resourceType,
		resourceLabel: resourceLabel,
	}

	return testData
}
