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

func TestAccPubsubLiteReservation_pubsubLiteReservationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckPubsubLiteReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubLiteReservation_pubsubLiteReservationBasicExample(context),
			},
			{
				ResourceName:            "google_pubsub_lite_reservation.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "name"},
			},
		},
	})
}

func testAccPubsubLiteReservation_pubsubLiteReservationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_lite_reservation" "example" {
  name = "tf-test-example-reservation%{random_suffix}"
  project = data.google_project.project.number
  throughput_capacity = 2
}

data "google_project" "project" {
}
`, context)
}

func testAccCheckPubsubLiteReservationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_pubsub_lite_reservation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PubsubLiteBasePath}}projects/{{project}}/locations/{{region}}/reservations/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("PubsubLiteReservation still exists at %s", url)
			}
		}

		return nil
	}
}
