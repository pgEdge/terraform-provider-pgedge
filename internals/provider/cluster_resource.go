package provider

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
								// "availability_zones": schema.ListAttribute{
								// 	ElementType: types.StringType,
								// 	Optional:    true,
								// 	Description: "Availability zones of the AWS node group",
								// },
								// "cidr": schema.StringAttribute{
								// 	Optional:    true,
								// 	Description: "CIDR of the AWS node group",
								// },
								// "public_subnets": schema.ListAttribute{
								// 	ElementType: types.StringType,
								// 	Optional:    true,
								// },
								// "private_subnets": schema.ListAttribute{
								// 	ElementType: types.StringType,
								// 	Optional:    true,
								// },
								// "nodes": schema.ListNestedAttribute{
								// 	Optional: true,
								// 	NestedObject: schema.NestedAttributeObject{
								// 		Attributes: map[string]schema.Attribute{
								// 			"display_name": schema.StringAttribute{
								// 				Optional:    true,
								// 				Description: "Display name of the node",
								// 			},
								// 			"ip_address": schema.StringAttribute{
								// 				Optional:    true,
								// 				Description: "IP address of the node",
								// 			},
								// 			"is_active": schema.BoolAttribute{
								// 				Optional:    true,
								// 				Description: "Is the node active",
								// 			},
								// 		},
								// 	},
								// },
								// "node_location": schema.StringAttribute{
								// 	Optional:    true,
								// 	Description: "Node location of the AWS node group",
								// },
								// "volume_size": schema.Int64Attribute{
								// 	Optional:    true,
								// 	Description: "Volume size of the AWS node group",
								// },
								// "volume_iops": schema.Int64Attribute{
								// 	Optional:    true,
								// 	Description: "Volume IOPS of the AWS node group",
								// },
								// "volume_type": schema.StringAttribute{
								// 	Optional:    true,
								// 	Description: "Volume type of the AWS node group",
								// },
								// "instance_type": schema.StringAttribute{
								// 	Optional:    true,
								// 	Description: "Instance type of the AWS node group",
								// },
							},
						},
					},
					// "azure": schema.ListNestedAttribute{
					// 	Optional: true,
					// 	NestedObject: schema.NestedAttributeObject{
					// 		Attributes: map[string]schema.Attribute{
					// 			"region": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Region of the AWS node group",
					// 			},
					// 			"availability_zones": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 				Description: "Availability zones of the AWS node group",
					// 			},
					// 			"cidr": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "CIDR of the AWS node group",
					// 			},
					// 			"public_subnets": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 			},
					// 			"private_subnets": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 			},
					// 			"nodes": schema.ListNestedAttribute{
					// 				Optional: true,
					// 				NestedObject: schema.NestedAttributeObject{
					// 					Attributes: map[string]schema.Attribute{
					// 						"display_name": schema.StringAttribute{
					// 							Optional:    true,
					// 							Description: "Display name of the node",
					// 						},
					// 						"ip_address": schema.StringAttribute{
					// 							Optional:    true,
					// 							Description: "IP address of the node",
					// 						},
					// 						"is_active": schema.BoolAttribute{
					// 							Optional:    true,
					// 							Description: "Is the node active",
					// 						},
					// 					},
					// 				},
					// 			},
					// 			"node_location": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Node location of the AWS node group",
					// 			},
					// 			"volume_size": schema.NumberAttribute{
					// 				Optional:    true,
					// 				Description: "Volume size of the AWS node group",
					// 			},
					// 			"volume_iops": schema.NumberAttribute{
					// 				Optional:    true,
					// 				Description: "Volume IOPS of the AWS node group",
					// 			},
					// 			"volume_type": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Volume type of the AWS node group",
					// 			},
					// 			"instance_type": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Instance type of the AWS node group",
					// 			},
					// 		},
					// 	},
					// },
					// "google": schema.ListNestedAttribute{
					// 	Optional: true,
					// 	NestedObject: schema.NestedAttributeObject{
					// 		Attributes: map[string]schema.Attribute{
					// 			"region": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Region of the AWS node group",
					// 			},
					// 			"availability_zones": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 				Description: "Availability zones of the AWS node group",
					// 			},
					// 			"cidr": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "CIDR of the AWS node group",
					// 			},
					// 			"public_subnets": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 			},
					// 			"private_subnets": schema.ListAttribute{
					// 				ElementType: types.StringType,
					// 				Optional:    true,
					// 			},
					// 			"nodes": schema.ListNestedAttribute{
					// 				Optional: true,
					// 				NestedObject: schema.NestedAttributeObject{
					// 					Attributes: map[string]schema.Attribute{
					// 						"display_name": schema.StringAttribute{
					// 							Optional:    true,
					// 							Description: "Display name of the node",
					// 						},
					// 						"ip_address": schema.StringAttribute{
					// 							Optional:    true,
					// 							Description: "IP address of the node",
					// 						},
					// 						"is_active": schema.BoolAttribute{
					// 							Optional:    true,
					// 							Description: "Is the node active",
					// 						},
					// 					},
					// 				},
					// 			},
					// 			"node_location": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Node location of the AWS node group",
					// 			},
					// 			"volume_size": schema.NumberAttribute{
					// 				Optional:    true,
					// 				Description: "Volume size of the AWS node group",
					// 			},
					// 			"volume_iops": schema.NumberAttribute{
					// 				Optional:    true,
					// 				Description: "Volume IOPS of the AWS node group",
					// 			},
					// 			"volume_type": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Volume type of the AWS node group",
					// 			},
					// 			"instance_type": schema.StringAttribute{
					// 				Optional:    true,
					// 				Description: "Instance type of the AWS node group",
					// 			},
					// 		},
					// 	},
					// },
				},
			},
		},
		Description: "Interface with the pgEdge service API for clusters.",
	}
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClusterDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterCreationRequest := &models.ClusterCreationRequest{
		Name:           plan.Name.ValueString(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		NodeGroups: &models.ClusterCreationRequestNodeGroups{
			Aws: []*models.NodeGroup{
				{
					InstanceType: "t4g.small",
					Region:       "us-east-1",
					Nodes: []*models.NodeGroupNodesItems0{
						{
							DisplayName: "Node1",
							IsActive:    true,
						},
					},
				},
			},
			Azure:  []*models.NodeGroup{},
			Google: []*models.NodeGroup{},
		},
		// Firewall: ,
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
		"port": types.Float64Type,
		"sources": types.ListType{
			ElemType: types.StringType,
		},
	}

	for _, firewall := range createdCluster.Firewall.Rules {
		var firewallSources []attr.Value
		firewallType := types.StringValue(firewall.Type)
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}

		firewallSourceList, _ := types.ListValue(types.StringType, firewallSources)

		firewallElements := map[string]attr.Value{
			"type":    firewallType,
			"port":    firewallPort,
			"sources": firewallSourceList,
		}
		firewallObjectValue, _ := types.ObjectValue(firewallElementsType, firewallElements)
		plan.Firewall = append(plan.Firewall, firewallObjectValue)
	}

	nodeGroupTypes := map[string]attr.Type{
		"aws": types.ListType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"region": types.StringType,
				},
			},
		},
	}

	var aws types.List
	for _, nodeGroup := range createdCluster.NodeGroups.Aws {
		AwsItemsValues, _ := types.ObjectValue(map[string]attr.Type{
			"region": types.StringType,
		}, map[string]attr.Value{
			"region": types.StringValue(nodeGroup.Region),
		})

		aws, _ = types.ListValue(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"region": types.StringType,
			},
		}, []attr.Value{
			AwsItemsValues,
		})
	}

	NodeGroupValues := map[string]attr.Value{
		"aws": aws,
	}

	nodeGroupObjectValue, _ := types.ObjectValue(nodeGroupTypes, NodeGroupValues)

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
		"port": types.Float64Type,
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

		firewallSourcesList, _ := types.ListValue(types.StringType, firewallSources)

		firewallElements := map[string]attr.Value{
			"type":    firewallType,
			"port":    firewallPort,
			"sources": firewallSourcesList,
		}
		firewallObjectValue, _ := types.ObjectValue(firewallElementsType, firewallElements)
		state.Firewall = append(state.Firewall, firewallObjectValue)
	}

	nodeGroupTypes := map[string]attr.Type{
		"aws": types.ListType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"region": types.StringType,
				},
			},
		},
	}

	var aws types.List
	for _, nodeGroup := range cluster.NodeGroups.Aws {
		AwsItemsValues, _ := types.ObjectValue(map[string]attr.Type{
			"region": types.StringType,
		}, map[string]attr.Value{
			"region": types.StringValue(nodeGroup.Region),
		})

		aws, _ = types.ListValue(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"region": types.StringType,
			},
		}, []attr.Value{
			AwsItemsValues,
		})
	}

	NodeGroupValues := map[string]attr.Value{
		"aws": aws,
	}

	nodeGroupObjectValue, _ := types.ObjectValue(nodeGroupTypes, NodeGroupValues)

	state.NodeGroups = nodeGroupObjectValue

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
