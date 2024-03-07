package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/awp_aws_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"strings"
)

func resourceDome9AWPAWSOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWPAWSOnboardingCreate,
		Read:   resourceAWPAWSOnboardingRead,
		Delete: resourceAWPAWSOnboardingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cloudguard_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cross_account_role_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cross_account_role_external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_guard_awp_stack_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scan_mode": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"inAccount",
					"saas",
					"in-account-hub",
					"in-account-sub",
				}, false),
			},
			"is_terraform": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"agentless_account_settings": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled_regions": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"scan_machine_interval_in_hours": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_concurrence_scans_per_region": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"skip_function_apps_scan": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"custom_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceAWPAWSOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	cloudguardAccountId := d.Get("cloudguard_account_id").(string)
	req := expandAWPOnboardingRequest(d)
	log.Printf("[INFO] Creating AWP AWS Onboarding request %+v\n", req)
	_, err := client.awpAwsOnboarding.CreateAWPOnboarding(cloudguardAccountId, req)
	if err != nil {
		return err
	}
	d.SetId(cloudguardAccountId)
	log.Printf("[INFO] Created AWP AWS Onboarding with CloudGuard Account ID: %v\n", cloudguardAccountId)
	d.SetId(cloudguardAccountId) // set the resource ID to the CloudGuard Account ID

	return resourceAWPAWSOnboardingRead(d, meta)
}

func expandAWPOnboardingRequest(d *schema.ResourceData) awp_aws_onboarding.CreateAWPOnboardingRequest {
	return awp_aws_onboarding.CreateAWPOnboardingRequest{
		// populate the request fields from the schema
		// replace the field names with the ones used in your schema
		Field1: d.Get("field1").(string),
		Field2: d.Get("field2").(int),
		// continue for all fields
	}
}

func resourceAWPAWSOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	resp, _, err := client.GetAWPOnboarding("aws", d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}
	// set the schema fields from the response
	return nil
}

func resourceAWPAWSOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	_, err := client.DeleteAWPOnboarding(d.Id(), true)
	if err != nil {
		return err
	}
	return nil
}
