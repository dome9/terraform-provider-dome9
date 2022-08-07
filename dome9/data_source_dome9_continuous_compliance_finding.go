package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceContinuousComplianceFinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousComplianceFindingSearch,

		Schema: map[string]*schema.Schema{
			"page_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"sorting": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"multi_sorting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"free_text_phrase": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fields": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: false,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: false,
									},
								},
							},
						},
						"only_ciem": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"included_features": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"creation_time": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:     schema.TypeString,
										Optional: false,
									},
									"to": {
										Type:     schema.TypeString,
										Optional: false,
									},
								},
							},
						},
					},
				},
			},
			"search_after": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceContinuousComplianceFindingSearch(d *schema.ResourceData, meta interface{}) error {
	//d9Client := meta.(*Client)
	//
	//policyID := d.Get("id").(string)
	//log.Printf("Getting data for Continuous Compliance Policy id: %s\n", policyID)
	//
	//resp, _, err := d9Client.continuousCompliancePolicy.Get(policyID)
	//if err != nil {
	//	return err
	//}
	//
	//d.SetId(resp.ID)
	//_ = d.Set("target_internal_id", resp.TargetInternalId)
	//_ = d.Set("target_external_id", resp.TargetExternalId)
	//_ = d.Set("target_type", resp.TargetType)
	//_ = d.Set("ruleset_id", resp.RulesetId)
	//if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
	//	return err
	//}

	return nil
}
