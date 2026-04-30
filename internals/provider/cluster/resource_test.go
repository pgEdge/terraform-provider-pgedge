package cluster_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccClusterResource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	// The cluster resource looks up its cloud account by ID at create time, so
	// seed one and inject its UUID into the HCL.
	ca := api.SeedCloudAccount(&models.CloudAccount{
		Name: strPtr("acceptance-account"),
		Type: strPtr("aws"),
	})

	config := common.ProviderConfig + fmt.Sprintf(`
		resource "pgedge_cluster" "test" {
			name             = "test-cluster"
			cloud_account_id = %q
			regions          = ["us-west-1", "us-west-2"]
			node_location    = "public"
			nodes = [
				{
					name              = "n1"
					region            = "us-west-1"
					instance_type     = "t4g.small"
					availability_zone = "us-west-1a"
					volume_size       = 20
					volume_type       = "gp2"
				},
				{
					name              = "n2"
					region            = "us-west-2"
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
	`, ca.ID.String())

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_cluster.test", "name", "test-cluster"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "cloud_account_id", ca.ID.String()),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.0", "us-west-1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "regions.1", "us-west-2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "node_location", "public"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.0.name", "n1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "nodes.1.name", "n2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "networks.#", "2"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "firewall_rules.#", "1"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "firewall_rules.0.name", "postgres"),
					resource.TestCheckResourceAttr("pgedge_cluster.test", "status", "available"),
					resource.TestCheckResourceAttrSet("pgedge_cluster.test", "id"),
					resource.TestCheckResourceAttrSet("pgedge_cluster.test", "created_at"),
				),
			},
		},
	})
}
