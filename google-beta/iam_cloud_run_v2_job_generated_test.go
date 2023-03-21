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

func TestAccCloudRunV2JobIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2JobIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_job_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudRunV2JobIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_job_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV2JobIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudRunV2JobIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_job_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer user:admin@hashicorptest.com", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV2JobIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV2JobIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_job_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudRunV2JobIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloud_run_v2_job_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s", GetTestProjectFromEnv(), GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudRunV2JobIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v2_job_iam_member" "foo" {
  project = google_cloud_run_v2_job.default.project
  location = google_cloud_run_v2_job.default.location
  name = google_cloud_run_v2_job.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudRunV2JobIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloud_run_v2_job_iam_policy" "foo" {
  project = google_cloud_run_v2_job.default.project
  location = google_cloud_run_v2_job.default.location
  name = google_cloud_run_v2_job.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudRunV2JobIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_cloud_run_v2_job_iam_policy" "foo" {
  project = google_cloud_run_v2_job.default.project
  location = google_cloud_run_v2_job.default.location
  name = google_cloud_run_v2_job.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudRunV2JobIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v2_job_iam_binding" "foo" {
  project = google_cloud_run_v2_job.default.project
  location = google_cloud_run_v2_job.default.location
  name = google_cloud_run_v2_job.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudRunV2JobIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"
  launch_stage = "BETA"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v2_job_iam_binding" "foo" {
  project = google_cloud_run_v2_job.default.project
  location = google_cloud_run_v2_job.default.location
  name = google_cloud_run_v2_job.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
