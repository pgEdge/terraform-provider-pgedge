package database_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccDatabaseResource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	ca := api.SeedCloudAccount(&models.CloudAccount{
		Name: strPtr("acceptance-account"), Type: strPtr("aws"),
	})
	bs := api.SeedBackupStore(&models.BackupStore{
		Name:           strPtr("acceptance-store"),
		CloudAccountID: strPtr(ca.ID.String()),
	})
	cluster := api.SeedCluster(&models.Cluster{
		Name:         strPtr("acceptance-cluster"),
		Regions:      []string{"us-west-2"},
		NodeLocation: strPtr("public"),
		CloudAccount: &models.CloudAccountProperties{ID: strPtr(ca.ID.String())},
		Nodes: []*models.ClusterNodeSettings{
			{Name: "n1", Region: strPtr("us-west-2"), InstanceType: "t4g.small"},
		},
	})

	config := common.ProviderConfig + fmt.Sprintf(`
		resource "pgedge_database" "tech" {
			name       = "acmedb"
			cluster_id = %q
			options    = []

			nodes = {
				n1 = {
					name = "n1"
				}
			}

			backups = {
				provider = "pgbackrest"
				config = [
					{
						id        = "default"
						schedules = []
						repositories = [
							{
								id              = "repo1"
								type            = "s3"
								backup_store_id = %q
							}
						]
					}
				]
			}
		}
	`, cluster.ID.String(), bs.ID.String())

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_database.tech", "name", "acmedb"),
					resource.TestCheckResourceAttr("pgedge_database.tech", "cluster_id", cluster.ID.String()),
					resource.TestCheckResourceAttr("pgedge_database.tech", "status", "available"),
					resource.TestCheckResourceAttrSet("pgedge_database.tech", "id"),
					resource.TestCheckResourceAttrSet("pgedge_database.tech", "domain"),
					resource.TestCheckResourceAttrSet("pgedge_database.tech", "created_at"),
				),
			},
		},
	})
}
