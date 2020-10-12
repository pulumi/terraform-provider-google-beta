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
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// validatePeerAddr returns false if a tunnel's peer_ip property
// is invalid. Currently, only addresses that collide with RFC
// 5735 (https://tools.ietf.org/html/rfc5735) fail validation.
func validatePeerAddr(i interface{}, val string) ([]string, []error) {
	ip := net.ParseIP(i.(string))
	if ip == nil {
		return nil, []error{fmt.Errorf("could not parse %q to IP address", val)}
	}
	for _, test := range invalidPeerAddrs {
		if bytes.Compare(ip, test.from) >= 0 && bytes.Compare(ip, test.to) <= 0 {
			return nil, []error{fmt.Errorf("address is invalid (is between %q and %q, conflicting with RFC5735)", test.from, test.to)}
		}
	}
	return nil, nil
}

// invalidPeerAddrs is a collection of IP address ranges that represent
// a conflict with RFC 5735 (https://tools.ietf.org/html/rfc5735#page-3).
// CIDR range notations in the RFC were converted to a (from, to) pair
// for easy checking with bytes.Compare.
var invalidPeerAddrs = []struct {
	from net.IP
	to   net.IP
}{
	{
		from: net.ParseIP("0.0.0.0"),
		to:   net.ParseIP("0.255.255.255"),
	},
	{
		from: net.ParseIP("10.0.0.0"),
		to:   net.ParseIP("10.255.255.255"),
	},
	{
		from: net.ParseIP("127.0.0.0"),
		to:   net.ParseIP("127.255.255.255"),
	},
	{
		from: net.ParseIP("169.254.0.0"),
		to:   net.ParseIP("169.254.255.255"),
	},
	{
		from: net.ParseIP("172.16.0.0"),
		to:   net.ParseIP("172.31.255.255"),
	},
	{
		from: net.ParseIP("192.0.0.0"),
		to:   net.ParseIP("192.0.0.255"),
	},
	{
		from: net.ParseIP("192.0.2.0"),
		to:   net.ParseIP("192.0.2.255"),
	},
	{
		from: net.ParseIP("192.88.99.0"),
		to:   net.ParseIP("192.88.99.255"),
	},
	{
		from: net.ParseIP("192.168.0.0"),
		to:   net.ParseIP("192.168.255.255"),
	},
	{
		from: net.ParseIP("198.18.0.0"),
		to:   net.ParseIP("198.19.255.255"),
	},
	{
		from: net.ParseIP("198.51.100.0"),
		to:   net.ParseIP("198.51.100.255"),
	},
	{
		from: net.ParseIP("203.0.113.0"),
		to:   net.ParseIP("203.0.113.255"),
	},
	{
		from: net.ParseIP("224.0.0.0"),
		to:   net.ParseIP("239.255.255.255"),
	},
	{
		from: net.ParseIP("240.0.0.0"),
		to:   net.ParseIP("255.255.255.255"),
	},
	{
		from: net.ParseIP("255.255.255.255"),
		to:   net.ParseIP("255.255.255.255"),
	},
}

func getVpnTunnelLink(config *Config, project, region, tunnel, userAgent string) (string, error) {
	if !strings.Contains(tunnel, "/") {
		// Tunnel value provided is just the name, lookup the tunnel SelfLink
		tunnelData, err := config.NewComputeClient(userAgent).VpnTunnels.Get(
			project, region, tunnel).Do()
		if err != nil {
			return "", fmt.Errorf("Error reading tunnel: %s", err)
		}
		tunnel = tunnelData.SelfLink
	}

	return tunnel, nil

}

func resourceComputeVpnTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeVpnTunnelCreate,
		Read:   resourceComputeVpnTunnelRead,
		Update: resourceComputeVpnTunnelUpdate,
		Delete: resourceComputeVpnTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeVpnTunnelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63
characters long and match the regular expression
'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
must be a lowercase letter, and all following characters must
be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"shared_secret": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Shared secret used to set the secure session between the Cloud VPN
gateway and the peer VPN gateway.`,
				Sensitive: true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"ike_version": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `IKE protocol version to use when establishing the VPN tunnel with
peer VPN gateway.
Acceptable IKE versions are 1 or 2. Default version is 2.`,
				Default: 2,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this VpnTunnel.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"local_traffic_selector": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Local traffic selector to use when establishing the VPN tunnel with
peer VPN gateway. The value should be a CIDR formatted string,
for example '192.168.0.0/16'. The ranges should be disjoint.
Only IPv4 is supported.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"peer_external_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the peer side external VPN gateway to which this VPN tunnel is connected.`,
			},
			"peer_external_gateway_interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `The interface ID of the external VPN gateway to which this VPN tunnel is connected.`,
			},
			"peer_gcp_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
ID in the peer GCP VPN gateway.
This field must reference a 'google_compute_ha_vpn_gateway' resource.`,
			},
			"peer_ip": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validatePeerAddr,
				Description:  `IP address of the peer VPN gateway. Only IPv4 is supported.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.`,
			},
			"remote_traffic_selector": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Remote traffic selector to use when establishing the VPN tunnel with
peer VPN gateway. The value should be a CIDR formatted string,
for example '192.168.0.0/16'. The ranges should be disjoint.
Only IPv4 is supported.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"router": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of router resource to be used for dynamic routing.`,
			},
			"target_vpn_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL of the Target VPN gateway with which this VPN tunnel is
associated.`,
			},
			"vpn_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL of the VPN gateway with which this VPN tunnel is associated.
This must be used if a High Availability VPN gateway resource is created.
This field must reference a 'google_compute_ha_vpn_gateway' resource.`,
			},
			"vpn_gateway_interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `The interface ID of the VPN gateway with which this VPN tunnel is associated.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"detailed_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Detailed status message for the VPN tunnel.`,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fingerprint used for optimistic locking of this resource.  Used
internally during updates.`,
			},
			"shared_secret_hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Hash of the shared secret.`,
			},
			"tunnel_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier for the resource. This identifier is defined by the server.`,
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

func resourceComputeVpnTunnelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeVpnTunnelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeVpnTunnelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	targetVpnGatewayProp, err := expandComputeVpnTunnelTargetVpnGateway(d.Get("target_vpn_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_vpn_gateway"); !isEmptyValue(reflect.ValueOf(targetVpnGatewayProp)) && (ok || !reflect.DeepEqual(v, targetVpnGatewayProp)) {
		obj["targetVpnGateway"] = targetVpnGatewayProp
	}
	vpnGatewayProp, err := expandComputeVpnTunnelVpnGateway(d.Get("vpn_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpn_gateway"); !isEmptyValue(reflect.ValueOf(vpnGatewayProp)) && (ok || !reflect.DeepEqual(v, vpnGatewayProp)) {
		obj["vpnGateway"] = vpnGatewayProp
	}
	vpnGatewayInterfaceProp, err := expandComputeVpnTunnelVpnGatewayInterface(d.Get("vpn_gateway_interface"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpn_gateway_interface"); ok || !reflect.DeepEqual(v, vpnGatewayInterfaceProp) {
		obj["vpnGatewayInterface"] = vpnGatewayInterfaceProp
	}
	peerExternalGatewayProp, err := expandComputeVpnTunnelPeerExternalGateway(d.Get("peer_external_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_external_gateway"); !isEmptyValue(reflect.ValueOf(peerExternalGatewayProp)) && (ok || !reflect.DeepEqual(v, peerExternalGatewayProp)) {
		obj["peerExternalGateway"] = peerExternalGatewayProp
	}
	peerExternalGatewayInterfaceProp, err := expandComputeVpnTunnelPeerExternalGatewayInterface(d.Get("peer_external_gateway_interface"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_external_gateway_interface"); ok || !reflect.DeepEqual(v, peerExternalGatewayInterfaceProp) {
		obj["peerExternalGatewayInterface"] = peerExternalGatewayInterfaceProp
	}
	peerGcpGatewayProp, err := expandComputeVpnTunnelPeerGcpGateway(d.Get("peer_gcp_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_gcp_gateway"); !isEmptyValue(reflect.ValueOf(peerGcpGatewayProp)) && (ok || !reflect.DeepEqual(v, peerGcpGatewayProp)) {
		obj["peerGcpGateway"] = peerGcpGatewayProp
	}
	routerProp, err := expandComputeVpnTunnelRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	peerIpProp, err := expandComputeVpnTunnelPeerIp(d.Get("peer_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_ip"); !isEmptyValue(reflect.ValueOf(peerIpProp)) && (ok || !reflect.DeepEqual(v, peerIpProp)) {
		obj["peerIp"] = peerIpProp
	}
	sharedSecretProp, err := expandComputeVpnTunnelSharedSecret(d.Get("shared_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("shared_secret"); !isEmptyValue(reflect.ValueOf(sharedSecretProp)) && (ok || !reflect.DeepEqual(v, sharedSecretProp)) {
		obj["sharedSecret"] = sharedSecretProp
	}
	ikeVersionProp, err := expandComputeVpnTunnelIkeVersion(d.Get("ike_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ike_version"); !isEmptyValue(reflect.ValueOf(ikeVersionProp)) && (ok || !reflect.DeepEqual(v, ikeVersionProp)) {
		obj["ikeVersion"] = ikeVersionProp
	}
	localTrafficSelectorProp, err := expandComputeVpnTunnelLocalTrafficSelector(d.Get("local_traffic_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_traffic_selector"); !isEmptyValue(reflect.ValueOf(localTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, localTrafficSelectorProp)) {
		obj["localTrafficSelector"] = localTrafficSelectorProp
	}
	remoteTrafficSelectorProp, err := expandComputeVpnTunnelRemoteTrafficSelector(d.Get("remote_traffic_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("remote_traffic_selector"); !isEmptyValue(reflect.ValueOf(remoteTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, remoteTrafficSelectorProp)) {
		obj["remoteTrafficSelector"] = remoteTrafficSelectorProp
	}
	labelsProp, err := expandComputeVpnTunnelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeVpnTunnelLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	regionProp, err := expandComputeVpnTunnelRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	obj, err = resourceComputeVpnTunnelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VpnTunnel: %#v", obj)
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
		return fmt.Errorf("Error creating VpnTunnel: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating VpnTunnel", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create VpnTunnel: %s", err)
	}

	log.Printf("[DEBUG] Finished creating VpnTunnel %q: %#v", d.Id(), res)

	if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		// Labels cannot be set in a create.  We'll have to set them here.
		err = resourceComputeVpnTunnelRead(d, meta)
		if err != nil {
			return err
		}

		obj := make(map[string]interface{})
		// d.Get("labels") will have been overridden by the Read call.
		labelsProp, err := expandComputeVpnTunnelLabels(v, d, config)
		if err != nil {
			return err
		}
		obj["labels"] = labelsProp
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp

		url, err = replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", project, url, userAgent, obj)
		if err != nil {
			return fmt.Errorf("Error adding labels to ComputeVpnTunnel %q: %s", d.Id(), err)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating ComputeVpnTunnel Labels", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}

	}

	return resourceComputeVpnTunnelRead(d, meta)
}

func resourceComputeVpnTunnelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
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
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeVpnTunnel %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}

	if err := d.Set("tunnel_id", flattenComputeVpnTunnelTunnelId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeVpnTunnelCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("name", flattenComputeVpnTunnelName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("description", flattenComputeVpnTunnelDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("target_vpn_gateway", flattenComputeVpnTunnelTargetVpnGateway(res["targetVpnGateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("vpn_gateway", flattenComputeVpnTunnelVpnGateway(res["vpnGateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("vpn_gateway_interface", flattenComputeVpnTunnelVpnGatewayInterface(res["vpnGatewayInterface"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("peer_external_gateway", flattenComputeVpnTunnelPeerExternalGateway(res["peerExternalGateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("peer_external_gateway_interface", flattenComputeVpnTunnelPeerExternalGatewayInterface(res["peerExternalGatewayInterface"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("peer_gcp_gateway", flattenComputeVpnTunnelPeerGcpGateway(res["peerGcpGateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("router", flattenComputeVpnTunnelRouter(res["router"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("peer_ip", flattenComputeVpnTunnelPeerIp(res["peerIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("shared_secret_hash", flattenComputeVpnTunnelSharedSecretHash(res["sharedSecretHash"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("ike_version", flattenComputeVpnTunnelIkeVersion(res["ikeVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("local_traffic_selector", flattenComputeVpnTunnelLocalTrafficSelector(res["localTrafficSelector"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("remote_traffic_selector", flattenComputeVpnTunnelRemoteTrafficSelector(res["remoteTrafficSelector"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("labels", flattenComputeVpnTunnelLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeVpnTunnelLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("detailed_status", flattenComputeVpnTunnelDetailedStatus(res["detailedStatus"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("region", flattenComputeVpnTunnelRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}

	return nil
}

func resourceComputeVpnTunnelUpdate(d *schema.ResourceData, meta interface{}) error {
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

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})

		labelsProp, err := expandComputeVpnTunnelLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeVpnTunnelLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}/setLabels")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating VpnTunnel %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating VpnTunnel %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating VpnTunnel", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeVpnTunnelRead(d, meta)
}

func resourceComputeVpnTunnelDelete(d *schema.ResourceData, meta interface{}) error {
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

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting VpnTunnel %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "VpnTunnel")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting VpnTunnel", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting VpnTunnel %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeVpnTunnelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/vpnTunnels/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeVpnTunnelTunnelId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelTargetVpnGateway(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelVpnGateway(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelVpnGatewayInterface(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeVpnTunnelPeerExternalGateway(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelPeerExternalGatewayInterface(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeVpnTunnelPeerGcpGateway(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelRouter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelPeerIp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelSharedSecretHash(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelIkeVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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

func flattenComputeVpnTunnelLocalTrafficSelector(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeVpnTunnelRemoteTrafficSelector(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeVpnTunnelLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelLabelFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelDetailedStatus(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnTunnelRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeVpnTunnelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelTargetVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("targetVpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelVpnGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelPeerExternalGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("externalVpnGateways", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for peer_external_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelPeerExternalGatewayInterface(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelPeerGcpGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for peer_gcp_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	}
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}"+f.RelativeLink())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func expandComputeVpnTunnelPeerIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelSharedSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelIkeVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelLocalTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRemoteTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeVpnTunnelLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeVpnTunnelEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	f, err := parseRegionalFieldValue("targetVpnGateways", d.Get("target_vpn_gateway").(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, err
	}
	if _, ok := d.GetOk("project"); !ok {
		if err := d.Set("project", f.Project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	if _, ok := d.GetOk("region"); !ok {
		if err := d.Set("region", f.Region); err != nil {
			return nil, fmt.Errorf("Error setting region: %s", err)
		}
	}
	return obj, nil
}
