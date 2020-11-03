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

func TestAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(context),
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-api-cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}
`, context)
}

func TestAccApiGatewayApiConfig_apigatewayApiConfigFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigFullExample(context),
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-api-cfg%{random_suffix}"
  display_name = "MM Dev API Config"
  labels = {
    environment = "dev"
  }

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}
`, context)
}

func testAccCheckApiGatewayApiConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_api_gateway_api_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ApiGatewayBasePath}}projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ApiGatewayApiConfig still exists at %s", url)
			}
		}

		return nil
	}
}
