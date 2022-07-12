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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataprocMetastoreServiceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreServiceIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccDataprocMetastoreServiceIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccDataprocMetastoreServiceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataprocMetastoreServiceIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccDataprocMetastoreServiceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreServiceIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccDataprocMetastoreServiceIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccDataprocMetastoreServiceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_member" "foo" {
  provider = google-beta
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataprocMetastoreServiceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataproc_metastore_service_iam_policy" "foo" {
  provider = google-beta
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocMetastoreServiceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_dataproc_metastore_service_iam_policy" "foo" {
  provider = google-beta
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataprocMetastoreServiceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_binding" "foo" {
  provider = google-beta
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataprocMetastoreServiceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}

resource "google_dataproc_metastore_service_iam_binding" "foo" {
  provider = google-beta
  project = google_dataproc_metastore_service.default.project
  location = google_dataproc_metastore_service.default.location
  service_id = google_dataproc_metastore_service.default.service_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
