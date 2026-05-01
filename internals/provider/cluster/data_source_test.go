package cluster_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccClustersDataSource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	ca := api.SeedCloudAccount(&models.CloudAccount{
		Name: strPtr("test-account"),
		Type: strPtr("aws"),
	})
	api.SeedCluster(&models.Cluster{
		Name:         strPtr("seeded-cluster"),
		Regions:      []string{"us-west-2"},
		NodeLocation: strPtr("public"),
		CloudAccount: &models.CloudAccountProperties{
			ID: strPtr(ca.ID.String()), Name: "test-account", Type: "aws",
		},
		Nodes: []*models.ClusterNodeSettings{
			{Name: "n1", Region: strPtr("us-west-2"), InstanceType: "t4g.small"},
		},
		Networks: []*models.ClusterNetworkSettings{
			{Region: strPtr("us-west-2"), Cidr: "10.1.0.0/16", PublicSubnets: []string{"10.1.1.0/24"}},
		},
		FirewallRules: []*models.ClusterFirewallRuleSettings{
			{Name: "postgres", Port: int64Ptr(5432), Sources: []string{"0.0.0.0/0"}},
		},
		BackupStoreIds: []string{},
		ResourceTags:   map[string]string{},
	})

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
					data "pgedge_clusters" "test" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.#", "1"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.name", "seeded-cluster"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.status", "available"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.node_location", "public"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.regions.#", "1"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.regions.0", "us-west-2"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.nodes.#", "1"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.networks.#", "1"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.firewall_rules.#", "1"),
					resource.TestCheckResourceAttr("data.pgedge_clusters.test", "clusters.0.cloud_account_id", ca.ID.String()),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.id"),
					resource.TestCheckResourceAttrSet("data.pgedge_clusters.test", "clusters.0.created_at"),
				),
			},
		},
	})
}

func strPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64 { return &i }
