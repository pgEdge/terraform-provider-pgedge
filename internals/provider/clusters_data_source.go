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

	Aws        AWS            `tfsdk:"aws"`
	// Database   Database       `tfsdk:"database"`
	// Firewall   []FirewallRule `tfsdk:"firewall"`
	// NodeGroups NodeGroups     `tfsdk:"node_groups"`
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
	Sources []types.String `tfsdk:"sources"`
}

type ClusterNode struct {
	DisplayName types.String `tfsdk:"display_name"`
	IPAddress   types.String `tfsdk:"ip_address"`
	IsActive    types.Bool   `tfsdk:"is_active"`
}

type NodeGroup struct {
	Region            types.String   `tfsdk:"region"`
	AvailabilityZones []types.String `tfsdk:"availability_zones"`
	Cidr              types.String   `tfsdk:"cidr"`
	PublicSubnets     []types.String `tfsdk:"public_subnets"`
	PrivateSubnets    []types.String `tfsdk:"private_subnets"`
	Nodes             []ClusterNode  `tfsdk:"nodes"`
	NodeLocation      types.String   `tfsdk:"node_location"`
	VolumeSize        types.Int64   `tfsdk:"volume_size"`
	VolumeIOPS        types.Int64   `tfsdk:"volume_iops"`
	VolumeType        types.String   `tfsdk:"volume_type"`
	InstanceType      types.String   `tfsdk:"instance_type"`
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
						"aws": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"role_arn": schema.StringAttribute{
									Computed:    true,
									Description: "Role ARN of the AWS cluster",
								},
								"key_pair": schema.StringAttribute{
									Computed:    true,
									Description: "Key pair of the AWS cluster",
								},
								"tags": schema.MapAttribute{
									ElementType: types.StringType,
									Computed:    true,
									Description: "Tags of the AWS cluster",
								},
							},
						},
						// "database": schema.SingleNestedAttribute{
						// 	Computed: true,
						// 	Attributes: map[string]schema.Attribute{
						// 		"pg_version": schema.StringAttribute{
						// 			Computed:    true,
						// 			Description: "PostgreSQL version of the database",
						// 		},
						// 		"username": schema.StringAttribute{
						// 			Computed:    true,
						// 			Description: "Username for the database",
						// 		},
						// 		"password": schema.StringAttribute{
						// 			Computed:    true,
						// 			Description: "Password for the database",
						// 		},
						// 		"name": schema.StringAttribute{
						// 			Computed:    true,
						// 			Description: "Name of the database",
						// 		},
						// 		"port": schema.Float64Attribute{
						// 			Computed:    true,
						// 			Description: "Port of the database",
						// 		},
						// 		"components": schema.ListAttribute{
						// 			ElementType: types.StringType,
						// 			Computed:    true,
						// 			Description: "Components of the database",
						// 		},
						// 		"scripts": schema.SingleNestedAttribute{
						// 			Computed: true,
						// 			Attributes: map[string]schema.Attribute{
						// 				"init": schema.StringAttribute{
						// 					Computed:    true,
						// 					Description: "Init script for the database",
						// 				},
						// 			},
						// 		},
						// 	},
						// },
						// "firewall": schema.ListNestedAttribute{
						// 	Computed: true,
						// 	NestedObject: schema.NestedAttributeObject{
						// 		Attributes: map[string]schema.Attribute{
						// 			"type": schema.StringAttribute{
						// 				Computed:    true,
						// 				Description: "Type of the firewall rule",
						// 			},
						// 			"port": schema.Int64Attribute{
						// 				Computed:    true,
						// 				Description: "Port for the firewall rule",
						// 			},
						// 			"sources": schema.ListAttribute{
						// 				ElementType: types.StringType,
						// 				Computed:    true,
						// 				Description: "Sources for the firewall rule",
						// 			},
						// 		},
						// 	},
						// },
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
		var clusterComponents []types.String
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

		// Populate AWS details
		clusterDetails.Aws.RoleARN = types.StringValue(cluster.Aws.RoleArn)
		clusterDetails.Aws.KeyPair = types.StringValue(cluster.Aws.KeyPair)
		clusterDetails.Aws.Tags, _ = types.MapValue(types.StringType, tagElements)

		// Populate Database details
		// clusterDetails.Database.PGVersion = types.StringValue(cluster.Database.PgVersion)
		// clusterDetails.Database.Username = types.StringValue(cluster.Database.Username)
		// clusterDetails.Database.Password = types.StringValue(cluster.Database.Password)
		// clusterDetails.Database.Name = types.StringValue(cluster.Database.Name)
		// clusterDetails.Database.Port = types.Float64Value(cluster.Database.Port)
		// clusterDetails.Database.Components = clusterComponents
		// clusterDetails.Database.Scripts.Init = types.StringValue(cluster.Database.Scripts.Init)

		// Populate Firewall details
		// for _, rule := range cluster.Firewall.Rules {
		// 	var firewallRule FirewallRule
		// 	for _, source := range rule.Sources {
		// 		firewallRule.Sources = append(firewallRule.Sources, types.StringValue(source))
		// 	}
		// 	firewallRule.Type = types.StringValue(rule.Type)
		// 	firewallRule.Port = types.Int64Value(rule.Port)
		// 	clusterDetails.Firewall = append(clusterDetails.Firewall, firewallRule)
		// }

		// Populate NodeGroups details
		// clusterDetails.NodeGroups.AWS = make([]NodeGroup, len(cluster.NodeGroups.Aws))
		for _, ng := range cluster.NodeGroups.Aws {
			var nodeGroup NodeGroup
			for _,availabilityZone := range ng.AvailabilityZones {
				nodeGroup.AvailabilityZones = append(nodeGroup.AvailabilityZones, types.StringValue(availabilityZone))
			}
			for _, privateSubnet := range ng.PrivateSubnets {
				nodeGroup.PrivateSubnets = append(nodeGroup.PrivateSubnets, types.StringValue(privateSubnet))
			}
			for _, publicSubnet := range ng.PublicSubnets {
				nodeGroup.PublicSubnets = append(nodeGroup.PublicSubnets, types.StringValue(publicSubnet))
			}
			nodeGroup.Region = types.StringValue(ng.Region)
			nodeGroup.Cidr = types.StringValue(ng.Cidr)
			nodeGroup.Nodes = make([]ClusterNode, len(ng.Nodes))
			for j, node := range ng.Nodes {
				var nodeDetails ClusterNode
				nodeDetails.DisplayName = types.StringValue(node.DisplayName)
				nodeDetails.IPAddress = types.StringValue(node.IPAddress)
				nodeDetails.IsActive = types.BoolValue(node.IsActive)
				nodeGroup.Nodes[j] = nodeDetails
			}
			nodeGroup.NodeLocation = types.StringValue(ng.NodeLocation)
			nodeGroup.VolumeSize = types.Int64Value(ng.VolumeSize)
			nodeGroup.VolumeIOPS = types.Int64Value(ng.VolumeIops)
			nodeGroup.VolumeType = types.StringValue(ng.VolumeType)
			nodeGroup.InstanceType = types.StringValue(ng.InstanceType)
			// clusterDetails.NodeGroups.AWS[i] = nodeGroup
		}

		state.Clusters = append(state.Clusters, clusterDetails)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
