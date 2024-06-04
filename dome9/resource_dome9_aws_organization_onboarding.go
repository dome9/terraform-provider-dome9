package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws_org"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"regexp"
)

func resourceAwsOrganizationOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsOrganizationOnboardingCreate,
		Read:   resourceAwsOrganizationOnboardingRead,
		Update: resourceAwsOrganizationOnboardingUpdate,
		Delete: resourceAwsOrganizationOnboardingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// OnboardingRequest object fields
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^arn:(aws|aws-cn):iam::\d{12}:role\/[A-Za-z0-9]+(?:-[A-Za-z0-9]+)+$`),
					"Role ARN should be in the format: arn:aws:iam:account_number:role_name.",
				),
			},
			"secret": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[\w+=,.@:\/-]{12,1224}$`),
					"ExternalId minimum length is 12 and maximum length is 1224.",
				),
			},
			"api_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_set_arn": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`put your stack set ARN regex pattern here`),
					"Invalid StackSet ARN format.",
				),
			},
			"aws_organization_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_stack_modify": {
				Type:     schema.TypeBool,
				Required: true,
				Default:  false,
			},
			// OrganizationManagementViewModel object fields
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceAwsOrganizationOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAwsOrganizationOnboardingRequest(d)
	log.Printf("[INFO] Creating Aws organization with request %+v\n", req)

	resp, _, err := d9Client.awsOrganizationOnboarding.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Aws organization. ID: %v\n", resp.Id)
	d.SetId(resp.Id)

	return resourceAwsOrganizationOnboardingRead(d, meta)
}

func resourceAwsOrganizationOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.awsOrganizationOnboarding.Get(d.Id())

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

func resourceAwsOrganizationOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Aws organization ID: %v\n", d.Id())
	if _, err := d9Client.awsOrganizationOnboarding.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceAwsOrganizationOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred")

	if d.HasChange("stack_set_arn") {
		log.Println("The StackSet ARN has been changed")

		if resp, err := d9Client.awsOrganizationOnboarding.UpdateStackSetArn(d.Id(), aws_org.UpdateStackSetArnRequest{
			StackSetArn: d.Get("stack_set_arn").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceAwsOrganizationOnboardingUpdate StackSetArn response is: %+v\n", resp)
		}
	}

	if d.HasChange("organization_root_ou_id") || d.HasChange("mapping_strategy") || d.HasChange("posture_management") {
		log.Println("The configuration has been changed")

		updateConfigReq := aws_org.UpdateConfigurationRequest{
			OrganizationRootOuId: getOptionalString(d, "organization_root_ou_id"),
			MappingStrategy:      aws_org.MappingStrategyType(d.Get("mapping_strategy").(string)),
			PostureManagement: aws_org.PostureManagementConfiguration{
				RulesetsIds:    getInt64Slice(d, "posture_management.rulesets_ids"),
				OnboardingMode: aws_org.OnboardingMode(d.Get("posture_management.onboarding_mode").(string)),
			},
		}

		if resp, err := d9Client.awsOrganizationOnboarding.UpdateConfiguration(d.Id(), updateConfigReq); err != nil {
			return err
		} else {
			log.Printf("resourceAwsOrganizationOnboardingUpdate Configuration response is: %+v\n", resp)
		}
	}

	return nil
}

func getInt64Slice(d *schema.ResourceData, key string) []int64 {
	if v, ok := d.GetOk(key); ok {
		values := v.([]interface{})
		result := make([]int64, len(values))
		for i, val := range values {
			result[i] = val.(int64)
		}
		return result
	}
	return nil
}

func expandAwsOrganizationOnboardingRequest(d *schema.ResourceData) aws_org.OnboardingRequest {
	req := aws_org.OnboardingRequest{
		ValidateStackSetArnRequest: aws_org.ValidateStackSetArnRequest{
			OnboardingPermissionRequest: aws_org.OnboardingPermissionRequest{
				RoleArn: d.Get("role_arn").(string),
				Secret:  d.Get("secret").(string),
				ApiKey:  getOptionalString(d, "api_key"),
				Type:    aws_org.CloudCredentialsType(d.Get("type").(string)),
			},
			StackSetArn: d.Get("stack_set_arn").(string),
		},
		AwsOrganizationName: getOptionalString(d, "aws_organization_name"),
		EnableStackModify:   d.Get("enable_stack_modify").(bool),
	}
	return req
}

func getOptionalString(d *schema.ResourceData, key string) *string {
	if v, ok := d.GetOk(key); ok {
		str := v.(string)
		return &str
	}
	return nil
}

func flattenAwsOrganizationOnboardingConfiguration(config aws_org.AwsOrganizationOnboardingConfiguration) map[string]interface{} {
	return map[string]interface{}{
		"organization_root_ou_id": config.OrganizationRootOuId,
		"mapping_strategy":        config.MappingStrategy,
		"posture_management":      flattenPostureManagementConfiguration(config.PostureManagement),
	}
}

func flattenPostureManagementConfiguration(config aws_org.PostureManagementConfiguration) map[string]interface{} {
	return map[string]interface{}{
		"rulesets_ids":    config.RulesetsIds,
		"onboarding_mode": config.OnboardingMode,
	}
}
