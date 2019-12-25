package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceIAMSafeEntity() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMSafeEntityCreate,
		Read:   resourceIAMSafeEntityRead,
		Update: resourceIAMSafeEntityUpdate,
		Delete: resourceIAMSafeEntityDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"protection_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(providerconst.IAMEntityProtectionMode, true),
			},
			"entity_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(providerconst.IAMEntityProtectType, true),
			},
			"aws_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dome9_users_id_to_protect": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_dome9_users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exists_in_aws": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIAMSafeEntityCreate(d *schema.ResourceData, meta interface{}) error {
	protectionMode := d.Get("protection_mode").(string)

	if protectionMode == providerconst.IAMSafeEntityProtect {
		return iamSafeEntityProtect(d, meta)
	} else {
		return iamSafeEntityProtectWithElevation(d, meta)
	}
}

func iamSafeEntityProtect(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandIAMSafeEntityRequest(d)
	cloudAccountID := d.Get("aws_cloud_account_id").(string)

	resp, _, err := d9Client.cloudaccountAWS.ProtectIAMSafeEntity(cloudAccountID, req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Protect IAM entity (user or role) for aws cloud account %s with arn: %v\n", cloudAccountID, *resp)
	// generate and set random id
	d.SetId(cloudAccountID)

	return resourceIAMSafeEntityRead(d, meta)
}

func iamSafeEntityProtectWithElevation(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	usersToProtectWithElevation := expandUsersToProtectWithElevation(d.Get("dome9_users_id_to_protect").([]interface{}))
	cloudAccountID := d.Get("aws_cloud_account_id").(string)
	entityName := d.Get("entity_name").(string)
	entityType := d.Get("entity_type").(string)

	_, err := d9Client.users.ProtectWithElevationIAMSafeEntity(cloudAccountID, entityName, entityType, usersToProtectWithElevation)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Protect with elevation IAM entities for aws cloud account")
	d.SetId(cloudAccountID)

	return resourceIAMSafeEntityRead(d, meta)
}

func resourceIAMSafeEntityRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	var err error
	var resp *aws.IAMSafeEntityResponse

	entityType := d.Get("entity_type").(string)
	cloudAccountID := d.Get("aws_cloud_account_id").(string)
	entityName := d.Get("entity_name").(string)
	resp, err = d9Client.cloudaccountAWS.GetProtectIAMSafeEntityStatusByName(cloudAccountID, entityName, entityType)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing protect IAM entity from state because aws cloud account %s no longer exists in Dome9", cloudAccountID)
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] IAM entity reading response: %+v\n", *resp)
	_ = d.Set("state", resp.State)
	_ = d.Set("attached_dome9_users", resp.AttachedDome9Users)
	_ = d.Set("exists_in_aws", resp.ExistsInAws)
	_ = d.Set("arn", resp.Arn)

	return nil
}

func resourceIAMSafeEntityDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	var err error

	protectionMode := d.Get("protection_mode").(string)
	entityType := d.Get("entity_type").(string)
	cloudAccountID := d.Get("aws_cloud_account_id").(string)
	entityName := d.Get("entity_name").(string)

	if protectionMode == providerconst.IAMSafeEntityProtect {
		_, err = d9Client.cloudaccountAWS.UnprotectIAMSafeEntity(cloudAccountID, entityName, entityType)
	} else {
		_, err = d9Client.users.UnprotectWithElevationIAMSafeEntity(cloudAccountID, entityName, entityType)
	}

	if err != nil {
		return err
	}
	log.Printf("[INFO] Unprotect IAM safe entity")

	return nil
}

func resourceIAMSafeEntityUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	var err error = nil
	protectionMode := d.Get("protection_mode")
	cloudAccountID := d.Get("aws_cloud_account_id").(string)
	entityType := d.Get("entity_type").(string)
	entityName := d.Get("entity_name").(string)
	usersToAttach := expandUsersToProtectWithElevation(d.Get("dome9_users_id_to_protect").([]interface{}))

	if d.HasChange("dome9_users_id_to_protect") && protectionMode == providerconst.IAMSafeEntityProtectWithElevation {
		log.Println("[INFO] Users to protect with elevation has been changed")
		_, err = d9Client.users.ProtectWithElevationIAMSafeEntityUpdate(cloudAccountID, entityType, entityName, usersToAttach)
		if len(usersToAttach) == 0 && err == nil {
			_ = d.Set("protection_mode", providerconst.IAMSafeEntityProtect)
		}

		if err != nil {
			return err
		}
	}

	if d.HasChange("protection_mode") {
		log.Println("[INFO] Protection mode changed to ", protectionMode)

		// it it was ProtectWithElevation and now Protect
		if protectionMode == providerconst.IAMSafeEntityProtect {
			_ = d.Set("dome9_users_id_to_protect", []string{})
			_, err = d9Client.users.ProtectWithElevationIAMSafeEntityUpdate(cloudAccountID, entityType, entityName, []string{})
		} else {
			_, err = d9Client.users.ProtectWithElevationIAMSafeEntityUpdate(cloudAccountID, entityType, entityName, usersToAttach)
		}
		if err != nil {
			return err
		}
	}

	return err
}

func expandUsersToProtectWithElevation(usersID []interface{}) []string {
	users := make([]string, len(usersID))
	for i, user := range usersID {
		users[i] = user.(string)
	}

	return users
}

func expandIAMSafeEntityRequest(d *schema.ResourceData) aws.RestrictedIamEntitiesRequest {
	return aws.RestrictedIamEntitiesRequest{
		EntityName: d.Get("entity_name").(string),
		EntityType: d.Get("entity_type").(string),
	}
}
