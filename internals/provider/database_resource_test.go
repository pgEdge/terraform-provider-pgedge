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
