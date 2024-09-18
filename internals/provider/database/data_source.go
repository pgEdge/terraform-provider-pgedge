package database

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
	// Nodes         types.List   `tfsdk:"nodes"`
	Options       types.List   `tfsdk:"options"`
	PgVersion     types.String `tfsdk:"pg_version"`
	StorageUsed   types.Int64  `tfsdk:"storage_used"`
	ConfigVersion types.String `tfsdk:"config_version"`

	// Backups     types.Object `tfsdk:"backups"`
	// Components types.List   `tfsdk:"components"`
	// Extensions types.Object `tfsdk:"extensions"`
	// Roles         types.List   `tfsdk:"roles"`
	// Tables        types.List   `tfsdk:"tables"`
}

type Node struct {
	Name       string         `tfsdk:"name"`
	Connection Connection     `tfsdk:"connection"`
	Location   Location       `tfsdk:"location"`
	Region     NodeRegion     `tfsdk:"region"`
	Extensions NodeExtensions `tfsdk:"extensions"`
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

type DatabaseComponentsItems struct {
	ID          string `tfsdk:"id"`
	Name        string `tfsdk:"name"`
	ReleaseDate string `tfsdk:"release_date"`
	Status      string `tfsdk:"status"`
	Version     string `tfsdk:"version"`
}

type DatabaseRolesItems struct {
	BypassRls       bool   `tfsdk:"bypass_rls"`
	ConnectionLimit int64  `tfsdk:"connection_limit"`
	CreateDb        bool   `tfsdk:"create_db"`
	CreateRole      bool   `tfsdk:"create_role"`
	Inherit         bool   `tfsdk:"inherit"`
	Login           bool   `tfsdk:"login"`
	Name            string `tfsdk:"name"`
	Replication     bool   `tfsdk:"replication"`
	Superuser       bool   `tfsdk:"superuser"`
}

type DatabaseTablesItems struct {
	Columns         []*DatabaseTablesItemsColumnsItems `tfsdk:"columns"`
	Name            string                             `tfsdk:"name"`
	PrimaryKey      []string                           `tfsdk:"primary_key"`
	ReplicationSets []string                           `tfsdk:"replication_sets"`
	Schema          string                             `tfsdk:"schema"`
	Status          []*DatabaseTablesItemsStatusItems  `tfsdk:"status"`
}

type DatabaseTablesItemsColumnsItems struct {
	DataType        string `tfsdk:"data_type"`
	Default         string `tfsdk:"default"`
	IsNullable      bool   `tfsdk:"is_nullable"`
	IsPrimaryKey    bool   `tfsdk:"is_primary_key"`
	Name            string `tfsdk:"name"`
	OrdinalPosition int64  `tfsdk:"ordinal_position"`
}

type DatabaseTablesItemsStatusItems struct {
	Aligned     bool   `tfsdk:"aligned"`
	NodeName    string `tfsdk:"node_name"`
	Present     bool   `tfsdk:"present"`
	Replicating bool   `tfsdk:"replicating"`
}

type DatabaseExtensions struct {
	AutoManage bool     `tfsdk:"auto_manage"`
	Available  []string `tfsdk:"available"`
	Requested  []string `tfsdk:"requested"`
}

var DatabaseExtensionsType = map[string]attr.Type{
	"auto_manage": types.BoolType,
	"available":   types.ListType{ElemType: types.StringType},
	"requested":   types.ListType{ElemType: types.StringType},
}

var DatabaseTablesItemsType = map[string]attr.Type{
	"columns":          types.ListType{ElemType: types.ObjectType{AttrTypes: DatabaseTablesItemsColumnsItemsType}},
	"name":             types.StringType,
	"primary_key":      types.ListType{ElemType: types.StringType},
	"replication_sets": types.ListType{ElemType: types.StringType},
	"schema":           types.StringType,
	"status":           types.ListType{ElemType: types.ObjectType{AttrTypes: DatabaseTablesItemsStatusItemsType}},
}

var DatabaseTablesItemsColumnsItemsType = map[string]attr.Type{
	"data_type":        types.StringType,
	"default":          types.StringType,
	"is_nullable":      types.BoolType,
	"is_primary_key":   types.BoolType,
	"name":             types.StringType,
	"ordinal_position": types.Int64Type,
}

var DatabaseTablesItemsStatusItemsType = map[string]attr.Type{
	"aligned":     types.BoolType,
	"node_name":   types.StringType,
	"present":     types.BoolType,
	"replicating": types.BoolType,
}

var DatabaseRolesItemsType = map[string]attr.Type{
	"bypass_rls":       types.BoolType,
	"connection_limit": types.Int64Type,
	"create_db":        types.BoolType,
	"create_role":      types.BoolType,
	"inherit":          types.BoolType,
	"login":            types.BoolType,
	"name":             types.StringType,
	"replication":      types.BoolType,
	"superuser":        types.BoolType,
}

var DatabaseComponentsItemsType = map[string]attr.Type{
	"id":           types.StringType,
	"name":         types.StringType,
	"release_date": types.StringType,
	"status":       types.StringType,
	"version":      types.StringType,
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
	"region_code": types.StringType,
}

var NodeRegionType = map[string]attr.Type{
	"active":             types.BoolType,
	"availability_zones": types.ListType{ElemType: types.StringType},
	"cloud":              types.StringType,
	"code":               types.StringType,
	"name":               types.StringType,
	"parent":             types.StringType,
}

var NodeType = map[string]attr.Type{
	"name":       types.StringType,
	"connection": types.ObjectType{AttrTypes: NodeConnectionType},
	"location":   types.ObjectType{AttrTypes: NodeLocationType},
	"region":     types.ObjectType{AttrTypes: NodeRegionType},
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
						"config_version": schema.StringAttribute{
							Optional:    true,
							Description: "Config version of the database",
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
									"region": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"active": schema.BoolAttribute{
												Optional: true,

												Computed:    true,
												Description: "Active status of the region",
											},
											"availability_zones": schema.ListAttribute{
												ElementType: types.StringType,
												Optional:    true,

												Computed:    true,
												Description: "Availability zones of the region",
											},
											"cloud": schema.StringAttribute{
												Optional: true,

												Computed:    true,
												Description: "Cloud provider of the region",
											},
											"code": schema.StringAttribute{
												Optional: true,

												Computed:    true,
												Description: "Code of the region",
											},
											"name": schema.StringAttribute{
												Optional: true,

												Computed:    true,
												Description: "Name of the region",
											},
											"parent": schema.StringAttribute{
												Optional: true,

												Computed:    true,
												Description: "Parent region",
											},
										},
									},
									"distance_measurement": schema.SingleNestedAttribute{

										Optional: true,
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"distance": schema.Float64Attribute{
												Optional:    true,
												Computed:    true,
												Description: "Distance from a reference point",
											},
											"from_latitude": schema.Float64Attribute{
												Optional:    true,
												Computed:    true,
												Description: "Latitude of the reference point",
											},
											"from_longitude": schema.Float64Attribute{
												Optional:    true,
												Computed:    true,
												Description: "Longitude of the reference point",
											},
											"unit": schema.StringAttribute{
												Optional:    true,
												Computed:    true,
												Description: "Unit of distance measurement",
											},
										},
									},
									"extensions": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"errors": schema.SingleNestedAttribute{
												Optional: true,
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"anim9ef": schema.StringAttribute{
														Optional:    true,
														Computed:    true,
														Description: "Error code anim9ef",
													},
													"enim3b": schema.StringAttribute{
														Optional:    true,
														Computed:    true,
														Description: "Error code enim3b",
													},
													"laborumd": schema.StringAttribute{
														Optional:    true,
														Computed:    true,
														Description: "Error code laborumd",
													},
													"mollit267": schema.StringAttribute{
														Optional:    true,
														Computed:    true,
														Description: "Error code mollit267",
													},
												},
											},
											"installed": schema.ListAttribute{
												ElementType: types.StringType,
												Computed:    true,
												Optional:    true,
												Description: "List of installed extensions",
											},
										},
									},
								},
							},
						},
						"components": schema.ListNestedAttribute{
							Computed: true,
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"status": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Status of the component",
									},
									"id": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Id of the component",
									},
									"version": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Version of the component",
									},
									"name": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Name of the component",
									},
									"release_date": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Release date of the component",
									},
								},
							},
						},
						"roles": schema.ListNestedAttribute{
							Computed: true,
							// Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"bypass_rls": schema.BoolAttribute{
										Computed:    true,
										Description: "Bypass RLS",
									},
									"connection_limit": schema.Int64Attribute{
										Computed:    true,
										Description: "Connection limit",
									},
									"create_db": schema.BoolAttribute{
										Computed:    true,
										Description: "Create database",
									},
									"create_role": schema.BoolAttribute{
										Computed:    true,
										Description: "Create role",
									},
									"inherit": schema.BoolAttribute{
										Computed:    true,
										Description: "Inherit",
									},
									"login": schema.BoolAttribute{
										Computed:    true,
										Description: "Login",
									},
									"name": schema.StringAttribute{
										Computed:    true,
										Description: "Name of the role",
									},
									"replication": schema.BoolAttribute{
										Computed:    true,
										Description: "Replication",
									},
									"superuser": schema.BoolAttribute{
										Computed:    true,
										Description: "Superuser",
									},
								},
							},
						},
						"tables": schema.ListNestedAttribute{
							Computed: true,
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Name of the table",
									},
									"schema": schema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Schema of the table",
									},
									"primary_key": schema.ListAttribute{ElementType: types.StringType,
										Optional: true,
										Computed: true, Description: "Primary key of the table"},
									"replication_sets": schema.ListAttribute{ElementType: types.StringType,
										Optional: true,
										Computed: true, Description: "Replication sets of the table"},
									"columns": schema.ListNestedAttribute{
										Optional: true,
										Computed: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Name of the column",
												},
												"data_type": schema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Data type of the column",
												},
												"default": schema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Default of the column",
												},
												"is_nullable": schema.BoolAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Is nullable of the column",
												},
												"is_primary_key": schema.BoolAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Is primary key of the column",
												},
												"ordinal_position": schema.Int64Attribute{
													Optional:    true,
													Computed:    true,
													Description: "Ordinal position of the column",
												},
											},
										},
									},
									"status": schema.ListNestedAttribute{
										Optional: true,
										Computed: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aligned": schema.BoolAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Aligned of the table",
												},
												"node_name": schema.StringAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Node name of the table",
												},
												"present": schema.BoolAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Present of the table",
												},
												"replicating": schema.BoolAttribute{
													Optional:    true,
													Computed:    true,
													Description: "Replicating of the table",
												},
											},
										},
									},
								},
							},
						},
						"extensions": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"auto_manage": schema.BoolAttribute{
									Computed:    true,
									Description: "Auto manage of the extension",
								},
								"available": schema.ListAttribute{ElementType: types.StringType, Computed: true, Description: "Available of the extension"},
								"requested": schema.ListAttribute{ElementType: types.StringType, Computed: true, Description: "Requested of the extension"},
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
		// var nodes []attr.Value
		database.ID = types.StringValue(db.ID.String())
		database.Name = types.StringValue(strings.Trim(strings.ToLower(db.Name), " "))
		database.Domain = types.StringValue(db.Domain)
		database.CreatedAt = types.StringValue(db.CreatedAt)
		database.UpdatedAt = types.StringValue(db.UpdatedAt)
		database.Status = types.StringValue(db.Status)
		database.ClusterID = types.StringValue(db.ClusterID.String())

		// for _, node := range db.Nodes {
		// 	nodeConnectionValue, _ := types.ObjectValue(NodeConnectionType, map[string]attr.Value{
		// 		"database":            types.StringValue(node.Connection.Database),
		// 		"host":                types.StringValue(node.Connection.Host),
		// 		"password":            types.StringValue(node.Connection.Password),
		// 		"port":                types.Int64Value(node.Connection.Port),
		// 		"username":            types.StringValue(node.Connection.Username),
		// 		"external_ip_address": types.StringValue(node.Connection.ExternalIPAddress),
		// 		"internal_ip_address": types.StringValue(node.Connection.InternalIPAddress),
		// 		"internal_host":       types.StringValue(node.Connection.InternalHost),
		// 	})

		// 	nodeLocationValue, _ := types.ObjectValue(NodeLocationType, map[string]attr.Value{
		// 		"code":      types.StringValue(node.Location.Code),
		// 		"country":   types.StringValue(node.Location.Country),
		// 		"latitude":  types.Float64Value(node.Location.Latitude),
		// 		"longitude": types.Float64Value(node.Location.Longitude),
		// 		"name":      types.StringValue(node.Location.Name),
		// 		"region":    types.StringValue(node.Location.Region),
		// 		"region_code": types.StringValue(
		// 			node.Location.RegionCode,
		// 		),
		// 		"timezone": types.StringValue(node.Location.Timezone),
		// 		"postal_code": types.StringValue(
		// 			node.Location.PostalCode,
		// 		),
		// 		"metro_code": types.StringValue(
		// 			node.Location.MetroCode,
		// 		),
		// 		"city": types.StringValue(
		// 			node.Location.City,
		// 		),
		// 	})

		// 	var nodeRegionValue attr.Value
		// 	if node.Region != nil {
		// 		nodeRegionValue, _ = types.ObjectValue(NodeRegionType, map[string]attr.Value{
		// 			"active": types.BoolValue(node.Region.Active),
		// 			"availability_zones": func() attr.Value {
		// 				var availability_zone []attr.Value
		// 				for _, region := range node.Region.AvailabilityZones {
		// 					availability_zone = append(availability_zone, types.StringValue(region))
		// 				}
		// 				availabilityZoneList, _ := types.ListValue(types.StringType, availability_zone)

		// 				if availabilityZoneList.IsNull() {
		// 					return types.ListNull(types.StringType)
		// 				}

		// 				return availabilityZoneList
		// 			}(),

		// 			"cloud":  types.StringValue(node.Region.Cloud),
		// 			"code":   types.StringValue(node.Region.Code),
		// 			"name":   types.StringValue(node.Region.Name),
		// 			"parent": types.StringValue(node.Region.Parent),
		// 		})
		// 	} else {
		// 		nodeRegionValue = types.ObjectNull(NodeRegionType)
		// 	}

		// 	resp.Diagnostics.Append(diags...)
		// 	if resp.Diagnostics.HasError() {
		// 		return
		// 	}

		// 	nodeValue := map[string]attr.Value{
		// 		"name":       types.StringValue(node.Name),
		// 		"connection": nodeConnectionValue,
		// 		"location":   nodeLocationValue,
		// 		"region":     nodeRegionValue,
		// 	}
		// 	node, diags := types.ObjectValue(NodeType, nodeValue)
		// 	resp.Diagnostics.Append(diags...)
		// 	if resp.Diagnostics.HasError() {
		// 		return
		// 	}
		// 	nodes = append(nodes, node)
		// }

		// database.Nodes, diags = types.ListValue(types.ObjectType{
		// 	AttrTypes: NodeType,
		// }, nodes)

		// resp.Diagnostics.Append(diags...)
		// if resp.Diagnostics.HasError() {
		// 	return
		// }

		var planOptions types.List

		var databaseOptionsAttr []attr.Value

		for _, option := range db.Options {
			databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
		}
		planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
		resp.Diagnostics.Append(diags...)
		database.Options = planOptions

		// var databaseComponents types.List
		// var databaseComponentsAttr []attr.Value
		// for _, components := range db.Components {
		// 	componentsValue, diags := types.ObjectValue(DatabaseComponentsItemsType, map[string]attr.Value{
		// 		"name":         types.StringValue(components.Name),
		// 		"id":           types.StringValue(components.ID),
		// 		"release_date": types.StringValue(components.ReleaseDate),
		// 		"version":      types.StringValue(components.Version),
		// 		"status":       types.StringValue(components.Status),
		// 	})
		// 	resp.Diagnostics.Append(diags...)
		// 	if resp.Diagnostics.HasError() {
		// 		return
		// 	}
		// 	databaseComponentsAttr = append(databaseComponentsAttr, componentsValue)
		// }
		// databaseComponents, diags = types.ListValue(types.ObjectType{
		// 	AttrTypes: DatabaseComponentsItemsType,
		// }, databaseComponentsAttr)
		// resp.Diagnostics.Append(diags...)
		// if resp.Diagnostics.HasError() {
		// 	return
		// }
		// database.Components = databaseComponents
		// var databaseRoles types.List
		// var databaseRolesAttr []attr.Value
		// for _, role := range db.Roles {
		// 	var rolesValue types.Object
		// 	if role != nil {
		// 		rolesValue, diags = types.ObjectValue(DatabaseRolesItemsType, map[string]attr.Value{
		// 			"bypass_rls":       types.BoolValue(role.BypassRls),
		// 			"connection_limit": types.Int64Value(role.ConnectionLimit),
		// 			"create_db":        types.BoolValue(role.CreateDb),
		// 			"create_role":      types.BoolValue(role.CreateRole),
		// 			"inherit":          types.BoolValue(role.Inherit),
		// 			"login":            types.BoolValue(role.Login),
		// 			"name":             types.StringValue(role.Name),
		// 			"replication":      types.BoolValue(role.Replication),
		// 			"superuser":        types.BoolValue(role.Superuser),
		// 		})

		// 		resp.Diagnostics.Append(diags...)
		// 		if resp.Diagnostics.HasError() {
		// 			return
		// 		}
		// 	} else {
		// 		rolesValue = types.ObjectNull(DatabaseRolesItemsType)
		// 	}

		// 	databaseRolesAttr = append(databaseRolesAttr, rolesValue)
		// }

		// databaseRoles, diags = types.ListValue(types.ObjectType{
		// 	AttrTypes: DatabaseRolesItemsType,
		// }, databaseRolesAttr)

		// resp.Diagnostics.Append(diags...)
		// if resp.Diagnostics.HasError() {
		// 	return
		// }

		// database.Roles = databaseRoles

		// var databaseTables types.List

		// var databaseTablesAttr []attr.Value

		// if db.Tables != nil {
		// 	for _, table := range db.Tables {
		// 		var tableColumns types.List

		// 		var tableColumnsAttr []attr.Value

		// 		for _, column := range table.Columns {
		// 			DatabaseTablesItemsColumnsItemsValue, diags := types.ObjectValue(DatabaseTablesItemsColumnsItemsType, map[string]attr.Value{
		// 				"name":             types.StringValue(column.Name),
		// 				"default":          types.StringValue(column.Default),
		// 				"is_nullable":      types.BoolValue(column.IsNullable),
		// 				"data_type":        types.StringValue(column.DataType),
		// 				"is_primary_key":   types.BoolValue(column.IsPrimaryKey),
		// 				"ordinal_position": types.Int64Value(column.OrdinalPosition),
		// 			})

		// 			resp.Diagnostics.Append(diags...)
		// 			if resp.Diagnostics.HasError() {
		// 				return
		// 			}

		// 			tableColumnsAttr = append(tableColumnsAttr, DatabaseTablesItemsColumnsItemsValue)
		// 		}

		// 		tableColumns, diags = types.ListValue(types.ObjectType{
		// 			AttrTypes: DatabaseTablesItemsColumnsItemsType,
		// 		}, tableColumnsAttr)

		// 		resp.Diagnostics.Append(diags...)
		// 		if resp.Diagnostics.HasError() {
		// 			return
		// 		}

		// 		var tableStatus types.List

		// 		var tableStatusAttr []attr.Value

		// 		for _, status := range table.Status {
		// 			DatabaseTablesItemsStatusItemsValue, diags := types.ObjectValue(DatabaseTablesItemsStatusItemsType, map[string]attr.Value{
		// 				"aligned":     types.BoolValue(status.Aligned),
		// 				"node_name":   types.StringValue(status.NodeName),
		// 				"present":     types.BoolValue(status.Present),
		// 				"replicating": types.BoolValue(status.Replicating),
		// 			})

		// 			resp.Diagnostics.Append(diags...)
		// 			if resp.Diagnostics.HasError() {
		// 				return
		// 			}

		// 			tableStatusAttr = append(tableStatusAttr, DatabaseTablesItemsStatusItemsValue)
		// 		}

		// 		tableStatus, diags = types.ListValue(types.ObjectType{
		// 			AttrTypes: DatabaseTablesItemsStatusItemsType,
		// 		}, tableStatusAttr)

		// 		resp.Diagnostics.Append(diags...)
		// 		if resp.Diagnostics.HasError() {
		// 			return
		// 		}

		// 		DatabaseTablesItemsValue, diags := types.ObjectValue(DatabaseTablesItemsType, map[string]attr.Value{
		// 			"columns": tableColumns,
		// 			"schema":  types.StringValue(table.Schema),
		// 			"primary_key": func() types.List {

		// 				var primaryKey types.List

		// 				var primaryKeyAttr []attr.Value

		// 				for _, pk := range table.PrimaryKey {
		// 					primaryKeyAttr = append(primaryKeyAttr, types.StringValue(pk))
		// 				}

		// 				primaryKey, diags = types.ListValue(types.StringType, primaryKeyAttr)

		// 				resp.Diagnostics.Append(diags...)
		// 				return primaryKey
		// 			}(),
		// 			"replication_sets": func() types.List {
		// 				var replicationSets types.List

		// 				var replicationSetsAttr []attr.Value

		// 				for _, rs := range table.ReplicationSets {
		// 					replicationSetsAttr = append(replicationSetsAttr, types.StringValue(rs))
		// 				}

		// 				replicationSets, diags = types.ListValue(types.StringType, replicationSetsAttr)

		// 				resp.Diagnostics.Append(diags...)
		// 				return replicationSets
		// 			}(),
		// 			"name":   types.StringValue(table.Name),
		// 			"status": tableStatus,
		// 		})

		// 		resp.Diagnostics.Append(diags...)
		// 		if resp.Diagnostics.HasError() {
		// 			return
		// 		}

		// 		databaseTablesAttr = append(databaseTablesAttr, DatabaseTablesItemsValue)
		// 	}

		// 	databaseTables, diags = types.ListValue(types.ObjectType{
		// 		AttrTypes: DatabaseTablesItemsType,
		// 	}, databaseTablesAttr)

		// 	resp.Diagnostics.Append(diags...)
		// 	if resp.Diagnostics.HasError() {
		// 		return
		// 	}
		// } else {
		// 	databaseTables = types.ListNull(types.ObjectType{
		// 		AttrTypes: DatabaseTablesItemsType,
		// 	})
		// }

		// database.Tables = databaseTables

		// DatabaseExtensionsValue, diags := types.ObjectValue(DatabaseExtensionsType, map[string]attr.Value{
		// 	"auto_manage": types.BoolValue(db.Extensions.AutoManage),
		// 	"available": func() types.List {
		// 		var available []attr.Value
		// 		if db.Extensions.Available != nil {
		// 			for _, extension := range db.Extensions.Available {
		// 				available = append(available, types.StringValue(extension))
		// 			}
		// 		} else {
		// 			available = append(available, types.StringValue(""))
		// 		}
		// 		availableList, diags := types.ListValue(types.StringType, available)
		// 		resp.Diagnostics.Append(diags...)
		// 		if availableList.IsNull() {
		// 			return types.ListNull(types.StringType)
		// 		}
		// 		return availableList
		// 	}(),
		// 	"requested": func() types.List {
		// 		var requested []attr.Value
		// 		if db.Extensions.Requested != nil {
		// 			for _, extension := range db.Extensions.Requested {
		// 				requested = append(requested, types.StringValue(extension))
		// 			}
		// 		} else {
		// 			requested = append(requested, types.StringValue(""))
		// 		}
		// 		requestedList, diags := types.ListValue(types.StringType, requested)
		// 		resp.Diagnostics.Append(diags...)
		// 		if requestedList.IsNull() {
		// 			return types.ListNull(types.StringType)
		// 		}
		// 		return requestedList
		// 	}(),
		// })

		// resp.Diagnostics.Append(diags...)
		// if resp.Diagnostics.HasError() {
		// 	return
		// }

		// database.Extensions = DatabaseExtensionsValue
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
