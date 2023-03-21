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

func TestAccFirebaseStorageBucket_firebasestorageBucketBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseStorageBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseStorageBucket_firebasestorageBucketBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_storage_bucket.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bucket_id"},
			},
		},
	})
}

func testAccFirebaseStorageBucket_firebasestorageBucketBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "default" {
  provider                    = google-beta
  name                        = "tf_test_test_bucket%{random_suffix}"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_firebase_storage_bucket" "default" {
  provider  = google-beta
  project   = "%{project_id}"
  bucket_id = google_storage_bucket.default.id
}
`, context)
}

func testAccCheckFirebaseStorageBucketDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_storage_bucket" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseStorageBasePath}}projects/{{project}}/buckets/{{bucket_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("FirebaseStorageBucket still exists at %s", url)
			}
		}

		return nil
	}
}
