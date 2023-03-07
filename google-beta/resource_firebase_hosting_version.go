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

func ResourceFirebaseHostingVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseHostingVersionCreate,
		Read:   resourceFirebaseHostingVersionRead,
		Delete: resourceFirebaseHostingVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseHostingVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"site_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. The ID of the site in which to create this Version.`,
			},
			"config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The configuration for the behavior of the site. This configuration exists in the 'firebase.json' file.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirects": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `An array of objects (called redirect rules), where each rule specifies a URL pattern that, if matched to the request URL path,
triggers Hosting to respond with a redirect to the specified destination path.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"location": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `The value to put in the HTTP location header of the response.
The location can contain capture group values from the pattern using a : prefix to identify
the segment and an optional * to capture the rest of the URL. For example:

'''hcl
redirects {
  glob = "/:capture*"
  status_code = 302
  location = "https://example.com/foo/:capture"
}
'''`,
									},
									"status_code": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: `The status HTTP code to return in the response. It must be a valid 3xx status code.`,
									},
									"glob": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										Description:  `The user-supplied glob to match against the request URL path.`,
										ExactlyOneOf: []string{},
									},
									"regex": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										Description:  `The user-supplied RE2 regular expression to match against the request URL path.`,
										ExactlyOneOf: []string{},
									},
								},
							},
						},
						"rewrites": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `An array of objects (called rewrite rules), where each rule specifies a URL pattern that, if matched to the
request URL path, triggers Hosting to respond as if the service were given the specified destination URL.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"function": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										Description:  `The function to proxy requests to. Must match the exported function name exactly.`,
										ExactlyOneOf: []string{},
									},
									"glob": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										Description:  `The user-supplied glob to match against the request URL path.`,
										ExactlyOneOf: []string{},
									},
									"regex": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										Description:  `The user-supplied RE2 regular expression to match against the request URL path.`,
										ExactlyOneOf: []string{},
									},
									"run": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: `The request will be forwarded to Cloud Run.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_id": {
													Type:        schema.TypeString,
													Required:    true,
													ForceNew:    true,
													Description: `User-defined ID of the Cloud Run service.`,
												},
												"region": {
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Description: `Optional. User-provided region where the Cloud Run service is hosted. Defaults to 'us-central1' if not supplied.`,
												},
											},
										},
										ExactlyOneOf: []string{},
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully-qualified resource name for the version, in the format:
sites/SITE_ID/versions/VERSION_ID`,
			},
			"version_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The ID for the version as in sites/SITE_ID/versions/VERSION_ID`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseHostingVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	configProp, err := expandFirebaseHostingVersionConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !isEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}

	url, err := replaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/versions")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Version: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Version: %s", err)
	}
	if err := d.Set("name", flattenFirebaseHostingVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "sites/{{site_id}}/versions/{{version_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Store the name as ID
	d.SetId(res["name"].(string))

	if err = d.Set("version_id", GetResourceNameFromSelfLink(res["name"].(string))); err != nil {
		return fmt.Errorf("Error setting version_id: %s", err)
	}

	obj = make(map[string]interface{})
	obj["status"] = "FINALIZED"

	url, err = replaceVars(d, config, "{{FirebaseHostingBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Version %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "status")
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	res, err = sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error finalizing Version %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished finalizing Version %q: %#v", d.Id(), res)
	}

	log.Printf("[DEBUG] Finished creating Version %q: %#v", d.Id(), res)

	return resourceFirebaseHostingVersionRead(d, meta)
}

func resourceFirebaseHostingVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FirebaseHostingBasePath}}sites/{{site_id}}/versions/{{version_id}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("FirebaseHostingVersion %q", d.Id()))
	}

	res, err = resourceFirebaseHostingVersionDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing FirebaseHostingVersion because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenFirebaseHostingVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("config", flattenFirebaseHostingVersionConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}

	return nil
}

func resourceFirebaseHostingVersionDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] FirebaseHosting Version resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseHostingVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"sites/(?P<site_id>[^/]+)/versions/(?P<version_id>[^/]+)",
		"(?P<site_id>[^/]+)/(?P<version_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "sites/{{site_id}}/versions/{{version_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseHostingVersionName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["rewrites"] =
		flattenFirebaseHostingVersionConfigRewrites(original["rewrites"], d, config)
	transformed["redirects"] =
		flattenFirebaseHostingVersionConfigRedirects(original["redirects"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseHostingVersionConfigRewrites(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"glob":     flattenFirebaseHostingVersionConfigRewritesGlob(original["glob"], d, config),
			"regex":    flattenFirebaseHostingVersionConfigRewritesRegex(original["regex"], d, config),
			"function": flattenFirebaseHostingVersionConfigRewritesFunction(original["function"], d, config),
			"run":      flattenFirebaseHostingVersionConfigRewritesRun(original["run"], d, config),
		})
	}
	return transformed
}
func flattenFirebaseHostingVersionConfigRewritesGlob(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRewritesRegex(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRewritesFunction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRewritesRun(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["service_id"] =
		flattenFirebaseHostingVersionConfigRewritesRunServiceId(original["serviceId"], d, config)
	transformed["region"] =
		flattenFirebaseHostingVersionConfigRewritesRunRegion(original["region"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseHostingVersionConfigRewritesRunServiceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRewritesRunRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRedirects(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"glob":        flattenFirebaseHostingVersionConfigRedirectsGlob(original["glob"], d, config),
			"regex":       flattenFirebaseHostingVersionConfigRedirectsRegex(original["regex"], d, config),
			"status_code": flattenFirebaseHostingVersionConfigRedirectsStatusCode(original["statusCode"], d, config),
			"location":    flattenFirebaseHostingVersionConfigRedirectsLocation(original["location"], d, config),
		})
	}
	return transformed
}
func flattenFirebaseHostingVersionConfigRedirectsGlob(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRedirectsRegex(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFirebaseHostingVersionConfigRedirectsStatusCode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenFirebaseHostingVersionConfigRedirectsLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFirebaseHostingVersionConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRewrites, err := expandFirebaseHostingVersionConfigRewrites(original["rewrites"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRewrites); val.IsValid() && !isEmptyValue(val) {
		transformed["rewrites"] = transformedRewrites
	}

	transformedRedirects, err := expandFirebaseHostingVersionConfigRedirects(original["redirects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRedirects); val.IsValid() && !isEmptyValue(val) {
		transformed["redirects"] = transformedRedirects
	}

	return transformed, nil
}

func expandFirebaseHostingVersionConfigRewrites(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGlob, err := expandFirebaseHostingVersionConfigRewritesGlob(original["glob"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGlob); val.IsValid() && !isEmptyValue(val) {
			transformed["glob"] = transformedGlob
		}

		transformedRegex, err := expandFirebaseHostingVersionConfigRewritesRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedFunction, err := expandFirebaseHostingVersionConfigRewritesFunction(original["function"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !isEmptyValue(val) {
			transformed["function"] = transformedFunction
		}

		transformedRun, err := expandFirebaseHostingVersionConfigRewritesRun(original["run"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRun); val.IsValid() && !isEmptyValue(val) {
			transformed["run"] = transformedRun
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirebaseHostingVersionConfigRewritesGlob(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRun(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceId, err := expandFirebaseHostingVersionConfigRewritesRunServiceId(original["service_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceId); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceId"] = transformedServiceId
	}

	transformedRegion, err := expandFirebaseHostingVersionConfigRewritesRunRegion(original["region"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegion); val.IsValid() && !isEmptyValue(val) {
		transformed["region"] = transformedRegion
	}

	return transformed, nil
}

func expandFirebaseHostingVersionConfigRewritesRunServiceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRewritesRunRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirects(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGlob, err := expandFirebaseHostingVersionConfigRedirectsGlob(original["glob"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGlob); val.IsValid() && !isEmptyValue(val) {
			transformed["glob"] = transformedGlob
		}

		transformedRegex, err := expandFirebaseHostingVersionConfigRedirectsRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedStatusCode, err := expandFirebaseHostingVersionConfigRedirectsStatusCode(original["status_code"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatusCode); val.IsValid() && !isEmptyValue(val) {
			transformed["statusCode"] = transformedStatusCode
		}

		transformedLocation, err := expandFirebaseHostingVersionConfigRedirectsLocation(original["location"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
			transformed["location"] = transformedLocation
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirebaseHostingVersionConfigRedirectsGlob(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsRegex(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsStatusCode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingVersionConfigRedirectsLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceFirebaseHostingVersionDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if err := d.Set("version_id", GetResourceNameFromSelfLink(res["name"].(string))); err != nil {
		return nil, err
	}

	return res, nil
}
