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
	// Backups     types.Object `tfsdk:"backups"`
	// Components  types.List   `tfsdk:"components"`
	// Extensions  types.Object `tfsdk:"extensions"`
	PgVersion types.String `tfsdk:"pg_version"`
	// Roles       types.List   `tfsdk:"roles"`
	// Tables      types.List   `tfsdk:"tables"`
	StorageUsed types.Int64 `tfsdk:"storage_used"`
}

type Node struct {
	Name                string                  `tfsdk:"name"`
	Connection          Connection              `tfsdk:"connection"`
	Location            Location                `tfsdk:"location"`
	// DistanceMeasurement NodeDistanceMeasurement `tfsdk:"distance_measurement"`
	// Region              NodeRegion              `tfsdk:"region"`
	// Extensions          NodeExtensions          `tfsdk:"extensions"`
}

type NodeExtensions struct {
	Errors    *NodeExtensionsErrors `tfsdk:"errors"`
	Installed []string              `tfsdk:"installed"`
}

type NodeExtensionsErrors struct {
	Anim9ef   string `tfsdk:"anim9ef"`
	Enim3b    string `tfsdk:"enim3b"`
	Laborumd  string `tfsdk:"laborum_d"`
	Mollit267 string `tfsdk:"mollit267"`
}

type NodeDistanceMeasurement struct {
	Distance      float64 `tfsdk:"distance"`
	FromLatitude  float64 `tfsdk:"from_latitude"`
	FromLongitude float64 `tfsdk:"from_longitude"`
	Unit          string  `tfsdk:"unit"`
}

type NodeRegion struct {
	Active            bool     `tfsdk:"active"`
	AvailabilityZones []string `tfsdk:"availability_zones"`
	Cloud             string   `tfsdk:"cloud"`
	Code              string   `tfsdk:"code"`
	Name              string   `tfsdk:"name"`
	Parent            string   `tfsdk:"parent"`
}

type Connection struct {
	Database          string `tfsdk:"database"`
	Host              string `tfsdk:"host"`
	Password          string `tfsdk:"password"`
	Port              int64  `tfsdk:"port"`
	Username          string `tfsdk:"username"`
	ExternalIPAddress string `tfsdk:"external_ip_address"`
	InternalIPAddress string `tfsdk:"internal_ip_address"`
	InternalHost      string `tfsdk:"internal_host"`
}

type Location struct {
	Code       string  `tfsdk:"code"`
	Country    string  `tfsdk:"country"`
	Latitude   float64 `tfsdk:"latitude"`
	Longitude  float64 `tfsdk:"longitude"`
	Name       string  `tfsdk:"name"`
	Region     string  `tfsdk:"region"`
	Timezone   string  `tfsdk:"timezone"`
	RegionCode string  `tfsdk:"region_code"`
	PostalCode string  `tfsdk:"postal_code"`
	MetroCode  string  `tfsdk:"metro_code"`
	City       string  `tfsdk:"city"`
}

var NodeConnectionType = map[string]attr.Type{
	"database":            types.StringType,
	"host":                types.StringType,
	"password":            types.StringType,
	"port":                types.Int64Type,
	"username":            types.StringType,
	"external_ip_address": types.StringType,
	"internal_ip_address": types.StringType,
	"internal_host":       types.StringType,
}

var NodeLocationType = map[string]attr.Type{
	"code":        types.StringType,
	"country":     types.StringType,
	"latitude":    types.Float64Type,
	"longitude":   types.Float64Type,
	"name":        types.StringType,
	"region":      types.StringType,
	"timezone":    types.StringType,
	"region_code": types.StringType,
	"postal_code": types.StringType,
	"metro_code":  types.StringType,
	"city":        types.StringType,
}

var NodeRegionType = map[string]attr.Type{
	"active": types.BoolType,
	"availability_zones": types.ListType{
		ElemType: types.StringType,
	},
	"cloud":  types.StringType,
	"code":   types.StringType,
	"name":   types.StringType,
	"parent": types.StringType,
}

var NodeDistanceMeasurementType = map[string]attr.Type{
	"distance":       types.Float64Type,
	"from_latitude":  types.Float64Type,
	"from_longitude": types.Float64Type,
	"unit":           types.StringType,
}

var NodeExtensionsType = map[string]attr.Type{
	"errors": types.ObjectType{
		AttrTypes: NodeExtensionsErrorsType,
	},
	"installed": types.ListType{
		ElemType: types.StringType,
	},
}

var NodeExtensionsErrorsType = map[string]attr.Type{
	"anim9ef":   types.StringType,
	"enim3b":    types.StringType,
	"laborumd":  types.StringType,
	"mollit267": types.StringType,
}

var NodeType = map[string]attr.Type{
	"name": types.StringType,
	"connection": types.ObjectType{
		AttrTypes: NodeConnectionType,
	},
	"location": types.ObjectType{
		AttrTypes: NodeLocationType,
	},
	// "region": types.ObjectType{
	// 	AttrTypes: NodeRegionType,
	// },
	// "distance_measurement": types.ObjectType{
	// 	AttrTypes: NodeDistanceMeasurementType,
	// },
	// "extensions": types.ObjectType{
	// 	AttrTypes: NodeExtensionsType,
	// },
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
							Computed:    true,
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
												Sensitive:   true,
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
											"external_ip_address": schema.StringAttribute{
												Computed:    true,
												Description: "External IP of the node",
											},
											"internal_ip_address": schema.StringAttribute{
												Computed:    true,
												Description: "Internal IP of the node",
											},
											"internal_host": schema.StringAttribute{
												Computed:    true,
												Description: "Internal Host of the node",
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
											"region_code": schema.StringAttribute{
												Computed:    true,
												Description: "Region code of the location",
											},
											"timezone": schema.StringAttribute{
												Computed:    true,
												Description: "Timezone of the location",
											},
											"postal_code": schema.StringAttribute{
												Computed:    true,
												Description: "Postal code of the location",
											},
											"metro_code": schema.StringAttribute{
												Computed:    true,
												Description: "Metro code of the location",
											},
											"city": schema.StringAttribute{
												Computed:    true,
												Description: "City of the location",
											},
										},
									},
								},
							},
						},
						"pg_version": schema.StringAttribute{
							Computed:    true,
							Description: "Postgres version of the database",
						},
						"storage_used": schema.Int64Attribute{
							Computed:    true,
							Description: "Storage used of the database",
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

	for _, db := range databases {
		var database DatabaseDetails
		var nodes []attr.Value
		database.ID = types.StringValue(db.ID.String())
		database.Name = types.StringValue(strings.Trim(strings.ToLower(db.Name), " "))
		database.Domain = types.StringValue(db.Domain)
		database.CreatedAt = types.StringValue(db.CreatedAt)
		database.UpdatedAt = types.StringValue(db.UpdatedAt)
		database.Status = types.StringValue(db.Status)
		database.ClusterID = types.StringValue(db.ClusterID.String())

		for _, node := range db.Nodes {
			nodeConnectionValue, _ := types.ObjectValue(NodeConnectionType, map[string]attr.Value{
				"database": types.StringValue(node.Connection.Database),
				"host":     types.StringValue(node.Connection.Host),
				"password": types.StringValue(node.Connection.Password),
				"port":     types.Int64Value(node.Connection.Port),
				"username": types.StringValue(node.Connection.Username),
				"external_ip_address": types.StringValue(node.Connection.ExternalIPAddress),
				"internal_ip_address": types.StringValue(node.Connection.InternalIPAddress),
				"internal_host":       types.StringValue(node.Connection.InternalHost),
			})

			nodeLocationValue, _ := types.ObjectValue(NodeLocationType, map[string]attr.Value{
				"code":      types.StringValue(node.Location.Code),
				"country":   types.StringValue(node.Location.Country),
				"latitude":  types.Float64Value(node.Location.Latitude),
				"longitude": types.Float64Value(node.Location.Longitude),
				"name":      types.StringValue(node.Location.Name),
				"region":    types.StringValue(node.Location.Region),
				"region_code": types.StringValue(
					node.Location.RegionCode,
				),
				"timezone": types.StringValue(node.Location.Timezone),
				"postal_code": types.StringValue(
					node.Location.PostalCode,
				),
				"metro_code": types.StringValue(
					node.Location.MetroCode,
				),
				"city": types.StringValue(
					node.Location.City,
				),
			})

			nodeValue := map[string]attr.Value{
				"name":       types.StringValue(node.Name),
				"connection": nodeConnectionValue,
				"location":   nodeLocationValue,
				// "region": 
				// "distance_measurement":
				// extensions:
			}
			node, _ := types.ObjectValue(NodeType, nodeValue)
			nodes = append(nodes, node)
		}

		database.Nodes, _ = types.ListValue(types.ObjectType{
			AttrTypes: NodeType,
		}, nodes)


		var planOptions types.List

		var databaseOptionsAttr []attr.Value


	// if len(db.Options) > 0 {
		for _, option := range db.Options {
			databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
		}
		planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
		resp.Diagnostics.Append(diags...)

		database.Options = planOptions
	// }
		

		database.PgVersion = types.StringValue(db.PgVersion)
		database.StorageUsed = types.Int64Value(db.StorageUsed)

		state.Databases = append(state.Databases, database)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
