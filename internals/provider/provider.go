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
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/cloud-account"
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
		Blocks:      map[string]schema.Block{},
		Description: "Interface with the pgEdge service API.",
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
		baseUrl = "https://api.pgedge.com"
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
		NewDatabaseResource,
		NewClusterResource,
	}
}

func (p *PgEdgeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDatabasesDataSource,
		NewClustersDataSource,
		cloudaccount.NewCloudAccountsDataSource,
	}
}
