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

func resourceIdentityPlatformTenantDefaultSupportedIdpConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformTenantDefaultSupportedIdpConfigCreate,
		Read:   resourceIdentityPlatformTenantDefaultSupportedIdpConfigRead,
		Update: resourceIdentityPlatformTenantDefaultSupportedIdpConfigUpdate,
		Delete: resourceIdentityPlatformTenantDefaultSupportedIdpConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformTenantDefaultSupportedIdpConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client ID`,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `OAuth client secret`,
			},
			"idp_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `ID of the IDP. Possible values include:

* 'apple.com'

* 'facebook.com'

* 'gc.apple.com'

* 'github.com'

* 'google.com'

* 'linkedin.com'

* 'microsoft.com'

* 'playgames.google.com'

* 'twitter.com'

* 'yahoo.com'`,
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the tenant where this DefaultSupportedIdpConfig resource exists`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this IDP allows the user to sign in`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the default supported IDP config resource`,
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

func resourceIdentityPlatformTenantDefaultSupportedIdpConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !isEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !isEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs?idpId={{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TenantDefaultSupportedIdpConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantDefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TenantDefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformTenantDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TenantDefaultSupportedIdpConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformTenantDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformTenantDefaultSupportedIdpConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantDefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformTenantDefaultSupportedIdpConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TenantDefaultSupportedIdpConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformTenantDefaultSupportedIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantDefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_id", flattenIdentityPlatformTenantDefaultSupportedIdpConfigClientId(res["clientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantDefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("client_secret", flattenIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantDefaultSupportedIdpConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantDefaultSupportedIdpConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformTenantDefaultSupportedIdpConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantDefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TenantDefaultSupportedIdpConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("client_id") {
		updateMask = append(updateMask, "clientId")
	}

	if d.HasChange("client_secret") {
		updateMask = append(updateMask, "clientSecret")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
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
		return fmt.Errorf("Error updating TenantDefaultSupportedIdpConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TenantDefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformTenantDefaultSupportedIdpConfigRead(d, meta)
}

func resourceIdentityPlatformTenantDefaultSupportedIdpConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantDefaultSupportedIdpConfig: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TenantDefaultSupportedIdpConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TenantDefaultSupportedIdpConfig")
	}

	log.Printf("[DEBUG] Finished deleting TenantDefaultSupportedIdpConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformTenantDefaultSupportedIdpConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/tenants/(?P<tenant>[^/]+)/defaultSupportedIdpConfigs/(?P<idp_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<tenant>[^/]+)/(?P<idp_id>[^/]+)",
		"(?P<tenant>[^/]+)/(?P<idp_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformTenantDefaultSupportedIdpConfigName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantDefaultSupportedIdpConfigClientId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
