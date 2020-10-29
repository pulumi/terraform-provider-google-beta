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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMonitoringService() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringServiceCreate,
		Read:   resourceMonitoringServiceRead,
		Update: resourceMonitoringServiceUpdate,
		Delete: resourceMonitoringServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name used for UI elements listing this Service.`,
			},
			"service_id": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z0-9\-]+$`),
				Description: `An optional service ID to use. If not given, the server will generate a
service ID.`,
			},
			"telemetry": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for how to query telemetry on a Service.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_name": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The full name of the resource that defines this service.
Formatted as described in
https://cloud.google.com/apis/design/resource_names.`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The full resource name for this service. The syntax is:
projects/[PROJECT_ID]/services/[SERVICE_ID].`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	telemetryProp, err := expandMonitoringServiceTelemetry(d.Get("telemetry"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("telemetry"); !isEmptyValue(reflect.ValueOf(telemetryProp)) && (ok || !reflect.DeepEqual(v, telemetryProp)) {
		obj["telemetry"] = telemetryProp
	}
	nameProp, err := expandMonitoringServiceServiceId(d.Get("service_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_id"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceMonitoringServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/services?serviceId={{service_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Service: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isMonitoringConcurrentEditError)
	if err != nil {
		return fmt.Errorf("Error creating Service: %s", err)
	}
	if err := d.Set("name", flattenMonitoringServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Service %q: %#v", d.Id(), res)

	return resourceMonitoringServiceRead(d, meta)
}

func resourceMonitoringServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil, isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringService %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	if err := d.Set("name", flattenMonitoringServiceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringServiceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("telemetry", flattenMonitoringServiceTelemetry(res["telemetry"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}
	if err := d.Set("service_id", flattenMonitoringServiceServiceId(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Service: %s", err)
	}

	return nil
}

func resourceMonitoringServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	telemetryProp, err := expandMonitoringServiceTelemetry(d.Get("telemetry"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("telemetry"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, telemetryProp)) {
		obj["telemetry"] = telemetryProp
	}

	obj, err = resourceMonitoringServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Service %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("telemetry") {
		updateMask = append(updateMask, "telemetry")
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

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate), isMonitoringConcurrentEditError)

	if err != nil {
		return fmt.Errorf("Error updating Service %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Service %q: %#v", d.Id(), res)
	}

	return resourceMonitoringServiceRead(d, meta)
}

func resourceMonitoringServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Service %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isMonitoringConcurrentEditError)
	if err != nil {
		return handleNotFoundError(err, d, "Service")
	}

	log.Printf("[DEBUG] Finished deleting Service %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringServiceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceTelemetry(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_name"] =
		flattenMonitoringServiceTelemetryResourceName(original["resourceName"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringServiceTelemetryResourceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMonitoringServiceServiceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandMonitoringServiceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceTelemetry(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceName, err := expandMonitoringServiceTelemetryResourceName(original["resource_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceName); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceName"] = transformedResourceName
	}

	return transformed, nil
}

func expandMonitoringServiceTelemetryResourceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringServiceServiceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceMonitoringServiceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Currently only CUSTOM service types can be created, but the
	// custom identifier block does not actually have fields right now.
	// Set to empty to indicate manually-created service type is CUSTOM.
	if _, ok := obj["custom"]; !ok {
		obj["custom"] = map[string]interface{}{}
	}
	// Name/Service ID is a query parameter only
	delete(obj, "name")

	return obj, nil
}
