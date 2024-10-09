package cloudaccount_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	common "github.com/pgEdge/terraform-provider-pgedge/internals/provider/common/test-helper"
)

func TestAccCloudAccountResource(t *testing.T) {
	roleArn := os.Getenv("PGEDGE_ROLE_ARN")
	if roleArn == "" {
		t.Skip("PGEDGE_ROLE_ARN= environment variable is not set")
	}


	resource.Test(t, resource.TestCase{
		PreCheck:                 func() {  },
		ProtoV6ProviderFactories: common.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: fmt.Sprintf(`
				resource "pgedge_cloud_account" "aws_account" {
					name = "test_account"
					type = "aws"
  					description = "My AWS test Cloud Account"
					
					credentials = {
						role_arn = "%s"
					}
				}
				`, roleArn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pgedge_cloud_account.aws_account", "name", "test_account"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
