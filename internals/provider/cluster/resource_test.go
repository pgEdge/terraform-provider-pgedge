package cluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

func TestAccClusterResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
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
