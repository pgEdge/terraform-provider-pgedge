package cloudaccount_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccCloudAccountsDataSource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	api.SeedCloudAccount(&models.CloudAccount{
		Name:        strPtr("acme-aws"),
		Type:        strPtr("aws"),
		Description: "primary aws account",
		CreatedAt:   strPtr("2025-01-01T00:00:00Z"),
		UpdatedAt:   strPtr("2025-01-01T00:00:00Z"),
	})
	api.SeedCloudAccount(&models.CloudAccount{
		Name:        strPtr("widgets-gcp"),
		Type:        strPtr("gcp"),
		Description: "gcp account",
		CreatedAt:   strPtr("2025-02-01T00:00:00Z"),
		UpdatedAt:   strPtr("2025-02-01T00:00:00Z"),
	})

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
				data "pgedge_cloud_accounts" "all" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.#", "2"),
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.0.name", "acme-aws"),
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.0.type", "aws"),
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.1.name", "widgets-gcp"),
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.1.type", "gcp"),
				),
			},
		},
	})
}

func strPtr(s string) *string { return &s }
