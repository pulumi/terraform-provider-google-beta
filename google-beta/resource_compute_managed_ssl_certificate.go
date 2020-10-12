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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeManagedSslCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeManagedSslCertificateCreate,
		Read:   resourceComputeManagedSslCertificateRead,
		Delete: resourceComputeManagedSslCertificateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeManagedSslCertificateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"managed": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Properties relevant to a managed certificate.  These will be used if the
certificate is managed (as indicated by a value of 'MANAGED' in 'type').`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domains": {
							Type:             schema.TypeList,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: absoluteDomainSuppress,
							Description: `Domains for which a managed SSL certificate will be valid.  Currently,
there can be up to 100 domains in this list.`,
							MaxItems: 100,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.


These are in the same namespace as the managed SSL certificates.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"MANAGED", ""}, false),
				Description: `Enum field whose value is always 'MANAGED' - used to signal to the API
which type this is. Default value: "MANAGED" Possible values: ["MANAGED"]`,
				Default: "MANAGED",
			},
			"certificate_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: `The unique identifier for the resource.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Expire time of the certificate.`,
			},
			"subject_alternative_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Domains associated with the certificate via Subject Alternative Name.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeManagedSslCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeManagedSslCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeManagedSslCertificateName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	managedProp, err := expandComputeManagedSslCertificateManaged(d.Get("managed"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("managed"); !isEmptyValue(reflect.ValueOf(managedProp)) && (ok || !reflect.DeepEqual(v, managedProp)) {
		obj["managed"] = managedProp
	}
	typeProp, err := expandComputeManagedSslCertificateType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagedSslCertificate: %#v", obj)
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
		return fmt.Errorf("Error creating ManagedSslCertificate: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating ManagedSslCertificate", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ManagedSslCertificate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ManagedSslCertificate %q: %#v", d.Id(), res)

	return resourceComputeManagedSslCertificateRead(d, meta)
}

func resourceComputeManagedSslCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeManagedSslCertificate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeManagedSslCertificateCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("description", flattenComputeManagedSslCertificateDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("certificate_id", flattenComputeManagedSslCertificateCertificateId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("name", flattenComputeManagedSslCertificateName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("managed", flattenComputeManagedSslCertificateManaged(res["managed"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("type", flattenComputeManagedSslCertificateType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("subject_alternative_names", flattenComputeManagedSslCertificateSubjectAlternativeNames(res["subjectAlternativeNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("expire_time", flattenComputeManagedSslCertificateExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading ManagedSslCertificate: %s", err)
	}

	return nil
}

func resourceComputeManagedSslCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ManagedSslCertificate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ManagedSslCertificate")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting ManagedSslCertificate", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ManagedSslCertificate %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeManagedSslCertificateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/sslCertificates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/sslCertificates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeManagedSslCertificateCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateCertificateId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeManagedSslCertificateName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateManaged(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["domains"] =
		flattenComputeManagedSslCertificateManagedDomains(original["domains"], d, config)
	return []interface{}{transformed}
}
func flattenComputeManagedSslCertificateManagedDomains(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateSubjectAlternativeNames(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeManagedSslCertificateExpireTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeManagedSslCertificateDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeManagedSslCertificateName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeManagedSslCertificateManaged(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDomains, err := expandComputeManagedSslCertificateManagedDomains(original["domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomains); val.IsValid() && !isEmptyValue(val) {
		transformed["domains"] = transformedDomains
	}

	return transformed, nil
}

func expandComputeManagedSslCertificateManagedDomains(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeManagedSslCertificateType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
