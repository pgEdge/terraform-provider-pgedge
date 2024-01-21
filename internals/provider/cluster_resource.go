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

type clusterResourceModel struct {
	Cluster ClusterDetails `tfsdk:"cluster"`
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
			"cluster": schema.SingleNestedAttribute{
				Optional: true,
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional:    true,
							Description: "ID of the cluster",
						},
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the cluster",
						},
						"cloud_account_id": schema.StringAttribute{
							Optional:    true,
							Description: "Cloud account ID of the cluster",
						},
						"created_at": schema.StringAttribute{
							Optional:    true,
							Description: "Created at of the cluster",
						},
						"status": schema.StringAttribute{
							Optional:    true,
							Description: "Status of the cluster",
						},
						// "aws": schema.SingleNestedAttribute{
						// 	Optional: true,
						// 	Attributes: map[string]schema.Attribute{
						// 		"role_arn": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "Role ARN of the AWS cluster",
						// 		},
						// 		"key_pair": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "Key pair of the AWS cluster",
						// 		},
						// 		"tags": schema.MapAttribute{
						// 			ElementType: types.StringType,
						// 			Optional:    true,
						// 			Description: "Tags of the AWS cluster",
						// 		},
						// 	},
						// },
						// "database": schema.SingleNestedAttribute{
						// 	Optional: true,
						// 	Attributes: map[string]schema.Attribute{
						// 		"pg_version": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "PostgreSQL version of the database",
						// 		},
						// 		"username": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "Username for the database",
						// 		},
						// 		"password": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "Password for the database",
						// 		},
						// 		"name": schema.StringAttribute{
						// 			Optional:    true,
						// 			Description: "Name of the database",
						// 		},
						// 		"port": schema.Float64Attribute{
						// 			Optional:    true,
						// 			Description: "Port of the database",
						// 		},
						// 		"components": schema.ListAttribute{
						// 			ElementType: types.StringType,
						// 			Optional:    true,
						// 			Description: "Components of the database",
						// 		},
						// 		"scripts": schema.SingleNestedAttribute{
						// 			Optional: true,
						// 			Attributes: map[string]schema.Attribute{
						// 				"init": schema.StringAttribute{
						// 					Optional:    true,
						// 					Description: "Init script for the database",
						// 				},
						// 			},
						// 		},
						// 	},
						// },
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
			},
		},
		Description: "Interface with the pgEdge service API for clusters.",
	}
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	fmt.Println("Read---------------------------------------------------")

	var plan clusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	fmt.Println("Read1---------------------------------------------------")

    clusterCreationRequest := &models.ClusterCreationRequest{
        Name: plan.Cluster.Name.ValueString(),
		CloudAccountID: plan.Cluster.CloudAccountID.ValueString(),
		NodeGroups: &models.ClusterCreationRequestNodeGroups{
			Aws: []*models.NodeGroup{
				{
					InstanceType: "t4g.small",
					Region: "us-east-1",
					Nodes: []*models.NodeGroupNodesItems0{
						{
							DisplayName: "Node1",
							IsActive:    true,
						},
					},
				},
			},
			Azure: []*models.NodeGroup{},
			Google: []*models.NodeGroup{},
		},
		// Firewall: ,
    }


    createdCluster, err := r.client.CreateCluster(ctx, clusterCreationRequest)
    if err != nil {
        resp.Diagnostics.AddError("Error creating pgEdge Cluster", "Could not create Cluster, unexpected error: "+err.Error())
        return
    }

	var clusterComponents []types.String
	for _, component := range createdCluster.Database.Components {
		clusterComponents = append(clusterComponents, types.StringValue(component))
	}
	tagElements := make(map[string]attr.Value)
	for k, v := range createdCluster.Aws.Tags {
		tagElements[k] = types.StringValue(v)
	}
	// awsTags, _:= types.MapValue(types.StringType, tagElements)

	plan.Cluster.ID = types.StringValue(createdCluster.ID)
	plan.Cluster.Name = types.StringValue(createdCluster.Name)
	plan.Cluster.CloudAccountID = types.StringValue(createdCluster.CloudAccountID)
	plan.Cluster.CreatedAt = types.StringValue(createdCluster.CreatedAt.String())
	plan.Cluster.Status = types.StringValue(createdCluster.Status)
	// plan.Cluster.Database = Database{
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
	// plan.Cluster.Aws = AWS{
	// 	RoleARN: types.StringValue(createdCluster.Aws.RoleArn),
	// 	KeyPair: types.StringValue(createdCluster.Aws.KeyPair),
	// 	Tags: awsTags,
	// }
	// plan.Cluster.NodeGroups = NodeGroups{
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
	var state clusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	

	

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state clusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCluster(ctx, strfmt.UUID(state.Cluster.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting pgEdge Database",
			"Could not delete Database, unexpected error: "+err.Error(),
		)
		return
	}
}
