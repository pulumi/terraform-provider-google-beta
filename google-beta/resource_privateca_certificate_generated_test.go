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

func TestAccPrivatecaCertificate_privatecaCertificateConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateConfigExample(context),
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateConfigExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "test-ca" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}


resource "google_privateca_certificate" "default" {
  provider = google-beta
  project = "%{project}"
  location = "us-central1"
  certificate_authority = google_privateca_certificate_authority.test-ca.certificate_authority_id
  lifetime = "860s"
  name = "tf-test-my-certificate%{random_suffix}"
  config {
      reusable_config {
        reusable_config= "projects/568668481468/locations/us-central1/reusableConfigs/leaf-server-tls"
      } 
      subject_config  {
        common_name = "san1.example.com"
        subject {
          country_code = "us"
          organization = "google"
          organizational_unit = "enterprise"
          locality = "mountain view"
          province = "california"
          street_address = "1600 amphitheatre parkway"
          postal_code = "94109"
        } 
        subject_alt_name {
          dns_names = ["hashicorp.com"]
          email_addresses = ["email@example.com"]
          ip_addresses = ["127.0.0.1"]
          uris = ["http://www.ietf.org/rfc/rfc3986.txt"]
        }
      }

    public_key {
      type = "PEM_RSA_KEY"
      key = filebase64("test-fixtures/rsa_public.pem")
    }    
  }
}
`, context)
}

func TestAccPrivatecaCertificate_privatecaCertificateCsrExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateCsrExample(context),
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateCsrExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_authority" "test-ca" {
  provider = google-beta
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "us-central1"
  tier = "ENTERPRISE"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
      }
      common_name = "my-certificate-authority"
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    reusable_config {
      reusable_config = "projects/568668481468/locations/us-central1/reusableConfigs/root-unconstrained"
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  disable_on_delete = true
}




resource "google_privateca_certificate" "default" {
  provider = google-beta
  project = "%{project}"
  location = "us-central1"
  certificate_authority = google_privateca_certificate_authority.test-ca.certificate_authority_id
  lifetime = "860s"
  name = "tf-test-my-certificate%{random_suffix}"
  pem_csr = file("test-fixtures/rsa_csr.pem")
}
`, context)
}

func testAccCheckPrivatecaCertificateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_certificate" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/certificateAuthorities/{{certificate_authority}}/certificates/{{name}}")

			if err != nil {
				return err
			}

			res, err := sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err != nil {
				return err
			}

			if _, ok := res["revocationDetails"]; !ok {
				return fmt.Errorf("CertificateAuthority.Certificate Revocation expected %s got %s, want revocationDetails.revocationTime", url, s)
			}

			return nil
		}

		return nil
	}
}
