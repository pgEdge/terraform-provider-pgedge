package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/models"
)

var (
	_ resource.Resource              = &databaseResource{}
	_ resource.ResourceWithConfigure = &databaseResource{}
)

func NewDatabaseResource() resource.Resource {
	return &databaseResource{}
}

type databaseResource struct {
	client *pgEdge.Client
}

type databaseResourceModel struct {
	// ID          types.String    `tfsdk:"id"`
	Database   DatabaseDetails `tfsdk:"database"`
	// LastUpdated types.String    `tfsdk:"last_updated"`
}

func (r *databaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database"
}

func (r *databaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *databaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// "id": schema.StringAttribute{
			// 	Computed: true,
			// },
			// "last_updated": schema.StringAttribute{
			// 	Computed: true,
			// },
			"database": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: "ID of the database",
					},
					"name": schema.StringAttribute{
						Required:    true,
						Description: "Name of the database",
					},
					"cluster_id": schema.StringAttribute{
						Required:    true,
						Description: "Cluster Id of the database",
					},
					"domain": schema.StringAttribute{
						Computed:    true,
						Description: "Domain of the database",
					},
					"status": schema.StringAttribute{
						Computed:    true,
						Description: "Status of the database",
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: "Created at of the database",
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: "Updated at of the database",
					},
					// 	"options": schema.ListAttribute{
					// 		Optional:    true,
					// 		Description: "Options for creating the database",
					// 		ElementType: types.StringType,
					// },
					// 	"nodes": schema.ListNestedAttribute{
					// 		Computed: true,
					// 		NestedObject: schema.NestedAttributeObject{
					// 			Attributes: map[string]schema.Attribute{
					// 				"name": schema.StringAttribute{
					// 					Computed:    true,
					// 					Description: "Name of the node",
					// 				},
					// 				"connection": schema.SingleNestedAttribute{
					// 					Computed: true,
					// 					Attributes: map[string]schema.Attribute{
					// 						"database": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Database of the node",
					// 						},
					// 						"host": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Host of the node",
					// 						},
					// 						"password": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Password of the node",
					// 						},
					// 						"port": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Port of the node",
					// 						},
					// 						"username": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Username of the node",
					// 						},
					// 					},
					// 				},
					// 				"location": schema.SingleNestedAttribute{
					// 					Computed: true,
					// 					Attributes: map[string]schema.Attribute{
					// 						"code": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Code of the location",
					// 						},
					// 						"country": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Country of the location",
					// 						},
					// 						"latitude": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Latitude of the location",
					// 						},
					// 						"longitude": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Longitude of the location",
					// 						},
					// 						"name": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Name of the location",
					// 						},
					// 						"region": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Region of the location",
					// 						},
					// 					},
					// 				},
					// 			},
					// 		},
					// 	},
				},
			},
		},
		Description: "Interface with the pgEdge service API.",
	}
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan databaseResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	databaseName := plan.Database.Name.ValueString()
	items := &models.DatabaseCreationRequest{
		Name:      plan.Database.Name.ValueString(),
		ClusterID: plan.Database.ClusterID.ValueString(),
		// Options:   []string{"install:northwind"}, //database.Options[0]
	}

	database, err := r.client.CreateDatabase(ctx, items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating database",
			"Could not create database, unexpected error: "+err.Error(),
		)
		return
	}

	if strings.ToLower(databaseName) != types.StringValue(strings.ToLower(database.Name)).ValueString() {
		databaseName = database.Name
	}

	// plan.ID = types.StringValue(database.ID.String())
	plan.Database = DatabaseDetails{
		ID:        types.StringValue(database.ID.String()),
		Name:      types.StringValue(databaseName),
		Domain:    types.StringValue(database.Domain),
		Status:    types.StringValue(database.Status),
		ClusterID: plan.Database.ClusterID,
		CreatedAt: types.StringValue(database.CreatedAt.String()),
		UpdatedAt: types.StringValue(database.UpdatedAt.String()),
		// Options:  nil, //database.Options[0]
	}
	// plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// for nodeIndex, node := range orderItem.Nodes {
	// 	plan.Database[orderItemIndex].Nodes[nodeIndex] = Node{
	// 		Name: node.Name,
	// 		Connection: Connection{
	// 			Database: node.Connection.Database,
	// 			Host:     node.Connection.Host,
	// 			Password: node.Connection.Password,
	// 			Port:     node.Connection.Port,
	// 			Username: node.Connection.Username,
	// 		},
	// 		Location: Location{
	// 			Code:      node.Location.Code,
	// 			Country:   node.Location.Country,
	// 			Latitude:  node.Location.Latitude,
	// 			Longitude: node.Location.Longitude,
	// 			Name:      node.Location.Name,
	// 			Region:    node.Location.Region,
	// 		},
	// 	}
	// }

	// plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	// }

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *databaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state databaseResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	database, err := r.client.GetDatabase(ctx, strfmt.UUID(state.Database.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading pgEdge Database",
			"Could not read pgEdge database ID "+state.Database.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Database = DatabaseDetails{}
	state.Database.ID = types.StringValue(database.ID.String())
	state.Database.Name = types.StringValue(database.Name)
	state.Database.Status = types.StringValue(database.Status)
	state.Database.CreatedAt = types.StringValue(database.CreatedAt.String())
	state.Database.UpdatedAt = types.StringValue(database.UpdatedAt.String())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *databaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state databaseResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    err := r.client.DeleteDatabase(ctx, strfmt.UUID(state.Database.ID.ValueString()))
    if err != nil {
        resp.Diagnostics.AddError(
            "Error Deleting pgEdge Database",
            "Could not delete Database, unexpected error: "+err.Error(),
        )
        return
    }
}
