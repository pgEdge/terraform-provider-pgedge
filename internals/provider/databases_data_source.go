package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/pgEdge/terraform-provider-pgedge/client"
)

var (
	_ datasource.DataSource              = &databasesDataSource{}
	_ datasource.DataSourceWithConfigure = &databasesDataSource{}
)

func NewDatabasesDataSource() datasource.DataSource {
	return &databasesDataSource{}
}

type databasesDataSource struct {
	client *client.Client
}

func (d *databasesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_databases"
}

func (d *databasesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

type DatabasesDataSourceModel struct {
	Databases []DatabaseDetails `tfsdk:"databases"`
}

type DatabaseDetails struct {
	ID        string `tfsdk:"id"`
	Name      string `tfsdk:"name"`
	Domain    string `tfsdk:"domain"`
	CreatedAt string `tfsdk:"created_at"`
	UpdatedAt string `tfsdk:"updated_at"`
	Status    string `tfsdk:"status"`
	Nodes     []Node `tfsdk:"nodes"`
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

// Schema defines the schema for the data source.
func (d *databasesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	fmt.Println("Inside Schema function")
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"databases": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "ID of the database",
							// Optional:    true,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the database",
							// Optional:    true,
						},
						"domain": schema.StringAttribute{
							Computed:    true,
							Description: "Domain of the database",
							// Optional:    true,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Status of the database",
							// Optional:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "Created at of the database",
							// Optional:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: "Updated at of the database",
							// Optional:    true,
						},
						"nodes": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "Name of the node",
										// Optional:    true,
									},
									"connection": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"database": schema.StringAttribute{
												Computed:    true,
												Description: "Database of the node",
												// Optional:    true,
											},
											"host": schema.StringAttribute{
												Computed:    true,
												Description: "Host of the node",
												// Optional:    true,
											},
											"password": schema.StringAttribute{
												Computed:    true,
												Description: "Password of the node",
												// Optional:    true,
											},
											"port": schema.NumberAttribute{
												Computed:    true,
												Description: "Port of the node",
												// Optional:    true,
											},
											"username": schema.StringAttribute{
												Computed:    true,
												Description: "Username of the node",
												// Optional:    true,
											},
										},
									},
									"location": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"code": schema.StringAttribute{
												Computed:    true,
												Description: "Code of the location",
												// Optional:    true,
											},
											"country": schema.StringAttribute{
												Computed:    true,
												Description: "Country of the location",
												// Optional:    true,
											},
											"latitude": schema.NumberAttribute{
												Computed:    true,
												Description: "Latitude of the location",
												// Optional:    true,
											},
											"longitude": schema.NumberAttribute{
												Computed:    true,
												Description: "Longitude of the location",
												// Optional:    true,
											},
											"name": schema.StringAttribute{
												Computed:    true,
												Description: "Name of the location",
												// Optional:    true,
											},
											"region": schema.StringAttribute{
												Computed:    true,
												Description: "Region of the location",
												// Optional:    true,
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
	fmt.Println("Inside Read function")
	var state DatabasesDataSourceModel

	databases, err := d.client.GetDatabases(ctx)
	fmt.Println("databases: ", databases[0].Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read pgEdge Databases",
			err.Error(),
		)
		return
	}

	// Iterate over each database from the client response
	for _, db := range databases {
		// Create a new DatabaseDetails instance
		var database DatabaseDetails

		// Populate DatabaseDetails fields
		database.ID = db.ID.String()
		database.Name = db.Name
		database.Domain = db.Domain
		database.CreatedAt = db.CreatedAt.String()
		database.UpdatedAt = db.UpdatedAt.String()
		database.Status = db.Status

		// Populate Nodes
		for _, node := range db.Nodes {
			var n Node
			n.Name = node.Name

			// Populate Connection
			n.Connection.Database = node.Connection.Database
			n.Connection.Host = node.Connection.Host
			n.Connection.Password = node.Connection.Password
			n.Connection.Port = node.Connection.Port
			n.Connection.Username = node.Connection.Username

			// Populate Location
			n.Location.Code = node.Location.Code
			n.Location.Country = node.Location.Country
			n.Location.Latitude = node.Location.Latitude
			n.Location.Longitude = node.Location.Longitude
			n.Location.Name = node.Location.Name
			n.Location.Region = node.Location.Region

			// Append the populated Node to the DatabaseDetails Nodes
			database.Nodes = append(database.Nodes, n)
		}

		// Append the populated DatabaseDetails to the state
		state.Databases = append(state.Databases, database)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
