package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/integrations"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
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

// Expansion functions
func expandIntegrationUpdateRequest(id string, d *schema.ResourceData) (integrations.IntegrationUpdateRequestModel, error) {
	putModel := integrations.IntegrationUpdateRequestModel{
		Id:            id,
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Configuration: []byte(d.Get("configuration").(string)),
	}

	return putModel, nil
}

func expandIntegrationCreateRequest(d *schema.ResourceData) (integrations.IntegrationPostRequestModel, error) {

	postModel := integrations.IntegrationPostRequestModel{
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Configuration: []byte(d.Get("configuration").(string)),
	}

	return postModel, nil
}

// CRUD API

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
	_ = d.Set("configuration", string(resp.Configuration))

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
