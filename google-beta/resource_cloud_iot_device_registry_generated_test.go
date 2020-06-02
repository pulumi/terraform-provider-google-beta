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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudIotDeviceRegistry_cloudiotDeviceRegistryBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIotDeviceRegistryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDeviceRegistry_cloudiotDeviceRegistryBasicExample(context),
			},
			{
				ResourceName:            "google_cloudiot_registry.test-registry",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudIotDeviceRegistry_cloudiotDeviceRegistryBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}
`, context)
}

func TestAccCloudIotDeviceRegistry_cloudiotDeviceRegistrySingleEventNotificationConfigsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIotDeviceRegistryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDeviceRegistry_cloudiotDeviceRegistrySingleEventNotificationConfigsExample(context),
			},
			{
				ResourceName:            "google_cloudiot_registry.test-registry",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudIotDeviceRegistry_cloudiotDeviceRegistrySingleEventNotificationConfigsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "default-telemetry" {
  name = "tf-test-default-telemetry%{random_suffix}"
}

resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"

  event_notification_configs {
    pubsub_topic_name = google_pubsub_topic.default-telemetry.id
    subfolder_matches = ""
  }

}
`, context)
}

func TestAccCloudIotDeviceRegistry_cloudiotDeviceRegistryFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudIotDeviceRegistryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDeviceRegistry_cloudiotDeviceRegistryFullExample(context),
			},
			{
				ResourceName:            "google_cloudiot_registry.test-registry",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccCloudIotDeviceRegistry_cloudiotDeviceRegistryFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_pubsub_topic" "default-devicestatus" {
  name = "tf-test-default-devicestatus%{random_suffix}"
}

resource "google_pubsub_topic" "default-telemetry" {
  name = "tf-test-default-telemetry%{random_suffix}"
}

resource "google_pubsub_topic" "additional-telemetry" {
  name = "tf-test-additional-telemetry%{random_suffix}"
}

resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"

  event_notification_configs {
    pubsub_topic_name = google_pubsub_topic.additional-telemetry.id
    subfolder_matches = "test/path%{random_suffix}"
  }

  event_notification_configs {
    pubsub_topic_name = google_pubsub_topic.default-telemetry.id
    subfolder_matches = ""
  }

  state_notification_config = {
    pubsub_topic_name = google_pubsub_topic.default-devicestatus.id
  }

  mqtt_config = {
    mqtt_enabled_state = "MQTT_ENABLED"
  }

  http_config = {
    http_enabled_state = "HTTP_ENABLED"
  }

  log_level = "INFO"

  credentials {
    public_key_certificate = {
      format      = "X509_CERTIFICATE_PEM"
      certificate = file("test-fixtures/rsa_cert.pem")
    }
  }
}
`, context)
}

func testAccCheckCloudIotDeviceRegistryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudiot_registry" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudIotBasePath}}projects/{{project}}/locations/{{region}}/registries/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("CloudIotDeviceRegistry still exists at %s", url)
			}
		}

		return nil
	}
}
