package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pgEdge/terraform-provider-pgedge/client"
	backupstore "github.com/pgEdge/terraform-provider-pgedge/internals/provider/backup-store"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/cloud-account"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/cluster"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/database"
	sshkey "github.com/pgEdge/terraform-provider-pgedge/internals/provider/ssh-key"
)

var _ provider.Provider = &PgEdgeProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PgEdgeProvider{
			version: version,
		}
	}
}

type PgEdgeProvider struct {
	version string
}

type PgEdgeProviderModel struct {
	BaseUrl      types.String `tfsdk:"base_url"`
}

func (p *PgEdgeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pgedge"
	resp.Version = p.version
}

func (p *PgEdgeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Optional:    true,
				Description: "Base Url to use when connecting to the PgEdge service.",
			},
		},
		Blocks: map[string]schema.Block{},
        Description: `
The pgEdge provider is used to interact with the resources supported by pgEdge. 
It allows you to manage various aspects of your pgEdge infrastructure, including databases, clusters, cloud accounts, SSH keys, and backup stores.

## Authentication

The provider needs to be configured with the proper credentials before it can be used. 
You can provide your credentials via environment variables:

- Set the PGEDGE_CLIENT_ID environment variable for your pgEdge Client ID.
- Set the PGEDGE_CLIENT_SECRET environment variable for your pgEdge Client Secret.

Example provider configuration:

` + "```hcl" + `
# Configure the pgEdge Provider
provider "pgedge" {}

# Set environment variables
# export PGEDGE_CLIENT_ID="your-client-id"
# export PGEDGE_CLIENT_SECRET="your-client-secret"
` + "```" + `

## Example Usage

### Managing a Cluster

` + "```hcl" + `
resource "pgedge_cluster" "example" {
  name             = "example-cluster"
  cloud_account_id = "your-cloud-account-id"
  regions          = ["us-west-2", "us-east-1"]
  node_location    = "public"

  nodes = [
    {
      name          = "node1"
      region        = "us-west-2"
      instance_type = "r6g.medium"
    },
    {
      name          = "node2"
      region        = "us-east-1"
      instance_type = "r6g.medium"
    }
  ]

  networks = [
    {
      region         = "us-west-2"
      cidr           = "10.1.0.0/16"
      public_subnets = ["10.1.1.0/24"]
    },
    {
      region         = "us-east-1"
      cidr           = "10.2.0.0/16"
      public_subnets = ["10.2.1.0/24"]
    }
  ]
}
` + "```" + `

### Managing a Database

` + "```hcl" + `
resource "pgedge_database" "example" {
  name       = "example-db"
  cluster_id = "your-cluster-id"

  options = [
    "install:northwind",
    "rest:enabled"
  ]

  extensions = {
    auto_manage = true
    requested   = ["postgis"]
  }

  nodes = {
	n1 = {
	  name = "n1"
	},
	n2 = {
	  name = "n2"
	}
  }
}
` + "```" + `

For more information on the available resources and their configurations, please refer to the documentation for each resource and data source.
`,
    }
}

func (p *PgEdgeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring pgEdge client")

	var config PgEdgeProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	baseUrl := os.Getenv("PGEDGE_BASE_URL")
	ClientId := os.Getenv("PGEDGE_CLIENT_ID")
	GrantType := os.Getenv("PGEDGE_GRANT_TYPE")
	ClientSecret := os.Getenv("PGEDGE_CLIENT_SECRET")

	if !config.BaseUrl.IsNull() {
		baseUrl = config.BaseUrl.ValueString()
	}

	if baseUrl == "" {
		baseUrl = "https://api.pgedge.com/v1"
	}

	if ClientId == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing pgEdge API client_id",
			"The provider cannot create the pgEdge API client as there is a missing or empty value for the pgEdge API client_id. "+
				"Set the client_id value in the configuration or use the PGEDGE_CLIENT_ID environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if ClientSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing pgEdge API client_secret",
			"The provider cannot create the pgEdge API client as there is a missing or empty value for the pgEdge API client_secret. "+
				"Set the client_secret value in the configuration or use the PGEDGE_CLIENT_SECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "pgEdge_base_url", baseUrl)
	ctx = tflog.SetField(ctx, "pgEdge_client_id", ClientId)
	ctx = tflog.SetField(ctx, "pgEdge_grant_type", GrantType)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "pgEdge_client_secret", ClientSecret)

	tflog.Debug(ctx, "Creating pgEdge client")

	mockClient := client.NewClient(baseUrl, "")
	token, err := mockClient.OAuthToken(context.Background(), ClientId, ClientSecret, GrantType)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create pgEdge API Client",
			"An unexpected error occurred when creating the pgEdge API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"pgEdge Client Error: "+err.Error(),
		)
		return
	}

	pgEdgeClient := client.NewClient(baseUrl, "Bearer "+token.AccessToken)
	resp.DataSourceData = pgEdgeClient
	resp.ResourceData = pgEdgeClient

	tflog.Info(ctx, "Configured pgEdge client", map[string]any{"success": true})
}

func (p *PgEdgeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		database.NewDatabaseResource,
		cluster.NewClusterResource,
		cloudaccount.NewCloudAccountResource,
		sshkey.NewSSHKeyResource,
		backupstore.NewBackupStoreResource,
	}
}

func (p *PgEdgeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		database.NewDatabasesDataSource,
		cluster.NewClustersDataSource,
		cloudaccount.NewCloudAccountsDataSource,
		sshkey.NewSSHKeysDataSource,
		backupstore.NewBackupStoresDataSource,
	}
}
