package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
)

func resourceAssessment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssessmentCreate,
		Read:   resourceAssessmentRead,
		Update: resourceAssessmentUpdate,
		Delete: resourceAssessmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloudAccountId": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloudAccountType": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(providerconst.AssessmentCloudAccountType, false),
			},
			"shouldMinimizeResult": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tested_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"relevant_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"non_complying_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"exclusion_stats": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tested_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"relevant_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"non_complying_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"entity_results": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"validation_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_relevant": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_valid": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_excluded": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"exclusion_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remediation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"test_obj": {
										Type:     schema.TypeSet,
										Computed: true,
									},
								},
							},
						},
						"rule": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"logic": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remediation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cloudbots": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"complianceTag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"controlTitle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ruleId": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"labels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"logic_hash": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"test_passed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"location_metadata": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"srl": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"test_entities": {
				Type:     schema.TypeSet,
				Computed: true,
			},
			"data_sync_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recently_successful_sync": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"general_fetch_permission_issues": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entities_with_permission_issues": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cloud_vendor_identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"assessment_passed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_errors": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAssessmentCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountAlibabaRequest(d)
	log.Printf("[INFO] Creating Alibaba Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountAlibaba.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Alibaba CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceAssessmentRead(d, meta)
}

func resourceAssessmentRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.assessment.RunBundle(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Alibaba cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("alibaba_account_id", resp.AlibabaAccountId)
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("creation_date", resp.CreationDate)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)

	return nil
}

func resourceAssessmentDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Alibaba Cloud Account ID: %v\n", d.Id())
	if _, err := d9Client.cloudaccountAlibaba.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceAssessmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("An update can not be made to an assessment")
	return nil
}

