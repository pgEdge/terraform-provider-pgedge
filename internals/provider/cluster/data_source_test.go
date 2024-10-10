package cluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
)

func TestAccClustersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
					data "pgedge_clusters" "test" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check that we have at least one cluster
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.#"),
					// Check the first cluster's attributes
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.id"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.name"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.cloud_account_id"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.created_at"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.status"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.regions.#"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.node_location"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.capacity"),
					// Check for nested attributes
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.firewall_rules.#"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.nodes.#"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.networks.#"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.backup_store_ids.#"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.resource_tags.%"),
				),
			},
		},
	})
}
