package provider

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
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	CloudAccountID types.String `tfsdk:"cloud_account_id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	Status         types.String `tfsdk:"status"`

	// Aws        types.Object   `tfsdk:"aws"`
	// Database   types.Object   `tfsdk:"database"`
	Firewall   []types.Object `tfsdk:"firewall"`
	// NodeGroups types.Object   `tfsdk:"node_groups"`
}

type AWS struct {
	RoleARN types.String `tfsdk:"role_arn"`
	KeyPair types.String `tfsdk:"key_pair"`
	Tags    types.Map    `tfsdk:"tags"`
}

type Database struct {
	PGVersion  types.String    `tfsdk:"pg_version"`
	Username   types.String    `tfsdk:"username"`
	Password   types.String    `tfsdk:"password"`
	Name       types.String    `tfsdk:"name"`
	Port       types.Float64   `tfsdk:"port"`
	Components []types.String  `tfsdk:"components"`
	Scripts    DatabaseScripts `tfsdk:"scripts"`
}

type DatabaseScripts struct {
	Init types.String `tfsdk:"init"`
}

type FirewallRule struct {
	Type    types.String   `tfsdk:"type,omitempty"`
	Port    types.Int64    `tfsdk:"port,omitempty"`
	Sources types.List `tfsdk:"sources"`
}

type ClusterNode struct {
	DisplayName types.String `tfsdk:"display_name"`
	IPAddress   types.String `tfsdk:"ip_address"`
	IsActive    types.Bool   `tfsdk:"is_active"`
}

type NodeGroup struct {
	Region            types.String `tfsdk:"region"`
	AvailabilityZones types.List   `tfsdk:"availability_zones"`
	Cidr              types.String `tfsdk:"cidr"`
	PublicSubnets     types.List   `tfsdk:"public_subnets"`
	PrivateSubnets    types.List   `tfsdk:"private_subnets"`
	Nodes             types.List   `tfsdk:"nodes"`
	NodeLocation      types.String `tfsdk:"node_location"`
	VolumeSize        types.Int64  `tfsdk:"volume_size"`
	VolumeIOPS        types.Int64  `tfsdk:"volume_iops"`
	VolumeType        types.String `tfsdk:"volume_type"`
	InstanceType      types.String `tfsdk:"instance_type"`
}

type NodeGroups struct {
	AWS    []NodeGroup `tfsdk:"aws"`
	Azure  []NodeGroup `tfsdk:"azure"`
	Google []NodeGroup `tfsdk:"google"`
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
						"firewall": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Computed:    true,
										Description: "Type of the firewall rule",
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
						// "node_groups": schema.SingleNestedAttribute{
						// 	Computed: true,
						// 	Attributes: map[string]schema.Attribute{
						// 		"aws": schema.ListNestedAttribute{
						// 			Computed: true,
						// 			NestedObject: schema.NestedAttributeObject{
						// 				Attributes: map[string]schema.Attribute{
						// 					"region": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Region of the AWS node group",
						// 					},
						// 					"availability_zones": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 						Description: "Availability zones of the AWS node group",
						// 					},
						// 					"cidr": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "CIDR of the AWS node group",
						// 					},
						// 					"public_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"private_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"nodes": schema.ListNestedAttribute{
						// 						Computed: true,
						// 						NestedObject: schema.NestedAttributeObject{
						// 							Attributes: map[string]schema.Attribute{
						// 								"display_name": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "Display name of the node",
						// 								},
						// 								"ip_address": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "IP address of the node",
						// 								},
						// 								"is_active": schema.BoolAttribute{
						// 									Computed:    true,
						// 									Description: "Is the node active",
						// 								},
						// 							},
						// 						},
						// 					},
						// 					"node_location": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Node location of the AWS node group",
						// 					},
						// 					"volume_size": schema.Int64Attribute{
						// 						Computed:    true,
						// 						Description: "Volume size of the AWS node group",
						// 					},
						// 					"volume_iops": schema.Int64Attribute{
						// 						Computed:    true,
						// 						Description: "Volume IOPS of the AWS node group",
						// 					},
						// 					"volume_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Volume type of the AWS node group",
						// 					},
						// 					"instance_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Instance type of the AWS node group",
						// 					},
						// 				},
						// 			},
						// 		},
						// 		"azure": schema.ListNestedAttribute{
						// 			Computed: true,
						// 			NestedObject: schema.NestedAttributeObject{
						// 				Attributes: map[string]schema.Attribute{
						// 					"region": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Region of the AWS node group",
						// 					},
						// 					"availability_zones": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 						Description: "Availability zones of the AWS node group",
						// 					},
						// 					"cidr": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "CIDR of the AWS node group",
						// 					},
						// 					"public_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"private_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"nodes": schema.ListNestedAttribute{
						// 						Computed: true,
						// 						NestedObject: schema.NestedAttributeObject{
						// 							Attributes: map[string]schema.Attribute{
						// 								"display_name": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "Display name of the node",
						// 								},
						// 								"ip_address": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "IP address of the node",
						// 								},
						// 								"is_active": schema.BoolAttribute{
						// 									Computed:    true,
						// 									Description: "Is the node active",
						// 								},
						// 							},
						// 						},
						// 					},
						// 					"node_location": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Node location of the AWS node group",
						// 					},
						// 					"volume_size": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Volume size of the AWS node group",
						// 					},
						// 					"volume_iops": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Volume IOPS of the AWS node group",
						// 					},
						// 					"volume_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Volume type of the AWS node group",
						// 					},
						// 					"instance_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Instance type of the AWS node group",
						// 					},
						// 				},
						// 			},
						// 		},
						// 		"google": schema.ListNestedAttribute{
						// 			Computed: true,
						// 			NestedObject: schema.NestedAttributeObject{
						// 				Attributes: map[string]schema.Attribute{
						// 					"region": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Region of the AWS node group",
						// 					},
						// 					"availability_zones": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 						Description: "Availability zones of the AWS node group",
						// 					},
						// 					"cidr": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "CIDR of the AWS node group",
						// 					},
						// 					"public_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"private_subnets": schema.ListAttribute{
						// 						ElementType: types.StringType,
						// 						Computed:    true,
						// 					},
						// 					"nodes": schema.ListNestedAttribute{
						// 						Computed: true,
						// 						NestedObject: schema.NestedAttributeObject{
						// 							Attributes: map[string]schema.Attribute{
						// 								"display_name": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "Display name of the node",
						// 								},
						// 								"ip_address": schema.StringAttribute{
						// 									Computed:    true,
						// 									Description: "IP address of the node",
						// 								},
						// 								"is_active": schema.BoolAttribute{
						// 									Computed:    true,
						// 									Description: "Is the node active",
						// 								},
						// 							},
						// 						},
						// 					},
						// 					"node_location": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Node location of the AWS node group",
						// 					},
						// 					"volume_size": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Volume size of the AWS node group",
						// 					},
						// 					"volume_iops": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Volume IOPS of the AWS node group",
						// 					},
						// 					"volume_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Volume type of the AWS node group",
						// 					},
						// 					"instance_type": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Instance type of the AWS node group",
						// 					},
						// 				},
						// 			},
						// 		},
						// 	},
						// },
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
		var clusterComponents []attr.Value
		tagElements := make(map[string]attr.Value)
		for k, v := range cluster.Aws.Tags {
			tagElements[k] = types.StringValue(v)
		}

		for _, component := range cluster.Database.Components {
			clusterComponents = append(clusterComponents, types.StringValue(component))
		}

		clusterDetails.ID = types.StringValue(cluster.ID)
		clusterDetails.Name = types.StringValue(cluster.Name)
		clusterDetails.CloudAccountID = types.StringValue(cluster.CloudAccountID)
		clusterDetails.CreatedAt = types.StringValue(cluster.CreatedAt.String())
		clusterDetails.Status = types.StringValue(cluster.Status)

		firewallElementTypes := map[string]attr.Type{
			"type": types.StringType,
			"port": types.Int64Type,
			"sources": types.ListType{
				ElemType: types.StringType,
			},
		}

		for _, rule := range cluster.Firewall.Rules {
			var firewallRule FirewallRule
			var firewallSources []attr.Value

			for _, source := range rule.Sources {
				firewallSources = append(firewallSources, types.StringValue(source))
			}

			firewallRule.Type = types.StringValue(rule.Type)
			firewallRule.Port = types.Int64Value(rule.Port)
			firewallRule.Sources, _ = types.ListValue(types.StringType, firewallSources)

			firewallElements := map[string]attr.Value{
				"type": types.StringValue(rule.Type),
				"port": types.Int64Value(rule.Port),
				"sources": firewallRule.Sources,
			}

			firewallObjectValue, _ := types.ObjectValue(firewallElementTypes, firewallElements)

			clusterDetails.Firewall = append(clusterDetails.Firewall, firewallObjectValue)
		}

		// nodeGroupElementTypes := map[string]attr.Type{
		// 	"aws": types.ListType{
		// 		ElemType: types.ObjectType{
		// 			AttrTypes: map[string]attr.Type{
		// 				"region": types.StringType,
		// 				"availability_zones": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"cidr": types.StringType,
		// 				"public_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"private_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"nodes": types.ListType{
		// 					ElemType: types.ObjectType{
		// 						AttrTypes: map[string]attr.Type{
		// 							"display_name": types.StringType,
		// 							"ip_address":   types.StringType,
		// 							"is_active":    types.BoolType,
		// 						},
		// 					},
		// 				},
		// 				"node_location": types.StringType,
		// 				"volume_size":   types.Int64Type,
		// 				"volume_iops":   types.Int64Type,
		// 				"volume_type":   types.StringType,
		// 				"instance_type": types.StringType,
		// 			},
		// 		},
		// 	},

		// 	"azure": types.ListType{
		// 		ElemType: types.ObjectType{
		// 			AttrTypes: map[string]attr.Type{
		// 				"region": types.StringType,
		// 				"availability_zones": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"cidr": types.StringType,
		// 				"public_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"private_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"nodes": types.ListType{
		// 					ElemType: types.ObjectType{
		// 						AttrTypes: map[string]attr.Type{
		// 							"display_name": types.StringType,
		// 							"ip_address":   types.StringType,
		// 							"is_active":    types.BoolType,
		// 						},
		// 					},
		// 				},
		// 				"node_location": types.StringType,
		// 				"volume_size":   types.Int64Type,
		// 				"volume_iops":   types.Int64Type,
		// 				"volume_type":   types.StringType,
		// 				"instance_type": types.StringType,
		// 			},
		// 		},
		// 	},

		// 	"google": types.ListType{
		// 		ElemType: types.ObjectType{
		// 			AttrTypes: map[string]attr.Type{
		// 				"region": types.StringType,
		// 				"availability_zones": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"cidr": types.StringType,
		// 				"public_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"private_subnets": types.ListType{
		// 					ElemType: types.StringType,
		// 				},
		// 				"nodes": types.ListType{
		// 					ElemType: types.ObjectType{
		// 						AttrTypes: map[string]attr.Type{
		// 							"display_name": types.StringType,
		// 							"ip_address":   types.StringType,
		// 							"is_active":    types.BoolType,
		// 						},
		// 					},
		// 				},
		// 				"node_location": types.StringType,
		// 				"volume_size":   types.Int64Type,
		// 				"volume_iops":   types.Int64Type,
		// 				"volume_type":   types.StringType,
		// 				"instance_type": types.StringType,
		// 			},
		// 		},
		// 	},
		// }

		// for _, nodeGroup := range cluster.NodeGroups.Aws {
		// 	var nodeDetails []attr.Value

		// 	for _, node := range nodeGroup.Nodes {
		// 		var nodeDetail ClusterNode

		// 		nodeDetail.DisplayName = types.StringValue(node.DisplayName)
		// 		nodeDetail.IPAddress = types.StringValue(node.IPAddress)
		// 		nodeDetail.IsActive = types.BoolValue(node.IsActive)

		// 		nodeElements := map[string]attr.Value{
		// 			"display_name": types.StringValue(node.DisplayName),
		// 			"ip_address":   types.StringValue(node.IPAddress),
		// 			"is_active":    types.BoolValue(node.IsActive),
		// 		}

		// 		nodeObjectValue, _ := types.ObjectValue(map[string]attr.Type{
		// 			"display_name": types.StringType,
		// 			"ip_address":   types.StringType,
		// 			"is_active":    types.BoolType,
		// 		}, nodeElements)

		// 		nodeDetails = append(nodeDetails, nodeObjectValue)
		// 	}

		// 	var availabilityZones []attr.Value
		// 	for _, availability_zones := range nodeGroup.AvailabilityZones {
		// 		availabilityZones = append(availabilityZones, types.StringValue(availability_zones))
		// 	}

		// 	availabilityZonesList, _ := types.ListValue(types.StringType, availabilityZones)

		// 	var publicSubnets []attr.Value
		// 	for _, public_subnets := range nodeGroup.PublicSubnets {
		// 		publicSubnets = append(publicSubnets, types.StringValue(public_subnets))
		// 	}

		// 	publicSubnetsList, _ := types.ListValue(types.StringType, publicSubnets)

		// 	var privateSubnets []attr.Value
		// 	for _, private_subnets := range nodeGroup.PrivateSubnets {
		// 		privateSubnets = append(privateSubnets, types.StringValue(private_subnets))
		// 	}

		// 	privateSubnetsList, _ := types.ListValue(types.StringType, privateSubnets)

		// 	nodes, _ := types.ListValue(types.ObjectType{
		// 		AttrTypes: map[string]attr.Type{
		// 			"display_name": types.StringType,
		// 			"ip_address":   types.StringType,
		// 			"is_active":    types.BoolType,
		// 		},
		// 	}, nodeDetails)

		// 	nodeGroupElements := map[string]attr.Value{
		// 		"region":            types.StringValue(nodeGroup.Region),
		// 		"availability_zones": availabilityZonesList,
		// 		"cidr":              types.StringValue(nodeGroup.Cidr),
		// 		"public_subnets":    publicSubnetsList,
		// 		"private_subnets":   privateSubnetsList,
		// 		"nodes":            nodes,
		// 		"node_location": types.StringValue(nodeGroup.NodeLocation),
		// 		"volume_size":   types.Int64Value(nodeGroup.VolumeSize),
		// 		"volume_iops":   types.Int64Value(nodeGroup.VolumeIops),
		// 		"volume_type":   types.StringValue(nodeGroup.VolumeType),
		// 		"instance_type": types.StringValue(nodeGroup.InstanceType),
		// 	}

		// 	nodeGroupObjectValue, _ := types.ObjectValue(nodeGroupElementTypes, nodeGroupElements)

		// 	clusterDetails.NodeGroups = nodeGroupObjectValue
		// }

		state.Clusters = append(state.Clusters, clusterDetails)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
