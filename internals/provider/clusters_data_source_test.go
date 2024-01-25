package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClustersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
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
