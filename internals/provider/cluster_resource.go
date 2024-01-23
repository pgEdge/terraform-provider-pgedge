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
			"aws": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"role_arn": schema.StringAttribute{
						Optional: true,

						Description: "Role ARN of the AWS cluster",
					},
					"key_pair": schema.StringAttribute{
						Optional: true,

						Description: "Key pair of the AWS cluster",
					},
					"tags": schema.MapAttribute{
						ElementType: types.StringType,
						Optional:    true,

						Description: "Tags of the AWS cluster",
					},
				},
			},
			"database": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"pg_version": schema.StringAttribute{
						Optional:    true,
						Description: "PostgreSQL version of the database",
					},
					"username": schema.StringAttribute{
						Optional:    true,
						Description: "Username for the database",
					},
					"password": schema.StringAttribute{
						Optional:    true,
						Description: "Password for the database",
					},
					"name": schema.StringAttribute{
						Optional:    true,
						Description: "Name of the database",
					},
					"port": schema.Float64Attribute{
						Optional:    true,
						Description: "Port of the database",
					},
					"components": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Description: "Components of the database",
					},
					"scripts": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"init": schema.StringAttribute{
								Optional:    true,
								Description: "Init script for the database",
							},
						},
					},
				},
			},
			// "firewall": schema.ListNestedAttribute{
			// 	Optional: true,
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"type": schema.StringAttribute{
			// 				Optional:    true,
			// 				Description: "Type of the firewall rule",
			// 			},
			// 			"port": schema.Int64Attribute{
			// 				Optional:    true,
			// 				Description: "Port for the firewall rule",
			// 			},
			// 			"sources": schema.ListAttribute{
			// 				ElementType: types.StringType,
			// 				Optional:    true,
			// 				Description: "Sources for the firewall rule",
			// 			},
			// 		},
			// 	},
			// },
			// "node_groups": schema.SingleNestedAttribute{
			// 	Optional: true,
			// 	Attributes: map[string]schema.Attribute{
			// 		"aws": schema.ListNestedAttribute{
			// 			Optional: true,
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"region": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Region of the AWS node group",
			// 					},
			// 					"availability_zones": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 						Description: "Availability zones of the AWS node group",
			// 					},
			// 					"cidr": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "CIDR of the AWS node group",
			// 					},
			// 					"public_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"private_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"nodes": schema.ListNestedAttribute{
			// 						Optional: true,
			// 						NestedObject: schema.NestedAttributeObject{
			// 							Attributes: map[string]schema.Attribute{
			// 								"display_name": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "Display name of the node",
			// 								},
			// 								"ip_address": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "IP address of the node",
			// 								},
			// 								"is_active": schema.BoolAttribute{
			// 									Optional:    true,
			// 									Description: "Is the node active",
			// 								},
			// 							},
			// 						},
			// 					},
			// 					"node_location": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Node location of the AWS node group",
			// 					},
			// 					"volume_size": schema.Int64Attribute{
			// 						Optional:    true,
			// 						Description: "Volume size of the AWS node group",
			// 					},
			// 					"volume_iops": schema.Int64Attribute{
			// 						Optional:    true,
			// 						Description: "Volume IOPS of the AWS node group",
			// 					},
			// 					"volume_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Volume type of the AWS node group",
			// 					},
			// 					"instance_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Instance type of the AWS node group",
			// 					},
			// 				},
			// 			},
			// 		},
			// 		"azure": schema.ListNestedAttribute{
			// 			Optional: true,
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"region": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Region of the AWS node group",
			// 					},
			// 					"availability_zones": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 						Description: "Availability zones of the AWS node group",
			// 					},
			// 					"cidr": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "CIDR of the AWS node group",
			// 					},
			// 					"public_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"private_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"nodes": schema.ListNestedAttribute{
			// 						Optional: true,
			// 						NestedObject: schema.NestedAttributeObject{
			// 							Attributes: map[string]schema.Attribute{
			// 								"display_name": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "Display name of the node",
			// 								},
			// 								"ip_address": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "IP address of the node",
			// 								},
			// 								"is_active": schema.BoolAttribute{
			// 									Optional:    true,
			// 									Description: "Is the node active",
			// 								},
			// 							},
			// 						},
			// 					},
			// 					"node_location": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Node location of the AWS node group",
			// 					},
			// 					"volume_size": schema.NumberAttribute{
			// 						Optional:    true,
			// 						Description: "Volume size of the AWS node group",
			// 					},
			// 					"volume_iops": schema.NumberAttribute{
			// 						Optional:    true,
			// 						Description: "Volume IOPS of the AWS node group",
			// 					},
			// 					"volume_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Volume type of the AWS node group",
			// 					},
			// 					"instance_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Instance type of the AWS node group",
			// 					},
			// 				},
			// 			},
			// 		},
			// 		"google": schema.ListNestedAttribute{
			// 			Optional: true,
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"region": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Region of the AWS node group",
			// 					},
			// 					"availability_zones": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 						Description: "Availability zones of the AWS node group",
			// 					},
			// 					"cidr": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "CIDR of the AWS node group",
			// 					},
			// 					"public_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"private_subnets": schema.ListAttribute{
			// 						ElementType: types.StringType,
			// 						Optional:    true,
			// 					},
			// 					"nodes": schema.ListNestedAttribute{
			// 						Optional: true,
			// 						NestedObject: schema.NestedAttributeObject{
			// 							Attributes: map[string]schema.Attribute{
			// 								"display_name": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "Display name of the node",
			// 								},
			// 								"ip_address": schema.StringAttribute{
			// 									Optional:    true,
			// 									Description: "IP address of the node",
			// 								},
			// 								"is_active": schema.BoolAttribute{
			// 									Optional:    true,
			// 									Description: "Is the node active",
			// 								},
			// 							},
			// 						},
			// 					},
			// 					"node_location": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Node location of the AWS node group",
			// 					},
			// 					"volume_size": schema.NumberAttribute{
			// 						Optional:    true,
			// 						Description: "Volume size of the AWS node group",
			// 					},
			// 					"volume_iops": schema.NumberAttribute{
			// 						Optional:    true,
			// 						Description: "Volume IOPS of the AWS node group",
			// 					},
			// 					"volume_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Volume type of the AWS node group",
			// 					},
			// 					"instance_type": schema.StringAttribute{
			// 						Optional:    true,
			// 						Description: "Instance type of the AWS node group",
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
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

	var clusterComponents []attr.Value
	for _, component := range createdCluster.Database.Components {
		clusterComponents = append(clusterComponents, types.StringValue(component))
	}
	tagElements := make(map[string]attr.Value)
	for k, v := range createdCluster.Aws.Tags {
		tagElements[k] = types.StringValue(v)
	}

	plan.ID = types.StringValue(createdCluster.ID)
	plan.Name = types.StringValue(createdCluster.Name)
	plan.CloudAccountID = types.StringValue(createdCluster.CloudAccountID)
	plan.CreatedAt = types.StringValue(createdCluster.CreatedAt.String())
	plan.Status = types.StringValue(createdCluster.Status)
	databaseElementTypes := map[string]attr.Type{
		"pg_version": types.StringType,
		"username":   types.StringType,
		"password":   types.StringType,
		"name":       types.StringType,
		"port":       types.Float64Type,
		"components": types.ListType{
			ElemType: types.StringType,
		},
		"scripts":    types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"init": types.StringType,
			},
		},
	}

	databaseComponent, _ := types.ListValue(types.StringType, clusterComponents)
	databaseScripts, _ := types.ObjectValue(map[string]attr.Type{
		"init": types.StringType,
	}, map[string]attr.Value{
		"init": types.StringValue(createdCluster.Database.Scripts.Init),
	})

	databaseElements := map[string]attr.Value{
		"pg_version": types.StringValue(createdCluster.Database.PgVersion),
		"username":   types.StringValue(createdCluster.Database.Username),
		"password":   types.StringValue(createdCluster.Database.Password),
		"name":       types.StringValue(createdCluster.Database.Name),
		"port":       types.Float64Value(createdCluster.Database.Port),
		"components": databaseComponent,
		"scripts":    databaseScripts,
	}

	databaseObjectValue, _ := types.ObjectValue(databaseElementTypes, databaseElements)

	plan.Database = databaseObjectValue
	// plan.Database = Database{
	// 	PGVersion: types.StringValue(createdCluster.Database.PgVersion),
	// 	Username: types.StringValue(createdCluster.Database.Username),
	// 	Password: types.StringValue(createdCluster.Database.Password),
	// 	Name: types.StringValue(createdCluster.Database.Name),
	// 	Port: types.Float64Value(createdCluster.Database.Port),
	// 	Components: clusterComponents,
	// 	Scripts: DatabaseScripts{
	// 		Init: types.StringValue(createdCluster.Database.Scripts.Init),
	// 	},
	// }

	awsElementTypes := map[string]attr.Type{
		"role_arn": types.StringType,
		"key_pair": types.StringType,
		"tags": types.MapType{
			ElemType: types.StringType,
		},
	}

	tags, _ := types.MapValue(types.StringType, tagElements)

	awsElements := map[string]attr.Value{
		"role_arn": types.StringValue(createdCluster.Aws.RoleArn),
		"key_pair": types.StringValue(createdCluster.Aws.KeyPair),
		"tags":     tags,
	}

	awsObjectValue, _ := types.ObjectValue(awsElementTypes, awsElements)

	plan.Aws = awsObjectValue
	// plan.NodeGroups = NodeGroups{
	// 	AWS: []NodeGroup{},
	// 	Azure: []NodeGroup{},
	// 	Google: []NodeGroup{},
	// }

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
	// awsTags, _:= types.MapValue(types.StringType, tagElements)

	state.ID = types.StringValue(cluster.ID)
	state.Name = types.StringValue(cluster.Name)
	state.CloudAccountID = types.StringValue(cluster.CloudAccountID)
	state.CreatedAt = types.StringValue(cluster.CreatedAt.String())
	state.Status = types.StringValue(cluster.Status)

	databaseElementTypes := map[string]attr.Type{
		"pg_version": types.StringType,
		"username":   types.StringType,
		"password":   types.StringType,
		"name":       types.StringType,
		"port":       types.Float64Type,
		"components": types.ListType{
			ElemType: types.StringType,
		},
		"scripts":    types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"init": types.StringType,
			},
		},
	}

	databaseComponent, _ := types.ListValue(types.StringType, clusterComponents)
	databaseScripts, _ := types.ObjectValue(map[string]attr.Type{
		"init": types.StringType,
	}, map[string]attr.Value{
		"init": types.StringValue(cluster.Database.Scripts.Init),
	})

	databaseElements := map[string]attr.Value{
		"pg_version": types.StringValue(cluster.Database.PgVersion),
		"username":   types.StringValue(cluster.Database.Username),
		"password":   types.StringValue(cluster.Database.Password),
		"name":       types.StringValue(cluster.Database.Name),
		"port":       types.Float64Value(cluster.Database.Port),
		"components": databaseComponent,
		"scripts":    databaseScripts,
	}

	databaseObjectValue, _ := types.ObjectValue(databaseElementTypes, databaseElements)

	state.Database = databaseObjectValue

	// state.Database = Database{
	// 	PGVersion: types.StringValue(cluster.Database.PgVersion),
	// 	Username: types.StringValue(cluster.Database.Username),
	// 	Password: types.StringValue(cluster.Database.Password),
	// 	Name: types.StringValue(cluster.Database.Name),
	// 	Port: types.Float64Value(cluster.Database.Port),
	// 	Components: clusterComponents,
	// 	Scripts: DatabaseScripts{
	// 		Init: types.StringValue(cluster.Database.Scripts.Init),
	// 	},
	// }

	awsElementTypes := map[string]attr.Type{
		"role_arn": types.StringType,
		"key_pair": types.StringType,
		"tags": types.MapType{
			ElemType: types.StringType,
		},
	}

	tags, _ := types.MapValue(types.StringType, tagElements)

	awsElements := map[string]attr.Value{
		"role_arn": types.StringValue(cluster.Aws.RoleArn),
		"key_pair": types.StringValue(cluster.Aws.KeyPair),
		"tags":     tags,
	}

	awsObjectValue, _ := types.ObjectValue(awsElementTypes, awsElements)

	state.Aws = awsObjectValue
	// state.NodeGroups = NodeGroups{
	// 	AWS: []NodeGroup{},
	// 	Azure: []NodeGroup{},
	// 	Google: []NodeGroup{},
	// }

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
