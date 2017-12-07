package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func TestAccSite_basic(t *testing.T) {
	var site models.Site

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSiteConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSiteExists("netlify_site.test", &site),
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

var testAccSiteConfig = `
resource "netlify_site" "test" {}
`
