package provider

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &databasesResource{}
	_ resource.ResourceWithConfigure = &databasesResource{}
)

// NewDatabasesResource is a helper function to simplify the provider implementation.
func NewDatabasesResource() resource.Resource {
	return &databasesResource{}
}

// databasesResource is the resource implementation.
type databasesResource struct {
	client *pgEdge.Client
}

type databasesResourceModel struct {
	ID          types.String    `tfsdk:"id"`
	Databases   DatabaseDetails `tfsdk:"databases"`
	LastUpdated types.String    `tfsdk:"last_updated"`
}

// Metadata returns the resource type name.
func (r *databasesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_databases"
}

// Configure adds the provider configured client to the resource.
func (r *databasesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Schema defines the schema for the resource.
func (r *databasesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"databases": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: "ID of the database",
						// Computed:    true,
					},
					"name": schema.StringAttribute{
						Required:    true,
						Description: "Name of the database",
						// Computed:    true,
					},
					"domain": schema.StringAttribute{
						Computed:    true,
						Description: "Domain of the database",
						// Computed:    true,
					},
					"status": schema.StringAttribute{
						Computed:    true,
						Description: "Status of the database",
						// Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: "Created at of the database",
						// Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: "Updated at of the database",
						// Computed:    true,
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
					// 					// Computed:    true,
					// 				},
					// 				"connection": schema.SingleNestedAttribute{
					// 					Computed: true,
					// 					Attributes: map[string]schema.Attribute{
					// 						"database": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Database of the node",
					// 							// Computed:    true,
					// 						},
					// 						"host": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Host of the node",
					// 							// Computed:    true,
					// 						},
					// 						"password": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Password of the node",
					// 							// Computed:    true,
					// 						},
					// 						"port": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Port of the node",
					// 							// Computed:    true,
					// 						},
					// 						"username": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Username of the node",
					// 							// Computed:    true,
					// 						},
					// 					},
					// 				},
					// 				"location": schema.SingleNestedAttribute{
					// 					Computed: true,
					// 					Attributes: map[string]schema.Attribute{
					// 						"code": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Code of the location",
					// 							// Computed:    true,
					// 						},
					// 						"country": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Country of the location",
					// 							// Computed:    true,
					// 						},
					// 						"latitude": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Latitude of the location",
					// 							// Computed:    true,
					// 						},
					// 						"longitude": schema.NumberAttribute{
					// 							Computed:    true,
					// 							Description: "Longitude of the location",
					// 							// Computed:    true,
					// 						},
					// 						"name": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Name of the location",
					// 							// Computed:    true,
					// 						},
					// 						"region": schema.StringAttribute{
					// 							Computed:    true,
					// 							Description: "Region of the location",
					// 							// Computed:    true,
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

// Create creates the resource and sets the initial Terraform state.
func (r *databasesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan databasesResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	databaseName := plan.Databases.Name.ValueString()
	items := &models.DatabaseCreationRequest{
		Name:      plan.Databases.Name.ValueString(),
		ClusterID: r.client.ClusterID,
		// Options:   []string{"install:northwind"}, //database.Options[0]
	}

	order, err := r.client.CreateDatabase(ctx, items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"Could not create order, unexpected error: "+err.Error(),
		)
		return
	}

	if strings.ToLower(databaseName) != types.StringValue(strings.ToLower(order.Name)).ValueString() {
		databaseName = order.Name
	}

	plan.ID = types.StringValue(order.ID.String())
	plan.Databases = DatabaseDetails{
		ID:        types.StringValue(order.ID.String()),
		Name:      types.StringValue(databaseName),
		Domain:    types.StringValue(order.Domain),
		Status:    types.StringValue(order.Status),
		CreatedAt: types.StringValue(order.CreatedAt.String()),
		UpdatedAt: types.StringValue(order.UpdatedAt.String()),
		// Options:  nil, //database.Options[0]
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// for nodeIndex, node := range orderItem.Nodes {
	// 	plan.Databases[orderItemIndex].Nodes[nodeIndex] = Node{
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

// Read refreshes the Terraform state with the latest data.
func (r *databasesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state databasesResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	order, err := r.client.GetDatabase(ctx, strfmt.UUID(state.Databases.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading HashiCups Order",
			"Could not read HashiCups order ID "+state.Databases.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Databases = DatabaseDetails{}
	state.Databases.ID = types.StringValue(order.ID.String())
	state.Databases.Name = types.StringValue(order.Name)
	state.Databases.Status = types.StringValue(order.Status)
	state.Databases.CreatedAt = types.StringValue(order.CreatedAt.String())
	state.Databases.UpdatedAt = types.StringValue(order.UpdatedAt.String())

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *databasesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *databasesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
