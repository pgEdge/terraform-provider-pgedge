package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
)

var (
	_ datasource.DataSource              = &databasesDataSource{}
	_ datasource.DataSourceWithConfigure = &databasesDataSource{}
)

func NewDatabasesDataSource() datasource.DataSource {
	return &databasesDataSource{}
}

type databasesDataSource struct {
	client *pgEdge.Client
}

func (d *databasesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_databases"
}

func (d *databasesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

type DatabasesDataSourceModel struct {
	Databases []DatabaseDetails `tfsdk:"databases"`
}

type DatabaseDetails struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Domain    types.String `tfsdk:"domain"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
	Status    types.String `tfsdk:"status"`
	ClusterID types.String `tfsdk:"cluster_id"`
	Nodes     types.List   `tfsdk:"nodes"`
	Options   types.List   `tfsdk:"options"`
}

type Node struct {
	Name       string     `tfsdk:"name"`
	Connection Connection `tfsdk:"connection"`
	Location   Location   `tfsdk:"location"`
}

type Connection struct {
	Database string `tfsdk:"database"`
	Host     string `tfsdk:"host"`
	Password string `tfsdk:"password"`
	Port     int64  `tfsdk:"port"`
	Username string `tfsdk:"username"`
}

type Location struct {
	Code      string  `tfsdk:"code"`
	Country   string  `tfsdk:"country"`
	Latitude  float64 `tfsdk:"latitude"`
	Longitude float64 `tfsdk:"longitude"`
	Name      string  `tfsdk:"name"`
	Region    string  `tfsdk:"region"`
}

func (d *databasesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"databases": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "ID of the database",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the database",
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
						"cluster_id": schema.StringAttribute{
							Computed:    true,
							Description: "Updated at of the database",
						},
						"options": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Description: "Options for creating the database",
						},
						"nodes": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "Name of the node",
									},
									"connection": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"database": schema.StringAttribute{
												Computed:    true,
												Description: "Database of the node",
											},
											"host": schema.StringAttribute{
												Computed:    true,
												Description: "Host of the node",
											},
											"password": schema.StringAttribute{
												Computed:    true,
												Description: "Password of the node",
											},
											"port": schema.Int64Attribute{
												Computed:    true,
												Description: "Port of the node",
											},
											"username": schema.StringAttribute{
												Computed:    true,
												Description: "Username of the node",
											},
										},
									},
									"location": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"code": schema.StringAttribute{
												Computed:    true,
												Description: "Code of the location",
											},
											"country": schema.StringAttribute{
												Computed:    true,
												Description: "Country of the location",
											},
											"latitude": schema.Float64Attribute{
												Computed:    true,
												Description: "Latitude of the location",
											},
											"longitude": schema.Float64Attribute{
												Computed:    true,
												Description: "Longitude of the location",
											},
											"name": schema.StringAttribute{
												Computed:    true,
												Description: "Name of the location",
											},
											"region": schema.StringAttribute{
												Computed:    true,
												Description: "Region of the location",
											},
										},
									},
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

func (d *databasesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DatabasesDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	databases, err := d.client.GetDatabases(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read pgEdge Databases",
			err.Error(),
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

	for _, db := range databases {
		var database DatabaseDetails
		var nodes []attr.Value
		database.ID = types.StringValue(db.ID.String())
		database.Name = types.StringValue(strings.Trim(strings.ToLower(db.Name), " "))
		database.Domain = types.StringValue(db.Domain)
		database.CreatedAt = types.StringValue(db.CreatedAt.String())
		database.UpdatedAt = types.StringValue(db.UpdatedAt.String())
		database.Status = types.StringValue(db.Status)

		for _, node := range db.Nodes {
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

		database.Nodes, _ = types.ListValue(types.ObjectType{
			AttrTypes: nodeType,
		}, nodes)

		database.ClusterID = types.StringValue(db.ClusterID.String())

		var planOptions types.List

		var databaseOptionsAttr []attr.Value

		for _, option := range db.Options {
			databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
		}

		planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
		resp.Diagnostics.Append(diags...)

		database.Options = planOptions

		state.Databases = append(state.Databases, database)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
