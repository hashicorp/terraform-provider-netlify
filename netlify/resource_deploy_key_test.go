package netlify

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func TestAccDeployKey_basic(t *testing.T) {
	var key models.DeployKey
	resourceName := "netlify_deploy_key.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeployKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDeployKeyConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeployKeyExists(resourceName, &key),
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

func TestAccDeployKey_disappears(t *testing.T) {
	var key models.DeployKey

	destroy := func(*terraform.State) error {
		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewDeleteDeployKeyParams()
		params.KeyID = key.ID
		_, err := meta.Netlify.Operations.DeleteDeployKey(params, meta.AuthInfo)
		return err
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeployKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDeployKeyConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeployKeyExists("netlify_deploy_key.test", &key),
					destroy,
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckDeployKeyExists(n string, key *models.DeployKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetDeployKeyParams()
		params.KeyID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetDeployKey(params, meta.AuthInfo)
		if err != nil {
			return err
		}

		*key = *resp.Payload
		return nil
	}
}

func testAccCheckDeployKeyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netlify_deploy_key" {
			continue
		}

		meta := testAccProvider.Meta().(*Meta)
		params := operations.NewGetDeployKeyParams()
		params.KeyID = rs.Primary.ID
		resp, err := meta.Netlify.Operations.GetDeployKey(params, meta.AuthInfo)
		if err == nil && resp.Payload != nil {
			return fmt.Errorf("Resource still exists: %s", rs.Primary.ID)
		}

		if err != nil {
			if v, ok := err.(*operations.GetDeployKeyDefault); ok && v.Code() == 404 {
				return nil
			}
		}

		return err
	}

	return nil
}

var testAccDeployKeyConfig = `resource "netlify_deploy_key" "test" {}`
