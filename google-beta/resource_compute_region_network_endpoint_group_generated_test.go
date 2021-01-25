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

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"zip_path":      createZIPArchiveForCloudFunctionSource(t, testHTTPTriggerPath),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.function_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(context map[string]interface{}) string {
	return Nprintf(`
// Cloud Functions Example
resource "google_compute_region_network_endpoint_group" "function_neg" {
  name                  = "tf-test-function-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_function {
    function = google_cloudfunctions_function.function_neg.name
  }
}

resource "google_cloudfunctions_function" "function_neg" {
  name        = "tf-test-function-neg%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = true
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_storage_bucket" "bucket" {
  name       = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
}

resource "google_storage_bucket_object" "archive" { 
  name       = "index.zip"
  bucket     = google_storage_bucket.bucket.name
  source     = "%{zip_path}"
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.cloudrun_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(context map[string]interface{}) string {
	return Nprintf(`
// Cloud Run Example
resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  name                  = "tf-test-cloudrun-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_run {
    service = google_cloud_run_service.cloudrun_neg.name
  }
}

resource "google_cloud_run_service" "cloudrun_neg" {
  name     = "tf-test-cloudrun-neg%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.appengine_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(context map[string]interface{}) string {
	return Nprintf(`
// App Engine Example
resource "google_compute_region_network_endpoint_group" "appengine_neg" {
  name                  = "tf-test-appengine-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  app_engine {
    service = google_app_engine_flexible_app_version.appengine_neg.service
    version = google_app_engine_flexible_app_version.appengine_neg.version_id
  }
}

resource "google_app_engine_flexible_app_version" "appengine_neg" {
  version_id = "v1"
  service    = "default"
  runtime    = "nodejs"

  entrypoint {
    shell = "node ./app.js"
  }

  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.appengine_neg.name}/${google_storage_bucket_object.appengine_neg.name}"
    }
  }

  liveness_check {
    path = "/"
  }

  readiness_check {
    path = "/"
  }

  env_variables = {
    port = "8080"
  }

  handlers {
    url_regex        = ".*\\/my-path\\/*"
    security_level   = "SECURE_ALWAYS"
    login            = "LOGIN_REQUIRED"
    auth_fail_action = "AUTH_FAIL_ACTION_REDIRECT"

    static_files {
      path = "my-other-path"
      upload_path_regex = ".*\\/my-path\\/*"
    }
  }

  automatic_scaling {
    cool_down_period = "120s"
    cpu_utilization {
      target_utilization = 0.5
    }
  }

  noop_on_destroy = true
}

resource "google_storage_bucket" "appengine_neg" {
  name       = "tf-test-appengine-neg%{random_suffix}"
}

resource "google_storage_bucket_object" "appengine_neg" {
  name      = "hello-world.zip"
  bucket    = google_storage_bucket.appengine_neg.name
  source    = "./test-fixtures/appengine/hello-world.zip"
}
`, context)
}

func testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_network_endpoint_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionNetworkEndpointGroup still exists at %s", url)
			}
		}

		return nil
	}
}
