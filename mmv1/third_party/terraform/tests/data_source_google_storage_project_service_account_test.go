package google_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleStorageProjectServiceAccount_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.google_storage_project_service_account.gcs_account"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:  func() { acctest.TestAccPreCheck(t) },
		Providers: acctest.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleStorageProjectServiceAccount_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "email_address"),
					resource.TestCheckResourceAttrSet(resourceName, "member"),
				),
			},
		},
	})
}

const testAccCheckGoogleStorageProjectServiceAccount_basic = `
data "google_storage_project_service_account" "gcs_account" {
}
`
