package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountAWSBasic(t *testing.T) {
	var cloudAccountResponse aws.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)
	originalArn := os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn)
	updatedArn := os.Getenv(environmentvariable.CloudAccountUpdatedAWSEnvVarArn)
	originalGroupBehavior := variable.CloudAccountAWSReadOnlyGroupBehavior
	updatedGroupBehavior := variable.CloudAccountAWSFullManageGroupBehavior

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, variable.CloudAccountAWSOriginalAccountName, originalArn, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAWSExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountAWSVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "allow_read_only", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "credentials.0.arn", originalArn),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAWSOriginalAccountName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.#", "1"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.#", fmt.Sprint(len(providerconst.AWSRegions))),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.0.region", variable.CloudAccountAWSFetchedRegion),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.0.new_group_behavior", originalGroupBehavior),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.1.new_group_behavior", originalGroupBehavior),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.2.new_group_behavior", originalGroupBehavior),
				),
			},
			{
				Config: testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, variable.CloudAccountAWSUpdatedAccountName, updatedArn, testAccCloudAccountAWSNetsecConfig(updatedGroupBehavior)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAWSExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAWSUpdatedAccountName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "credentials.0.arn", updatedArn),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.#", fmt.Sprint(len(providerconst.AWSRegions))),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.0.new_group_behavior", updatedGroupBehavior),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.1.new_group_behavior", updatedGroupBehavior),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.2.new_group_behavior", originalGroupBehavior),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAWSExists(resource string, cloudAccount *aws.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*cloudAccount = *receivedCloudAccountResponse

		return nil
	}
}

func testAccCloudAccountAWSEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAWSEnvVarArn)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAWSEnvVarSecret); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAWSEnvVarSecret)
	}
	if v := os.Getenv(environmentvariable.CloudAccountUpdatedAWSEnvVarArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountUpdatedAWSEnvVarArn)
	}
}

func testAccCheckCloudAccountDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAWS {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("cloudaccount with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, resourceName, arn, additionalBlock string) string {
	return fmt.Sprintf(`
// aws cloud account creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}

`,
		// aws cloud account
		getCloudAccountAWSResourceHCL(generatedName, resourceName, arn, additionalBlock),

		// data source variables
		resourcetype.CloudAccountAWS,
		generatedName,
		resourceTypeAndName,
	)
}

func getCloudAccountAWSResourceHCL(generatedName, resourceName, arn, additionalBlock string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name        = "%s"
  credentials {
    arn      = "%s"
    secret   = "%s"
    type     = "RoleBased"
  }

  %s

}
`,
		// aws cloud account variables
		resourcetype.CloudAccountAWS,
		generatedName,
		resourceName,
		arn,
		os.Getenv(environmentvariable.CloudAccountAWSEnvVarSecret),
		additionalBlock,
	)
}

func testAccCloudAccountAWSNetsecConfig(groupBehavior string) string {
	return fmt.Sprintf(`
net_sec {
    regions {
      new_group_behavior = "%s"
      region             = "us_east_1"
    }
    regions {
      new_group_behavior = "%s"
      region             = "us_west_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_southeast_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_northeast_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_west_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "sa_east_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_southeast_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_central_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_northeast_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_south_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_east_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ca_central_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_3"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_north_1"
    }
  }
`,
		groupBehavior,
		groupBehavior,
	)
}
