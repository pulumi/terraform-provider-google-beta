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

func TestAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckStorageDefaultObjectAccessControlDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(context),
			},
			{
				ResourceName:            "google_storage_default_object_access_control.public_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bucket"},
			},
		},
	})
}

func testAccStorageDefaultObjectAccessControl_storageDefaultObjectAccessControlPublicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_default_object_access_control" "public_rule" {
  bucket = google_storage_bucket.bucket.name
  role   = "READER"
  entity = "allUsers"
}

resource "google_storage_bucket" "bucket" {
  name = "tf-test-static-content-bucket%{random_suffix}"
}
`, context)
}

func testAccCheckStorageDefaultObjectAccessControlDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_storage_default_object_access_control" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{StorageBasePath}}b/{{bucket}}/defaultObjectAcl/{{entity}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("StorageDefaultObjectAccessControl still exists at %s", url)
			}
		}

		return nil
	}
}
