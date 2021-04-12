package provider

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/accounts"
	"github.com/hashicorp/boundary/testing/controller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	fooAccountDesc       = "test account"
	fooAccountDescUpdate = "test account update"
)

var (
	fooAccount = fmt.Sprintf(`
resource "boundary_auth_method" "foo" {
	name        = "test"
	description = "test account"
	type        = "password"
	scope_id    = boundary_scope.org1.id
	depends_on = [boundary_role.org1_admin]
}

resource "boundary_account" "foo" {
	name           = "test"
	description    = "%s"
	type           = "password"
	login_name     = "foo"
	password       = "foofoofoo"
	auth_method_id = boundary_auth_method.foo.id
}`, fooAccountDesc)

	fooAccountUpdate = fmt.Sprintf(`
resource "boundary_auth_method" "foo" {
	name        = "test"
	description = "test account"
	type        = "password"
	scope_id    = boundary_scope.org1.id
	depends_on = [boundary_role.org1_admin]
}

resource "boundary_account" "foo" {
	name           = "test"
	description    = "%s"
	type           = "password"
	login_name     = "foo"
	password       = "foofoofoo"
	auth_method_id = boundary_auth_method.foo.id
}`, fooAccountDescUpdate)
)

var (
	fooAccountAttr = fmt.Sprintf(`
resource "boundary_auth_method" "foo" {
	name        = "test"
	description = "test account"
	type        = "password"
	scope_id    = boundary_scope.org1.id
	depends_on = [boundary_role.org1_admin]
}

resource "boundary_account" "foo" {
	name           = "test"
	description    = "%s"
	type           = "password"
	auth_method_id = boundary_auth_method.foo.id

  attributes = {
    "login_name" = "foo"
    "password" = "foofoofoo"
	}
}`, fooAccountDesc)

	fooAccountAttrUpdate = fmt.Sprintf(`
resource "boundary_auth_method" "foo" {
	name        = "test"
	description = "test account"
	type        = "password"
	scope_id    = boundary_scope.org1.id
	depends_on = [boundary_role.org1_admin]
}

resource "boundary_account" "foo" {
	name           = "test"
	description    = "%s"
	type           = "password"
	auth_method_id = boundary_auth_method.foo.id

  attributes = {
    "login_name" = "fooUpdate"
    "password" = "foofoofooUpdate"
	}
}`, fooAccountDescUpdate)
)

func TestAccAccountBase(t *testing.T) {
	tc := controller.NewTestController(t, tcConfig...)
	defer tc.Shutdown()
	url := tc.ApiAddrs()[0]

	var provider *schema.Provider

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories(&provider),
		CheckDestroy:      testAccCheckAccountResourceDestroy(t, provider),
		Steps: []resource.TestStep{
			{
				// create
				Config: testConfig(url, fooOrg, fooAccount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("boundary_account.foo", "description", fooAccountDesc),
					resource.TestCheckResourceAttr("boundary_account.foo", "name", "test"),
					resource.TestCheckResourceAttr("boundary_account.foo", "type", "password"),
					resource.TestCheckResourceAttr("boundary_account.foo", "login_name", "foo"),
					resource.TestCheckResourceAttr("boundary_account.foo", "password", "foofoofoo"),
					testAccCheckAccountResourceExists(provider, "boundary_account.foo"),
				),
			},
			importStep("boundary_account.foo", "password"),
			{
				// update
				Config: testConfig(url, fooOrg, fooAccountUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("boundary_account.foo", "description", fooAccountDescUpdate),
					resource.TestCheckResourceAttr("boundary_account.foo", "name", "test"),
					resource.TestCheckResourceAttr("boundary_account.foo", "type", "password"),
					resource.TestCheckResourceAttr("boundary_account.foo", "login_name", "foo"),
					resource.TestCheckResourceAttr("boundary_account.foo", "password", "foofoofoo"),
					testAccCheckAccountResourceExists(provider, "boundary_account.foo"),
				),
			},
			importStep("boundary_account.foo", "password"),
		},
	})
}

func TestAccAccountBaseAttr(t *testing.T) {
	tc := controller.NewTestController(t, tcConfig...)
	defer tc.Shutdown()
	url := tc.ApiAddrs()[0]

	var provider *schema.Provider

	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories(&provider),
		CheckDestroy:      testAccCheckAccountResourceDestroy(t, provider),
		Steps: []resource.TestStep{
			{
				// create
				Config: testConfig(url, fooOrg, fooAccountAttr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("boundary_account.foo", "description", fooAccountDesc),
					resource.TestCheckResourceAttr("boundary_account.foo", "name", "test"),
					resource.TestCheckResourceAttr("boundary_account.foo", "type", "password"),
					testAccCheckAccountAttrSet(
						provider,
						"boundary_account.foo",
						map[string]interface{}{
							"login_name": "foo",
							"passwrod":   "foofoofoo"}),

					testAccCheckAccountResourceExists(provider, "boundary_account.foo"),
				),
			},
			importStep("boundary_account.foo", "password"),
			{
				// update
				Config: testConfig(url, fooOrg, fooAccountAttrUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("boundary_account.foo", "description", fooAccountDescUpdate),
					resource.TestCheckResourceAttr("boundary_account.foo", "name", "test"),
					resource.TestCheckResourceAttr("boundary_account.foo", "type", "password"),
					testAccCheckAccountAttrSet(
						provider,
						"boundary_account.foo",
						map[string]interface{}{
							"login_name": "fooUpdate",
							"passwrod":   "foofoofooUpdate"}),

					testAccCheckAccountResourceExists(provider, "boundary_account.foo"),
				),
			},
			importStep("boundary_account.foo", "password"),
		},
	})
}
func testAccCheckAccountResourceExists(testProvider *schema.Provider, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("No ID is set")
		}

		md := testProvider.Meta().(*metaData)

		amClient := accounts.NewClient(md.client)

		if _, err := amClient.Read(context.Background(), id); err != nil {
			return fmt.Errorf("Got an error when reading account %q: %v", id, err)
		}

		return nil
	}
}

func testAccCheckAccountResourceDestroy(t *testing.T, testProvider *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if testProvider.Meta() == nil {
			t.Fatal("got nil provider metadata")
		}
		md := testProvider.Meta().(*metaData)

		for _, rs := range s.RootModule().Resources {
			switch rs.Type {
			case "boundary_account":
				id := rs.Primary.ID

				amClient := accounts.NewClient(md.client)

				_, err := amClient.Read(context.Background(), id)
				if apiErr := api.AsServerError(err); apiErr == nil || apiErr.Response().StatusCode() != http.StatusNotFound {
					return fmt.Errorf("didn't get a 404 when reading destroyed account %q: %v", id, err)
				}

			default:
				continue
			}
		}
		return nil
	}
}

func testAccCheckAccountAttrSet(testProvider *schema.Provider, resourceName string, attrs map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("account resource not found: %s", resourceName)
		}

		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("account resource ID is not set")
		}

		md := testProvider.Meta().(*metaData)
		aClient := accounts.NewClient(md.client)

		ar, err := aClient.Read(context.Background(), id)
		if err != nil {
			return fmt.Errorf("Got an error when reading account %q: %v", id, err)
		}

		if fmt.Sprint(ar.Item.Attributes) == fmt.Sprint(attrs) {
			return nil
		}

		return fmt.Errorf("attrs not equal, got %+v, want %+v\n", ar.Item.Attributes, attrs)
	}
}
