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
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFilestoreInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceFilestoreInstanceCreate,
		Read:   resourceFilestoreInstanceRead,
		Update: resourceFilestoreInstanceUpdate,
		Delete: resourceFilestoreInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFilestoreInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"file_shares": {
				Type:     schema.TypeList,
				Required: true,
				Description: `File system shares on the instance. For this version, only a
single file share is supported.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capacity_gb": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `File share capacity in GiB. This must be at least 1024 GiB
for the standard tier, or 2560 GiB for the premium tier.`,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The name of the fileshare (16 characters or less)`,
						},
						"nfs_export_options": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Nfs Export Options. There is a limit of 10 export options per file share.`,
							MaxItems:    10,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"READ_ONLY", "READ_WRITE", ""}, false),
										Description: `Either READ_ONLY, for allowing only read requests on the exported directory,
or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]`,
										Default: "READ_WRITE",
									},
									"anon_gid": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `An integer representing the anonymous group id with a default value of 65534.
Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
if this field is specified for other squashMode settings.`,
									},
									"anon_uid": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `An integer representing the anonymous user id with a default value of 65534.
Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
if this field is specified for other squashMode settings.`,
									},
									"ip_ranges": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"squash_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"NO_ROOT_SQUASH", "ROOT_SQUASH", ""}, false),
										Description: `Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
for not allowing root access. The default is NO_ROOT_SQUASH. Default value: "NO_ROOT_SQUASH" Possible values: ["NO_ROOT_SQUASH", "ROOT_SQUASH"]`,
										Default: "NO_ROOT_SQUASH",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource name of the instance.`,
			},
			"networks": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `VPC networks to which the instance is connected. For this version,
only a single network is supported.`,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modes": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Description: `IP versions for which the instance has
IP addresses assigned. Possible values: ["ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"]`,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"}, false),
							},
						},
						"network": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The name of the GCE VPC network to which the
instance is connected.`,
						},
						"reserved_ip_range": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `A /29 CIDR block that identifies the range of IP
addresses reserved for this instance.`,
						},
						"ip_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `A list of IPv4 or IPv6 addresses.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"tier": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"TIER_UNSPECIFIED", "STANDARD", "PREMIUM", "BASIC_HDD", "BASIC_SSD", "HIGH_SCALE_SSD"}, false),
				Description:  `The service tier of the instance. Possible values: ["TIER_UNSPECIFIED", "STANDARD", "PREMIUM", "BASIC_HDD", "BASIC_SSD", "HIGH_SCALE_SSD"]`,
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the Filestore zone of the instance.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the instance.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Server-specified ETag for the instance resource to prevent
simultaneous updates from overwriting each other.`,
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

func resourceFilestoreInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tierProp, err := expandFilestoreInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	labelsProp, err := expandFilestoreInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("file_shares"); !isEmptyValue(reflect.ValueOf(fileSharesProp)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}
	networksProp, err := expandFilestoreInstanceNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); !isEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	lockName, err := replaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{zone}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isNotFilestoreQuotaError)
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = filestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceFilestoreInstanceRead(d, meta)
}

func resourceFilestoreInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil, isNotFilestoreQuotaError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("FilestoreInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("description", flattenFilestoreInstanceDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenFilestoreInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("tier", flattenFilestoreInstanceTier(res["tier"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenFilestoreInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("file_shares", flattenFilestoreInstanceFileShares(res["fileShares"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("networks", flattenFilestoreInstanceNetworks(res["networks"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("etag", flattenFilestoreInstanceEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceFilestoreInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandFilestoreInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("file_shares"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}

	lockName, err := replaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("file_shares") {
		updateMask = append(updateMask, "fileShares")
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

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate), isNotFilestoreQuotaError)

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = filestoreOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceFilestoreInstanceRead(d, meta)
}

func resourceFilestoreInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	lockName, err := replaceVars(d, config, "filestore/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{FilestoreBasePath}}projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isNotFilestoreQuotaError)
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = filestoreOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceFilestoreInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{zone}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFilestoreInstanceDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceTier(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceFileShares(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"name":               flattenFilestoreInstanceFileSharesName(original["name"], d, config),
			"capacity_gb":        flattenFilestoreInstanceFileSharesCapacityGb(original["capacityGb"], d, config),
			"nfs_export_options": flattenFilestoreInstanceFileSharesNfsExportOptions(original["nfsExportOptions"], d, config),
		})
	}
	return transformed
}
func flattenFilestoreInstanceFileSharesName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceFileSharesCapacityGb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenFilestoreInstanceFileSharesNfsExportOptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"ip_ranges":   flattenFilestoreInstanceFileSharesNfsExportOptionsIpRanges(original["ipRanges"], d, config),
			"access_mode": flattenFilestoreInstanceFileSharesNfsExportOptionsAccessMode(original["accessMode"], d, config),
			"squash_mode": flattenFilestoreInstanceFileSharesNfsExportOptionsSquashMode(original["squashMode"], d, config),
			"anon_uid":    flattenFilestoreInstanceFileSharesNfsExportOptionsAnonUid(original["anonUid"], d, config),
			"anon_gid":    flattenFilestoreInstanceFileSharesNfsExportOptionsAnonGid(original["anonGid"], d, config),
		})
	}
	return transformed
}
func flattenFilestoreInstanceFileSharesNfsExportOptionsIpRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceFileSharesNfsExportOptionsAccessMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceFileSharesNfsExportOptionsSquashMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceFileSharesNfsExportOptionsAnonUid(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenFilestoreInstanceFileSharesNfsExportOptionsAnonGid(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenFilestoreInstanceNetworks(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"network":           flattenFilestoreInstanceNetworksNetwork(original["network"], d, config),
			"modes":             flattenFilestoreInstanceNetworksModes(original["modes"], d, config),
			"reserved_ip_range": flattenFilestoreInstanceNetworksReservedIpRange(original["reservedIpRange"], d, config),
			"ip_addresses":      flattenFilestoreInstanceNetworksIpAddresses(original["ipAddresses"], d, config),
		})
	}
	return transformed
}
func flattenFilestoreInstanceNetworksNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksModes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksReservedIpRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceNetworksIpAddresses(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenFilestoreInstanceEtag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandFilestoreInstanceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFilestoreInstanceFileShares(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandFilestoreInstanceFileSharesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedCapacityGb, err := expandFilestoreInstanceFileSharesCapacityGb(original["capacity_gb"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityGb); val.IsValid() && !isEmptyValue(val) {
			transformed["capacityGb"] = transformedCapacityGb
		}

		transformedNfsExportOptions, err := expandFilestoreInstanceFileSharesNfsExportOptions(original["nfs_export_options"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNfsExportOptions); val.IsValid() && !isEmptyValue(val) {
			transformed["nfsExportOptions"] = transformedNfsExportOptions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesCapacityGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpRanges, err := expandFilestoreInstanceFileSharesNfsExportOptionsIpRanges(original["ip_ranges"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpRanges); val.IsValid() && !isEmptyValue(val) {
			transformed["ipRanges"] = transformedIpRanges
		}

		transformedAccessMode, err := expandFilestoreInstanceFileSharesNfsExportOptionsAccessMode(original["access_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAccessMode); val.IsValid() && !isEmptyValue(val) {
			transformed["accessMode"] = transformedAccessMode
		}

		transformedSquashMode, err := expandFilestoreInstanceFileSharesNfsExportOptionsSquashMode(original["squash_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSquashMode); val.IsValid() && !isEmptyValue(val) {
			transformed["squashMode"] = transformedSquashMode
		}

		transformedAnonUid, err := expandFilestoreInstanceFileSharesNfsExportOptionsAnonUid(original["anon_uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAnonUid); val.IsValid() && !isEmptyValue(val) {
			transformed["anonUid"] = transformedAnonUid
		}

		transformedAnonGid, err := expandFilestoreInstanceFileSharesNfsExportOptionsAnonGid(original["anon_gid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAnonGid); val.IsValid() && !isEmptyValue(val) {
			transformed["anonGid"] = transformedAnonGid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsIpRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAccessMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsSquashMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAnonUid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAnonGid(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandFilestoreInstanceNetworksNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !isEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		transformedModes, err := expandFilestoreInstanceNetworksModes(original["modes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedModes); val.IsValid() && !isEmptyValue(val) {
			transformed["modes"] = transformedModes
		}

		transformedReservedIpRange, err := expandFilestoreInstanceNetworksReservedIpRange(original["reserved_ip_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReservedIpRange); val.IsValid() && !isEmptyValue(val) {
			transformed["reservedIpRange"] = transformedReservedIpRange
		}

		transformedIpAddresses, err := expandFilestoreInstanceNetworksIpAddresses(original["ip_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !isEmptyValue(val) {
			transformed["ipAddresses"] = transformedIpAddresses
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceNetworksNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksModes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksIpAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
