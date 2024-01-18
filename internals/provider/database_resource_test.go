package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatabaseResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "pgedge_database" "tech" {
				  database = {
					name       = "newdatabase101",
					cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
				  }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pgedge_database.tech", "database.name", "newdatabase101"),
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
