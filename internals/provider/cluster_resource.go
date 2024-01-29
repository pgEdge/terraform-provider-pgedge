package provider

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/models"
)

var (
	_ resource.Resource              = &clusterResource{}
	_ resource.ResourceWithConfigure = &clusterResource{}
)

func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

type clusterResource struct {
	client *pgEdge.Client
}

func (r *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *clusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (r *clusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
				Required:    true,
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
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional:    true,
							Description: "Type of the firewall rule",
						},
						"port": schema.Int64Attribute{
							Optional:    true,
							Description: "Port for the firewall rule",
						},
						"sources": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Description: "Sources for the firewall rule",
						},
					},
				},
			},
			"node_groups": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"aws": schema.ListNestedAttribute{
						Optional: true,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"region": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Region of the AWS node group",
								},
								"availability_zones": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Description: "Availability zones of the AWS node group",
								},
								"cidr": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "CIDR of the AWS node group",
								},
								"public_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"private_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"nodes": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"display_name": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Display name of the node",
											},
											"ip_address": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "IP address of the node",
											},
											"is_active": schema.BoolAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Is the node active",
											},
										},
									},
								},
								"node_location": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Node location of the AWS node group",
								},
								"volume_size": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume size of the AWS node group",
								},
								"volume_iops": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume IOPS of the AWS node group",
								},
								"volume_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume type of the AWS node group",
								},
								"instance_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Instance type of the AWS node group",
								},
							},
						},
					},
					"azure": schema.ListNestedAttribute{
						Optional: true,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"region": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Region of the AWS node group",
								},
								"availability_zones": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Description: "Availability zones of the AWS node group",
								},
								"cidr": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "CIDR of the AWS node group",
								},
								"public_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"private_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"nodes": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"display_name": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Display name of the node",
											},
											"ip_address": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "IP address of the node",
											},
											"is_active": schema.BoolAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Is the node active",
											},
										},
									},
								},
								"node_location": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Node location of the AWS node group",
								},
								"volume_size": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume size of the AWS node group",
								},
								"volume_iops": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume IOPS of the AWS node group",
								},
								"volume_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume type of the AWS node group",
								},
								"instance_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Instance type of the AWS node group",
								},
							},
						},
					},
					"google": schema.ListNestedAttribute{
						Optional: true,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"region": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Region of the AWS node group",
								},
								"availability_zones": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Description: "Availability zones of the AWS node group",
								},
								"cidr": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "CIDR of the AWS node group",
								},
								"public_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"private_subnets": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"nodes": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"display_name": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Display name of the node",
											},
											"ip_address": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "IP address of the node",
											},
											"is_active": schema.BoolAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Is the node active",
											},
										},
									},
								},
								"node_location": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Node location of the AWS node group",
								},
								"volume_size": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume size of the AWS node group",
								},
								"volume_iops": schema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume IOPS of the AWS node group",
								},
								"volume_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Volume type of the AWS node group",
								},
								"instance_type": schema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Instance type of the AWS node group",
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

func nodeGroupsClusterReq(ctx context.Context, resp *resource.CreateResponse, nodeGroupReq []NodeGroup) (*resource.CreateResponse, []*models.NodeGroup) {

	var (
		nodeGroupReqs []*models.NodeGroup
		nodes         []*models.NodeGroupNodesItems0
		nodeReq       []Nodes

		availability_zones []string
		public_subnets     []string
		private_subnets    []string
	)
	for _, nodeGroup := range nodeGroupReq {
		diags := nodeGroup.Nodes.ElementsAs(ctx, &nodeReq, false)
		resp.Diagnostics.Append(diags...)

		for _, node := range nodeReq {
			nodes = append(nodes, &models.NodeGroupNodesItems0{
				DisplayName: node.DisplayName.ValueString(),
				IPAddress:   node.IPAddress.ValueString(),
				IsActive:    node.IsActive.ValueBool(),
			})
		}

		nodeGroup.AvailabilityZones.ElementsAs(ctx, &availability_zones, false)
		nodeGroup.PublicSubnets.ElementsAs(ctx, &public_subnets, false)
		nodeGroup.PrivateSubnets.ElementsAs(ctx, &private_subnets, false)

		nodeGroupReqs = append(nodeGroupReqs, &models.NodeGroup{
			Region:            nodeGroup.Region.ValueString(),
			Cidr:              nodeGroup.Cidr.ValueString(),
			AvailabilityZones: availability_zones,
			PublicSubnets:     public_subnets,
			PrivateSubnets:    private_subnets,
			Nodes:             nodes,
			NodeLocation:      nodeGroup.NodeLocation.ValueString(),
			VolumeSize:        nodeGroup.VolumeSize.ValueInt64(),
			VolumeIops:        nodeGroup.VolumeIOPS.ValueInt64(),
			VolumeType:        nodeGroup.VolumeType.ValueString(),
			InstanceType:      nodeGroup.InstanceType.ValueString(),
		})

	}

	return resp, nodeGroupReqs
}

func fireWallRulesClusterReq(ctx context.Context, resp *resource.CreateResponse, firewallRuleReq []basetypes.ObjectValue) (*resource.CreateResponse, []*models.ClusterCreationRequestFirewallRulesItems0) {
	var (
		firewallRules    []*models.ClusterCreationRequestFirewallRulesItems0
		firewallRuleType FirewallRule
		sources          []string
	)

	for _, firewallRule := range firewallRuleReq {

		diags := firewallRule.As(ctx, &firewallRuleType, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		diags = firewallRuleType.Sources.ElementsAs(ctx, &sources, false)
		resp.Diagnostics.Append(diags...)
		firewallRules = append(firewallRules, &models.ClusterCreationRequestFirewallRulesItems0{
			Type:    firewallRuleType.Type.ValueString(),
			Port:    firewallRuleType.Port.ValueInt64(),
			Sources: sources,
		})
	}

	return resp, firewallRules
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClusterDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var nodeGroupsReq NodeGroups
	var awsNodeGroupReq []NodeGroup
	var azureNodeGroupReq []NodeGroup
	var googleNodeGroupReq []NodeGroup

	diags = plan.NodeGroups.As(ctx, &nodeGroupsReq, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	nodeGroupsReq.AWS.ElementsAs(ctx, &awsNodeGroupReq, false)
	nodeGroupsReq.Azure.ElementsAs(ctx, &azureNodeGroupReq, false)
	nodeGroupsReq.Google.ElementsAs(ctx, &googleNodeGroupReq, false)

	resp, firewallRules := fireWallRulesClusterReq(ctx, resp, plan.Firewall)
	if resp.Diagnostics.HasError() {
		return
	}

	resp, awsReq := nodeGroupsClusterReq(ctx, resp, awsNodeGroupReq)
	if resp.Diagnostics.HasError() {
		return
	}
	resp, azureReq := nodeGroupsClusterReq(ctx, resp, azureNodeGroupReq)
	if resp.Diagnostics.HasError() {
		return
	}
	resp, googleReq := nodeGroupsClusterReq(ctx, resp, googleNodeGroupReq)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterCreationRequest := &models.ClusterCreationRequest{
		Name:           plan.Name.ValueString(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		NodeGroups: &models.ClusterCreationRequestNodeGroups{
			Aws:    awsReq,
			Azure:  azureReq,
			Google: googleReq,
		},
		Firewall: &models.ClusterCreationRequestFirewall{
			Rules: firewallRules,
		},
	}

	createdCluster, err := r.client.CreateCluster(ctx, clusterCreationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error creating pgEdge Cluster", "Could not create Cluster, unexpected error: "+err.Error())
		return
	}

	plan.ID = types.StringValue(createdCluster.ID)
	plan.Name = types.StringValue(createdCluster.Name)
	plan.CloudAccountID = types.StringValue(createdCluster.CloudAccountID)
	plan.CreatedAt = types.StringValue(createdCluster.CreatedAt.String())
	plan.Status = types.StringValue(createdCluster.Status)

	firewallElementsType := map[string]attr.Type{
		"type": types.StringType,
		"port": types.Int64Type,
		"sources": types.ListType{
			ElemType: types.StringType,
		},
	}

	var firewallResp []types.Object
	for _, firewall := range createdCluster.Firewall.Rules {
		var firewallSources []attr.Value
		firewallType := types.StringValue(firewall.Type)
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}

		firewallSourceList, diags := types.ListValue(types.StringType, firewallSources)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		firewallElements := map[string]attr.Value{
			"type":    firewallType,
			"port":    firewallPort,
			"sources": firewallSourceList,
		}
		firewallObjectValue, diags := types.ObjectValue(firewallElementsType, firewallElements)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		firewallResp = append(firewallResp, firewallObjectValue)
	}

	plan.Firewall = firewallResp

	var aws types.List
	var awsItems []attr.Value
	for _, nodeGroup := range createdCluster.NodeGroups.Aws {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)

			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}

			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		AwsItemsValues, _ := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})

		awsItems = append(awsItems, AwsItemsValues)
	}

	aws, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, awsItems)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var azure types.List
	var azureItems []attr.Value
	for _, nodeGroup := range createdCluster.NodeGroups.Azure {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)
			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}
			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		AzureItemsValues, diags := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		azureItems = append(azureItems, AzureItemsValues)
	}

	azure, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, azureItems)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	var google types.List
	var googleItems []attr.Value
	for _, nodeGroup := range createdCluster.NodeGroups.Google {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)
			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}
			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		GoogleItemsValues, diags := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		googleItems = append(azureItems, GoogleItemsValues)
	}

	google, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, googleItems)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	NodeGroupsValues := map[string]attr.Value{
		"aws":    aws,
		"azure":  azure,
		"google": google,
	}

	nodeGroupObjectValue, diags := types.ObjectValue(NodeGroupsTypes, NodeGroupsValues)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	plan.NodeGroups = nodeGroupObjectValue

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ClusterDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cluster, err := r.client.GetCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading pgEdge Cluster",
			"Could not read Cluster, unexpected error: "+err.Error(),
		)
		return
	}

	var clusterComponents []attr.Value
	for _, component := range cluster.Database.Components {
		clusterComponents = append(clusterComponents, types.StringValue(component))
	}
	tagElements := make(map[string]attr.Value)
	for k, v := range cluster.Aws.Tags {
		tagElements[k] = types.StringValue(v)
	}

	state.ID = types.StringValue(cluster.ID)
	state.Name = types.StringValue(cluster.Name)
	state.CloudAccountID = types.StringValue(cluster.CloudAccountID)
	state.CreatedAt = types.StringValue(cluster.CreatedAt.String())
	state.Status = types.StringValue(cluster.Status)

	firewallElementsType := map[string]attr.Type{
		"type": types.StringType,
		"port": types.Int64Type,
		"sources": types.ListType{
			ElemType: types.StringType,
		},
	}

	for _, firewall := range cluster.Firewall.Rules {
		var firewallSources []attr.Value
		firewallType := types.StringValue(firewall.Type)
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}

		firewallSourcesList, diags := types.ListValue(types.StringType, firewallSources)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		firewallElements := map[string]attr.Value{
			"type":    firewallType,
			"port":    firewallPort,
			"sources": firewallSourcesList,
		}
		firewallObjectValue, diags := types.ObjectValue(firewallElementsType, firewallElements)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		state.Firewall = append(state.Firewall, firewallObjectValue)
	}

	var aws types.List
	var awsItems []attr.Value
	for _, nodeGroup := range cluster.NodeGroups.Aws {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)
			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}
			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		AwsItemsValues, diags := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		awsItems = append(awsItems, AwsItemsValues)
	}

	aws, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, awsItems)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	var azure types.List
	var azureItems []attr.Value
	for _, nodeGroup := range cluster.NodeGroups.Azure {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)
			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}
			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		AzureItemsValues, diags := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		azureItems = append(azureItems, AzureItemsValues)
	}

	azure, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, azureItems)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var google types.List
	var googleItems []attr.Value
	for _, nodeGroup := range cluster.NodeGroups.Google {
		var availabilityZones []attr.Value

		for _, zone := range nodeGroup.AvailabilityZones {
			availabilityZones = append(availabilityZones, types.StringValue(zone))
		}

		allAvailabilityZones, diags := types.ListValue(types.StringType, availabilityZones)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var publicSubnets []attr.Value

		for _, subnet := range nodeGroup.PublicSubnets {
			publicSubnets = append(publicSubnets, types.StringValue(subnet))
		}

		allPublicSubnets, diags := types.ListValue(types.StringType, publicSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var privateSubnets []attr.Value

		for _, subnet := range nodeGroup.PrivateSubnets {
			privateSubnets = append(privateSubnets, types.StringValue(subnet))
		}

		allPrivateSubnets, diags := types.ListValue(types.StringType, privateSubnets)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		var nodes []attr.Value

		for _, node := range nodeGroup.Nodes {
			nodeDetails := map[string]attr.Value{
				"display_name": types.StringValue(node.DisplayName),
				"ip_address":   types.StringValue(node.IPAddress),
				"is_active":    types.BoolValue(node.IsActive),
			}
			nodeObjectValue, diags := types.ObjectValue(NodesNodeGroupType, nodeDetails)
			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				return
			}
			nodes = append(nodes, nodeObjectValue)
		}

		allNodes, diags := types.ListValue(types.ObjectType{
			AttrTypes: NodesNodeGroupType,
		}, nodes)
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		GoogleItemsValues, diags := types.ObjectValue(NodeGroupTypes, map[string]attr.Value{
			"region":             types.StringValue(nodeGroup.Region),
			"cidr":               types.StringValue(nodeGroup.Cidr),
			"availability_zones": allAvailabilityZones,
			"public_subnets":     allPublicSubnets,
			"private_subnets":    allPrivateSubnets,
			"nodes":              allNodes,
			"node_location":      types.StringValue(nodeGroup.NodeLocation),
			"volume_size":        types.Int64Value(nodeGroup.VolumeSize),
			"volume_iops":        types.Int64Value(nodeGroup.VolumeIops),
			"volume_type":        types.StringValue(nodeGroup.VolumeType),
			"instance_type":      types.StringValue(nodeGroup.InstanceType),
		})
		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		googleItems = append(azureItems, GoogleItemsValues)
	}

	google, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeGroupTypes,
	}, googleItems)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	NodeGroupsValues := map[string]attr.Value{
		"aws":    aws,
		"azure":  azure,
		"google": google,
	}

	nodeGroupsObjectValue, diags := types.ObjectValue(NodeGroupsTypes, NodeGroupsValues)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	state.NodeGroups = nodeGroupsObjectValue

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ClusterDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting pgEdge Database",
			"Could not delete Database, unexpected error: "+err.Error(),
		)
		return
	}
}
