package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func TestAccSite_basic(t *testing.T) {
	var site models.Site
	resourceName := "netlify_site.test"

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckSiteDestroy,
		IDRefreshName: resourceName,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSiteConfig_repo,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSiteExists(resourceName, &site),
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

func TestAccSite_disappears(t *testing.T) {
	var site models.Site

	destroy := func(*terraform.State) error {
		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewDeleteSiteParams()
		params.SiteID = site.ID
		_, err := meta.Netlify.Operations.DeleteSite(params, meta.AuthInfo)
		return err
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSiteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSiteConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSiteExists("netlify_site.test", &site),
					destroy,
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSite_updateName(t *testing.T) {
	var site models.Site
	resourceName := "netlify_site.test"
	siteName := fmt.Sprintf("test-%s", acctest.RandStringFromCharSet(
		5, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSiteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSiteConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSiteExists(resourceName, &site),
					testAccAssert("has random name", func() bool {
						return site.Name != "" && site.Name != "tubes"
					}),
				),
			},

			resource.TestStep{
				Config: fmt.Sprintf(testAccSiteConfig_updateName, siteName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSiteExists(resourceName, &site),
					testAccAssert("has configured name", func() bool {
						return site.Name == siteName
					}),
				),
			},
		},
	})
}

func testAccCheckSiteExists(n string, site *models.Site) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetSiteParams()
		params.SiteID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetSite(params, meta.AuthInfo)
		if err != nil {
			return err
		}

		*site = *resp.Payload
		return nil
	}
}

func testAccCheckSiteDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netlify_site" {
			continue
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetSiteParams()
		params.SiteID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetSite(params, meta.AuthInfo)
		if err == nil && resp.Payload != nil {
			return fmt.Errorf("Site still exists: %s", rs.Primary.ID)
		}

		if err != nil {
			if v, ok := err.(*operations.GetSiteDefault); ok && v.Code() == 404 {
				return nil
			}
		}

		return err
	}

	return nil
}

var testAccSiteConfig = `
resource "netlify_site" "test" {}
`

var testAccSiteConfig_repo = `
resource "netlify_site" "test" {
	repo {
		provider = "github"
		repo_path = "mitchellh/fogli"
		repo_branch = "master"
	}
}
`

var testAccSiteConfig_updateName = `
resource "netlify_site" "test" {
	name = "%s"
}
`
