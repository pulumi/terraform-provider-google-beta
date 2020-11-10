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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRedisInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisInstanceCreate,
		Read:   resourceRedisInstanceRead,
		Update: resourceRedisInstanceUpdate,
		Delete: resourceRedisInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"memory_size_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Redis memory size in GiB.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z][a-z0-9-]{0,39}[a-z0-9]$`),
				Description:  `The ID of the instance or a fully qualified identifier for the instance.`,
			},
			"alternative_location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Only applicable to STANDARD_HA tier which protects the instance
against zonal failures by provisioning it across two zones.
If provided, it must be a different zone from the one provided in
[locationId].`,
			},
			"auth_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Optional. Indicates whether OSS Redis AUTH is enabled for the
instance. If set to "true" AUTH is enabled on the instance.
Default value is "false" meaning AUTH is disabled.`,
				Default: false,
			},
			"authorized_network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The full name of the Google Compute Engine network to which the
instance is connected. If left unspecified, the default network
will be used.`,
			},
			"connect_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS", ""}, false),
				Description:  `The connection mode of the Redis instance. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]`,
				Default:      "DIRECT_PEERING",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An arbitrary and optional user-provided name for the instance.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The zone where the instance will be provisioned. If not provided,
the service will choose a zone for the instance. For STANDARD_HA tier,
instances will be created across two zones for protection against
zonal failures. If [alternativeLocationId] is also provided, it must
be different from [locationId].`,
			},
			"redis_configs": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Redis configuration parameters, according to http://redis.io/topics/config.
Please check Memorystore documentation for the list of supported parameters:
https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"redis_version": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The version of Redis software. If not provided, latest supported
version will be used. Currently, the supported values are:

- REDIS_5_0 for Redis 5.0 compatibility
- REDIS_4_0 for Redis 4.0 compatibility
- REDIS_3_2 for Redis 3.2 compatibility`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the Redis region of the instance.`,
			},
			"reserved_ip_range": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The CIDR range of internal addresses that are reserved for this
instance. If not provided, the service will choose an unused /29
block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
unique and non-overlapping with existing subnets in an authorized
network.`,
			},
			"tier": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BASIC", "STANDARD_HA", ""}, false),
				Description: `The service tier of the instance. Must be one of these values:

- BASIC: standalone instance
- STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"]`,
				Default: "BASIC",
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time the instance was created in RFC3339 UTC "Zulu" format,
accurate to nanoseconds.`,
			},
			"current_location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current zone where the Redis endpoint is placed.
For Basic Tier instances, this will always be the same as the
[locationId] provided by the user at creation time. For Standard Tier
instances, this can be either [locationId] or [alternativeLocationId]
and can change after a failover event.`,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Hostname or IP address of the exposed Redis endpoint used by clients
to connect to the service.`,
			},
			"persistence_iam_identity": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Cloud IAM identity used by import / export operations
to transfer data to/from Cloud Storage. Format is "serviceAccount:".
The value may change over time for a given instance so should be
checked before each import/export operation.`,
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The port number of the exposed Redis endpoint.`,
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

func resourceRedisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alternative_location_id"); !isEmptyValue(reflect.ValueOf(alternativeLocationIdProp)) && (ok || !reflect.DeepEqual(v, alternativeLocationIdProp)) {
		obj["alternativeLocationId"] = alternativeLocationIdProp
	}
	authEnabledProp, err := expandRedisInstanceAuthEnabled(d.Get("auth_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auth_enabled"); !isEmptyValue(reflect.ValueOf(authEnabledProp)) && (ok || !reflect.DeepEqual(v, authEnabledProp)) {
		obj["authEnabled"] = authEnabledProp
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	connectModeProp, err := expandRedisInstanceConnectMode(d.Get("connect_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connect_mode"); !isEmptyValue(reflect.ValueOf(connectModeProp)) && (ok || !reflect.DeepEqual(v, connectModeProp)) {
		obj["connectMode"] = connectModeProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(memorySizeGbProp)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_version"); !isEmptyValue(reflect.ValueOf(redisVersionProp)) && (ok || !reflect.DeepEqual(v, redisVersionProp)) {
		obj["redisVersion"] = redisVersionProp
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}

	obj, err = resourceRedisInstanceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/instances?instanceId={{name}}")
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

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = redisOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	if err := d.Set("name", flattenRedisInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
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

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("RedisInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("alternative_location_id", flattenRedisInstanceAlternativeLocationId(res["alternativeLocationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("auth_enabled", flattenRedisInstanceAuthEnabled(res["authEnabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("authorized_network", flattenRedisInstanceAuthorizedNetwork(res["authorizedNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("connect_mode", flattenRedisInstanceConnectMode(res["connectMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenRedisInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("current_location_id", flattenRedisInstanceCurrentLocationId(res["currentLocationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("display_name", flattenRedisInstanceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("host", flattenRedisInstanceHost(res["host"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenRedisInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("redis_configs", flattenRedisInstanceRedisConfigs(res["redisConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("location_id", flattenRedisInstanceLocationId(res["locationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("name", flattenRedisInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memory_size_gb", flattenRedisInstanceMemorySizeGb(res["memorySizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("port", flattenRedisInstancePort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("persistence_iam_identity", flattenRedisInstancePersistenceIamIdentity(res["persistenceIamIdentity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("redis_version", flattenRedisInstanceRedisVersion(res["redisVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("reserved_ip_range", flattenRedisInstanceReservedIpRange(res["reservedIpRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("tier", flattenRedisInstanceTier(res["tier"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceRedisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
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
	authEnabledProp, err := expandRedisInstanceAuthEnabled(d.Get("auth_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auth_enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authEnabledProp)) {
		obj["authEnabled"] = authEnabledProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}

	obj, err = resourceRedisInstanceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("auth_enabled") {
		updateMask = append(updateMask, "authEnabled")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("redis_configs") {
		updateMask = append(updateMask, "redisConfigs")
	}

	if d.HasChange("memory_size_gb") {
		updateMask = append(updateMask, "memorySizeGb")
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
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = redisOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceDelete(d *schema.ResourceData, meta interface{}) error {
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

	url, err := replaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = redisOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceRedisInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenRedisInstanceAlternativeLocationId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceAuthEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceConnectMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceCurrentLocationId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceRedisConfigs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceLocationId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenRedisInstanceMemorySizeGb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenRedisInstancePort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenRedisInstancePersistenceIamIdentity(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceRedisVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceReservedIpRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenRedisInstanceTier(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandRedisInstanceAlternativeLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	fv, err := ParseNetworkFieldValue(v.(string), d, config)
	if err != nil {
		return nil, err
	}
	return fv.RelativeLink(), nil
}

func expandRedisInstanceConnectMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceRedisConfigs(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceLocationId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
}

func expandRedisInstanceMemorySizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRedisVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReservedIpRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTier(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceRedisInstanceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	region, err := getRegionFromSchema("region", "location_id", d, config)
	if err != nil {
		return nil, err
	}
	if err := d.Set("region", region); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	return obj, nil
}
