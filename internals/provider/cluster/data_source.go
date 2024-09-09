package cluster

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	ID             types.String   `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
	CloudAccountID types.String   `tfsdk:"cloud_account_id"`
	CreatedAt      types.String   `tfsdk:"created_at"`
	Status         types.String   `tfsdk:"status"`
	SSHKeyID       types.String   `tfsdk:"ssh_key_id"`
	Regions        types.List     `tfsdk:"regions"`
	NodeLocation   types.String   `tfsdk:"node_location"`
	Capacity       types.Int64    `tfsdk:"capacity"`
	FirewallRules  []types.Object `tfsdk:"firewall_rules"`
	Nodes          []types.Object `tfsdk:"nodes"`
	Networks       []types.Object `tfsdk:"networks"`
	BackupStoreIDs types.List     `tfsdk:"backup_store_ids"`
	ResourceTags   types.Map      `tfsdk:"resource_tags"`
}

type FirewallRule struct {
	Port    types.Int64 `tfsdk:"port"`
	Sources types.List  `tfsdk:"sources"`
}

type ClusterNode struct {
	Region           types.String `tfsdk:"region"`
	Name             types.String `tfsdk:"name"`
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	Options          types.List   `tfsdk:"options"`
	VolumeSize       types.Int64  `tfsdk:"volume_size"`
	VolumeIOPS       types.Int64  `tfsdk:"volume_iops"`
	VolumeType       types.String `tfsdk:"volume_type"`
	InstanceType     types.String `tfsdk:"instance_type"`
}

type ClusterNetworks struct {
	Cidr           types.String `tfsdk:"cidr"`
	PublicSubnets  types.List   `tfsdk:"public_subnets"`
	PrivateSubnets types.List   `tfsdk:"private_subnets"`
	Region         types.String `tfsdk:"region"`
	Name           types.String `tfsdk:"name"`
	External       types.Bool   `tfsdk:"external"`
	ExternalId     types.String `tfsdk:"external_id"`
}

var (
	FireWallType = map[string]attr.Type{
		"port": types.Int64Type,
		"sources": types.ListType{
			ElemType: types.StringType,
		},
	}
	NetworksType = map[string]attr.Type{
		"cidr":            types.StringType,
		"public_subnets":  types.ListType{ElemType: types.StringType},
		"private_subnets": types.ListType{ElemType: types.StringType},
		"region":          types.StringType,
		"name":            types.StringType,
		"external":        types.BoolType,
		"external_id":     types.StringType,
	}
	ClusterNodeTypes = map[string]attr.Type{
		"region":            types.StringType,
		"name":              types.StringType,
		"availability_zone": types.StringType,
		"options":           types.ListType{ElemType: types.StringType},
		"volume_size":       types.Int64Type,
		"volume_iops":       types.Int64Type,
		"volume_type":       types.StringType,
		"instance_type":     types.StringType,
	}
)

var ClusterNodeAttribute = schema.ListNestedAttribute{
	Computed: true,
	Optional: true,
	NestedObject: schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"region": schema.StringAttribute{
				Required:    true,
				Description: "Cloud provider region",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Node name",
			},
			"availability_zone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cloud provider availability zone name",
			},
			"options": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"volume_size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Volume size of the node data volume",
			},
			"volume_iops": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Volume IOPS of the node data volume",
			},
			"volume_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Volume type of the node data volume",
			},
			"instance_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Instance type used for the node",
			},
		},
	},
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
							Computed:    true,
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
						"ssh_key_id": schema.StringAttribute{
							Computed:    true,
							Description: "SSH key ID of the cluster",
						},
						"regions": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"node_location": schema.StringAttribute{
							Computed:    true,
							Description: "Node location of the cluster",
						},
						"capacity": schema.Int64Attribute{
							Computed:    true,
							Description: "Capacity of the cluster",
						},
						"firewall_rules": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "Name of the firewall rule",
									},
									"port": schema.Int64Attribute{
										Computed:    true,
										Description: "Port for the firewall rule",
									},
									"sources": schema.ListAttribute{
										ElementType: types.StringType,
										Computed:    true,
										Description: "Sources for the firewall rule",
									},
								},
							},
						},
						"nodes": ClusterNodeAttribute,
						"networks": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "Name of the network",
									},
									"region": schema.StringAttribute{
										Computed:    true,
										Description: "Region of the network",
									},
									"external": schema.BoolAttribute{
										Computed:    true,
										Description: "Is the network external",
									},
									"external_id": schema.StringAttribute{
										Computed:    true,
										Description: "External ID of the network",
									},
									"cidr": schema.StringAttribute{
										Computed:    true,
										Description: "CIDR of the network",
									},
									"public_subnets": schema.ListAttribute{
										ElementType: types.StringType,
										Computed:    true,
									},
									"private_subnets": schema.ListAttribute{
										ElementType: types.StringType,
										Computed:    true,
									},
								},
							},
						},
						"backup_store_ids": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Backup store IDs of the cluster",
						},
						"resource_tags": schema.MapAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Resource tags of the cluster",
						},
					},
				},
			},
		},
		Description: "Data source for pgEdge clusters.",
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
		clusterDetails.ID = types.StringValue(cluster.ID.String())
		clusterDetails.Name = types.StringValue(*cluster.Name)
		clusterDetails.CloudAccountID = types.StringValue(*cluster.CloudAccount.ID)
		clusterDetails.CreatedAt = types.StringValue(*cluster.CreatedAt)
		clusterDetails.Status = types.StringValue(*cluster.Status)
		clusterDetails.SSHKeyID = types.StringValue(cluster.SSHKeyID)
		clusterDetails.NodeLocation = types.StringValue(*cluster.NodeLocation)
		clusterDetails.Capacity = types.Int64Value(cluster.Capacity)

		// Set Regions
		regions := make([]attr.Value, len(cluster.Regions))
		for i, region := range cluster.Regions {
			regions[i] = types.StringValue(region)
		}
		clusterDetails.Regions = types.ListValueMust(types.StringType, regions)

		// Set FirewallRules
		firewallRules := make([]types.Object, len(cluster.FirewallRules))
		for i, rule := range cluster.FirewallRules {
			firewallRule, diags := types.ObjectValue(
				map[string]attr.Type{
					"name":    types.StringType,
					"port":    types.Int64Type,
					"sources": types.ListType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"name": types.StringValue(rule.Name),
					"port": types.Int64Value(*rule.Port),
					"sources": types.ListValueMust(
						types.StringType,
						func() []attr.Value {
							sources := make([]attr.Value, len(rule.Sources))
							for i, source := range rule.Sources {
								sources[i] = types.StringValue(source)
							}
							return sources
						}(),
					),
				},
			)
			resp.Diagnostics.Append(diags...)
			firewallRules[i] = firewallRule
		}
		clusterDetails.FirewallRules = firewallRules

		// Set Nodes
		nodes := make([]types.Object, len(cluster.Nodes))
		for i, node := range cluster.Nodes {
			nodeObj, diags := types.ObjectValue(
				map[string]attr.Type{
					"name":              types.StringType,
					"region":            types.StringType,
					"availability_zone": types.StringType,
					"instance_type":     types.StringType,
					"volume_size":       types.Int64Type,
					"volume_type":       types.StringType,
					"volume_iops":       types.Int64Type,
					"options":           types.ListType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"name":              types.StringValue(node.Name),
					"region":            types.StringValue(*node.Region),
					"availability_zone": types.StringValue(node.AvailabilityZone),
					"instance_type":     types.StringValue(node.InstanceType),
					"volume_size":       types.Int64Value(node.VolumeSize),
					"volume_type":       types.StringValue(node.VolumeType),
					"volume_iops":       types.Int64Value(node.VolumeIops),
					"options": types.ListValueMust(
						types.StringType,
						func() []attr.Value {
							options := make([]attr.Value, len(node.Options))
							for i, option := range node.Options {
								options[i] = types.StringValue(option)
							}
							return options
						}(),
					),
				},
			)
			resp.Diagnostics.Append(diags...)
			nodes[i] = nodeObj
		}
		clusterDetails.Nodes = nodes

		// Set Networks
		networks := make([]types.Object, len(cluster.Networks))
		for i, network := range cluster.Networks {
			networkObj, diags := types.ObjectValue(
				map[string]attr.Type{
					"name":            types.StringType,
					"region":          types.StringType,
					"external":        types.BoolType,
					"external_id":     types.StringType,
					"cidr":            types.StringType,
					"public_subnets":  types.ListType{ElemType: types.StringType},
					"private_subnets": types.ListType{ElemType: types.StringType},
				},
				map[string]attr.Value{
					"name":        types.StringValue(network.Name),
					"region":      types.StringValue(*network.Region),
					"external":    types.BoolValue(network.External),
					"external_id": types.StringValue(network.ExternalID),
					"cidr":        types.StringValue(network.Cidr),
					"public_subnets": types.ListValueMust(
						types.StringType,
						func() []attr.Value {
							subnets := make([]attr.Value, len(network.PublicSubnets))
							for i, subnet := range network.PublicSubnets {
								subnets[i] = types.StringValue(subnet)
							}
							return subnets
						}(),
					),
					"private_subnets": types.ListValueMust(
						types.StringType,
						func() []attr.Value {
							subnets := make([]attr.Value, len(network.PrivateSubnets))
							for i, subnet := range network.PrivateSubnets {
								subnets[i] = types.StringValue(subnet)
							}
							return subnets
						}(),
					),
				},
			)
			resp.Diagnostics.Append(diags...)
			networks[i] = networkObj
		}
		clusterDetails.Networks = networks

		// Set BackupStoreIDs
		backupStoreIDs := make([]attr.Value, len(cluster.BackupStoreIds))
		for i, id := range cluster.BackupStoreIds {
			backupStoreIDs[i] = types.StringValue(id)
		}
		clusterDetails.BackupStoreIDs = types.ListValueMust(types.StringType, backupStoreIDs)

		// Set ResourceTags
		resourceTags := make(map[string]attr.Value)
		for k, v := range cluster.ResourceTags {
			resourceTags[k] = types.StringValue(v)
		}
		clusterDetails.ResourceTags = types.MapValueMust(types.StringType, resourceTags)

		state.Clusters = append(state.Clusters, clusterDetails)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
