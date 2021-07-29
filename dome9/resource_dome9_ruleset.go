package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/rulebundles"
	"github.com/dome9/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceRuleSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleSetCreate,
		Read:   resourceRuleSetRead,
		Update: resourceRuleSetUpdate,
		Delete: resourceRuleSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_vendor": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(providerconst.CloudVendors, false),
			},
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hide_in_compliance": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_template": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"min_feature_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logic": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Low",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remediation": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"compliance_tag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"control_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"rule_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logic_hash": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}

func resourceRuleSetCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandRuleSetCreateRequest(d)
	log.Printf("[INFO] Creating dome9 rule set with request\n%+v\n", req)

	ruleSet, _, err := d9Client.ruleSet.Create(&req)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(ruleSet.ID))
	return resourceRuleSetRead(d, meta)

}

func resourceRuleSetRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.ruleSet.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing rule set %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting rule set:\n%+v\n", resp)
	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("cloud_vendor", resp.CloudVendor)
	_ = d.Set("language", resp.Language)
	_ = d.Set("hide_in_compliance", resp.HideInCompliance)
	_ = d.Set("is_template", resp.IsTemplate)
	_ = d.Set("min_feature_tier", resp.MinFeatureTier)
	_ = d.Set("created_time", resp.CreatedTime)
	_ = d.Set("updated_time", resp.UpdatedTime)

	if err := d.Set("rules", flattenRules(resp.Rules)); err != nil {
		return err
	}

	return nil
}

func resourceRuleSetUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	req := expandRuleSetCreateRequest(d)
	req.ID = id
	log.Printf("[INFO] Updating rule set with name %s\n", req.Name)

	if _, _, err := d9Client.ruleSet.Update(&req); err != nil {
		return err
	}

	return resourceRuleSetRead(d, meta)
}

func resourceRuleSetDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	log.Printf("[INFO] Deleting rule set with id %v\n", d.Id())

	if _, err := d9Client.ruleSet.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandRuleSetCreateRequest(d *schema.ResourceData) rulebundles.RuleBundleRequest {
	return rulebundles.RuleBundleRequest{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Rules:            expandRules(d),
		HideInCompliance: d.Get("hide_in_compliance").(bool),
		MinFeatureTier:   d.Get("min_feature_tier").(string),
		CloudVendor:      d.Get("cloud_vendor").(string),
		Language:         d.Get("language").(string),
	}
}

func expandRules(d *schema.ResourceData) *[]rulebundles.Rule {
	var rules []rulebundles.Rule
	if itemsInterface, ok := d.GetOk("rules"); ok {
		items := itemsInterface.([]interface{})
		rules = make([]rulebundles.Rule, len(items))
		for i, item := range items {
			rule := item.(map[string]interface{})
			rules[i] = rulebundles.Rule{
				Name:          rule["name"].(string),
				Severity:      rule["severity"].(string),
				Logic:         rule["logic"].(string),
				Description:   rule["description"].(string),
				Remediation:   rule["remediation"].(string),
				ComplianceTag: rule["compliance_tag"].(string),
				Domain:        rule["domain"].(string),
				Priority:      rule["priority"].(string),
				ControlTitle:  rule["control_title"].(string),
				RuleID:        rule["rule_id"].(string),
				LogicHash:     rule["logic_hash"].(string),
				IsDefault:     rule["is_default"].(bool),
			}
		}
	}

	return &rules
}

func flattenRules(responseRules []rulebundles.Rule) []interface{} {
	rules := make([]interface{}, len(responseRules))
	for i, val := range responseRules {
		rules[i] = map[string]interface{}{
			"name":           val.Name,
			"severity":       val.Severity,
			"logic":          val.Logic,
			"description":    val.Description,
			"remediation":    val.Remediation,
			"compliance_tag": val.ComplianceTag,
			"domain":         val.Domain,
			"priority":       val.Priority,
			"control_title":  val.ControlTitle,
			"rule_id":        val.RuleID,
			"logic_hash":     val.LogicHash,
			"is_default":     val.IsDefault,
		}
	}

	return rules
}
