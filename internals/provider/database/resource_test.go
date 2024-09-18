package database_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

func TestAccDatabaseResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
				resource "pgedge_database" "tech" {
					name       = "",
					cluster_id = ""
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pgedge_database.tech", "database.name", ""),
					resource.TestCheckResourceAttrSet(
						"pgedge_database.tech", "database.id"),
					resource.TestCheckResourceAttrSet(
						"pgedge_database.tech", "database.domain"),
					resource.TestCheckResourceAttrSet(
						"pgedge_database.tech", "database.created_at"),
					resource.TestCheckResourceAttrSet(
						"pgedge_database.tech", "database.updated_at"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
