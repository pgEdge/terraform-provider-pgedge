package provider

import (
	"context"
	"fmt"

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
	// Nodes     []Node `tfsdk:"nodes"`
	// Options   []types.String  `tfsdk:"options"`
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
							Required:    true,
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
						// "options": schema.StringAttribute{
						// 	Optional:    true,
						// 	Description: "Options for creating the database",
						// },
						// "nodes": schema.ListNestedAttribute{
						// 	Computed: true,
						// 	NestedObject: schema.NestedAttributeObject{
						// 		Attributes: map[string]schema.Attribute{
						// 			"name": schema.StringAttribute{
						// 				Computed:    true,
						// 				Description: "Name of the node",
						// 			},
						// 			"connection": schema.SingleNestedAttribute{
						// 				Computed: true,
						// 				Attributes: map[string]schema.Attribute{
						// 					"database": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Database of the node",
						// 					},
						// 					"host": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Host of the node",
						// 					},
						// 					"password": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Password of the node",
						// 					},
						// 					"port": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Port of the node",
						// 					},
						// 					"username": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Username of the node",
						// 					},
						// 				},
						// 			},
						// 			"location": schema.SingleNestedAttribute{
						// 				Computed: true,
						// 				Attributes: map[string]schema.Attribute{
						// 					"code": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Code of the location",
						// 					},
						// 					"country": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Country of the location",
						// 					},
						// 					"latitude": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Latitude of the location",
						// 					},
						// 					"longitude": schema.NumberAttribute{
						// 						Computed:    true,
						// 						Description: "Longitude of the location",
						// 					},
						// 					"name": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Name of the location",
						// 					},
						// 					"region": schema.StringAttribute{
						// 						Computed:    true,
						// 						Description: "Region of the location",
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
		Description: "Interface with the pgEdge service API.",
	}
}

func (d *databasesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DatabasesDataSourceModel

	databases, err := d.client.GetDatabases(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read pgEdge Databases",
			err.Error(),
		)
		return
	}

	for _, db := range databases {
		var database DatabaseDetails

		database.ID = types.StringValue(db.ID.String())
		database.Name = types.StringValue(db.Name)
		database.Domain = types.StringValue(db.Domain)
		database.CreatedAt = types.StringValue(db.CreatedAt.String())
		database.UpdatedAt = types.StringValue(db.UpdatedAt.String())
		database.Status = types.StringValue(db.Status)

		for _, node := range db.Nodes {
			var n Node
			n.Name = node.Name

			n.Connection.Database = node.Connection.Database
			n.Connection.Host = node.Connection.Host
			n.Connection.Password = node.Connection.Password
			n.Connection.Port = node.Connection.Port
			n.Connection.Username = node.Connection.Username

			n.Location.Code = node.Location.Code
			n.Location.Country = node.Location.Country
			n.Location.Latitude = node.Location.Latitude
			n.Location.Longitude = node.Location.Longitude
			n.Location.Name = node.Location.Name
			n.Location.Region = node.Location.Region

			// database.Nodes = append(database.Nodes, n)
		}

		state.Databases = append(state.Databases, database)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
