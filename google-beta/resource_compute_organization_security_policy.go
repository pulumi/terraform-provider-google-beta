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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeOrganizationSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeOrganizationSecurityPolicyCreate,
		Read:   resourceComputeOrganizationSecurityPolicyRead,
		Update: resourceComputeOrganizationSecurityPolicyUpdate,
		Delete: resourceComputeOrganizationSecurityPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeOrganizationSecurityPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A textual name of the security policy.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
Format: organizations/{organization_id} or folders/{folder_id}`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A textual description for the organization security policy.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"FIREWALL", ""}, false),
				Description: `The type indicates the intended use of the security policy.
For organization security policies, the only supported type
is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"]`,
				Default: "FIREWALL",
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. This field is used internally during
updates of this resource.`,
			},
			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
			},
		},
	}
}

func resourceComputeOrganizationSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandComputeOrganizationSecurityPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandComputeOrganizationSecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeOrganizationSecurityPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	typeProp, err := expandComputeOrganizationSecurityPolicyType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	parentProp, err := expandComputeOrganizationSecurityPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}locations/global/securityPolicies?parentId={{parent}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationSecurityPolicy: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating OrganizationSecurityPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "locations/global/securityPolicies/{{policy_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationSecurityPolicy %q: %#v", d.Id(), res)

	parent := d.Get("parent").(string)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "Creating OrganizationSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create OrganizationSecurityPolicy: %s", err)
	}

	policyId, ok := opRes["targetId"]
	if !ok {
		return fmt.Errorf("Create response didn't contain targetId. Create may not have succeeded.")
	}
	if err := d.Set("policy_id", policyId.(string)); err != nil {
		return fmt.Errorf("Error setting policy_id: %s", err)
	}

	// Store the ID now.
	id, err = replaceVars(d, config, "locations/global/securityPolicies/{{policy_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceComputeOrganizationSecurityPolicyRead(d, meta)
}

func resourceComputeOrganizationSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}locations/global/securityPolicies/{{policy_id}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeOrganizationSecurityPolicy %q", d.Id()))
	}

	if err := d.Set("display_name", flattenComputeOrganizationSecurityPolicyDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeOrganizationSecurityPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeOrganizationSecurityPolicyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}
	if err := d.Set("policy_id", flattenComputeOrganizationSecurityPolicyPolicyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}
	if err := d.Set("type", flattenComputeOrganizationSecurityPolicyType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}
	if err := d.Set("parent", flattenComputeOrganizationSecurityPolicyParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSecurityPolicy: %s", err)
	}

	return nil
}

func resourceComputeOrganizationSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeOrganizationSecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeOrganizationSecurityPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}locations/global/securityPolicies/{{policy_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationSecurityPolicy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating OrganizationSecurityPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating OrganizationSecurityPolicy %q: %#v", d.Id(), res)
	}

	parent := d.Get("parent").(string)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "Creating OrganizationSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually update
		return fmt.Errorf("Error waiting to update OrganizationSecurityPolicy: %s", err)
	}
	return resourceComputeOrganizationSecurityPolicyRead(d, meta)
}

func resourceComputeOrganizationSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	url, err := replaceVars(d, config, "{{ComputeBasePath}}locations/global/securityPolicies/{{policy_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting OrganizationSecurityPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "OrganizationSecurityPolicy")
	}

	parent := d.Get("parent").(string)
	var opRes map[string]interface{}
	err = computeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "Creating OrganizationSecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually delete
		return fmt.Errorf("Error waiting to delete OrganizationSecurityPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting OrganizationSecurityPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeOrganizationSecurityPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"locations/global/securityPolicies/(?P<policy_id>[^/]+)",
		"(?P<policy_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "locations/global/securityPolicies/{{policy_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeOrganizationSecurityPolicyDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyPolicyId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeOrganizationSecurityPolicyParent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeOrganizationSecurityPolicyDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
