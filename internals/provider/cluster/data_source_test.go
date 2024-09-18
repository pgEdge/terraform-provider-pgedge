package cluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

func TestAccClustersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
					data "pgedge_clusters" "tech" {}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pgedge_clusters.tech", "clusters.#", "5"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.tech", "clusters.0.name", "test15"),
				),
			},
		},
	})
}
