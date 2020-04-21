package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := TestFincloudProvider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
