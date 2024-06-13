package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceAwsOrganizationOnboarding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsOrganizationOnboardingRead,

		Schema: map[string]*schema.Schema{
			// OrganizationManagementViewModel object fields
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_organization_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_management_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_account_stack_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_account_stack_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"onboarding_configuration": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_root_ou_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mapping_strategy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"posture_management": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rulesets_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
									},
									"onboarding_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"organization_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stack_set_regions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"stack_set_organizational_unit_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceAwsOrganizationOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	orgId := d.Get("id").(string)
	resp, _, err := d9Client.awsOrganizationOnboarding.Get(orgId)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Aws organization %s from state because it no longer exists in CloudGuard", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("accountId", resp.AccountId)
	_ = d.Set("externalOrganizationId", resp.ExternalOrganizationId)
	_ = d.Set("externalManagementAccountId", resp.ExternalManagementAccountId)
	_ = d.Set("managementAccountStackId", resp.ManagementAccountStackId)
	_ = d.Set("managementAccountStackRegion", resp.ManagementAccountStackRegion)
	_ = d.Set("userId", resp.UserId)
	_ = d.Set("organizationName", resp.OrganizationName)
	_ = d.Set("updateTime", resp.UpdateTime)
	_ = d.Set("creationTime", resp.CreationTime)
	_ = d.Set("stackSetRegions", resp.StackSetRegions)
	_ = d.Set("stackSetOrganizationalUnitIds", resp.StackSetOrganizationalUnitIds)

	if err := d.Set("onboarding_configuration", flattenAwsOrganizationOnboardingConfiguration(resp.OnboardingConfiguration)); err != nil {
		return err
	}

	return nil
}
