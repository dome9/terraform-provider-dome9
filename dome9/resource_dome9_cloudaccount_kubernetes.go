package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudAccountKubernetes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountKubernetesCreate,
		Read:   resourceCloudAccountKubernetesRead,
		Update: resourceCloudAccountKubernetesUpdate,
		Delete: resourceCloudAccountKubernetesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organizational_unit_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organizational_unit_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organizational_unit_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"runtime_protection": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"admission_control": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"image_assurance": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceCloudAccountKubernetesCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := createKubernetesCloudAccountRequest(d)
	log.Printf("[INFO] Creating Kubernetes Cloud Account with request\n%+v\n", req)
	resp, _, err := d9Client.cloudaccountKubernetes.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Kubernetes CloudAccount. ID: %v\n", resp.ID)

	err = featuresCreate(d, d9Client, resp.ID)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)

	return resourceCloudAccountKubernetesRead(d, meta)
}

func resourceCloudAccountKubernetesRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	resp, _, err := d9Client.cloudaccountKubernetes.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() { // 404 response code
			log.Printf("[WARN] Removing Kubernetes cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading Kubernetes account response and settings states: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("cluster_version", resp.ClusterVersion)
	_ = d.Set("runtime_protection", expandRuntimeProtectionConfig(resp))
	_ = d.Set("admission_control", expandAdmissionControlConfig(resp))
	_ = d.Set("image_assurance", expandImageAssuranceConfig(resp))

	return nil
}

func resourceCloudAccountKubernetesDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Kubernetes Cloud Account ID: %v\n", d.Id())

	err := featuresDelete(d, d9Client)
	if err != nil {
		return err
	}

	if _, err := d9Client.cloudaccountKubernetes.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountKubernetesUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred for Kubernetes account")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if _, _, err := d9Client.cloudaccountKubernetes.UpdateName(d.Id(), k8s.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The Organizational Unit ID has been changed")

		if _, _, err := d9Client.cloudaccountKubernetes.UpdateOrganizationalID(d.Id(), k8s.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		}
	}

	err := featuresUpdate(d, d9Client)
	if err != nil {
		return err
	}

	return resourceCloudAccountKubernetesRead(d, meta)
}

func createKubernetesCloudAccountRequest(d *schema.ResourceData) k8s.CloudAccountRequest {
	return k8s.CloudAccountRequest{
		Name:                 d.Get("name").(string),
		OrganizationalUnitID: d.Get("organizational_unit_id").(string),
	}
}

func featuresCreate(d *schema.ResourceData, d9Client *Client, newId string) error {
	runtimeProtection, ok := d.GetOk("runtime_protection")
	if ok {
		if err := configureRuntimeProtection(runtimeProtection, newId, d9Client); err != nil {
			return err
		}
	}

	admissionControl, ok := d.GetOk("admission_control")
	if ok {
		if err := configureAdmissionControl(admissionControl, newId, d9Client); err != nil {
			return err
		}
	}

	imageAssurance, ok := d.GetOk("image_assurance")
	if ok {
		if err := configureImageAssurance(imageAssurance, newId, d9Client); err != nil {
			return err
		}
	}

	return nil
}

func featuresUpdate(d *schema.ResourceData, d9Client *Client) error {
	if d.HasChange("runtime_protection") {
		log.Println("Runtime Protection has been changed")

		runtimeProtection := d.Get("runtime_protection")
		if err := configureRuntimeProtection(runtimeProtection, d.Id(), d9Client); err != nil {
			return err
		}
	}

	if d.HasChange("admission_control") {
		log.Println("Admission Control has been changed")

		admissionControl := d.Get("admission_control")
		if err := configureAdmissionControl(admissionControl, d.Id(), d9Client); err != nil {
			return err
		}
	}

	if d.HasChange("image_assurance") {
		log.Println("Image Assurance has been changed")

		imageAssurance := d.Get("image_assurance")
		if err := configureImageAssurance(imageAssurance, d.Id(), d9Client); err != nil {
			return err
		}
	}
	return nil
}

func featuresDelete(d *schema.ResourceData, d9Client *Client) error {
	runtimeProtection, ok := d.GetOk("runtime_protection")
	if ok {
		if err := disableRuntimeProtectionIfEnabled(runtimeProtection, d.Id(), d9Client); err != nil {
			return err
		}
	}

	admissionControl, ok := d.GetOk("admission_control")
	if ok {
		if err := disableAdmissionControlIfEnabled(admissionControl, d.Id(), d9Client); err != nil {
			return err
		}
	}

	imageAssurance, ok := d.GetOk("image_assurance")
	if ok {
		if err := disableImageAssuranceIfEnabled(imageAssurance, d.Id(), d9Client); err != nil {
			return err
		}
	}

	return nil
}

func configureRuntimeProtection(runtimeProtection interface{}, clusterId string, d9Client *Client) error {
	runtimeProtectionConfig := runtimeProtection.([]interface{})[0].(map[string]interface{})
	req := createRuntimeProtectionEnableRequest(clusterId, runtimeProtectionConfig["enabled"].(bool))
	log.Println("[INFO] Configuring Runtime Protection for Kubernetes Cloud Account")
	if _, err := d9Client.cloudaccountKubernetes.EnableRuntimeProtection(req); err != nil {
		return err
	}

	return nil
}

func configureAdmissionControl(admissionControl interface{}, clusterId string, d9Client *Client) error {
	admissionControlConfig := admissionControl.([]interface{})[0].(map[string]interface{})
	log.Println("[INFO] Configuring Admission Control for Kubernetes Cloud Account")

	enableReq := createAdmissionControlEnableRequest(clusterId, admissionControlConfig["enabled"].(bool))
	if _, err := d9Client.cloudaccountKubernetes.EnableAdmissionControl(enableReq); err != nil {
		return err
	}

	return nil
}

func configureImageAssurance(ImageAssurance interface{}, clusterId string, d9Client *Client) error {
	ImageAssuranceConfig := ImageAssurance.([]interface{})[0].(map[string]interface{})
	req := createImageAssuranceEnableRequest(clusterId, ImageAssuranceConfig["enabled"].(bool))
	log.Println("[INFO] Configuring Image Assurance for Kubernetes Cloud Account")
	if _, err := d9Client.cloudaccountKubernetes.EnableImageAssurance(req); err != nil {
		return err
	}

	return nil
}

func createRuntimeProtectionEnableRequest(clusterId string, enabled bool) k8s.RuntimeProtectionEnableRequest {

	return k8s.RuntimeProtectionEnableRequest{
		CloudAccountId: clusterId,
		Enabled:        enabled,
	}
}

func createAdmissionControlEnableRequest(clusterId string, enabled bool) k8s.AdmissionControlEnableRequest {
	return k8s.AdmissionControlEnableRequest{
		CloudAccountId: clusterId,
		Enabled:        enabled,
	}
}

func createImageAssuranceEnableRequest(clusterId string, enabled bool) k8s.ImageAssuranceEnableRequest {
	return k8s.ImageAssuranceEnableRequest{
		CloudAccountId: clusterId,
		Enabled:        enabled,
	}
}
func expandRuntimeProtectionConfig(resp *k8s.CloudAccountResponse) []interface{} {
	runtimeProtectionConfig := make(map[string]interface{})

	runtimeProtectionConfig["enabled"] = resp.RuntimeProtectionEnabled

	return []interface{}{runtimeProtectionConfig}
}

func expandAdmissionControlConfig(resp *k8s.CloudAccountResponse) []interface{} {
	admissionControlConfig := make(map[string]interface{})

	admissionControlConfig["enabled"] = resp.AdmissionControlEnabled

	return []interface{}{admissionControlConfig}
}

func expandImageAssuranceConfig(resp *k8s.CloudAccountResponse) []interface{} {
	ImageAssuranceConfig := make(map[string]interface{})

	ImageAssuranceConfig["enabled"] = resp.ImageAssuranceEnabled

	return []interface{}{ImageAssuranceConfig}
}

func disableRuntimeProtectionIfEnabled(runtimeProtection interface{}, clusterId string, d9Client *Client) error {
	runtimeProtectionConfig := runtimeProtection.([]interface{})[0].(map[string]interface{})

	if runtimeProtectionConfig["enabled"].(bool) {
		req := createRuntimeProtectionEnableRequest(clusterId, false)
		log.Println("[INFO] Disabling Runtime Protection for Kubernetes Cloud Account")
		if _, err := d9Client.cloudaccountKubernetes.EnableRuntimeProtection(req); err != nil {
			return err
		}
	}

	return nil
}

func disableAdmissionControlIfEnabled(admissionControl interface{}, clusterId string, d9Client *Client) error {
	admissionControlConfig := admissionControl.([]interface{})[0].(map[string]interface{})

	if admissionControlConfig["enabled"].(bool) {
		req := createAdmissionControlEnableRequest(clusterId, false)
		log.Println("[INFO] Disabling Admission Control for Kubernetes Cloud Account")
		if _, err := d9Client.cloudaccountKubernetes.EnableAdmissionControl(req); err != nil {
			return err
		}
	}

	return nil
}

func disableImageAssuranceIfEnabled(ImageAssurance interface{}, clusterId string, d9Client *Client) error {
	ImageAssuranceConfig := ImageAssurance.([]interface{})[0].(map[string]interface{})

	if ImageAssuranceConfig["enabled"].(bool) {
		req := createImageAssuranceEnableRequest(clusterId, false)
		log.Println("[INFO] Disabling Image Assurance for Kubernetes Cloud Account")
		if _, err := d9Client.cloudaccountKubernetes.EnableImageAssurance(req); err != nil {
			return err
		}
	}

	return nil
}
