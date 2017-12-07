package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testRepo string = "test-repo"

func TestAccDeployKey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDeployKeyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeployKeyExists("netlify_deploy_key.test"),
				),
			},
		},
	})
}

func testAccCheckDeployKeyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		return nil
	}
}

var testAccDeployKeyConfig = `resource "netlify_deploy_key" "test" {}`
