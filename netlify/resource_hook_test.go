package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func TestAccHook(t *testing.T) {
	var hook models.Hook
	resourceName := "netlify_hook.test"

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     testAccProviders,
		IDRefreshName: resourceName,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccHookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHookExists(resourceName, &hook),
				),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccHook_disappears(t *testing.T) {
	var hook models.Hook
	resourceName := "netlify_hook.test"

	destroy := func(*terraform.State) error {
		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewDeleteHookBySiteIDParams()
		params.HookID = hook.ID
		_, err := meta.Netlify.Operations.DeleteHookBySiteID(params, meta.AuthInfo)
		return err
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccHookConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHookExists(resourceName, &hook),
					destroy,
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckHookExists(n string, hook *models.Hook) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetHookParams()
		params.HookID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetHook(params, meta.AuthInfo)
		if err != nil {
			return err
		}

		*hook = *resp.Payload
		return nil
	}
}

var testAccHookConfig = `
resource "netlify_site" "test" {}

resource "netlify_hook" "test" {
	site_id = "${netlify_site.test.id}"
	type  = "url"
	event = "deploy_locked"
	data  = {
		url = "http://www.example.com"
	}
}
`
