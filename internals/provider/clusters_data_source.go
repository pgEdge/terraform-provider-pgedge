package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
)

var (
	_ datasource.DataSource              = &clustersDataSource{}
	_ datasource.DataSourceWithConfigure = &clustersDataSource{}
)

func NewClustersDataSource() datasource.DataSource {
	return &clustersDataSource{}
}

type clustersDataSource struct {
	client *pgEdge.Client
}

func (c *clustersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusters"
}

func (c *clustersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pgEdge.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *pgEdge.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	c.client = client
}

type ClustersDataSourceModel struct {
	Clusters []ClusterDetails `tfsdk:"clusters"`
}

type ClusterDetails struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	CloudAccountID types.String `tfsdk:"cloud_account_id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	Status         types.String `tfsdk:"status"`

	Aws        ClusterDetailsAws        `tfsdk:"aws"`
	Database   ClusterDetailsDatabase   `tfsdk:"database"`
	Firewall   ClusterDetailsFirewall   `tfsdk:"firewall"`
	NodeGroups ClusterDetailsNodeGroups `tfsdk:"node_groups"`
}

type ClusterDetailsAws struct {
	RoleArn types.String `tfsdk:"role_arn"`
}

type ClusterDetailsDatabase struct {
	Name      types.String `tfsdk:"name"`
	PgVersion types.String `tfsdk:"pg_version"`
	Scripts   interface{}  `tfsdk:"scripts"`
	Username  types.String `tfsdk:"username"`
}

type ClusterDetailsFirewall struct {
	Rules []ClusterDetailsFirewallRulesItems0 `tfsdk:"rules"`
}

type ClusterDetailsFirewallRulesItems0 struct {
	Port    types.Int64    `tfsdk:"port"`
	Sources []types.String `tfsdk:"sources"`
	Type    types.String   `tfsdk:"type"`
}

type ClusterDetailsNodeGroups struct {
	Aws    []interface{} `tfsdk:"aws"`
	Azure  []interface{} `tfsdk:"azure"`
	Google []interface{} `tfsdk:"google"`
}

func (c *clustersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"clusters": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "ID of the cluster",
						},
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the cluster",
						},
						"cloud_account_id": schema.StringAttribute{
							Computed:    true,
							Description: "Cloud account ID of the cluster",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "Created at of the cluster",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Status of the cluster",
						},
						"aws": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"role_arn": schema.StringAttribute{
									Computed:    true,
									Description: "Role ARN of the AWS cluster",
								},
							},
						},
						"database": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Computed:    true,
									Description: "Name of the database",
								},
								"pg_version": schema.StringAttribute{
									Computed:    true,
									Description: "PostgreSQL version of the database",
								},
								"scripts": schema.MapAttribute{
									Computed:    true,
									Description: "Scripts for the database",
								},
								"username": schema.StringAttribute{
									Computed:    true,
									Description: "Username for the database",
								},
							},
						},
						"firewall": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"rules": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"port": schema.NumberAttribute{
												Computed:    true,
												Description: "Port for the firewall rule",
											},
											"sources": schema.ListAttribute{
												ElementType: types.StringType,
												Computed:    true,
												Description: "Sources for the firewall rule",
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: "Type of the firewall rule",
											},
										},
									},
								},
							},
						},
						"node_groups": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"aws": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
									Description: "AWS node groups",
								},
								"azure": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
									Description: "Azure node groups",
								},
								"google": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
									Description: "Google node groups",
								},
							},
						},
					},
				},
			},
		},
		Description: "Interface with the pgEdge service API for clusters.",
	}
}

func (c *clustersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ClustersDataSourceModel

	clusters, err := c.client.GetAllClusters(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read pgEdge Clusters",
			err.Error(),
		)
		return
	}

	for _, cluster := range clusters {
		var clusterDetails ClusterDetails

		clusterDetails.ID = types.StringValue(cluster.ID)
		clusterDetails.Name = types.StringValue(cluster.Name)
		clusterDetails.CloudAccountID = types.StringValue(cluster.CloudAccountID)
		clusterDetails.CreatedAt = types.StringValue(cluster.CreatedAt.String())
		clusterDetails.Status = types.StringValue(cluster.Status)

		// Populate AWS details
		clusterDetails.Aws.RoleArn = types.StringValue(cluster.Aws.RoleArn)

		// Populate Database details
		clusterDetails.Database.Name = types.StringValue(cluster.Database.Name)
		clusterDetails.Database.PgVersion = types.StringValue(cluster.Database.PgVersion)
		clusterDetails.Database.Scripts = cluster.Database.Scripts
		clusterDetails.Database.Username = types.StringValue(cluster.Database.Username)

		// Populate Firewall details
		for _, rule := range cluster.Firewall.Rules {
			var firewallRuleSources []types.String

			for _, rule := range rule.Sources {
				firewallRuleSources = append(firewallRuleSources, types.StringValue(rule))
			}
			var firewallRule ClusterDetailsFirewallRulesItems0
			firewallRule.Port = types.Int64Value(rule.Port)
			firewallRule.Sources = firewallRuleSources
			firewallRule.Type = types.StringValue(rule.Type)
			clusterDetails.Firewall.Rules = append(clusterDetails.Firewall.Rules, firewallRule)
		}

		// Populate NodeGroups details
		clusterDetails.NodeGroups.Aws = cluster.NodeGroups.Aws
		clusterDetails.NodeGroups.Azure = cluster.NodeGroups.Azure
		clusterDetails.NodeGroups.Google = cluster.NodeGroups.Google

		state.Clusters = append(state.Clusters, clusterDetails)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
