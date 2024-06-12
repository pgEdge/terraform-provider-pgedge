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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "ID of the database",
		},
		"name": schema.StringAttribute{
			Required: true,
			Validators: []validator.String{
				stringvalidator.RegexMatches(
					regexp.MustCompile(`^[a-z0-9]+$`),
					"must contain only lowercase alphanumeric characters",
				),
			},
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
			Required:    true,
			Description: "Updated at of the database",
		},
		"config_version": schema.StringAttribute{
			Optional:    true,
			Description: "Config version of the database",
		},
		"options": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			Description: "Options for creating the database",
		},
		"nodes": schema.ListNestedAttribute{
			Computed: true,
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
			Optional: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"bypass_rls": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: "Bypass RLS",
					},
					"connection_limit": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: "Connection limit",
					},
					"create_db": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: "Create database",
					},
					"create_role": schema.BoolAttribute{
						Computed: true,
						Optional: true,

						Description: "Create role",
					},
					"inherit": schema.BoolAttribute{
						Computed: true,
						Optional: true,

						Description: "Inherit",
					},
					"login": schema.BoolAttribute{
						Computed: true,
						Optional: true,

						Description: "Login",
					},
					"name": schema.StringAttribute{
						Computed: true,
						Optional: true,

						Description: "Name of the role",
					},
					"replication": schema.BoolAttribute{
						Computed: true,
						Optional: true,

						Description: "Replication",
					},
					"superuser": schema.BoolAttribute{
						Computed: true,
						Optional: true,

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
					"primary_key": schema.ListAttribute{ElementType: types.StringType,Optional:    true,
						Computed: true, Description: "Primary key of the table"},
					"replication_sets": schema.ListAttribute{ElementType: types.StringType,Optional:    true,
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
					Optional:    true,
					Description: "Auto manage of the extension",
				},
				"available": schema.ListAttribute{ElementType: types.StringType, Optional: true, Computed: true, Description: "Available of the extension"},
				"requested": schema.ListAttribute{ElementType: types.StringType, Optional: true, Computed: true, Description: "Requested of the extension"},
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
		Description: "Interface with the pgEdge service API.",
	}
}

func databaseExtensionsReq(ctx context.Context, resp *resource.CreateResponse, extensionsReq basetypes.ObjectValue) (*resource.CreateResponse, *models.DatabaseCreationRequestExtensions) {
	var extension DatabaseExtensions

	diags := extensionsReq.As(ctx, &extension, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)

	var available []string
	var requested []string
	if len(extension.Available) > 0 {
		available = extension.Available
	}

	if len(extension.Requested) > 0 {
		requested = extension.Requested
	}
	
	return resp, &models.DatabaseCreationRequestExtensions{
		AutoManage:   extension.AutoManage,
		Available:    available,
		Requested:    requested,
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
	if !plan.Options.IsNull() && !plan.Options.IsUnknown() {
		diags = plan.Options.ElementsAs(ctx, &databaseOptions, false)
		resp.Diagnostics.Append(diags...)
	}

	items := &models.DatabaseCreationRequest{
		Name:      plan.Name.ValueString(),
		ClusterID: strfmt.UUID(plan.ClusterID.ValueString()),
		Options:   databaseOptions,
		ConfigVersion: plan.ConfigVersion.ValueString(),
	}

	if !plan.Extensions.IsNull() && !plan.Extensions.IsUnknown() {
		resp, items.Extensions = databaseExtensionsReq(ctx, resp, plan.Extensions)
	}

	database, err := r.client.CreateDatabase(ctx, items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating database",
			"Could not create database, unexpected error: "+err.Error(),
		)
		return
	}

	plan = DatabaseDetails{
		ID:          types.StringValue(database.ID.String()),
		Name:        types.StringValue(strings.Trim(strings.ToLower(database.Name), " ")),
		Domain:      types.StringValue(database.Domain),
		Status:      types.StringValue(database.Status),
		ClusterID:   types.StringValue(database.ClusterID.String()),
		CreatedAt:   types.StringValue(database.CreatedAt),
		UpdatedAt:   types.StringValue(database.UpdatedAt),
		PgVersion:   types.StringValue(database.PgVersion),
		StorageUsed: types.Int64Value(int64(database.StorageUsed)),
	}

	var planOptions types.List

	var databaseOptionsAttr []attr.Value

	for _, option := range database.Options {
		databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
	}
	planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
	resp.Diagnostics.Append(diags...)

	plan.Options = planOptions

	var nodes []attr.Value
	for _, node := range database.Nodes {
		nodeConnectionValue, _ := types.ObjectValue(NodeConnectionType, map[string]attr.Value{
			"database":            types.StringValue(node.Connection.Database),
			"host":                types.StringValue(node.Connection.Host),
			"password":            types.StringValue(node.Connection.Password),
			"port":                types.Int64Value(node.Connection.Port),
			"username":            types.StringValue(node.Connection.Username),
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

		var nodeRegionValue attr.Value
		if node.Region != nil {
			nodeRegionValue, _ = types.ObjectValue(NodeRegionType, map[string]attr.Value{
				"active": types.BoolValue(node.Region.Active),
				"availability_zones": func() attr.Value {
					var availability_zone []attr.Value
					for _, region := range node.Region.AvailabilityZones {
						availability_zone = append(availability_zone, types.StringValue(region))
					}
					availabilityZoneList, _ := types.ListValue(types.StringType, availability_zone)

					if availabilityZoneList.IsNull() {
						return types.ListNull(types.StringType)
					}

					return availabilityZoneList
				}(),

				"cloud":  types.StringValue(node.Region.Cloud),
				"code":   types.StringValue(node.Region.Code),
				"name":   types.StringValue(node.Region.Name),
				"parent": types.StringValue(node.Region.Parent),
			})
		} else {
			nodeRegionValue = types.ObjectNull(NodeRegionType)
		}

		var nodeDistanceMeasurementValue attr.Value
		if node.DistanceMeasurement != nil {
			nodeDistanceMeasurementValue, _ = types.ObjectValue(NodeDistanceMeasurementType, map[string]attr.Value{
				"distance":       types.Float64Value(node.DistanceMeasurement.Distance),
				"from_latitude":  types.Float64Value(node.DistanceMeasurement.FromLatitude),
				"from_longitude": types.Float64Value(node.DistanceMeasurement.FromLongitude),
				"unit":           types.StringValue(node.DistanceMeasurement.Unit),
			})
		} else {
			nodeDistanceMeasurementValue = types.ObjectNull(NodeDistanceMeasurementType)
		}

		nodeExtensionsValue, diags := types.ObjectValue(NodeExtensionsType, map[string]attr.Value{
			"errors": func() types.Object {
				var NodeExtensionsErrorsValue = map[string]attr.Value{}
				if node.Extensions != nil {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(node.Extensions.Errors.Anim9ef),
						"enim3b":    types.StringValue(node.Extensions.Errors.Enim3b),
						"laborumd":  types.StringValue(node.Extensions.Errors.Laborumd),
						"mollit267": types.StringValue(node.Extensions.Errors.Mollit267),
					}
				} else {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(""),
						"enim3b":    types.StringValue(""),
						"laborumd":  types.StringValue(""),
						"mollit267": types.StringValue(""),
					}
				}

				item, diags := types.ObjectValue(NodeExtensionsErrorsType, NodeExtensionsErrorsValue)

				resp.Diagnostics.Append(diags...)

				if item.IsNull() {
					return types.ObjectNull(NodeExtensionsErrorsType)
				}

				return item
			}(),
			"installed": func() types.List {
				var installed []attr.Value
				if node.Extensions != nil {
					for _, extension := range node.Extensions.Installed {
						installed = append(installed, types.StringValue(extension))
					}
				} else {
					installed = append(installed, types.StringValue(""))
				}
				installedList, diags := types.ListValue(types.StringType, installed)
				resp.Diagnostics.Append(diags...)
				if installedList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return installedList
			}(),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		nodeValue := map[string]attr.Value{
			"name":                 types.StringValue(node.Name),
			"connection":           nodeConnectionValue,
			"location":             nodeLocationValue,
			"region":               nodeRegionValue,
			"distance_measurement": nodeDistanceMeasurementValue,
			"extensions":           nodeExtensionsValue,
		}

		node, diags := types.ObjectValue(NodeType, nodeValue)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		nodes = append(nodes, node)
	}

	plan.Nodes, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeType,
	}, nodes)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var databaseComponents types.List

	var databaseComponentsAttr []attr.Value

	for _, components := range database.Components {
		componentsValue, diags := types.ObjectValue(DatabaseComponentsItemsType, map[string]attr.Value{
			"name":         types.StringValue(components.Name),
			"id":           types.StringValue(components.ID),
			"release_date": types.StringValue(components.ReleaseDate),
			"version":      types.StringValue(components.Version),
			"status":       types.StringValue(components.Status),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		databaseComponentsAttr = append(databaseComponentsAttr, componentsValue)
	}

	databaseComponents, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseComponentsItemsType,
	}, databaseComponentsAttr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	plan.Components = databaseComponents

	var databaseRoles types.List

	var databaseRolesAttr []attr.Value

	for _, role := range database.Roles {
		var rolesValue types.Object
		if role != nil {
			rolesValue, diags = types.ObjectValue(DatabaseRolesItemsType, map[string]attr.Value{
				"bypass_rls":       types.BoolValue(role.BypassRls),
				"connection_limit": types.Int64Value(role.ConnectionLimit),
				"create_db":        types.BoolValue(role.CreateDb),
				"create_role":      types.BoolValue(role.CreateRole),
				"inherit":          types.BoolValue(role.Inherit),
				"login":            types.BoolValue(role.Login),
				"name":             types.StringValue(role.Name),
				"replication":      types.BoolValue(role.Replication),
				"superuser":        types.BoolValue(role.Superuser),
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}else{
			rolesValue = types.ObjectNull(DatabaseRolesItemsType)
		}

		databaseRolesAttr = append(databaseRolesAttr, rolesValue)
	}

	databaseRoles, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseRolesItemsType,
	}, databaseRolesAttr)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Roles = databaseRoles

	var databaseTables types.List

	var databaseTablesAttr []attr.Value

	if database.Tables != nil {
		for _, table := range database.Tables {
			var tableColumns types.List
	
			var tableColumnsAttr []attr.Value
	
			for _, column := range table.Columns {
				DatabaseTablesItemsColumnsItemsValue, diags := types.ObjectValue(DatabaseTablesItemsColumnsItemsType, map[string]attr.Value{
					"name":             types.StringValue(column.Name),
					"default":          types.StringValue(column.Default),
					"is_nullable":      types.BoolValue(column.IsNullable),
					"data_type":        types.StringValue(column.DataType),
					"is_primary_key":   types.BoolValue(column.IsPrimaryKey),
					"ordinal_position": types.Int64Value(column.OrdinalPosition),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableColumnsAttr = append(tableColumnsAttr, DatabaseTablesItemsColumnsItemsValue)
			}
	
			tableColumns, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsColumnsItemsType,
			}, tableColumnsAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			var tableStatus types.List
	
			var tableStatusAttr []attr.Value
	
			for _, status := range table.Status {
				DatabaseTablesItemsStatusItemsValue, diags := types.ObjectValue(DatabaseTablesItemsStatusItemsType, map[string]attr.Value{
					"aligned":     types.BoolValue(status.Aligned),
					"node_name":   types.StringValue(status.NodeName),
					"present":     types.BoolValue(status.Present),
					"replicating": types.BoolValue(status.Replicating),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableStatusAttr = append(tableStatusAttr, DatabaseTablesItemsStatusItemsValue)
			}
	
			tableStatus, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsStatusItemsType,
			}, tableStatusAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			DatabaseTablesItemsValue, diags := types.ObjectValue(DatabaseTablesItemsType, map[string]attr.Value{
				"columns": tableColumns,
				"schema":  types.StringValue(table.Schema),
				"primary_key": func() types.List {
	
					var primaryKey types.List
	
					var primaryKeyAttr []attr.Value
	
					for _, pk := range table.PrimaryKey {
						primaryKeyAttr = append(primaryKeyAttr, types.StringValue(pk))
					}
	
					primaryKey, diags = types.ListValue(types.StringType, primaryKeyAttr)
	
					resp.Diagnostics.Append(diags...)
					return primaryKey
				}(),
				"replication_sets": func() types.List {
					var replicationSets types.List
	
					var replicationSetsAttr []attr.Value
	
					for _, rs := range table.ReplicationSets {
						replicationSetsAttr = append(replicationSetsAttr, types.StringValue(rs))
					}
	
					replicationSets, diags = types.ListValue(types.StringType, replicationSetsAttr)
	
					resp.Diagnostics.Append(diags...)
					return replicationSets
				}(),
				"name":   types.StringValue(table.Name),
				"status": tableStatus,
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			databaseTablesAttr = append(databaseTablesAttr, DatabaseTablesItemsValue)
		}

		databaseTables, diags = types.ListValue(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		}, databaseTablesAttr)
	
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}else{
		databaseTables = types.ListNull(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		})
	}

	plan.Tables = databaseTables

	DatabaseExtensionsValue, diags := types.ObjectValue(DatabaseExtensionsType, map[string]attr.Value{
		"auto_manage": types.BoolValue(database.Extensions.AutoManage),
		"available": func() types.List {
			if database.Extensions == nil || database.Extensions.Available == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var available []attr.Value
			if database.Extensions.Available != nil {
				for _, extension := range database.Extensions.Available {
					available = append(available, types.StringValue(extension))
				}
			} else {
				available = append(available, types.StringValue(""))
			}
			availableList, diags := types.ListValue(types.StringType, available)
			resp.Diagnostics.Append(diags...)
			
			return availableList
		}(),
		"requested": func() types.List {
			if database.Extensions == nil || database.Extensions.Requested == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var requested []attr.Value
			if database.Extensions.Requested != nil {
				for _, extension := range database.Extensions.Requested {
					requested = append(requested, types.StringValue(extension))
				}
			} else {
				requested = append(requested, types.StringValue(""))
			}
			requestedList, diags := types.ListValue(types.StringType, requested)
			resp.Diagnostics.Append(diags...)
			
			return requestedList
		}(),
	})

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Extensions = DatabaseExtensionsValue

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

	state = DatabaseDetails{
		ID:          types.StringValue(database.ID.String()),
		Name:        types.StringValue(strings.Trim(strings.ToLower(database.Name), " ")),
		Domain:      types.StringValue(database.Domain),
		Status:      types.StringValue(database.Status),
		ClusterID:   types.StringValue(database.ClusterID.String()),
		CreatedAt:   types.StringValue(database.CreatedAt),
		UpdatedAt:   types.StringValue(database.UpdatedAt),
		PgVersion:   types.StringValue(database.PgVersion),
		StorageUsed: types.Int64Value(int64(database.StorageUsed)),
	}

	var stateOptions types.List

	var databaseOptionsAttr []attr.Value

	for _, option := range database.Options {
		databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
	}
	stateOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
	resp.Diagnostics.Append(diags...)

	state.Options = stateOptions

	var nodes []attr.Value
	for _, node := range database.Nodes {
		nodeConnectionValue, _ := types.ObjectValue(NodeConnectionType, map[string]attr.Value{
			"database":            types.StringValue(node.Connection.Database),
			"host":                types.StringValue(node.Connection.Host),
			"password":            types.StringValue(node.Connection.Password),
			"port":                types.Int64Value(node.Connection.Port),
			"username":            types.StringValue(node.Connection.Username),
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

		var nodeRegionValue attr.Value
		if node.Region != nil {
			nodeRegionValue, _ = types.ObjectValue(NodeRegionType, map[string]attr.Value{
				"active": types.BoolValue(node.Region.Active),
				"availability_zones": func() attr.Value {
					var availability_zone []attr.Value
					for _, region := range node.Region.AvailabilityZones {
						availability_zone = append(availability_zone, types.StringValue(region))
					}
					availabilityZoneList, _ := types.ListValue(types.StringType, availability_zone)

					if availabilityZoneList.IsNull() {
						return types.ListNull(types.StringType)
					}

					return availabilityZoneList
				}(),

				"cloud":  types.StringValue(node.Region.Cloud),
				"code":   types.StringValue(node.Region.Code),
				"name":   types.StringValue(node.Region.Name),
				"parent": types.StringValue(node.Region.Parent),
			})
		} else {
			nodeRegionValue = types.ObjectNull(NodeRegionType)
		}

		var nodeDistanceMeasurementValue attr.Value
		if node.DistanceMeasurement != nil {
			nodeDistanceMeasurementValue, _ = types.ObjectValue(NodeDistanceMeasurementType, map[string]attr.Value{
				"distance":       types.Float64Value(node.DistanceMeasurement.Distance),
				"from_latitude":  types.Float64Value(node.DistanceMeasurement.FromLatitude),
				"from_longitude": types.Float64Value(node.DistanceMeasurement.FromLongitude),
				"unit":           types.StringValue(node.DistanceMeasurement.Unit),
			})
		} else {
			nodeDistanceMeasurementValue = types.ObjectNull(NodeDistanceMeasurementType)
		}

		nodeExtensionsValue, diags := types.ObjectValue(NodeExtensionsType, map[string]attr.Value{
			"errors": func() types.Object {
				var NodeExtensionsErrorsValue = map[string]attr.Value{}
				if node.Extensions != nil {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(node.Extensions.Errors.Anim9ef),
						"enim3b":    types.StringValue(node.Extensions.Errors.Enim3b),
						"laborumd":  types.StringValue(node.Extensions.Errors.Laborumd),
						"mollit267": types.StringValue(node.Extensions.Errors.Mollit267),
					}
				} else {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(""),
						"enim3b":    types.StringValue(""),
						"laborumd":  types.StringValue(""),
						"mollit267": types.StringValue(""),
					}
				}

				item, diags := types.ObjectValue(NodeExtensionsErrorsType, NodeExtensionsErrorsValue)

				resp.Diagnostics.Append(diags...)

				if item.IsNull() {
					return types.ObjectNull(NodeExtensionsErrorsType)
				}

				return item
			}(),
			"installed": func() types.List {
				var installed []attr.Value
				if node.Extensions != nil {
					for _, extension := range node.Extensions.Installed {
						installed = append(installed, types.StringValue(extension))
					}
				} else {
					installed = append(installed, types.StringValue(""))
				}
				installedList, diags := types.ListValue(types.StringType, installed)
				resp.Diagnostics.Append(diags...)
				if installedList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return installedList
			}(),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		nodeValue := map[string]attr.Value{
			"name":                 types.StringValue(node.Name),
			"connection":           nodeConnectionValue,
			"location":             nodeLocationValue,
			"region":               nodeRegionValue,
			"distance_measurement": nodeDistanceMeasurementValue,
			"extensions":           nodeExtensionsValue,
		}

		node, diags := types.ObjectValue(NodeType, nodeValue)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		nodes = append(nodes, node)
	}

	state.Nodes, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeType,
	}, nodes)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var databaseComponents types.List

	var databaseComponentsAttr []attr.Value

	for _, components := range database.Components {
		componentsValue, diags := types.ObjectValue(DatabaseComponentsItemsType, map[string]attr.Value{
			"name":         types.StringValue(components.Name),
			"id":           types.StringValue(components.ID),
			"release_date": types.StringValue(components.ReleaseDate),
			"version":      types.StringValue(components.Version),
			"status":       types.StringValue(components.Status),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		databaseComponentsAttr = append(databaseComponentsAttr, componentsValue)
	}

	databaseComponents, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseComponentsItemsType,
	}, databaseComponentsAttr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	state.Components = databaseComponents

	var databaseRoles types.List

	var databaseRolesAttr []attr.Value

	for _, role := range database.Roles {
		var rolesValue types.Object
		if role != nil {
			rolesValue, diags = types.ObjectValue(DatabaseRolesItemsType, map[string]attr.Value{
				"bypass_rls":       types.BoolValue(role.BypassRls),
				"connection_limit": types.Int64Value(role.ConnectionLimit),
				"create_db":        types.BoolValue(role.CreateDb),
				"create_role":      types.BoolValue(role.CreateRole),
				"inherit":          types.BoolValue(role.Inherit),
				"login":            types.BoolValue(role.Login),
				"name":             types.StringValue(role.Name),
				"replication":      types.BoolValue(role.Replication),
				"superuser":        types.BoolValue(role.Superuser),
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}else{
			rolesValue = types.ObjectNull(DatabaseRolesItemsType)
		}

		databaseRolesAttr = append(databaseRolesAttr, rolesValue)
	}

	databaseRoles, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseRolesItemsType,
	}, databaseRolesAttr)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Roles = databaseRoles

	var databaseTables types.List

	var databaseTablesAttr []attr.Value

	if database.Tables != nil {
		for _, table := range database.Tables {
			var tableColumns types.List
	
			var tableColumnsAttr []attr.Value
	
			for _, column := range table.Columns {
				DatabaseTablesItemsColumnsItemsValue, diags := types.ObjectValue(DatabaseTablesItemsColumnsItemsType, map[string]attr.Value{
					"name":             types.StringValue(column.Name),
					"default":          types.StringValue(column.Default),
					"is_nullable":      types.BoolValue(column.IsNullable),
					"data_type":        types.StringValue(column.DataType),
					"is_primary_key":   types.BoolValue(column.IsPrimaryKey),
					"ordinal_position": types.Int64Value(column.OrdinalPosition),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableColumnsAttr = append(tableColumnsAttr, DatabaseTablesItemsColumnsItemsValue)
			}
	
			tableColumns, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsColumnsItemsType,
			}, tableColumnsAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			var tableStatus types.List
	
			var tableStatusAttr []attr.Value
	
			for _, status := range table.Status {
				DatabaseTablesItemsStatusItemsValue, diags := types.ObjectValue(DatabaseTablesItemsStatusItemsType, map[string]attr.Value{
					"aligned":     types.BoolValue(status.Aligned),
					"node_name":   types.StringValue(status.NodeName),
					"present":     types.BoolValue(status.Present),
					"replicating": types.BoolValue(status.Replicating),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableStatusAttr = append(tableStatusAttr, DatabaseTablesItemsStatusItemsValue)
			}
	
			tableStatus, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsStatusItemsType,
			}, tableStatusAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			DatabaseTablesItemsValue, diags := types.ObjectValue(DatabaseTablesItemsType, map[string]attr.Value{
				"columns": tableColumns,
				"schema":  types.StringValue(table.Schema),
				"primary_key": func() types.List {
	
					var primaryKey types.List
	
					var primaryKeyAttr []attr.Value
	
					for _, pk := range table.PrimaryKey {
						primaryKeyAttr = append(primaryKeyAttr, types.StringValue(pk))
					}
	
					primaryKey, diags = types.ListValue(types.StringType, primaryKeyAttr)
	
					resp.Diagnostics.Append(diags...)
					return primaryKey
				}(),
				"replication_sets": func() types.List {
					var replicationSets types.List
	
					var replicationSetsAttr []attr.Value
	
					for _, rs := range table.ReplicationSets {
						replicationSetsAttr = append(replicationSetsAttr, types.StringValue(rs))
					}
	
					replicationSets, diags = types.ListValue(types.StringType, replicationSetsAttr)
	
					resp.Diagnostics.Append(diags...)
					return replicationSets
				}(),
				"name":   types.StringValue(table.Name),
				"status": tableStatus,
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			databaseTablesAttr = append(databaseTablesAttr, DatabaseTablesItemsValue)
		}

		databaseTables, diags = types.ListValue(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		}, databaseTablesAttr)
	
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}else{
		databaseTables = types.ListNull(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		})
	}

	state.Tables = databaseTables

	DatabaseExtensionsValue, diags := types.ObjectValue(DatabaseExtensionsType, map[string]attr.Value{
		"auto_manage": types.BoolValue(database.Extensions.AutoManage),
		"available": func() types.List {
			if database.Extensions == nil || database.Extensions.Available == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var available []attr.Value
			if database.Extensions.Available != nil {
				for _, extension := range database.Extensions.Available {
					available = append(available, types.StringValue(extension))
				}
			} else {
				available = append(available, types.StringValue(""))
			}
			availableList, diags := types.ListValue(types.StringType, available)
			resp.Diagnostics.Append(diags...)
			
			return availableList
		}(),
		"requested": func() types.List {
			if database.Extensions == nil || database.Extensions.Requested == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var requested []attr.Value
			if database.Extensions.Requested != nil {
				for _, extension := range database.Extensions.Requested {
					requested = append(requested, types.StringValue(extension))
				}
			} else {
				requested = append(requested, types.StringValue(""))
			}
			requestedList, diags := types.ListValue(types.StringType, requested)
			resp.Diagnostics.Append(diags...)
			
			return requestedList
		}(),
	})

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Extensions = DatabaseExtensionsValue

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DatabaseDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state DatabaseDetails
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var databaseOptions []string
	if !plan.Options.IsNull() && !plan.Options.IsUnknown() {
		diags = plan.Options.ElementsAs(ctx, &databaseOptions, false)
		resp.Diagnostics.Append(diags...)
	}

	newResp := &resource.CreateResponse{}

	_, databaseExtensions := databaseExtensionsReq(ctx, newResp, plan.Extensions)

	*resp = resource.UpdateResponse(*newResp)

	items := &models.DatabaseUpdateRequest{
		Options: databaseOptions,
		Extensions: &models.DatabaseUpdateRequestExtensions{
			AutoManage: databaseExtensions.AutoManage,
			Available:  databaseExtensions.Available,
			Requested:  databaseExtensions.Requested,
		},
	}

	database, err := r.client.UpdateDatabase(ctx, strfmt.UUID(state.ID.ValueString()), items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating database",
			"Could not update database, unexpected error: "+err.Error(),
		)
		return
	}

	plan = DatabaseDetails{
		ID:          types.StringValue(database.ID.String()),
		Name:        types.StringValue(strings.Trim(strings.ToLower(database.Name), " ")),
		Domain:      types.StringValue(database.Domain),
		Status:      types.StringValue(database.Status),
		ClusterID:   types.StringValue(database.ClusterID.String()),
		CreatedAt:   types.StringValue(database.CreatedAt),
		UpdatedAt:   types.StringValue(database.UpdatedAt),
		PgVersion:   types.StringValue(database.PgVersion),
		StorageUsed: types.Int64Value(int64(database.StorageUsed)),
	}

	var planOptions types.List

	var databaseOptionsAttr []attr.Value

	for _, option := range database.Options {
		databaseOptionsAttr = append(databaseOptionsAttr, types.StringValue(option))
	}
	planOptions, diags = types.ListValue(types.StringType, databaseOptionsAttr)
	resp.Diagnostics.Append(diags...)

	plan.Options = planOptions

	var nodes []attr.Value
	for _, node := range database.Nodes {
		nodeConnectionValue, _ := types.ObjectValue(NodeConnectionType, map[string]attr.Value{
			"database":            types.StringValue(node.Connection.Database),
			"host":                types.StringValue(node.Connection.Host),
			"password":            types.StringValue(node.Connection.Password),
			"port":                types.Int64Value(node.Connection.Port),
			"username":            types.StringValue(node.Connection.Username),
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

		var nodeRegionValue attr.Value
		if node.Region != nil {
			nodeRegionValue, _ = types.ObjectValue(NodeRegionType, map[string]attr.Value{
				"active": types.BoolValue(node.Region.Active),
				"availability_zones": func() attr.Value {
					var availability_zone []attr.Value
					for _, region := range node.Region.AvailabilityZones {
						availability_zone = append(availability_zone, types.StringValue(region))
					}
					availabilityZoneList, _ := types.ListValue(types.StringType, availability_zone)

					if availabilityZoneList.IsNull() {
						return types.ListNull(types.StringType)
					}

					return availabilityZoneList
				}(),

				"cloud":  types.StringValue(node.Region.Cloud),
				"code":   types.StringValue(node.Region.Code),
				"name":   types.StringValue(node.Region.Name),
				"parent": types.StringValue(node.Region.Parent),
			})
		} else {
			nodeRegionValue = types.ObjectNull(NodeRegionType)
		}

		var nodeDistanceMeasurementValue attr.Value
		if node.DistanceMeasurement != nil {
			nodeDistanceMeasurementValue, _ = types.ObjectValue(NodeDistanceMeasurementType, map[string]attr.Value{
				"distance":       types.Float64Value(node.DistanceMeasurement.Distance),
				"from_latitude":  types.Float64Value(node.DistanceMeasurement.FromLatitude),
				"from_longitude": types.Float64Value(node.DistanceMeasurement.FromLongitude),
				"unit":           types.StringValue(node.DistanceMeasurement.Unit),
			})
		} else {
			nodeDistanceMeasurementValue = types.ObjectNull(NodeDistanceMeasurementType)
		}

		nodeExtensionsValue, diags := types.ObjectValue(NodeExtensionsType, map[string]attr.Value{
			"errors": func() types.Object {
				var NodeExtensionsErrorsValue = map[string]attr.Value{}
				if node.Extensions != nil {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(node.Extensions.Errors.Anim9ef),
						"enim3b":    types.StringValue(node.Extensions.Errors.Enim3b),
						"laborumd":  types.StringValue(node.Extensions.Errors.Laborumd),
						"mollit267": types.StringValue(node.Extensions.Errors.Mollit267),
					}
				} else {
					NodeExtensionsErrorsValue = map[string]attr.Value{
						"anim9ef":   types.StringValue(""),
						"enim3b":    types.StringValue(""),
						"laborumd":  types.StringValue(""),
						"mollit267": types.StringValue(""),
					}
				}

				item, diags := types.ObjectValue(NodeExtensionsErrorsType, NodeExtensionsErrorsValue)

				resp.Diagnostics.Append(diags...)

				if item.IsNull() {
					return types.ObjectNull(NodeExtensionsErrorsType)
				}

				return item
			}(),
			"installed": func() types.List {
				var installed []attr.Value
				if node.Extensions != nil {
					for _, extension := range node.Extensions.Installed {
						installed = append(installed, types.StringValue(extension))
					}
				} else {
					installed = append(installed, types.StringValue(""))
				}
				installedList, diags := types.ListValue(types.StringType, installed)
				resp.Diagnostics.Append(diags...)
				if installedList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return installedList
			}(),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		nodeValue := map[string]attr.Value{
			"name":                 types.StringValue(node.Name),
			"connection":           nodeConnectionValue,
			"location":             nodeLocationValue,
			"region":               nodeRegionValue,
			"distance_measurement": nodeDistanceMeasurementValue,
			"extensions":           nodeExtensionsValue,
		}

		node, diags := types.ObjectValue(NodeType, nodeValue)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}
		nodes = append(nodes, node)
	}

	plan.Nodes, diags = types.ListValue(types.ObjectType{
		AttrTypes: NodeType,
	}, nodes)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var databaseComponents types.List

	var databaseComponentsAttr []attr.Value

	for _, components := range database.Components {
		componentsValue, diags := types.ObjectValue(DatabaseComponentsItemsType, map[string]attr.Value{
			"name":         types.StringValue(components.Name),
			"id":           types.StringValue(components.ID),
			"release_date": types.StringValue(components.ReleaseDate),
			"version":      types.StringValue(components.Version),
			"status":       types.StringValue(components.Status),
		})

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		databaseComponentsAttr = append(databaseComponentsAttr, componentsValue)
	}

	databaseComponents, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseComponentsItemsType,
	}, databaseComponentsAttr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	plan.Components = databaseComponents

	var databaseRoles types.List

	var databaseRolesAttr []attr.Value

	for _, role := range database.Roles {
		var rolesValue types.Object
		if role != nil {
			rolesValue, diags = types.ObjectValue(DatabaseRolesItemsType, map[string]attr.Value{
				"bypass_rls":       types.BoolValue(role.BypassRls),
				"connection_limit": types.Int64Value(role.ConnectionLimit),
				"create_db":        types.BoolValue(role.CreateDb),
				"create_role":      types.BoolValue(role.CreateRole),
				"inherit":          types.BoolValue(role.Inherit),
				"login":            types.BoolValue(role.Login),
				"name":             types.StringValue(role.Name),
				"replication":      types.BoolValue(role.Replication),
				"superuser":        types.BoolValue(role.Superuser),
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}else{
			rolesValue = types.ObjectNull(DatabaseRolesItemsType)
		}

		databaseRolesAttr = append(databaseRolesAttr, rolesValue)
	}

	databaseRoles, diags = types.ListValue(types.ObjectType{
		AttrTypes: DatabaseRolesItemsType,
	}, databaseRolesAttr)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Roles = databaseRoles

	var databaseTables types.List

	var databaseTablesAttr []attr.Value

	if database.Tables != nil {
		for _, table := range database.Tables {
			var tableColumns types.List
	
			var tableColumnsAttr []attr.Value
	
			for _, column := range table.Columns {
				DatabaseTablesItemsColumnsItemsValue, diags := types.ObjectValue(DatabaseTablesItemsColumnsItemsType, map[string]attr.Value{
					"name":             types.StringValue(column.Name),
					"default":          types.StringValue(column.Default),
					"is_nullable":      types.BoolValue(column.IsNullable),
					"data_type":        types.StringValue(column.DataType),
					"is_primary_key":   types.BoolValue(column.IsPrimaryKey),
					"ordinal_position": types.Int64Value(column.OrdinalPosition),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableColumnsAttr = append(tableColumnsAttr, DatabaseTablesItemsColumnsItemsValue)
			}
	
			tableColumns, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsColumnsItemsType,
			}, tableColumnsAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			var tableStatus types.List
	
			var tableStatusAttr []attr.Value
	
			for _, status := range table.Status {
				DatabaseTablesItemsStatusItemsValue, diags := types.ObjectValue(DatabaseTablesItemsStatusItemsType, map[string]attr.Value{
					"aligned":     types.BoolValue(status.Aligned),
					"node_name":   types.StringValue(status.NodeName),
					"present":     types.BoolValue(status.Present),
					"replicating": types.BoolValue(status.Replicating),
				})
	
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
	
				tableStatusAttr = append(tableStatusAttr, DatabaseTablesItemsStatusItemsValue)
			}
	
			tableStatus, diags = types.ListValue(types.ObjectType{
				AttrTypes: DatabaseTablesItemsStatusItemsType,
			}, tableStatusAttr)
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			DatabaseTablesItemsValue, diags := types.ObjectValue(DatabaseTablesItemsType, map[string]attr.Value{
				"columns": tableColumns,
				"schema":  types.StringValue(table.Schema),
				"primary_key": func() types.List {
	
					var primaryKey types.List
	
					var primaryKeyAttr []attr.Value
	
					for _, pk := range table.PrimaryKey {
						primaryKeyAttr = append(primaryKeyAttr, types.StringValue(pk))
					}
	
					primaryKey, diags = types.ListValue(types.StringType, primaryKeyAttr)
	
					resp.Diagnostics.Append(diags...)
					return primaryKey
				}(),
				"replication_sets": func() types.List {
					var replicationSets types.List
	
					var replicationSetsAttr []attr.Value
	
					for _, rs := range table.ReplicationSets {
						replicationSetsAttr = append(replicationSetsAttr, types.StringValue(rs))
					}
	
					replicationSets, diags = types.ListValue(types.StringType, replicationSetsAttr)
	
					resp.Diagnostics.Append(diags...)
					return replicationSets
				}(),
				"name":   types.StringValue(table.Name),
				"status": tableStatus,
			})
	
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
	
			databaseTablesAttr = append(databaseTablesAttr, DatabaseTablesItemsValue)
		}

		databaseTables, diags = types.ListValue(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		}, databaseTablesAttr)
	
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}else{
		databaseTables = types.ListNull(types.ObjectType{
			AttrTypes: DatabaseTablesItemsType,
		})
	}

	plan.Tables = databaseTables

	DatabaseExtensionsValue, diags := types.ObjectValue(DatabaseExtensionsType, map[string]attr.Value{
		"auto_manage": types.BoolValue(database.Extensions.AutoManage),
		"available": func() types.List {
			if database.Extensions == nil || database.Extensions.Available == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var available []attr.Value
			if database.Extensions.Available != nil {
				for _, extension := range database.Extensions.Available {
					available = append(available, types.StringValue(extension))
				}
			} else {
				available = append(available, types.StringValue(""))
			}
			availableList, diags := types.ListValue(types.StringType, available)
			resp.Diagnostics.Append(diags...)
			
			return availableList
		}(),
		"requested": func() types.List {
			if database.Extensions == nil || database.Extensions.Requested == nil {
				val, _ := types.ListValue(types.StringType, nil)
				return val
			}
			var requested []attr.Value
			if database.Extensions.Requested != nil {
				for _, extension := range database.Extensions.Requested {
					requested = append(requested, types.StringValue(extension))
				}
			} else {
				requested = append(requested, types.StringValue(""))
			}
			requestedList, diags := types.ListValue(types.StringType, requested)
			resp.Diagnostics.Append(diags...)
			
			return requestedList
		}(),
	})

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Extensions = DatabaseExtensionsValue

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
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
