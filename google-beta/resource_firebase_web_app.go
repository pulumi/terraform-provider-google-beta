// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFirebaseWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseWebAppCreate,
		Read:   resourceFirebaseWebAppRead,
		Update: resourceFirebaseWebAppUpdate,
		Delete: resourceFirebaseWebAppDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseWebAppImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user-assigned display name of the App.`,
			},
			"app_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The globally unique, Firebase-assigned identifier of the App.
This identifier should be treated as an opaque token, as the data format is not specified.`,
			},
			"app_urls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The URLs where the 'WebApp' is hosted.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully qualified resource name of the App, for example:
projects/projectId/webApps/appId`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ABANDON",
				Description: `Set to 'ABANDON' to allow the WebApp to be untracked from terraform state
rather than deleted upon 'terraform destroy'. This is useful becaue the WebApp may be
serving traffic. Set to 'DELETE' to delete the WebApp. Default to 'ABANDON'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseWebAppCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/webApps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new WebApp: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating WebApp: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = firebaseOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating WebApp", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating WebApp %q: %#v", d.Id(), res)

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseWebApp %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "ABANDON"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	if err := d.Set("name", flattenFirebaseWebAppName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("display_name", flattenFirebaseWebAppDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("app_id", flattenFirebaseWebAppAppId(res["appId"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}
	if err := d.Set("app_urls", flattenFirebaseWebAppAppUrls(res["appUrls"], d, config)); err != nil {
		return fmt.Errorf("Error reading WebApp: %s", err)
	}

	return nil
}

func resourceFirebaseWebAppUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for WebApp: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseWebAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating WebApp %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating WebApp %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating WebApp %q: %#v", d.Id(), res)
	}

	return resourceFirebaseWebAppRead(d, meta)
}

func resourceFirebaseWebAppDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	// Handwritten
	obj := make(map[string]interface{})
	if d.Get("deletion_policy") == "DELETE" {
		obj["immediate"] = true
	} else {
		fmt.Printf("Skip deleting App %q due to deletion_policy: %q\n", d.Id(), d.Get("deletion_policy"))
		return nil
	}
	// End of Handwritten
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{FirebaseBasePath}}{{name}}:remove")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting App %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "App")
	}

	err = firebaseOperationWaitTime(
		config, res, project, "Deleting App", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting App %q: %#v", d.Id(), res)

	// This is useful if the Delete operation returns before the Get operation
	// during post-test destroy shows the completed state of the resource.
	time.Sleep(5 * time.Second)

	return nil
}

func resourceFirebaseWebAppImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseWebAppName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseWebAppDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseWebAppAppId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseWebAppAppUrls(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirebaseWebAppDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
