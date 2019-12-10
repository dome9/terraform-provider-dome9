package dome9

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"dome9": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv(providerconst.ProviderAccessIDEnvVariable); v == "" {
		t.Fatal(providerconst.ProviderAccessIDEnvVariable, "must be set for acceptance tests")
	}
	if v := os.Getenv(providerconst.ProviderSecretKeyEnvVariable); v == "" {
		t.Fatal(providerconst.ProviderSecretKeyEnvVariable, "must be set for acceptance tests")
	}
}
