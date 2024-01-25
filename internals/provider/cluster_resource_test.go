package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClusterResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				resource "pgedge_cluster" "tech" {
					name       = ""
					cloud_account_id = ""
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pgedge_cluster.tech", "name", "test123"),
					resource.TestCheckResourceAttrSet(
						"pgedge_cluster.tech", "id"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
