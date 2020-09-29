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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataFusionInstance_dataFusionInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceBasicExample(context),
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  provider = google-beta
  name = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type = "BASIC"
}
`, context)
}

func TestAccDataFusionInstance_dataFusionInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDataFusionInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_dataFusionInstanceFullExample(context),
			},
		},
	})
}

func testAccDataFusionInstance_dataFusionInstanceFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_fusion_instance" "extended_instance" {
  provider = google-beta
  name = "tf-test-my-instance%{random_suffix}"
  description = "My Data Fusion instance"
  region = "us-central1"
  type = "BASIC"
  enable_stackdriver_logging = true
  enable_stackdriver_monitoring = true
  labels = {
    example_key = "example_value"
  }
  private_instance = true
  network_config {
    network = "default"
    ip_allocation = "10.89.48.0/22"
  }
  version = "6.1.1"
}
`, context)
}

func testAccCheckDataFusionInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_fusion_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataFusionBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("DataFusionInstance still exists at %s", url)
			}
		}

		return nil
	}
}
