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

func resourceHealthcareDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareDatasetCreate,
		Read:   resourceHealthcareDatasetRead,
		Update: resourceHealthcareDatasetUpdate,
		Delete: resourceHealthcareDatasetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareDatasetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the Dataset.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name for the Dataset.`,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
"America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
(e.g., HL7 messages) where no explicit timezone is specified.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully qualified name of this dataset`,
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

func resourceHealthcareDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareDatasetName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeZoneProp, err := expandHealthcareDatasetTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets?datasetId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Dataset: %#v", obj)
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

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Dataset: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Dataset %q: %#v", d.Id(), res)

	return resourceHealthcareDatasetRead(d, meta)
}

func resourceHealthcareDatasetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
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

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("HealthcareDataset %q", d.Id()))
	}

	res, err = resourceHealthcareDatasetDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing HealthcareDataset because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	if err := d.Set("name", flattenHealthcareDatasetName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("time_zone", flattenHealthcareDatasetTimeZone(res["timeZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	return nil
}

func resourceHealthcareDatasetUpdate(d *schema.ResourceData, meta interface{}) error {
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
	timeZoneProp, err := expandHealthcareDatasetTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Dataset %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("time_zone") {
		updateMask = append(updateMask, "timeZone")
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
		return fmt.Errorf("Error updating Dataset %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Dataset %q: %#v", d.Id(), res)
	}

	return resourceHealthcareDatasetRead(d, meta)
}

func resourceHealthcareDatasetDelete(d *schema.ResourceData, meta interface{}) error {
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

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Dataset %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Dataset")
	}

	log.Printf("[DEBUG] Finished deleting Dataset %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareDatasetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/datasets/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareDatasetName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenHealthcareDatasetTimeZone(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandHealthcareDatasetName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareDatasetTimeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareDatasetDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	if err := d.Set("self_link", res["name"].(string)); err != nil {
		return nil, fmt.Errorf("Error setting self_link: %s", err)
	}
	res["name"] = d.Get("name").(string)
	return res, nil
}
