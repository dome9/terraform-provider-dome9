package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceContinuousComplianceFinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousComplianceFindingSearch,

		Schema: map[string]*schema.Schema{
			"pageSize": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"sorting": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fieldName": {
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
			"multiSorting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fieldName": {
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
						"freeTextPhrase": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fields": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"Name": {
										Type:     schema.TypeString,
										Optional: false,
									},
									"Value": {
										Type:     schema.TypeString,
										Optional: false,
									},
								},
							},
						},
						"onlyCIEM": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"includedFeatures": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"creationTime": {
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
			"searchAfter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dataSource": {
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
