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

func TestAccGameServicesRealm_gameServiceRealmBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckGameServicesRealmDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGameServicesRealm_gameServiceRealmBasicExample(context),
			},
			{
				ResourceName:            "google_game_services_realm.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "realm_id"},
			},
		},
	})
}

func testAccGameServicesRealm_gameServiceRealmBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_game_services_realm" "default" {
  realm_id  = "tf-test-tf-test-realm%{random_suffix}"
  time_zone = "EST"
  location  = "global"

  description = "one of the nine"
}
`, context)
}

func testAccCheckGameServicesRealmDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_game_services_realm" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("GameServicesRealm still exists at %s", url)
			}
		}

		return nil
	}
}
