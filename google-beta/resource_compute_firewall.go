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
	"bytes"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceComputeFirewallRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["protocol"].(string))))

	// We need to make sure to sort the strings below so that we always
	// generate the same hash code no matter what is in the set.
	if v, ok := m["ports"]; ok && v != nil {
		s := convertStringArr(v.([]interface{}))
		sort.Strings(s)

		for _, v := range s {
			buf.WriteString(fmt.Sprintf("%s-", v))
		}
	}

	return hashcode.String(buf.String())
}

func compareCaseInsensitive(k, old, new string, d *schema.ResourceData) bool {
	return strings.ToLower(old) == strings.ToLower(new)
}

func diffSuppressEnableLogging(k, old, new string, d *schema.ResourceData) bool {
	if k == "log_config.#" {
		if new == "0" && d.Get("enable_logging").(bool) {
			return true
		}
	}

	return false
}

func resourceComputeFirewallEnableLoggingCustomizeDiff(diff *schema.ResourceDiff, v interface{}) error {
	enableLogging, enableExists := diff.GetOkExists("enable_logging")
	if !enableExists {
		return nil
	}

	logConfigExists := diff.Get("log_config.#").(int) != 0
	if logConfigExists && enableLogging == false {
		return fmt.Errorf("log_config cannot be defined when enable_logging is false")
	}

	return nil
}

func resourceComputeFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeFirewallCreate,
		Read:   resourceComputeFirewallRead,
		Update: resourceComputeFirewallUpdate,
		Delete: resourceComputeFirewallDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeFirewallImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		SchemaVersion: 1,
		MigrateState:  resourceComputeFirewallMigrateState,
		CustomizeDiff: resourceComputeFirewallEnableLoggingCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name or self_link of the network to attach this firewall to.`,
			},
			"allow": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `The list of ALLOW rules specified by this firewall. Each rule
specifies a protocol and port-range tuple that describes a permitted
connection.`,
				Elem:         computeFirewallAllowSchema(),
				Set:          resourceComputeFirewallRuleHash,
				ExactlyOneOf: []string{"allow", "deny"},
			},
			"deny": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `The list of DENY rules specified by this firewall. Each rule specifies
a protocol and port-range tuple that describes a denied connection.`,
				Elem:         computeFirewallDenySchema(),
				Set:          resourceComputeFirewallRuleHash,
				ExactlyOneOf: []string{"allow", "deny"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"destination_ranges": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `If destination ranges are specified, the firewall will apply only to
traffic that has destination IP address in these ranges. These ranges
must be expressed in CIDR format. Only IPv4 is supported.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"source_ranges", "source_tags"},
			},
			"direction": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"INGRESS", "EGRESS", ""}, false),
				Description: `Direction of traffic to which this firewall applies; default is
INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
destinationRanges; For EGRESS traffic, it is NOT supported to specify
sourceRanges OR sourceTags. Possible values: ["INGRESS", "EGRESS"]`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Denotes whether the firewall rule is disabled, i.e not applied to the
network it is associated with. When set to true, the firewall rule is
not enforced and the network behaves as if it did not exist. If this
is unspecified, the firewall rule will be enabled.`,
			},
			"log_config": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: diffSuppressEnableLogging,
				Description: `This field denotes the logging options for a particular firewall rule.
If defined, logging is enabled, and logs will be exported to Cloud Logging.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"}, false),
							Description:  `This field denotes whether to include or exclude metadata for firewall logs. Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"]`,
						},
					},
				},
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description: `Priority for this rule. This is an integer between 0 and 65535, both
inclusive. When not specified, the value assumed is 1000. Relative
priorities determine precedence of conflicting rules. Lower value of
priority implies higher precedence (eg, a rule with priority 0 has
higher precedence than a rule with priority 1). DENY rules take
precedence over ALLOW rules having equal priority.`,
				Default: 1000,
			},
			"source_ranges": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `If source ranges are specified, the firewall will apply only to
traffic that has source IP address in these ranges. These ranges must
be expressed in CIDR format. One or both of sourceRanges and
sourceTags may be set. If both properties are set, the firewall will
apply to traffic that has source IP address within sourceRanges OR the
source IP that belongs to a tag listed in the sourceTags property. The
connection does not need to match both properties for the firewall to
apply. Only IPv4 is supported.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"destination_ranges"},
			},
			"source_service_accounts": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `If source service accounts are specified, the firewall will apply only
to traffic originating from an instance with a service account in this
list. Source service accounts cannot be used to control traffic to an
instance's external IP address because service accounts are associated
with an instance, not an IP address. sourceRanges can be set at the
same time as sourceServiceAccounts. If both are set, the firewall will
apply to traffic that has source IP address within sourceRanges OR the
source IP belongs to an instance with service account listed in
sourceServiceAccount. The connection does not need to match both
properties for the firewall to apply. sourceServiceAccounts cannot be
used at the same time as sourceTags or targetTags.`,
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"source_tags", "target_tags"},
			},
			"source_tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `If source tags are specified, the firewall will apply only to traffic
with source IP that belongs to a tag listed in source tags. Source
tags cannot be used to control traffic to an instance's external IP
address. Because tags are associated with an instance, not an IP
address. One or both of sourceRanges and sourceTags may be set. If
both properties are set, the firewall will apply to traffic that has
source IP address within sourceRanges OR the source IP that belongs to
a tag listed in the sourceTags property. The connection does not need
to match both properties for the firewall to apply.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"destination_ranges", "source_service_accounts", "target_service_accounts"},
			},
			"target_service_accounts": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `A list of service accounts indicating sets of instances located in the
network that may make network connections as specified in allowed[].
targetServiceAccounts cannot be used at the same time as targetTags or
sourceTags. If neither targetServiceAccounts nor targetTags are
specified, the firewall rule applies to all instances on the specified
network.`,
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"source_tags", "target_tags"},
			},
			"target_tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `A list of instance tags indicating sets of instances located in the
network that may make network connections as specified in allowed[].
If no targetTags are specified, the firewall rule applies to all
instances on the specified network.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"source_service_accounts", "target_service_accounts"},
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"enable_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Deprecated:  "Deprecated in favor of log_config",
				Description: "This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.",
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

func computeFirewallAllowSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareCaseInsensitive,
				Description: `The IP protocol to which this rule applies. The protocol type is
required when creating a firewall rule. This value can either be
one of the following well known protocol strings (tcp, udp,
icmp, esp, ah, sctp, ipip), or the IP protocol number.`,
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `An optional list of ports to which this rule applies. This field
is only applicable for UDP or TCP protocol. Each entry must be
either an integer or a range. If not specified, this rule
applies to connections through any port.

Example inputs include: ["22"], ["80","443"], and
["12345-12349"].`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func computeFirewallDenySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareCaseInsensitive,
				Description: `The IP protocol to which this rule applies. The protocol type is
required when creating a firewall rule. This value can either be
one of the following well known protocol strings (tcp, udp,
icmp, esp, ah, sctp, ipip), or the IP protocol number.`,
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `An optional list of ports to which this rule applies. This field
is only applicable for UDP or TCP protocol. Each entry must be
either an integer or a range. If not specified, this rule
applies to connections through any port.

Example inputs include: ["22"], ["80","443"], and
["12345-12349"].`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceComputeFirewallCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	allowedProp, err := expandComputeFirewallAllow(d.Get("allow"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow"); !isEmptyValue(reflect.ValueOf(allowedProp)) && (ok || !reflect.DeepEqual(v, allowedProp)) {
		obj["allowed"] = allowedProp
	}
	deniedProp, err := expandComputeFirewallDeny(d.Get("deny"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deny"); !isEmptyValue(reflect.ValueOf(deniedProp)) && (ok || !reflect.DeepEqual(v, deniedProp)) {
		obj["denied"] = deniedProp
	}
	descriptionProp, err := expandComputeFirewallDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	destinationRangesProp, err := expandComputeFirewallDestinationRanges(d.Get("destination_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_ranges"); !isEmptyValue(reflect.ValueOf(destinationRangesProp)) && (ok || !reflect.DeepEqual(v, destinationRangesProp)) {
		obj["destinationRanges"] = destinationRangesProp
	}
	directionProp, err := expandComputeFirewallDirection(d.Get("direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("direction"); !isEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	disabledProp, err := expandComputeFirewallDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); ok || !reflect.DeepEqual(v, disabledProp) {
		obj["disabled"] = disabledProp
	}
	logConfigProp, err := expandComputeFirewallLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
		obj["logConfig"] = logConfigProp
	}
	nameProp, err := expandComputeFirewallName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeFirewallNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeFirewallPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	sourceRangesProp, err := expandComputeFirewallSourceRanges(d.Get("source_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_ranges"); !isEmptyValue(reflect.ValueOf(sourceRangesProp)) && (ok || !reflect.DeepEqual(v, sourceRangesProp)) {
		obj["sourceRanges"] = sourceRangesProp
	}
	sourceServiceAccountsProp, err := expandComputeFirewallSourceServiceAccounts(d.Get("source_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_service_accounts"); !isEmptyValue(reflect.ValueOf(sourceServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, sourceServiceAccountsProp)) {
		obj["sourceServiceAccounts"] = sourceServiceAccountsProp
	}
	sourceTagsProp, err := expandComputeFirewallSourceTags(d.Get("source_tags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_tags"); !isEmptyValue(reflect.ValueOf(sourceTagsProp)) && (ok || !reflect.DeepEqual(v, sourceTagsProp)) {
		obj["sourceTags"] = sourceTagsProp
	}
	targetServiceAccountsProp, err := expandComputeFirewallTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !isEmptyValue(reflect.ValueOf(targetServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}
	targetTagsProp, err := expandComputeFirewallTargetTags(d.Get("target_tags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_tags"); !isEmptyValue(reflect.ValueOf(targetTagsProp)) && (ok || !reflect.DeepEqual(v, targetTagsProp)) {
		obj["targetTags"] = targetTagsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/firewalls")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Firewall: %#v", obj)
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

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Firewall: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating Firewall",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Firewall: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Firewall %q: %#v", d.Id(), res)

	return resourceComputeFirewallRead(d, meta)
}

func resourceComputeFirewallRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/firewalls/{{name}}")
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

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeFirewall %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}

	if err := d.Set("allow", flattenComputeFirewallAllow(res["allowed"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeFirewallCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("deny", flattenComputeFirewallDeny(res["denied"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("description", flattenComputeFirewallDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("destination_ranges", flattenComputeFirewallDestinationRanges(res["destinationRanges"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("direction", flattenComputeFirewallDirection(res["direction"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("disabled", flattenComputeFirewallDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("log_config", flattenComputeFirewallLogConfig(res["logConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("name", flattenComputeFirewallName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("network", flattenComputeFirewallNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("priority", flattenComputeFirewallPriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("source_ranges", flattenComputeFirewallSourceRanges(res["sourceRanges"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("source_service_accounts", flattenComputeFirewallSourceServiceAccounts(res["sourceServiceAccounts"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("source_tags", flattenComputeFirewallSourceTags(res["sourceTags"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("target_service_accounts", flattenComputeFirewallTargetServiceAccounts(res["targetServiceAccounts"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("target_tags", flattenComputeFirewallTargetTags(res["targetTags"], d, config)); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Firewall: %s", err)
	}

	return nil
}

func resourceComputeFirewallUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	allowedProp, err := expandComputeFirewallAllow(d.Get("allow"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowedProp)) {
		obj["allowed"] = allowedProp
	}
	deniedProp, err := expandComputeFirewallDeny(d.Get("deny"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deny"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deniedProp)) {
		obj["denied"] = deniedProp
	}
	descriptionProp, err := expandComputeFirewallDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	destinationRangesProp, err := expandComputeFirewallDestinationRanges(d.Get("destination_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_ranges"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, destinationRangesProp)) {
		obj["destinationRanges"] = destinationRangesProp
	}
	disabledProp, err := expandComputeFirewallDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); ok || !reflect.DeepEqual(v, disabledProp) {
		obj["disabled"] = disabledProp
	}
	logConfigProp, err := expandComputeFirewallLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
		obj["logConfig"] = logConfigProp
	}
	networkProp, err := expandComputeFirewallNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeFirewallPriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	sourceRangesProp, err := expandComputeFirewallSourceRanges(d.Get("source_ranges"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_ranges"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceRangesProp)) {
		obj["sourceRanges"] = sourceRangesProp
	}
	sourceServiceAccountsProp, err := expandComputeFirewallSourceServiceAccounts(d.Get("source_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_service_accounts"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceServiceAccountsProp)) {
		obj["sourceServiceAccounts"] = sourceServiceAccountsProp
	}
	sourceTagsProp, err := expandComputeFirewallSourceTags(d.Get("source_tags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_tags"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceTagsProp)) {
		obj["sourceTags"] = sourceTagsProp
	}
	targetServiceAccountsProp, err := expandComputeFirewallTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}
	targetTagsProp, err := expandComputeFirewallTargetTags(d.Get("target_tags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_tags"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetTagsProp)) {
		obj["targetTags"] = targetTagsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Firewall %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Firewall %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Firewall %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating Firewall",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeFirewallRead(d, meta)
}

func resourceComputeFirewallDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Firewall %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Firewall")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting Firewall",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Firewall %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeFirewallImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/firewalls/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/firewalls/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeFirewallAllow(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallAllowProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallAllowPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeFirewallAllowProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallAllowPorts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallDeny(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallDenyProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallDenyPorts(original["ports"], d, config),
		})
	}
	return transformed
}
func flattenComputeFirewallDenyProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallDenyPorts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallDestinationRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallDirection(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallDisabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallLogConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}

	v, ok := original["enable"]
	if ok && !v.(bool) {
		return nil
	}

	transformed := make(map[string]interface{})
	transformed["metadata"] = original["metadata"]
	return []interface{}{transformed}
}

func flattenComputeFirewallName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeFirewallNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeFirewallPriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeFirewallSourceRanges(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceServiceAccounts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceTags(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetServiceAccounts(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetTags(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func expandComputeFirewallAllow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProtocol, err := expandComputeFirewallAllowProtocol(original["protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["IPProtocol"] = transformedProtocol
		}

		transformedPorts, err := expandComputeFirewallAllowPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallAllowProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallAllowPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDeny(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProtocol, err := expandComputeFirewallDenyProtocol(original["protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProtocol); val.IsValid() && !isEmptyValue(val) {
			transformed["IPProtocol"] = transformedProtocol
		}

		transformedPorts, err := expandComputeFirewallDenyPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !isEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallDenyProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDenyPorts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDestinationRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallDirection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallLogConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	transformed := make(map[string]interface{})

	if len(l) == 0 || l[0] == nil {
		// send enable = enable_logging value to ensure correct logging status if there is no config
		transformed["enable"] = d.Get("enable_logging").(bool)
		return transformed, nil
	}

	raw := l[0]
	original := raw.(map[string]interface{})

	// The log_config block is specified, so logging should be enabled
	transformed["enable"] = true
	transformed["metadata"] = original["metadata"]

	return transformed, nil
}

func expandComputeFirewallName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeFirewallPriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallSourceRanges(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallSourceServiceAccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallSourceTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallTargetServiceAccounts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeFirewallTargetTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}
