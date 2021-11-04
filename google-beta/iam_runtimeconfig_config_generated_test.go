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

func TestAccRuntimeConfigConfigIamBindingGenerated(t *testing.T) {
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
				Config: testAccRuntimeConfigConfigIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccRuntimeConfigConfigIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccRuntimeConfigConfigIamMemberGenerated(t *testing.T) {
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
				Config: testAccRuntimeConfigConfigIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccRuntimeConfigConfigIamPolicyGenerated(t *testing.T) {
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
				Config: testAccRuntimeConfigConfigIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccRuntimeConfigConfigIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccRuntimeConfigConfigIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_member" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccRuntimeConfigConfigIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_runtimeconfig_config_iam_policy" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccRuntimeConfigConfigIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_runtimeconfig_config_iam_policy" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccRuntimeConfigConfigIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_binding" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccRuntimeConfigConfigIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_runtimeconfig_config" "config" {
  provider    = google-beta
  name        = "tf-test-my-config%{random_suffix}"
  description = "Runtime configuration values for my service"
}

resource "google_runtimeconfig_config_iam_binding" "foo" {
  provider = google-beta
  project = google_runtimeconfig_config.config.project
  config = google_runtimeconfig_config.config.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
