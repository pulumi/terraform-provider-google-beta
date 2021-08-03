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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeDiskIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDiskIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_disk_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/disks/%s roles/viewer", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-test-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeDiskIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_disk_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/disks/%s roles/viewer", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-test-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeDiskIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeDiskIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_disk_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/disks/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-test-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeDiskIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeDiskIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_disk_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/disks/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-test-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeDiskIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_disk_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/zones/%s/disks/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-test-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeDiskIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-8-jessie-v20170523"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

resource "google_compute_disk_iam_member" "foo" {
  project = google_compute_disk.default.project
  zone = google_compute_disk.default.zone
  name = google_compute_disk.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeDiskIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-8-jessie-v20170523"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_disk_iam_policy" "foo" {
  project = google_compute_disk.default.project
  zone = google_compute_disk.default.zone
  name = google_compute_disk.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeDiskIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-8-jessie-v20170523"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

data "google_iam_policy" "foo" {
}

resource "google_compute_disk_iam_policy" "foo" {
  project = google_compute_disk.default.project
  zone = google_compute_disk.default.zone
  name = google_compute_disk.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeDiskIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-8-jessie-v20170523"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

resource "google_compute_disk_iam_binding" "foo" {
  project = google_compute_disk.default.project
  zone = google_compute_disk.default.zone
  name = google_compute_disk.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeDiskIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_disk" "default" {
  name  = "tf-test-test-disk%{random_suffix}"
  type  = "pd-ssd"
  zone  = "us-central1-a"
  image = "debian-8-jessie-v20170523"
  labels = {
    environment = "dev"
  }
  physical_block_size_bytes = 4096
}

resource "google_compute_disk_iam_binding" "foo" {
  project = google_compute_disk.default.project
  zone = google_compute_disk.default.zone
  name = google_compute_disk.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
