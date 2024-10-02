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
The official Terraform provider for [pgEdge Cloud](https://www.pgedge.com/cloud), designed to simplify the management of pgEdge Cloud resources for both **Developer** and **Enterprise** edition.

## Authentication

Before using the provider, you must create an API Client in [pgEdge Cloud](https://app.pgedge.com) and configure the following environment variables:

` + "```sh" + `
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
` + "```" + `

These credentials authenticate the Terraform provider with your pgEdge Cloud account.

## Example Usage

### Developer Edition Configuration

For Developer Edition, pgEdge offers access to manage databases. Here's an example setup for Developer Edition:

` + "```hcl" + `
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}

provider "pgedge" {}

# Define a database
resource "pgedge_database" "defaultdb" {
  name       = "defaultdb"
  cluster_id = "f12239ddq-df9d-4ded-adqwead9-3e2bvhe6d6ee"

  options = [
    "rest:enabled",
    "install:northwind"
  ]
}
` + "```" + `

### Enterprise Edition Configuration

Enterprise Edition users can manage Cloud Accounts, SSH keys, Backup Stores, and Clusters. Here's an Enterprise Edition example:

` + "```hcl" + `
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}

provider "pgedge" {
  base_url = "https://api.pgedge.com"
}

# SSH Key resource
resource "pgedge_ssh_key" "example" {
  name       = "example-key"
  public_key = "ssh-ed25519 AAAAC3NzaC1wes241mmT63i04t5fvvsdwqVG7DkyxvyXbYQNhKP/rSeLY user@example.com"
}

# Cloud Account resource
resource "pgedge_cloud_account" "example" {
  name        = "my-aws-account"
  type        = "aws"
  description = "My AWS Cloud Account"

  credentials = {
    role_arn = "arn:aws:iam::0123456789:role/pgedge-13fe3332c"
  }

  depends_on = [pgedge_ssh_key.example]
}

# Backup Store resource
resource "pgedge_backup_store" "test_store" {
  name             = "test-store"
  cloud_account_id = pgedge_cloud_account.example.id
  region           = "us-east-1"

  depends_on = [pgedge_cloud_account.example]
}

# Cluster resource
resource "pgedge_cluster" "example" {
  name             = "example"
  cloud_account_id = pgedge_cloud_account.example.id
  regions          = ["us-west-2", "us-east-1", "eu-central-1"]
  ssh_key_id       = pgedge_ssh_key.example.id
  backup_store_ids = [pgedge_backup_store.test_store.id]
  node_location    = "public"

  nodes = [
    {
      name          = "n1"
      region        = "us-west-2"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    },
    {
      name          = "n2"
      region        = "us-east-1"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    },
    {
      name          = "n3"
      region        = "eu-central-1"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    }
  ]

  networks = [
    {
      region         = "us-west-2"
      cidr           = "10.1.0.0/16"
      public_subnets = ["10.1.0.0/24"]
    },
    {
      region         = "us-east-1"
      cidr           = "10.2.0.0/16"
      public_subnets = ["10.2.0.0/24"]
    },
    {
      region         = "eu-central-1"
      cidr           = "10.3.0.0/16"
      public_subnets = ["10.3.0.0/24"]
    }
  ]

  firewall_rules = [
    {
      name    = "postgres"
      port    = 5432
      sources = ["192.0.2.44/32"]
    },
  ]

  depends_on = [pgedge_cloud_account.example]
}

# Database Resource
resource "pgedge_database" "example_db" {
  name       = "exampledb"
  cluster_id = pgedge_cluster.example.id

  options = [
    "autoddl:enabled",
  ]

  extensions = {
    auto_manage = true
    requested = [
      "postgis",
      "vector"
    ]
  }

  nodes = {
    n1 = {
      name = "n1"
    },
    n2 = {
      name = "n2"
    },
    n3 = {
      name = "n3"
    }
  }

  backups = {
    provider = "pgbackrest"
    config = [
      {
        id        = "default"
        schedules = [
          {
            type            = "full"
            cron_expression = "0 6 * * ?"
            id              = "daily-full-backup"
          },
          {
            type            = "incr"
            cron_expression = "0 * * * ?"
            id              = "hourly-incr-backup"
          }
        ]
      }
    ]
  }

  depends_on = [pgedge_cluster.example]
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
				"If the error is not clear, please contact the provider developer.\n\n"+
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
