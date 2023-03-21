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

func ResourceCloudIdsEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIdsEndpointCreate,
		Read:   resourceCloudIdsEndpointRead,
		Update: resourceCloudIdsEndpointUpdate,
		Delete: resourceCloudIdsEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIdsEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the endpoint.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.`,
			},
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net").`,
			},
			"severity": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"}),
				Description:  `The minimum alert severity level that is reported by the endpoint. Possible values: ["INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of the endpoint.`,
			},
			"threat_exceptions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC 3339 text format.`,
			},
			"endpoint_forwarding_rule": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `URL of the endpoint's network address to which traffic is to be sent by Packet Mirroring.`,
			},
			"endpoint_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Internal IP address of the endpoint's network entry point.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Last update timestamp in RFC 3339 text format.`,
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

func resourceCloudIdsEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandCloudIdsEndpointName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandCloudIdsEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	descriptionProp, err := expandCloudIdsEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	severityProp, err := expandCloudIdsEndpointSeverity(d.Get("severity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("severity"); !isEmptyValue(reflect.ValueOf(severityProp)) && (ok || !reflect.DeepEqual(v, severityProp)) {
		obj["severity"] = severityProp
	}
	threatExceptionsProp, err := expandCloudIdsEndpointThreatExceptions(d.Get("threat_exceptions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_exceptions"); !isEmptyValue(reflect.ValueOf(threatExceptionsProp)) && (ok || !reflect.DeepEqual(v, threatExceptionsProp)) {
		obj["threatExceptions"] = threatExceptionsProp
	}

	url, err := replaceVars(d, config, "{{CloudIdsBasePath}}projects/{{project}}/locations/{{location}}/endpoints?endpointId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Endpoint: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Endpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Endpoint: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = CloudIdsOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Endpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Endpoint: %s", err)
	}

	if err := d.Set("name", flattenCloudIdsEndpointName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Endpoint %q: %#v", d.Id(), res)

	return resourceCloudIdsEndpointRead(d, meta)
}

func resourceCloudIdsEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudIdsBasePath}}projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Endpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudIdsEndpoint %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}

	if err := d.Set("name", flattenCloudIdsEndpointName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("create_time", flattenCloudIdsEndpointCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("update_time", flattenCloudIdsEndpointUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("network", flattenCloudIdsEndpointNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("description", flattenCloudIdsEndpointDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("endpoint_forwarding_rule", flattenCloudIdsEndpointEndpointForwardingRule(res["endpointForwardingRule"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("endpoint_ip", flattenCloudIdsEndpointEndpointIp(res["endpointIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("severity", flattenCloudIdsEndpointSeverity(res["severity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("threat_exceptions", flattenCloudIdsEndpointThreatExceptions(res["threatExceptions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}

	return nil
}

func resourceCloudIdsEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Endpoint: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	threatExceptionsProp, err := expandCloudIdsEndpointThreatExceptions(d.Get("threat_exceptions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_exceptions"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, threatExceptionsProp)) {
		obj["threatExceptions"] = threatExceptionsProp
	}

	url, err := replaceVars(d, config, "{{CloudIdsBasePath}}projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Endpoint %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("threat_exceptions") {
		updateMask = append(updateMask, "threatExceptions")
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

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Endpoint %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Endpoint %q: %#v", d.Id(), res)
	}

	err = CloudIdsOperationWaitTime(
		config, res, project, "Updating Endpoint", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceCloudIdsEndpointRead(d, meta)
}

func resourceCloudIdsEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Endpoint: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{CloudIdsBasePath}}projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Endpoint %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Endpoint")
	}

	err = CloudIdsOperationWaitTime(
		config, res, project, "Deleting Endpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Endpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIdsEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/endpoints/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudIdsEndpointName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	parts := strings.Split(d.Get("name").(string), "/")
	return parts[len(parts)-1]
}

func flattenCloudIdsEndpointCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointEndpointForwardingRule(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointEndpointIp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointSeverity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIdsEndpointThreatExceptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudIdsEndpointName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
}

func expandCloudIdsEndpointNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointSeverity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointThreatExceptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
