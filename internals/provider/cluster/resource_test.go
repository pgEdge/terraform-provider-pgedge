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
				resource "pgedge_cluster" "test" {
					name             = "test-cluster"
					cloud_account_id = "dc8wewec1a-sddffsdf-wsdad-23qqwe-wsadmn8123a"
					regions          = ["us-west-1", "us-west-2"]
					node_location    = "public"
					nodes = [
						{
							name              = "n1"
							region            = "us-west-1"
							image             = "postgres"
							instance_type     = "t4g.small"
							availability_zone = "us-west-1a"
							volume_size       = 20
							volume_type       = "gp2"
						},
						{
							name              = "n2"
							region            = "us-west-2"
							image             = "postgres"
							instance_type     = "t4g.medium"
							availability_zone = "us-west-2a"
							volume_size       = 20
							volume_type       = "gp2"
						}
					]
					networks = [
						{
							region         = "us-west-1"
							cidr           = "10.1.0.0/16"
							public_subnets = ["10.1.1.0/24"]
						},
						{
							region         = "us-west-2"
							cidr           = "10.2.0.0/16"
							public_subnets = ["10.2.1.0/24"]
						}
					]
					firewall_rules = [
						{
							name    = "postgres"
							port    = 5432
							sources = ["0.0.0.0/0"]
						}
					]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_cluster.test", "name", "test-cluster"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "cloud_account_id", "dc8wewec1a-sddffsdf-wsdad-23qqwe-wsadmn8123a"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.0", "us-west-1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.1", "us-west-2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "node_location", "public"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.0.name", "n1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.0.region", "us-west-1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.1.name", "n2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.1.region", "us-west-2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "networks.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "networks.0.region", "us-west-1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "networks.1.region", "us-west-2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "firewall_rules.#", "1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "firewall_rules.0.name", "postgres"),
					resource.TestCheckResourceAttrSet("pgedge_cluster.test", "id"),
					resource.TestCheckResourceAttrSet("pgedge_cluster.test", "status"),
					resource.TestCheckResourceAttrSet("pgedge_cluster.test", "created_at"),
				),
			},
		},
	})
}