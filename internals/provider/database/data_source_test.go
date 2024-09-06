package database_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

func TestAccDatabasesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
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
