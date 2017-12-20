package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func TestAccBuildHook(t *testing.T) {
	var hook models.BuildHook
	resourceName := "netlify_build_hook.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBuildHookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBuildHookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBuildHookExists(resourceName, &hook),
				),
			},
		},
	})
}

func TestAccBuildHook_disappears(t *testing.T) {
	var hook models.BuildHook
	resourceName := "netlify_build_hook.test"

	destroy := func(*terraform.State) error {
		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewDeleteSiteBuildHookParams()
		params.ID = hook.ID
		params.SiteID = hook.SiteID
		_, err := meta.Netlify.Operations.DeleteSiteBuildHook(params, meta.AuthInfo)
		return err
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBuildHookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBuildHookConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckBuildHookExists(resourceName, &hook),
					destroy,
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccBuildHook_updateBranch(t *testing.T) {
	var hook models.BuildHook
	resourceName := "netlify_build_hook.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBuildHookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBuildHookConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBuildHookExists(resourceName, &hook),
					testAccAssert("has default ur", func() bool {
						return hook.Branch == "master"
					}),
				),
			},

			resource.TestStep{
				Config: testAccBuildHookConfig_updateBranch,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBuildHookExists(resourceName, &hook),
					testAccAssert("has changed url", func() bool {
						return hook.Branch == "changed"
					}),
				),
			},
		},
	})
}

func testAccCheckBuildHookExists(n string, hook *models.BuildHook) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetSiteBuildHookParams()
		params.ID = rs.Primary.ID
		params.SiteID = rs.Primary.Attributes["site_id"]
		resp, err := meta.Netlify.Operations.GetSiteBuildHook(params, meta.AuthInfo)
		if err != nil {
			return err
		}

		*hook = *resp.Payload
		return nil
	}
}

func testAccCheckBuildHookDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netlify_build_hook" {
			continue
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetSiteBuildHookParams()
		params.ID = rs.Primary.ID
		params.SiteID = rs.Primary.Attributes["site_id"]
		resp, err := meta.Netlify.Operations.GetSiteBuildHook(params, meta.AuthInfo)
		if err == nil && resp.Payload != nil {
			return fmt.Errorf("BuildHook still exists: %s", rs.Primary.ID)
		}

		if err != nil {
			if v, ok := err.(*operations.GetSiteBuildHookDefault); ok && v.Code() == 404 {
				return nil
			}
		}

		return err
	}

	return nil
}

var testAccBuildHookConfig = `
resource "netlify_site" "test" {}

resource "netlify_build_hook" "test" {
	site_id = "${netlify_site.test.id}"
	branch = "master"
	title = "tubes"
}
`
var testAccBuildHookConfig_updateBranch = `
resource "netlify_site" "test" {}

resource "netlify_build_hook" "test" {
	site_id = "${netlify_site.test.id}"
	branch = "changed"
	title = "tubes"
}
`
