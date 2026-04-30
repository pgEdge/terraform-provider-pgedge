package cloudaccount_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper/fakeapi"
)

func TestAccCloudAccountResource(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
				resource "pgedge_cloud_account" "aws_account" {
					name        = "test_account"
					type        = "aws"
					description = "My AWS test Cloud Account"

					credentials = {
						role_arn = "arn:aws:iam::000000000000:role/test"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_cloud_account.aws_account", "name", "test_account"),
					resource.TestCheckResourceAttr("pgedge_cloud_account.aws_account", "type", "aws"),
					resource.TestCheckResourceAttrSet("pgedge_cloud_account.aws_account", "id"),
				),
			},
		},
	})
}

// Description is Optional (no Computed). When the user omits it, the API
// returns "" — the Read flow used to lift that into a known string, which
// drifted from a null config and triggered a non-empty plan on every refresh.
// Update is unimplemented, so the user would be stuck. This test would fail
// (non-empty plan after apply) if that regression returned.
func TestAccCloudAccountResource_OmittedDescription(t *testing.T) {
	api := fakeapi.New(t)
	api.ConfigureProvider(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: common.ProviderConfig + `
				resource "pgedge_cloud_account" "no_desc" {
					name = "no_desc_account"
					type = "aws"
					credentials = {
						role_arn = "arn:aws:iam::000000000000:role/test"
					}
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_cloud_account.no_desc", "name", "no_desc_account"),
					resource.TestCheckNoResourceAttr("pgedge_cloud_account.no_desc", "description"),
				),
			},
		},
	})
}
