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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// waitForAttachmentToBeProvisioned waits for an attachment to leave the
// "UNPROVISIONED" state, to indicate that it's either ready or awaiting partner
// activity.
func waitForAttachmentToBeProvisioned(d *schema.ResourceData, config *Config, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		if err := resourceComputeInterconnectAttachmentRead(d, config); err != nil {
			return resource.NonRetryableError(err)
		}

		name := d.Get("name").(string)
		state := d.Get("state").(string)
		if state == "UNPROVISIONED" {
			return resource.RetryableError(fmt.Errorf("InterconnectAttachment %q has state %q.", name, state))
		}
		log.Printf("InterconnectAttachment %q has state %q.", name, state)
		return nil
	})
}

func resourceComputeInterconnectAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInterconnectAttachmentCreate,
		Read:   resourceComputeInterconnectAttachmentRead,
		Update: resourceComputeInterconnectAttachmentUpdate,
		Delete: resourceComputeInterconnectAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeInterconnectAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
				Description: `Name of the resource. Provided by the client when the resource is created. The
name must be 1-63 characters long, and comply with RFC1035. Specifically, the
name must be 1-63 characters long and match the regular expression
'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a
lowercase letter, and all following characters must be a dash, lowercase
letter, or digit, except the last character, which cannot be a dash.`,
			},
			"router": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL of the cloud router to be used for dynamic routing. This router must be in
the same region as this InterconnectAttachment. The InterconnectAttachment will
automatically connect the Interconnect to the network & region within which the
Cloud Router is configured.`,
			},
			"admin_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether the VLAN attachment is enabled or disabled.  When using
PARTNER type this will Pre-Activate the interconnect attachment`,
				Default: true,
			},
			"bandwidth": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G", ""}, false),
				Description: `Provisioned bandwidth capacity for the interconnect attachment.
For attachments of type DEDICATED, the user can set the bandwidth.
For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"]`,
			},
			"candidate_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Up to 16 candidate prefixes that can be used to restrict the allocation
of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
All prefixes must be within link-local address space (169.254.0.0/16)
and must be /29 or shorter (/28, /27, etc). Google will attempt to select
an unused /29 from the supplied candidate prefix(es). The request will
fail if all possible /29s are in use on Google's edge. If not supplied,
Google will randomly select an unused /29 from all of link-local space.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"edge_availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Desired availability domain for the attachment. Only available for type
PARTNER, at creation time. For improved reliability, customers should
configure a pair of attachments with one per availability domain. The
selected availability domain will be provided to the Partner via the
pairing key so that the provisioned circuit will lie in the specified
domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.`,
			},
			"encryption": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "IPSEC", ""}, false),
				Description: `Indicates the user-supplied encryption option of this interconnect
attachment:

NONE is the default value, which means that the attachment carries
unencrypted traffic. VMs can send traffic to, or receive traffic
from, this type of attachment.

IPSEC indicates that the attachment carries only traffic encrypted by
an IPsec device such as an HA VPN gateway. VMs cannot directly send
traffic to, or receive traffic from, such an attachment. To use
IPsec-encrypted Cloud Interconnect create the attachment using this
option.

Not currently available publicly. Default value: "NONE" Possible values: ["NONE", "IPSEC"]`,
				Default: "NONE",
			},
			"interconnect": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL of the underlying Interconnect object that this attachment's
traffic will traverse through. Required if type is DEDICATED, must not
be set if type is PARTNER.`,
			},
			"ipsec_internal_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `URL of addresses that have been reserved for the interconnect
attachment, Used only for interconnect attachment that has the
encryption option as IPSEC.

The addresses must be RFC 1918 IP address ranges. When creating HA
VPN gateway over the interconnect attachment, if the attachment is
configured to use an RFC 1918 IP address, then the VPN gateway's IP
address will be allocated from the IP address range specified
here.

For example, if the HA VPN gateway's interface 0 is paired to this
interconnect attachment, then an RFC 1918 IP address for the VPN
gateway interface 0 will be allocated from the IP address specified
for this interconnect attachment.

If this field is not specified for interconnect attachment that has
encryption option as IPSEC, later on when creating HA VPN gateway on
this interconnect attachment, the HA VPN gateway's IP address will be
allocated from regional external IP address pool.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"mtu": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `Maximum Transmission Unit (MTU), in bytes, of packets passing through
this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Region where the regional interconnect attachment resides.`,
			},
			"type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DEDICATED", "PARTNER", "PARTNER_PROVIDER", ""}, false),
				Description: `The type of InterconnectAttachment you wish to create. Defaults to
DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"]`,
			},
			"vlan_tag8021q": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
using PARTNER type this will be managed upstream.`,
			},
			"cloud_router_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `IPv4 address + prefix length to be configured on Cloud Router
Interface for this interconnect attachment.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"customer_router_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `IPv4 address + prefix length to be configured on the customer
router subinterface for this interconnect attachment.`,
			},
			"google_reference_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Google reference ID, to be used when raising support tickets with
Google or otherwise to debug backend connectivity issues.`,
			},
			"pairing_key": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `[Output only for type PARTNER. Not present for DEDICATED]. The opaque
identifier of an PARTNER attachment used to initiate provisioning with
a selected partner. Of the form "XXXXX/region/domain"`,
			},
			"partner_asn": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `[Output only for type PARTNER. Not present for DEDICATED]. Optional
BGP ASN for the router that should be supplied by a layer 3 Partner if
they configured BGP on behalf of the customer.`,
			},
			"private_interconnect_info": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Information specific to an InterconnectAttachment. This property
is populated if the interconnect that this is attached to is of type DEDICATED.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag8021q": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: `802.1q encapsulation tag to be used for traffic between
Google and the customer, going to and from this network and region.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `[Output Only] The current state of this attachment's functionality.`,
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

func resourceComputeInterconnectAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	adminEnabledProp, err := expandComputeInterconnectAttachmentAdminEnabled(d.Get("admin_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admin_enabled"); ok || !reflect.DeepEqual(v, adminEnabledProp) {
		obj["adminEnabled"] = adminEnabledProp
	}
	interconnectProp, err := expandComputeInterconnectAttachmentInterconnect(d.Get("interconnect"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interconnect"); !isEmptyValue(reflect.ValueOf(interconnectProp)) && (ok || !reflect.DeepEqual(v, interconnectProp)) {
		obj["interconnect"] = interconnectProp
	}
	descriptionProp, err := expandComputeInterconnectAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	mtuProp, err := expandComputeInterconnectAttachmentMtu(d.Get("mtu"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtu"); !isEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	bandwidthProp, err := expandComputeInterconnectAttachmentBandwidth(d.Get("bandwidth"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bandwidth"); !isEmptyValue(reflect.ValueOf(bandwidthProp)) && (ok || !reflect.DeepEqual(v, bandwidthProp)) {
		obj["bandwidth"] = bandwidthProp
	}
	edgeAvailabilityDomainProp, err := expandComputeInterconnectAttachmentEdgeAvailabilityDomain(d.Get("edge_availability_domain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edge_availability_domain"); !isEmptyValue(reflect.ValueOf(edgeAvailabilityDomainProp)) && (ok || !reflect.DeepEqual(v, edgeAvailabilityDomainProp)) {
		obj["edgeAvailabilityDomain"] = edgeAvailabilityDomainProp
	}
	typeProp, err := expandComputeInterconnectAttachmentType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	routerProp, err := expandComputeInterconnectAttachmentRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	nameProp, err := expandComputeInterconnectAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	candidateSubnetsProp, err := expandComputeInterconnectAttachmentCandidateSubnets(d.Get("candidate_subnets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("candidate_subnets"); !isEmptyValue(reflect.ValueOf(candidateSubnetsProp)) && (ok || !reflect.DeepEqual(v, candidateSubnetsProp)) {
		obj["candidateSubnets"] = candidateSubnetsProp
	}
	vlanTag8021qProp, err := expandComputeInterconnectAttachmentVlanTag8021q(d.Get("vlan_tag8021q"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vlan_tag8021q"); !isEmptyValue(reflect.ValueOf(vlanTag8021qProp)) && (ok || !reflect.DeepEqual(v, vlanTag8021qProp)) {
		obj["vlanTag8021q"] = vlanTag8021qProp
	}
	ipsecInternalAddressesProp, err := expandComputeInterconnectAttachmentIpsecInternalAddresses(d.Get("ipsec_internal_addresses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ipsec_internal_addresses"); !isEmptyValue(reflect.ValueOf(ipsecInternalAddressesProp)) && (ok || !reflect.DeepEqual(v, ipsecInternalAddressesProp)) {
		obj["ipsecInternalAddresses"] = ipsecInternalAddressesProp
	}
	encryptionProp, err := expandComputeInterconnectAttachmentEncryption(d.Get("encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption"); !isEmptyValue(reflect.ValueOf(encryptionProp)) && (ok || !reflect.DeepEqual(v, encryptionProp)) {
		obj["encryption"] = encryptionProp
	}
	regionProp, err := expandComputeInterconnectAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InterconnectAttachment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterconnectAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating InterconnectAttachment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating InterconnectAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InterconnectAttachment: %s", err)
	}

	if err := waitForAttachmentToBeProvisioned(d, config, d.Timeout(schema.TimeoutCreate)); err != nil {
		return fmt.Errorf("Error waiting for InterconnectAttachment %q to be provisioned: %q", d.Get("name").(string), err)
	}

	log.Printf("[DEBUG] Finished creating InterconnectAttachment %q: %#v", d.Id(), res)

	return resourceComputeInterconnectAttachmentRead(d, meta)
}

func resourceComputeInterconnectAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterconnectAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeInterconnectAttachment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}

	if err := d.Set("admin_enabled", flattenComputeInterconnectAttachmentAdminEnabled(res["adminEnabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("cloud_router_ip_address", flattenComputeInterconnectAttachmentCloudRouterIpAddress(res["cloudRouterIpAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("customer_router_ip_address", flattenComputeInterconnectAttachmentCustomerRouterIpAddress(res["customerRouterIpAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("interconnect", flattenComputeInterconnectAttachmentInterconnect(res["interconnect"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("description", flattenComputeInterconnectAttachmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("mtu", flattenComputeInterconnectAttachmentMtu(res["mtu"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("bandwidth", flattenComputeInterconnectAttachmentBandwidth(res["bandwidth"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("edge_availability_domain", flattenComputeInterconnectAttachmentEdgeAvailabilityDomain(res["edgeAvailabilityDomain"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("pairing_key", flattenComputeInterconnectAttachmentPairingKey(res["pairingKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("partner_asn", flattenComputeInterconnectAttachmentPartnerAsn(res["partnerAsn"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("private_interconnect_info", flattenComputeInterconnectAttachmentPrivateInterconnectInfo(res["privateInterconnectInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("type", flattenComputeInterconnectAttachmentType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("state", flattenComputeInterconnectAttachmentState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("google_reference_id", flattenComputeInterconnectAttachmentGoogleReferenceId(res["googleReferenceId"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("router", flattenComputeInterconnectAttachmentRouter(res["router"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeInterconnectAttachmentCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("name", flattenComputeInterconnectAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("vlan_tag8021q", flattenComputeInterconnectAttachmentVlanTag8021q(res["vlanTag8021q"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("ipsec_internal_addresses", flattenComputeInterconnectAttachmentIpsecInternalAddresses(res["ipsecInternalAddresses"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("encryption", flattenComputeInterconnectAttachmentEncryption(res["encryption"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("region", flattenComputeInterconnectAttachmentRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}

	return nil
}

func resourceComputeInterconnectAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterconnectAttachment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	adminEnabledProp, err := expandComputeInterconnectAttachmentAdminEnabled(d.Get("admin_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admin_enabled"); ok || !reflect.DeepEqual(v, adminEnabledProp) {
		obj["adminEnabled"] = adminEnabledProp
	}
	descriptionProp, err := expandComputeInterconnectAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	mtuProp, err := expandComputeInterconnectAttachmentMtu(d.Get("mtu"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mtu"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	bandwidthProp, err := expandComputeInterconnectAttachmentBandwidth(d.Get("bandwidth"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bandwidth"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bandwidthProp)) {
		obj["bandwidth"] = bandwidthProp
	}
	regionProp, err := expandComputeInterconnectAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating InterconnectAttachment %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating InterconnectAttachment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating InterconnectAttachment %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating InterconnectAttachment", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeInterconnectAttachmentRead(d, meta)
}

func resourceComputeInterconnectAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InterconnectAttachment: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	if err := waitForAttachmentToBeProvisioned(d, config, d.Timeout(schema.TimeoutCreate)); err != nil {
		return fmt.Errorf("Error waiting for InterconnectAttachment %q to be provisioned: %q", d.Get("name").(string), err)
	}
	log.Printf("[DEBUG] Deleting InterconnectAttachment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "InterconnectAttachment")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting InterconnectAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InterconnectAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeInterconnectAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/interconnectAttachments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeInterconnectAttachmentAdminEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentCloudRouterIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentCustomerRouterIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentInterconnect(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentMtu(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return fmt.Sprintf("%d", int(floatVal))
	}
	return v
}

func flattenComputeInterconnectAttachmentBandwidth(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentEdgeAvailabilityDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPairingKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPartnerAsn(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPrivateInterconnectInfo(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["tag8021q"] =
		flattenComputeInterconnectAttachmentPrivateInterconnectInfoTag8021q(original["tag8021q"], d, config)
	return []interface{}{transformed}
}
func flattenComputeInterconnectAttachmentPrivateInterconnectInfoTag8021q(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeInterconnectAttachmentType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentGoogleReferenceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentRouter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeInterconnectAttachmentCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentVlanTag8021q(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeInterconnectAttachmentIpsecInternalAddresses(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeInterconnectAttachmentEncryption(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeInterconnectAttachmentAdminEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentInterconnect(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentMtu(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentBandwidth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentEdgeAvailabilityDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeInterconnectAttachmentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentCandidateSubnets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentVlanTag8021q(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentIpsecInternalAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ipsec_internal_addresses: nil")
		}
		f, err := parseRegionalFieldValue("addresses", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ipsec_internal_addresses: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeInterconnectAttachmentEncryption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
