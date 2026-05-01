package database_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccDatabasesDataSource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	ca := api.SeedCloudAccount(&models.CloudAccount{
		Name: strPtr("test-account"), Type: strPtr("aws"),
	})
	cluster := api.SeedCluster(&models.Cluster{
		Name:         strPtr("c"),
		Regions:      []string{"us-west-2"},
		NodeLocation: strPtr("public"),
		CloudAccount: &models.CloudAccountProperties{ID: strPtr(ca.ID.String())},
	})
	api.SeedDatabase(&models.Database{
		Name:      strPtr("acme"),
		ClusterID: cluster.ID,
		CreatedAt: strPtr("2025-01-01T00:00:00Z"),
		UpdatedAt: strPtr("2025-01-01T00:00:00Z"),
	})
	api.SeedDatabase(&models.Database{
		Name:      strPtr("widgets"),
		ClusterID: cluster.ID,
		CreatedAt: strPtr("2025-02-01T00:00:00Z"),
		UpdatedAt: strPtr("2025-02-01T00:00:00Z"),
	})

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
					data "pgedge_databases" "tech" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.#", "2"),
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.0.name", "acme"),
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.0.status", "available"),
					resource.TestCheckResourceAttr("data.pgedge_databases.tech", "databases.1.name", "widgets"),
				),
			},
		},
	})
}

func strPtr(s string) *string { return &s }
