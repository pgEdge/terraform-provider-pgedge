package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatabasesResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "pgedge_databases" "tech" {
				  databases = {
					name       = "newdatabase101",
				  }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pgedge_databases.tech", "databases.name", "newdatabase101"),
					resource.TestCheckResourceAttrSet(
						"pgedge_databases.tech", "databases.id"),
            resource.TestCheckResourceAttrSet(
              "pgedge_databases.tech", "databases.domain"),
					resource.TestCheckResourceAttrSet(
						"pgedge_databases.tech", "databases.created_at"),
					resource.TestCheckResourceAttrSet(
						"pgedge_databases.tech", "databases.updated_at"),
				),
			},
			
			// Delete testing automatically occurs in TestCase
		},
	})
}
