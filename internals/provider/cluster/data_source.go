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
	Firewall       []types.Object `tfsdk:"firewall_rules"`
	Nodes          []types.Object `tfsdk:"nodes"`
	Networks       []types.Object `tfsdk:"networks"`
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
						// "last_updated": schema.StringAttribute{
						// 	Computed: true,
						// 	Optional: true,
						// },
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
						"ssh_key_id": schema.StringAttribute{
							Computed:    true,
							Description: "SSH key ID of the cluster",
						},
						// "resource_tags": schema.MapAttribute{
						// 	ElementType: types.StringType,
						// 	Computed:    true,
						// },
						"regions": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"node_location": schema.StringAttribute{
							Computed:    true,
							Description: "Node location of the cluster",
						},
						"cloud_account": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed:    true,
									Description: "Display name of the node",
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: "IP address of the node",
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: "Type of the node",
								},
							},
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
										Description: "CIDR of the AWS node group",
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
					},
				},
			},
		},
		Description: "Interface with the pgEdge service API for clusters.",
	}
}

func (c *clustersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ClustersDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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
		clusterDetails.Name = types.StringValue(cluster.Name)
		clusterDetails.CloudAccountID = types.StringValue(cluster.CloudAccount.ID.String())
		clusterDetails.CreatedAt = types.StringValue(cluster.CreatedAt.String())
		clusterDetails.Status = types.StringValue(cluster.Status)

		for _, rule := range cluster.FirewallRules {
			var firewallRule FirewallRule
			var firewallSources []attr.Value
			for _, source := range rule.Sources {
				firewallSources = append(firewallSources, types.StringValue(source))
			}
			firewallRule.Port = types.Int64Value(rule.Port)
			firewallRule.Sources, diags = types.ListValue(types.StringType, firewallSources)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			firewallElements := map[string]attr.Value{
				"port":    types.Int64Value(rule.Port),
				"sources": firewallRule.Sources,
			}
			firewallObjectValue, diags := types.ObjectValue(FireWallType, firewallElements)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			clusterDetails.Firewall = append(clusterDetails.Firewall, firewallObjectValue)
		}

		for _, node := range cluster.Nodes {
			var clusterNode ClusterNode
			clusterNode.Region = types.StringValue(node.Region)
			clusterNode.Name = types.StringValue(node.Name)
			clusterNode.AvailabilityZone = types.StringValue(node.AvailabilityZone)

			var options []attr.Value
			for _, option := range node.Options {
				options = append(options, types.StringValue(option))
			}
			clusterNode.Options, diags = types.ListValue(types.StringType, options)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			clusterNode.VolumeSize = types.Int64Value(node.VolumeSize)
			clusterNode.VolumeIOPS = types.Int64Value(node.VolumeIops)
			clusterNode.VolumeType = types.StringValue(node.VolumeType)
			clusterNode.InstanceType = types.StringValue(node.InstanceType)

			nodeMap := map[string]attr.Value{
				"region":            clusterNode.Region,
				"name":              clusterNode.Name,
				"availability_zone": clusterNode.AvailabilityZone,
				"options":           clusterNode.Options,
				"volume_size":       clusterNode.VolumeSize,
				"volume_iops":       clusterNode.VolumeIOPS,
				"volume_type":       clusterNode.VolumeType,
				"instance_type":     clusterNode.InstanceType,
			}

			nodeValue, diags := types.ObjectValue(ClusterNodeTypes, nodeMap)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			clusterDetails.Nodes = append(clusterDetails.Nodes, nodeValue)
		}

		for _, network := range cluster.Networks {
			var clusterNetwork ClusterNetworks
			clusterNetwork.Cidr = types.StringValue(network.Cidr)

			var publicSubnets []attr.Value
			for _, subnet := range network.PublicSubnets {
				publicSubnets = append(publicSubnets, types.StringValue(subnet))
			}
			clusterNetwork.PublicSubnets, diags = types.ListValue(types.StringType, publicSubnets)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			var privateSubnets []attr.Value
			for _, subnet := range network.PrivateSubnets {
				privateSubnets = append(privateSubnets, types.StringValue(subnet))
			}
			clusterNetwork.PrivateSubnets, diags = types.ListValue(types.StringType, privateSubnets)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			clusterNetwork.Region = types.StringValue(network.Region)
			clusterNetwork.Name = types.StringValue(network.Name)
			clusterNetwork.External = types.BoolValue(network.External)
			clusterNetwork.ExternalId = types.StringValue(network.ExternalID)

			networkMap := map[string]attr.Value{
				"cidr":            clusterNetwork.Cidr,
				"public_subnets":  clusterNetwork.PublicSubnets,
				"private_subnets": clusterNetwork.PrivateSubnets,
				"region":          clusterNetwork.Region,
				"name":            clusterNetwork.Name,
				"external":        clusterNetwork.External,
				"external_id":     clusterNetwork.ExternalId,
			}

			networkValue, diags := types.ObjectValue(NetworksType, networkMap)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			clusterDetails.Networks = append(clusterDetails.Networks, networkValue)
		}

		clusterDetails.SSHKeyID = types.StringValue(cluster.SSHKeyID)
		// var resourceTags []attr.Value
		// for key, value := range cluster.ResourceTags {
		// 	resourceTags = append(resourceTags, types.StringValue(key+":"+value))
		// }
		// clusterDetails.ResourceTags, diags = types.MapValue(types.StringType, resourceTags)
		// resp.Diagnostics.Append(diags...)
		// if resp.Diagnostics.HasError() {
		// 	return
		// }

		var regions []attr.Value
		for _, region := range cluster.Regions {
			regions = append(regions, types.StringValue(region))
		}
		clusterDetails.Regions, diags = types.ListValue(types.StringType, regions)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		clusterDetails.NodeLocation = types.StringValue(cluster.NodeLocation)

		state.Clusters = append(state.Clusters, clusterDetails)
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
