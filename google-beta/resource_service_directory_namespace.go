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

func resourceServiceDirectoryNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDirectoryNamespaceCreate,
		Read:   resourceServiceDirectoryNamespaceRead,
		Update: resourceServiceDirectoryNamespaceUpdate,
		Delete: resourceServiceDirectoryNamespaceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceDirectoryNamespaceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The location for the Namespace.
A full list of valid locations can be found by running
'gcloud beta service-directory locations list'.`,
			},
			"namespace_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRFC1035Name(2, 63),
				Description: `The Resource ID must be 1-63 characters long, including digits,
lowercase letters or the hyphen character.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Resource labels associated with this Namespace. No more than 64 user
labels can be associated with a given resource. Label keys and values can
be no longer than 63 characters.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name for the namespace
in the format 'projects/*/locations/*/namespaces/*'.`,
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

func resourceServiceDirectoryNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandServiceDirectoryNamespaceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}projects/{{project}}/locations/{{location}}/namespaces?namespaceId={{namespace_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Namespace: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Namespace: %s", err)
	}
	if err := d.Set("name", flattenServiceDirectoryNamespaceName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Namespace %q: %#v", d.Id(), res)

	return resourceServiceDirectoryNamespaceRead(d, meta)
}

func resourceServiceDirectoryNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ServiceDirectoryNamespace %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}

	if err := d.Set("name", flattenServiceDirectoryNamespaceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}
	if err := d.Set("labels", flattenServiceDirectoryNamespaceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Namespace: %s", err)
	}

	return nil
}

func resourceServiceDirectoryNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandServiceDirectoryNamespaceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Namespace %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
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
		return fmt.Errorf("Error updating Namespace %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Namespace %q: %#v", d.Id(), res)
	}

	return resourceServiceDirectoryNamespaceRead(d, meta)
}

func resourceServiceDirectoryNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Namespace: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Namespace %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Namespace")
	}

	log.Printf("[DEBUG] Finished deleting Namespace %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceDirectoryNamespaceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 6 {
		// `projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}`
		if err := d.Set("project", nameParts[1]); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("location", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
		if err := d.Set("namespace_id", nameParts[5]); err != nil {
			return nil, fmt.Errorf("Error setting namespace_id: %s", err)
		}
	} else if len(nameParts) == 3 {
		// `{{project}}/{{location}}/{{namespace_id}}`
		if err := d.Set("project", nameParts[0]); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("location", nameParts[1]); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
		if err := d.Set("namespace_id", nameParts[2]); err != nil {
			return nil, fmt.Errorf("Error setting namespace_id: %s", err)
		}
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", nameParts[0], nameParts[1], nameParts[2])
		if err := d.Set("name", id); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		d.SetId(id)
	} else if len(nameParts) == 2 {
		// `{{location}}/{{namespace_id}}`
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("location", nameParts[0]); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
		if err := d.Set("namespace_id", nameParts[1]); err != nil {
			return nil, fmt.Errorf("Error setting namespace_id: %s", err)
		}
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", project, nameParts[0], nameParts[1])
		if err := d.Set("name", id); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		d.SetId(id)
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s, %s or %s",
			d.Get("name"),
			"projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}",
			"{{project}}/{{location}}/{{namespace_id}}",
			"{{location}}/{{namespace_id}}")
	}
	return []*schema.ResourceData{d}, nil
}

func flattenServiceDirectoryNamespaceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenServiceDirectoryNamespaceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandServiceDirectoryNamespaceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
