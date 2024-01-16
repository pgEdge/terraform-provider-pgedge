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
)

// Ensure PgEdgeProvider satisfies various provider interfaces.
var _ provider.Provider = &PgEdgeProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PgEdgeProvider{
			version: version,
		}
	}
}

// PgEdgeProvider defines the provider implementation.
type PgEdgeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PgEdgeProviderModel describes the provider data model.
type PgEdgeProviderModel struct {
	AuthHeader types.String `tfsdk:"auth_header"`
	ClusterID  types.String `tfsdk:"cluster_id"`
}

func (p *PgEdgeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pgedge"
	resp.Version = p.version
}

func (p *PgEdgeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth_header": schema.StringAttribute{
				Required:    true,
				Description: "Authentication header to use when connecting to the PgEdge service.",
				Sensitive:   true,
			},
			"cluster_id": schema.StringAttribute{
				Required:    true,
				Description: "The Cluster ID to use when connecting to the PgEdge service.",
			},
		},
		Blocks:      map[string]schema.Block{},
		Description: "Interface with the pgEdge service API.",
	}
}
func (p *PgEdgeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    tflog.Info(ctx, "Configuring pgEdge client")
    
    // Retrieve provider data from configuration
    var config PgEdgeProviderModel
    diags := req.Config.Get(ctx, &config)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // If practitioner provided a configuration value for any of the
    // attributes, it must be a known value.

    if config.AuthHeader.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("auth_header"),
            "Unknown PgEdge API Auth Header",
            "The provider cannot create the pgEdge API client as there is an unknown configuration value for the pgEdge API Auth Header. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the PGEDGE_AUTHHEADER environment variable.",
        )
    }

    if config.ClusterID.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("cluster_id"),
            "Unknown pgEdge API Cluster ID",
            "The provider cannot create the pgEdge API client as there is an unknown configuration value for the pgEdge API Cluster ID. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the PGEDGE_CLUSTERID environment variable.",
        )
    }
    if resp.Diagnostics.HasError() {
        return
    }

    // Default values to environment variables, but override
    // with Terraform configuration value if set.

    authHeader := os.Getenv("PGEDGE_AUTHHEADER")
    clusterId := os.Getenv("PGEDGE_CLUSTERID")

    if !config.AuthHeader.IsNull() {
        authHeader = config.AuthHeader.ValueString()
    }

    if !config.ClusterID.IsNull() {
        clusterId = config.ClusterID.ValueString()
    }

    // If any of the expected configurations are missing, return
    // errors with provider-specific guidance.

    if authHeader == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("auth_header"),
            "Missing pgEdge API auth_header",
            "The provider cannot create the pgEdge API client as there is a missing or empty value for the pgEdge API auth_header. "+
                "Set the auth_header value in the configuration or use the PGEDGE_AUTHHEADER environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if clusterId == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("cluster_id"),
            "Missing pgEdge API cluster_id",
            "The provider cannot create the pgEdge API client as there is a missing or empty value for the pgEdge API cluster_id. "+
                "Set the cluster_id value in the configuration or use the PGEDGE_CLUSTERID environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    ctx = tflog.SetField(ctx, "pgEdge_cluster_id", clusterId)
    ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "pgEdge_auth_header", authHeader)

    tflog.Debug(ctx, "Creating pgEdge client")

    // Create a new pgEdge client using the configuration values
    client := client.NewClient("https://devapi.pgedge.com", authHeader, clusterId)
    // if err != nil {
    //     resp.Diagnostics.AddError(
    //         "Unable to Create pgEdge API Client",
    //         "An unexpected error occurred when creating the pgEdge API client. "+
    //             "If the error is not clear, please contact the provider developers.\n\n"+
    //             "pgEdge Client Error: "+err.Error(),
    //     )
    //     return
    // }

    // Make the pgEdge client available during DataSource and Resource
    // type Configure methods.
    resp.DataSourceData = client
    resp.ResourceData = client

    tflog.Info(ctx, "Configured pgEdge client", map[string]any{"success": true})
}

func (p *PgEdgeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDatabasesResource,
	}
}

func (p *PgEdgeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
        NewDatabasesDataSource,
	}
}
