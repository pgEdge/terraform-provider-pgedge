package cloudaccount_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

func TestAccCloudAccountsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() {},
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
				data "pgedge_cloud_accounts" "all" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of cloud accounts returned
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.#", "1"),
					// Verify an attribute of the first cloud account
					resource.TestCheckResourceAttr("data.pgedge_cloud_accounts.all", "cloud_accounts.0.name", "pgedge"),
				),
			},
		},
	})
}

