// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApigeeEnvgroupAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvgroupAttachmentCreate,
		Read:   resourceApigeeEnvgroupAttachmentRead,
		Delete: resourceApigeeEnvgroupAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvgroupAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"envgroup_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee environment group associated with the Apigee environment,
in the format 'organizations/{{org_name}}/envgroups/{{envgroup_name}}'.`,
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource ID of the environment.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the newly created  attachment (output parameter).`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvgroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	environmentProp, err := expandApigeeEnvgroupAttachmentEnvironment(d.Get("environment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("environment"); !isEmptyValue(reflect.ValueOf(environmentProp)) && (ok || !reflect.DeepEqual(v, environmentProp)) {
		obj["environment"] = environmentProp
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{envgroup_id}}/attachments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EnvgroupAttachment: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EnvgroupAttachment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = apigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating EnvgroupAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EnvgroupAttachment: %s", err)
	}

	if err := d.Set("name", flattenApigeeEnvgroupAttachmentName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EnvgroupAttachment %q: %#v", d.Id(), res)

	return resourceApigeeEnvgroupAttachmentRead(d, meta)
}

func resourceApigeeEnvgroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvgroupAttachment %q", d.Id()))
	}

	if err := d.Set("environment", flattenApigeeEnvgroupAttachmentEnvironment(res["environment"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvgroupAttachment: %s", err)
	}
	if err := d.Set("name", flattenApigeeEnvgroupAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvgroupAttachment: %s", err)
	}

	return nil
}

func resourceApigeeEnvgroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EnvgroupAttachment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EnvgroupAttachment")
	}

	err = apigeeOperationWaitTime(
		config, res, "Deleting EnvgroupAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EnvgroupAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvgroupAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{
		"(?P<envgroup_id>.+)/attachments/(?P<name>.+)",
		"(?P<envgroup_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{envgroup_id}}/attachments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvgroupAttachmentEnvironment(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEnvgroupAttachmentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApigeeEnvgroupAttachmentEnvironment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
