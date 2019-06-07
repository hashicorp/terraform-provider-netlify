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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHookExists(resourceName, &hook),
				),
			},

			{
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
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

func TestAccHook_updateData(t *testing.T) {
	var hook models.Hook
	resourceName := "netlify_hook.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHookExists(resourceName, &hook),
					testAccAssert("has default ur", func() bool {
						m := hook.Data.(map[string]interface{})
						return m["url"] == "http://www.example.com"
					}),
				),
			},

			{
				Config: testAccHookConfig_updateData,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHookExists(resourceName, &hook),
					testAccAssert("has changed url", func() bool {
						m := hook.Data.(map[string]interface{})
						return m["url"] == "http://www.example.com/tubes"
					}),
				),
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

func testAccCheckHookDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netlify_hook" {
			continue
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetHookParams()
		params.HookID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetHook(params, meta.AuthInfo)
		if err == nil && resp.Payload != nil {
			return fmt.Errorf("Hook still exists: %s", rs.Primary.ID)
		}

		if err != nil {
			if v, ok := err.(*operations.GetHookDefault); ok && v.Code() == 404 {
				return nil
			}
		}

		return err
	}

	return nil
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
var testAccHookConfig_updateData = `
resource "netlify_site" "test" {}

resource "netlify_hook" "test" {
	site_id = "${netlify_site.test.id}"
	type  = "url"
	event = "deploy_locked"
	data  = {
		url = "http://www.example.com/tubes"
	}
}
`
