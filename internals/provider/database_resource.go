package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

			"id": schema.StringAttribute{
				Computed:    true,
				Description: "ID of the database",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-z0-9]+$`),
						"must contain only lowercase alphanumeric characters",
					),
				},
				Description: "Name of the database",
			},
			"cluster_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
			"options": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Options for creating the database",
			},
			"nodes": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:    true,
							Description: "Name of the node",
						},
						"connection": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"database": schema.StringAttribute{
									Optional:    true,
									Description: "Database of the node",
								},
								"host": schema.StringAttribute{
									Optional:    true,
									Description: "Host of the node",
								},
								"password": schema.StringAttribute{
									Optional:    true,
									Description: "Password of the node",
								},
								"port": schema.Int64Attribute{
									Optional:    true,
									Description: "Port of the node",
								},
								"username": schema.StringAttribute{
									Optional:    true,
									Description: "Username of the node",
								},
							},
						},
						"location": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"code": schema.StringAttribute{
									Optional:    true,
									Description: "Code of the location",
								},
								"country": schema.StringAttribute{
									Optional:    true,
									Description: "Country of the location",
								},
								"latitude": schema.Float64Attribute{
									Optional:    true,
									Description: "Latitude of the location",
								},
								"longitude": schema.Float64Attribute{
									Optional:    true,
									Description: "Longitude of the location",
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "Name of the location",
								},
								"region": schema.StringAttribute{
									Optional:    true,
									Description: "Region of the location",
								},
							},
						},
					},
				},
			},
		},
		Description: "Interface with the pgEdge service API.",
	}
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DatabaseDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var databaseOptions []string
	diags = plan.Options.ElementsAs(ctx, &databaseOptions, false)
	resp.Diagnostics.Append(diags...)

	items := &models.DatabaseCreationRequest{
		Name:      plan.Name.ValueString(),
		ClusterID: plan.ClusterID.ValueString(),
		Options:   databaseOptions,
	}

	database, err := r.client.CreateDatabase(ctx, items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating database",
			"Could not create database, unexpected error: "+err.Error(),
		)
		return
	}

	nodeConnectionType := map[string]attr.Type{
		"database": types.StringType,
		"host":     types.StringType,
		"password": types.StringType,
		"port":     types.Int64Type,
		"username": types.StringType,
	}

	NodeLocationType := map[string]attr.Type{
		"code":      types.StringType,
		"country":   types.StringType,
		"latitude":  types.Float64Type,
		"longitude": types.Float64Type,
		"name":      types.StringType,
		"region":    types.StringType,
	}

	nodeType := map[string]attr.Type{
		"name": types.StringType,
		"connection": types.ObjectType{
			AttrTypes: nodeConnectionType,
		},
		"location": types.ObjectType{
			AttrTypes: NodeLocationType,
		},
	}

	var planOptions types.List

	var databaseOptionsAttr []attr.Value

	for _, option := range database.Options {
		databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
	}

	planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
	resp.Diagnostics.Append(diags...)
	plan = DatabaseDetails{
		ID:        types.StringValue(database.ID.String()),
		Name:      types.StringValue(strings.Trim(strings.ToLower(database.Name), " ")),
		Domain:    types.StringValue(database.Domain),
		Status:    types.StringValue(database.Status),
		ClusterID: types.StringValue(database.ClusterID.String()),
		CreatedAt: types.StringValue(database.CreatedAt.String()),
		UpdatedAt: types.StringValue(database.UpdatedAt.String()),
		Options:   planOptions,
	}

	var nodes []attr.Value
	for _, node := range database.Nodes {
		nodeConnectionValue, _ := types.ObjectValue(nodeConnectionType, map[string]attr.Value{
			"database": types.StringValue(node.Connection.Database),
			"host":     types.StringValue(node.Connection.Host),
			"password": types.StringValue(node.Connection.Password),
			"port":     types.Int64Value(node.Connection.Port),
			"username": types.StringValue(node.Connection.Username),
		})

		nodeLocationValue, _ := types.ObjectValue(NodeLocationType, map[string]attr.Value{
			"code":      types.StringValue(node.Location.Code),
			"country":   types.StringValue(node.Location.Country),
			"latitude":  types.Float64Value(node.Location.Latitude),
			"longitude": types.Float64Value(node.Location.Longitude),
			"name":      types.StringValue(node.Location.Name),
			"region":    types.StringValue(node.Location.Region),
		})

		nodeValue := map[string]attr.Value{
			"name":       types.StringValue(node.Name),
			"connection": nodeConnectionValue,
			"location":   nodeLocationValue,
		}

		node, _ := types.ObjectValue(nodeType, nodeValue)
		nodes = append(nodes, node)
	}

	plan.Nodes, _ = types.ListValue(types.ObjectType{
		AttrTypes: nodeType,
	}, nodes)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *databaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DatabaseDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	database, err := r.client.GetDatabase(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading pgEdge Database",
			"Could not read pgEdge database ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state = DatabaseDetails{}
	state.ID = types.StringValue(database.ID.String())
	state.Name = types.StringValue(strings.Trim(strings.ToLower(database.Name), " "))
	state.Status = types.StringValue(database.Status)
	state.CreatedAt = types.StringValue(database.CreatedAt.String())
	state.UpdatedAt = types.StringValue(database.UpdatedAt.String())

	nodeConnectionType := map[string]attr.Type{
		"database": types.StringType,
		"host":     types.StringType,
		"password": types.StringType,
		"port":     types.Int64Type,
		"username": types.StringType,
	}

	NodeLocationType := map[string]attr.Type{
		"code":      types.StringType,
		"country":   types.StringType,
		"latitude":  types.Float64Type,
		"longitude": types.Float64Type,
		"name":      types.StringType,
		"region":    types.StringType,
	}

	nodeType := map[string]attr.Type{
		"name": types.StringType,
		"connection": types.ObjectType{
			AttrTypes: nodeConnectionType,
		},
		"location": types.ObjectType{
			AttrTypes: NodeLocationType,
		},
	}

	var nodes []attr.Value
	for _, node := range database.Nodes {
		nodeConnectionValue, _ := types.ObjectValue(nodeConnectionType, map[string]attr.Value{
			"database": types.StringValue(node.Connection.Database),
			"host":     types.StringValue(node.Connection.Host),
			"password": types.StringValue(node.Connection.Password),
			"port":     types.Int64Value(node.Connection.Port),
			"username": types.StringValue(node.Connection.Username),
		})

		nodeLocationValue, _ := types.ObjectValue(NodeLocationType, map[string]attr.Value{
			"code":      types.StringValue(node.Location.Code),
			"country":   types.StringValue(node.Location.Country),
			"latitude":  types.Float64Value(node.Location.Latitude),
			"longitude": types.Float64Value(node.Location.Longitude),
			"name":      types.StringValue(node.Location.Name),
			"region":    types.StringValue(node.Location.Region),
		})

		nodeValue := map[string]attr.Value{
			"name":       types.StringValue(node.Name),
			"connection": nodeConnectionValue,
			"location":   nodeLocationValue,
		}

		node, _ := types.ObjectValue(nodeType, nodeValue)
		nodes = append(nodes, node)
	}

	var planOptions types.List

	var databaseOptionsAttr []attr.Value

	for _, option := range database.Options {
		databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
	}

	planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
	resp.Diagnostics.Append(diags...)

	state.Options = planOptions

	state.ClusterID = types.StringValue(database.ClusterID.String())

	state.Nodes, _ = types.ListValue(types.ObjectType{
		AttrTypes: nodeType,
	}, nodes)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *databaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DatabaseDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteDatabase(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting pgEdge Database",
			"Could not delete Database, unexpected error: "+err.Error(),
		)
		return
	}
}
