package dome9

import (
	"encoding/json"
	"github.com/dome9/dome9-sdk-go/services/integrations"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"sort"
	"strings"
)

func resourceIntegration() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegrationCreate,
		Read:   resourceIntegrationRead,
		Update: resourceIntegrationUpdate,
		Delete: resourceIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
			},
		},
	}
}

// Helper functions
func trimAndSortJSONKeys(rawMessage json.RawMessage) (json.RawMessage, error) {
	// Unmarshal the JSON into a map
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(rawMessage, &jsonMap); err != nil {
		return nil, err
	}

	// Create a new map with trimmed keys
	trimmedMap := make(map[string]interface{})
	for k, v := range jsonMap {
		trimmedKey := strings.TrimSpace(k)
		trimmedMap[trimmedKey] = v
	}

	// Get the keys and sort them
	keys := make([]string, 0, len(trimmedMap))
	for k := range trimmedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create a sorted map
	sortedMap := make(map[string]interface{}, len(trimmedMap))
	for _, k := range keys {
		sortedMap[k] = trimmedMap[k]
	}

	// Marshal the sorted map back into a json.RawMessage
	sortedJSON, err := json.Marshal(sortedMap)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(sortedJSON), nil
}

// Expansion functions
func expandIntegrationUpdateRequest(id string, d *schema.ResourceData) (integrations.IntegrationUpdateRequestModel, error) {
	putModel := integrations.IntegrationUpdateRequestModel{
		Id:            id,
		Name:          d.Get("name").(string),
		Type:          integrations.IntegrationType(d.Get("type").(string)),
		Configuration: []byte(d.Get("configuration").(string)),
	}

	return putModel, nil
}

func expandIntegrationCreateRequest(d *schema.ResourceData) (integrations.IntegrationPostRequestModel, error) {

	postModel := integrations.IntegrationPostRequestModel{
		Name:          d.Get("name").(string),
		Type:          integrations.IntegrationType(d.Get("type").(string)),
		Configuration: []byte(d.Get("configuration").(string)),
	}

	return postModel, nil
}

func resourceIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req, err := expandIntegrationCreateRequest(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating integration request\n%+v\n", req)
	resp, _, err := d9Client.integration.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created integration. ID: %v\n", resp.Id)
	d.SetId(resp.Id)

	return resourceIntegrationRead(d, meta)
}

func resourceIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Reading integration ID: %v", d.Id())

	resp, _, err := d9Client.integration.GetById(d.Id())
	if err != nil {
		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("name", resp.Name)
	_ = d.Set("type", resp.Type)

	trimmedConfiguration, err := trimAndSortJSONKeys(resp.Configuration)
	if err != nil {
		return err
	}
	_ = d.Set("configuration", string(trimmedConfiguration))

	return nil
}

func resourceIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req, err := expandIntegrationUpdateRequest(d.Id(), d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Updating integration request\n%+v\n", req)
	resp, _, err := d9Client.integration.Update(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Updated integration. ID: %v\n", resp.Id)

	return resourceIntegrationRead(d, meta)
}

func resourceIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting integration ID: %v", d.Id())

	if _, err := d9Client.integration.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}
