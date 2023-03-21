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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeHaVpnGateway_haVpnGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckComputeHaVpnGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHaVpnGateway_haVpnGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_compute_ha_vpn_gateway.ha_gateway1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeHaVpnGateway_haVpnGatewayBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ha_vpn_gateway" "ha_gateway1" {
  region   = "us-central1"
  name     = "tf-test-ha-vpn-1%{random_suffix}"
  network  = google_compute_network.network1.id
}

resource "google_compute_network" "network1" {
  name                    = "network1%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeHaVpnGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_ha_vpn_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeHaVpnGateway still exists at %s", url)
			}
		}

		return nil
	}
}
