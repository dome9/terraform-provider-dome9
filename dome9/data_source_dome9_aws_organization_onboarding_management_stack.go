package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"regexp"
)

func dataSourceAwsOrganizationOnboardingManagementStack() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsOrganizationOnboardingManagementStackRead,

		Schema: map[string]*schema.Schema{
			"aws_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_cft_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_management_onboarded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsOrganizationOnboardingManagementStackRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	accountId := d.Get("aws_account_id").(string)
	resp, _, err := d9Client.awsOrganizationOnboarding.GetOnboardingConfiguration(accountId)
	if err != nil {
		return err
	}

	d.SetId(accountId)
	_ = d.Set("external_id", resp.ExternalId)
	_ = d.Set("content", resp.Content)
	_ = d.Set("is_management_onboarded", resp.IsManagementOnboarded)

	templateUrl, err := getTemplateUrl(resp.ManagementCftUrl)
	if err != nil {
		return err
	}
	_ = d.Set("management_cft_url", templateUrl)
	return nil
}

func getTemplateUrl(clickableUrl string) (string, error) {
	re := regexp.MustCompile(`templateURL=(?P<templateUrl>https://.+\.yml)(&|$)`)
	match := re.FindStringSubmatch(clickableUrl)

	for i, name := range re.SubexpNames() {
		if name == "templateUrl" {
			return match[i], nil
		}
	}

	return "", fmt.Errorf("no match found")
}
