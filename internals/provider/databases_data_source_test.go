package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatabasesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					data "pgedge_databases" "tech" {}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.#", "47"),
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.0.name", "newdatabase101"),
				),
			},
		},
	})
}
