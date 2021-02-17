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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSnapshotCreate,
		Read:   resourceComputeSnapshotRead,
		Update: resourceComputeSnapshotUpdate,
		Delete: resourceComputeSnapshotDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSnapshotImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"source_disk": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the disk used to create this snapshot.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this Snapshot.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The customer-supplied encryption key of the snapshot. Required if the
source snapshot is protected by a customer-supplied encryption key.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_self_link": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The name of the encryption key that is stored in Google Cloud KMS.`,
						},
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account used for the encryption request for the given KMS key.
If absent, the Compute Engine Service Agent service account is used.`,
						},
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Specifies a 256-bit customer-supplied encryption key, encoded in
RFC 4648 base64 to either encrypt or decrypt this resource.`,
							Sensitive: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
encryption key that protects this resource.`,
						},
					},
				},
			},
			"source_disk_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The customer-supplied encryption key of the source snapshot. Required
if the source snapshot is protected by a customer-supplied encryption
key.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account used for the encryption request for the given KMS key.
If absent, the Compute Engine Service Agent service account is used.`,
						},
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Specifies a 256-bit customer-supplied encryption key, encoded in
RFC 4648 base64 to either encrypt or decrypt this resource.`,
							Sensitive: true,
						},
					},
				},
			},
			"storage_locations": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Cloud Storage bucket storage location of the snapshot (regional or multi-regional).`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the disk is hosted.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"disk_size_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Size of the snapshot, specified in GB.`,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fingerprint used for optimistic locking of this resource. Used
internally during updates.`,
			},
			"licenses": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `A list of public visible licenses that apply to this snapshot. This
can be because the original image had licenses attached (such as a
Windows image).  snapshotEncryptionKey nested object Encrypts the
snapshot using a customer-supplied encryption key.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"snapshot_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"storage_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `A size of the storage used by the snapshot. As snapshots share
storage, this number is expected to change with snapshot
creation/deletion.`,
			},
			"source_disk_link": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "Deprecated in favor of source_disk, which contains a compatible value. This field will be removed in the next major release of the provider.",
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
		UseJSONNumber: true,
	}
}

func resourceComputeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeSnapshotName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	storageLocationsProp, err := expandComputeSnapshotStorageLocations(d.Get("storage_locations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_locations"); !isEmptyValue(reflect.ValueOf(storageLocationsProp)) && (ok || !reflect.DeepEqual(v, storageLocationsProp)) {
		obj["storageLocations"] = storageLocationsProp
	}
	labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	sourceDiskProp, err := expandComputeSnapshotSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk"); !isEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	zoneProp, err := expandComputeSnapshotZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	snapshotEncryptionKeyProp, err := expandComputeSnapshotSnapshotEncryptionKey(d.Get("snapshot_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot_encryption_key"); !isEmptyValue(reflect.ValueOf(snapshotEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, snapshotEncryptionKeyProp)) {
		obj["snapshotEncryptionKey"] = snapshotEncryptionKeyProp
	}
	sourceDiskEncryptionKeyProp, err := expandComputeSnapshotSourceDiskEncryptionKey(d.Get("source_disk_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk_encryption_key"); !isEmptyValue(reflect.ValueOf(sourceDiskEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, sourceDiskEncryptionKeyProp)) {
		obj["sourceDiskEncryptionKey"] = sourceDiskEncryptionKeyProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{source_disk}}/createSnapshot")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Snapshot: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Snapshot: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating Snapshot", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Snapshot: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Snapshot %q: %#v", d.Id(), res)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSnapshot %q", d.Id()))
	}

	res, err = resourceComputeSnapshotDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeSnapshot because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeSnapshotCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_id", flattenComputeSnapshotSnapshotId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("disk_size_gb", flattenComputeSnapshotDiskSizeGb(res["diskSizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("name", flattenComputeSnapshotName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("description", flattenComputeSnapshotDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("storage_bytes", flattenComputeSnapshotStorageBytes(res["storageBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("storage_locations", flattenComputeSnapshotStorageLocations(res["storageLocations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("licenses", flattenComputeSnapshotLicenses(res["licenses"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("labels", flattenComputeSnapshotLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeSnapshotLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("source_disk", flattenComputeSnapshotSourceDisk(res["sourceDisk"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("snapshot_encryption_key", flattenComputeSnapshotSnapshotEncryptionKey(res["snapshotEncryptionKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Snapshot: %s", err)
	}

	return nil
}

func resourceComputeSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})

		labelsProp, err := expandComputeSnapshotLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeSnapshotLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}/setLabels")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Snapshot %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Snapshot %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating Snapshot", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeSnapshotRead(d, meta)
}

func resourceComputeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snapshot: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Snapshot %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Snapshot")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting Snapshot", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Snapshot %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSnapshotImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/snapshots/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/snapshots/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSnapshotCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeSnapshotDiskSizeGb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeSnapshotName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotStorageBytes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeSnapshotStorageLocations(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotLicenses(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeSnapshotLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotLabelFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotSourceDisk(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeSnapshotSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeSnapshotSnapshotEncryptionKeyRawKey(original["rawKey"], d, config)
	transformed["sha256"] =
		flattenComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	transformed["kms_key_self_link"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return d.Get("snapshot_encryption_key.0.raw_key")
}

func flattenComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeSnapshotName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotStorageLocations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeSnapshotLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDisk(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSnapshotSnapshotEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSnapshotEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeSnapshotSnapshotEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !isEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeySelfLink, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeySha256(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeSnapshotSourceDiskEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !isEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedKmsKeyServiceAccount, err := expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyRawKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSnapshotSourceDiskEncryptionKeyKmsKeyServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeSnapshotDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["snapshotEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("snapshot_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]

		if kmsKeyName, ok := original["kmsKeyName"]; ok {
			// The response for crypto keys often includes the version of the key which needs to be removed
			// format: projects/<project>/locations/<region>/keyRings/<keyring>/cryptoKeys/<key>/cryptoKeyVersions/1
			transformed["kmsKeyName"] = strings.Split(kmsKeyName.(string), "/cryptoKeyVersions")[0]
		}

		if kmsKeyServiceAccount, ok := original["kmsKeyServiceAccount"]; ok {
			transformed["kmsKeyServiceAccount"] = kmsKeyServiceAccount
		}

		res["snapshotEncryptionKey"] = transformed
	}

	if v, ok := res["sourceDiskEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_disk_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]

		if kmsKeyName, ok := original["kmsKeyName"]; ok {
			// The response for crypto keys often includes the version of the key which needs to be removed
			// format: projects/<project>/locations/<region>/keyRings/<keyring>/cryptoKeys/<key>/cryptoKeyVersions/1
			transformed["kmsKeyName"] = strings.Split(kmsKeyName.(string), "/cryptoKeyVersions")[0]
		}

		if kmsKeyServiceAccount, ok := original["kmsKeyServiceAccount"]; ok {
			transformed["kmsKeyServiceAccount"] = kmsKeyServiceAccount
		}

		res["sourceDiskEncryptionKey"] = transformed
	}

	if err := d.Set("source_disk_link", ConvertSelfLinkToV1(res["sourceDisk"].(string))); err != nil {
		return nil, fmt.Errorf("Error setting source_disk_link: %s", err)
	}
	return res, nil
}
