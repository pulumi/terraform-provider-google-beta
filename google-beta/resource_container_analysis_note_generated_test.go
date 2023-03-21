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

func TestAccContainerAnalysisNote_containerAnalysisNoteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNote_containerAnalysisNoteBasicExample(context),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNote_containerAnalysisNoteBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}
`, context)
}

func TestAccContainerAnalysisNote_containerAnalysisNoteAttestationFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckContainerAnalysisNoteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAnalysisNote_containerAnalysisNoteAttestationFullExample(context),
			},
			{
				ResourceName:      "google_container_analysis_note.note",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccContainerAnalysisNote_containerAnalysisNoteAttestationFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_analysis_note" "note" {
  name = "tf-test-attestor-note%{random_suffix}"

  short_description = "test note"
  long_description = "a longer description of test note"
  expiration_time = "2120-10-02T15:01:23.045123456Z"

  related_url {
    url = "some.url"
    label = "foo"
  }

  related_url {
    url = "google.com"
  }

  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}
`, context)
}

func testAccCheckContainerAnalysisNoteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_container_analysis_note" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ContainerAnalysisNote still exists at %s", url)
			}
		}

		return nil
	}
}
