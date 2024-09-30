package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
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
	Databases []DatabaseModel `tfsdk:"databases"`
}

type DatabaseModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	ClusterID      types.String `tfsdk:"cluster_id"`
	Status         types.String `tfsdk:"status"`
	CreatedAt      types.String `tfsdk:"created_at"`
	// UpdatedAt      types.String `tfsdk:"updated_at"`
	PgVersion      types.String `tfsdk:"pg_version"`
	// StorageUsed    types.Int64  `tfsdk:"storage_used"`
	Domain         types.String `tfsdk:"domain"`
	ConfigVersion  types.String `tfsdk:"config_version"`
	Options        types.List   `tfsdk:"options"`
	Backups        types.Object `tfsdk:"backups"`
	Components     types.List   `tfsdk:"components"`
	Extensions     types.Object `tfsdk:"extensions"`
	Nodes          types.Map    `tfsdk:"nodes"`
	Roles          types.List   `tfsdk:"roles"`
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
						"cluster_id": schema.StringAttribute{
							Computed:    true,
							Description: "ID of the cluster this database belongs to",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Status of the database",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "Creation timestamp of the database",
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: "Last update timestamp of the database",
						},
						"pg_version": schema.StringAttribute{
							Computed:    true,
							Description: "PostgreSQL version of the database",
						},
						// "storage_used": schema.Int64Attribute{
						// 	Computed:    true,
						// 	Description: "Storage used by the database in bytes",
						// },
						"domain": schema.StringAttribute{
							Computed:    true,
							Description: "Domain of the database",
						},
						"config_version": schema.StringAttribute{
							Computed:    true,
							Description: "Configuration version of the database",
						},
						"options": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "Options for the database",
						},
						"backups": schema.SingleNestedAttribute{
							Computed:    true,
							Description: "Backup configuration for the database",
							Attributes:  d.backupsSchema(),
						},
						"components": schema.ListNestedAttribute{
							Computed:    true,
							Description: "Components of the database",
							NestedObject: schema.NestedAttributeObject{
								Attributes: d.componentSchema(),
							},
						},
						"extensions": schema.SingleNestedAttribute{
							Computed:    true,
							Description: "Extensions configuration for the database",
							Attributes:  d.extensionsSchema(),
						},
						"nodes": schema.MapNestedAttribute{
    				Computed:    true,
    Description: "Map of nodes in the database",
    NestedObject: schema.NestedAttributeObject{
        Attributes: d.nodeSchema(),
    },
},
						"roles": schema.ListNestedAttribute{
							Computed:    true,
							Description: "Roles in the database",
							NestedObject: schema.NestedAttributeObject{
								Attributes: d.roleSchema(),
							},
						},
					},
				},
			},
		},
		Description: "Data source for pgEdge databases",
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
		databaseModel := DatabaseModel{
			ID:             types.StringValue(db.ID.String()),
			Name:           types.StringPointerValue(db.Name),
			ClusterID:      types.StringValue(db.ClusterID.String()),
			Status:         types.StringPointerValue(db.Status),
			CreatedAt:      types.StringPointerValue(db.CreatedAt),
			// UpdatedAt:      types.StringPointerValue(db.UpdatedAt),
			PgVersion:      types.StringValue(db.PgVersion),
			// StorageUsed:    types.Int64Value(db.StorageUsed),
			Domain:         types.StringValue(db.Domain),
			ConfigVersion:  types.StringValue(db.ConfigVersion),
			Options:        d.convertToListValue(db.Options),
			Backups:        d.mapBackupsToModel(db.Backups),
			Components:     d.mapComponentsToModel(db.Components),
			Extensions:     d.mapExtensionsToModel(db.Extensions),
			Nodes:          d.mapNodesToModel(db.Nodes),
			Roles:          d.mapRolesToModel(db.Roles),
		}

		state.Databases = append(state.Databases, databaseModel)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Helper methods for schema definition

func (d *databasesDataSource) backupsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"provider": schema.StringAttribute{
			Computed:    true,
			Description: "Backup provider",
		},
		"config": schema.ListNestedAttribute{
			Computed:    true,
			Description: "Backup configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: "Backup configuration ID",
					},
					"node_name": schema.StringAttribute{
						Computed:    true,
						Description: "Node name",
					},
					"repositories": schema.ListNestedAttribute{
                        Computed:    true,
                        Description: "Backup repositories",
                        NestedObject: schema.NestedAttributeObject{
                            Attributes: d.backupRepositorySchema(),
                        },
                    },
                    "schedules": schema.ListNestedAttribute{
                        Computed:    true,
                        Description: "Backup schedules",
                        NestedObject: schema.NestedAttributeObject{
                            Attributes: d.backupScheduleSchema(),
                        },
                    },
				},
			},
		},
	}
}

func (d *databasesDataSource) backupRepositorySchema() map[string]schema.Attribute {
    return map[string]schema.Attribute{
        "id":                   schema.StringAttribute{Computed: true},
        "type":                 schema.StringAttribute{Computed: true},
        "backup_store_id":      schema.StringAttribute{Computed: true},
        "base_path":            schema.StringAttribute{Computed: true},
        "s3_bucket":            schema.StringAttribute{Computed: true},
        "s3_region":            schema.StringAttribute{Computed: true},
        "s3_endpoint":          schema.StringAttribute{Computed: true},
        "gcs_bucket":           schema.StringAttribute{Computed: true},
        "gcs_endpoint":         schema.StringAttribute{Computed: true},
        "azure_account":        schema.StringAttribute{Computed: true},
        "azure_container":      schema.StringAttribute{Computed: true},
        "azure_endpoint":       schema.StringAttribute{Computed: true},
        "retention_full":       schema.Int64Attribute{Computed: true},
        "retention_full_type":  schema.StringAttribute{Computed: true},
    }
}

func (d *databasesDataSource) backupScheduleSchema() map[string]schema.Attribute {
    return map[string]schema.Attribute{
        "id":               schema.StringAttribute{Computed: true},
        "type":             schema.StringAttribute{Computed: true},
        "cron_expression":  schema.StringAttribute{Computed: true},
    }
}


func (d *databasesDataSource) componentSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Component ID",
		},
		"name": schema.StringAttribute{
			Computed:    true,
			Description: "Component name",
		},
		"version": schema.StringAttribute{
			Computed:    true,
			Description: "Component version",
		},
		"release_date": schema.StringAttribute{
			Computed:    true,
			Description: "Component release date",
		},
		"status": schema.StringAttribute{
			Computed:    true,
			Description: "Component status",
		},
	}
}

func (d *databasesDataSource) extensionsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"auto_manage": schema.BoolAttribute{
			Computed:    true,
			Description: "Auto-manage extensions",
		},
		"available": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
			Description: "Available extensions",
		},
		"requested": schema.ListAttribute{
			Computed:    true,
			ElementType: types.StringType,
			Description: "Requested extensions",
		},
	}
}

func (d *databasesDataSource) nodeSchema() map[string]schema.Attribute {
    return map[string]schema.Attribute{
        "name": schema.StringAttribute{
            Computed:    true,
            Description: "Name of the node",
        },
        "connection": schema.SingleNestedAttribute{
            Computed:    true,
            Description: "Node connection details",
            Attributes:  d.connectionSchema(),
        },
        "location": schema.SingleNestedAttribute{
            Computed:    true,
            Description: "Node location",
            Attributes:  d.locationSchema(),
        },
        "region": schema.SingleNestedAttribute{
            Computed:    true,
            Description: "Node region",
            Attributes:  d.regionSchema(),
        },
        "extensions": schema.SingleNestedAttribute{
            Computed:    true,
            Description: "Node extensions",
            Attributes:  d.nodeExtensionsSchema(),
        },
    }
}



func (d *databasesDataSource) connectionSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"database":            schema.StringAttribute{Computed: true},
		"host":                schema.StringAttribute{Computed: true},
		"password":            schema.StringAttribute{Computed: true},
		"port":                schema.Int64Attribute{Computed: true},
		"username":            schema.StringAttribute{Computed: true},
		"external_ip_address": schema.StringAttribute{Computed: true},
		"internal_ip_address": schema.StringAttribute{Computed: true},
		"internal_host":       schema.StringAttribute{Computed: true},
	}
}

func (d *databasesDataSource) locationSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"code":        schema.StringAttribute{Computed: true},
		"country":     schema.StringAttribute{Computed: true},
		"latitude":    schema.Float64Attribute{Computed: true},
		"longitude":   schema.Float64Attribute{Computed: true},
		"name":        schema.StringAttribute{Computed: true},
		"region":      schema.StringAttribute{Computed: true},
		"region_code": schema.StringAttribute{Computed: true},
		"timezone":    schema.StringAttribute{Computed: true},
		"postal_code": schema.StringAttribute{Computed: true},
		"metro_code":  schema.StringAttribute{Computed: true},
		"city":        schema.StringAttribute{Computed: true},
	}
}

func (d *databasesDataSource) regionSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"active":             schema.BoolAttribute{Computed: true},
		"availability_zones": schema.ListAttribute{Computed: true, ElementType: types.StringType},
		"cloud":              schema.StringAttribute{Computed: true},
		"code":               schema.StringAttribute{Computed: true},
		"name":               schema.StringAttribute{Computed: true},
		"parent":             schema.StringAttribute{Computed: true},
	}
}

func (d *databasesDataSource) nodeExtensionsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"errors":    schema.MapAttribute{Computed: true, ElementType: types.StringType},
		"installed": schema.ListAttribute{Computed: true, ElementType: types.StringType},
	}
}

func (d *databasesDataSource) roleSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name":             schema.StringAttribute{Computed: true},
		"bypass_rls":       schema.BoolAttribute{Computed: true},
		"connection_limit": schema.Int64Attribute{Computed: true},
		"create_db":        schema.BoolAttribute{Computed: true},
		"create_role":      schema.BoolAttribute{Computed: true},
		"inherit":          schema.BoolAttribute{Computed: true},
		"login":            schema.BoolAttribute{Computed: true},
		"replication":      schema.BoolAttribute{Computed: true},
		"superuser":        schema.BoolAttribute{Computed: true},
	}
}


// Helper methods for mapping API responses to model

func (d *databasesDataSource) convertToListValue(slice []string) types.List {
    elements := make([]attr.Value, len(slice))
    for i, s := range slice {
        elements[i] = types.StringValue(s)
    }
    return types.ListValueMust(types.StringType, elements)
}

func (d *databasesDataSource) mapBackupsToModel(backups *models.Backups) types.Object {
    if backups == nil {
        return types.ObjectNull(map[string]attr.Type{
            "provider": types.StringType,
            "config":   types.ListType{ElemType: types.ObjectType{AttrTypes: d.backupConfigType()}},
        })
    }

    configList := []attr.Value{}
    for _, config := range backups.Config {
        configObj, _ := types.ObjectValue(
            d.backupConfigType(),
            map[string]attr.Value{
                "id":           types.StringPointerValue(config.ID),
                "node_name":    types.StringValue(config.NodeName),
                "repositories": d.mapBackupRepositoriesToModel(config.Repositories),
                "schedules":    d.mapBackupSchedulesToModel(config.Schedules),
            },
        )
        configList = append(configList, configObj)
    }

    backupsObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "provider": types.StringType,
            "config":   types.ListType{ElemType: types.ObjectType{AttrTypes: d.backupConfigType()}},
        },
        map[string]attr.Value{
            "provider": types.StringPointerValue(backups.Provider),
            "config":   types.ListValueMust(types.ObjectType{AttrTypes: d.backupConfigType()}, configList),
        },
    )

    return backupsObj
}

func (d *databasesDataSource) backupConfigType() map[string]attr.Type {
    return map[string]attr.Type{
        "id":           types.StringType,
        "node_name":    types.StringType,
        "repositories": types.ListType{ElemType: types.ObjectType{AttrTypes: d.backupRepositoryType()}},
        "schedules":    types.ListType{ElemType: types.ObjectType{AttrTypes: d.backupScheduleType()}},
    }
}

func (d *databasesDataSource) backupRepositoryType() map[string]attr.Type {
    return map[string]attr.Type{
        "id":                   types.StringType,
        "type":                 types.StringType,
        "backup_store_id":      types.StringType,
        "base_path":            types.StringType,
        "s3_bucket":            types.StringType,
        "s3_region":            types.StringType,
        "s3_endpoint":          types.StringType,
        "gcs_bucket":           types.StringType,
        "gcs_endpoint":         types.StringType,
        "azure_account":        types.StringType,
        "azure_container":      types.StringType,
        "azure_endpoint":       types.StringType,
        "retention_full":       types.Int64Type,
        "retention_full_type":  types.StringType,
    }
}

func (d *databasesDataSource) backupScheduleType() map[string]attr.Type {
    return map[string]attr.Type{
        "id":               types.StringType,
        "type":             types.StringType,
        "cron_expression":  types.StringType,
    }
}

func (d *databasesDataSource) mapBackupRepositoriesToModel(repositories []*models.BackupRepository) types.List {
    repoList := []attr.Value{}
    for _, repo := range repositories {
        repoObj, _ := types.ObjectValue(
            d.backupRepositoryType(),
            map[string]attr.Value{
                "id":                   types.StringValue(repo.ID),
                "type":                 types.StringValue(repo.Type),
                "backup_store_id":      types.StringValue(repo.BackupStoreID),
                "base_path":            types.StringValue(repo.BasePath),
                "s3_bucket":            types.StringValue(repo.S3Bucket),
                "s3_region":            types.StringValue(repo.S3Region),
                "s3_endpoint":          types.StringValue(repo.S3Endpoint),
                "gcs_bucket":           types.StringValue(repo.GcsBucket),
                "gcs_endpoint":         types.StringValue(repo.GcsEndpoint),
                "azure_account":        types.StringValue(repo.AzureAccount),
                "azure_container":      types.StringValue(repo.AzureContainer),
                "azure_endpoint":       types.StringValue(repo.AzureEndpoint),
                "retention_full":       types.Int64Value(repo.RetentionFull),
                "retention_full_type":  types.StringValue(repo.RetentionFullType),
            },
        )
        repoList = append(repoList, repoObj)
    }
    return types.ListValueMust(types.ObjectType{AttrTypes: d.backupRepositoryType()}, repoList)
}

func (d *databasesDataSource) mapBackupSchedulesToModel(schedules []*models.BackupSchedule) types.List {
    scheduleList := []attr.Value{}
    for _, schedule := range schedules {
        scheduleObj, _ := types.ObjectValue(
            d.backupScheduleType(),
            map[string]attr.Value{
                "id":               types.StringPointerValue(schedule.ID),
                "type":             types.StringPointerValue(schedule.Type),
                "cron_expression":  types.StringPointerValue(schedule.CronExpression),
            },
        )
        scheduleList = append(scheduleList, scheduleObj)
    }
    return types.ListValueMust(types.ObjectType{AttrTypes: d.backupScheduleType()}, scheduleList)
}

func (d *databasesDataSource) mapComponentsToModel(components []*models.DatabaseComponentVersion) types.List {
    componentsList := []attr.Value{}
    componentAttrTypes := map[string]attr.Type{
        "id":           types.StringType,
        "name":         types.StringType,
        "version":      types.StringType,
        "release_date": types.StringType,
        "status":       types.StringType,
    }

    for _, component := range components {
        componentObj, _ := types.ObjectValue(
            componentAttrTypes,
            map[string]attr.Value{
                "id":           types.StringValue(component.ID.String()),
                "name":         types.StringPointerValue(component.Name),
                "version":      types.StringPointerValue(component.Version),
                "release_date": types.StringPointerValue(component.ReleaseDate),
                "status":       types.StringPointerValue(component.Status),
            },
        )
        componentsList = append(componentsList, componentObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: componentAttrTypes}, componentsList)
}

func (d *databasesDataSource) mapExtensionsToModel(extensions *models.Extensions) types.Object {
    if extensions == nil {
        return types.ObjectNull(map[string]attr.Type{
            "auto_manage": types.BoolType,
            "available":   types.ListType{ElemType: types.StringType},
            "requested":   types.ListType{ElemType: types.StringType},
        })
    }

    extensionsObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "auto_manage": types.BoolType,
            "available":   types.ListType{ElemType: types.StringType},
            "requested":   types.ListType{ElemType: types.StringType},
        },
        map[string]attr.Value{
            "auto_manage": types.BoolValue(extensions.AutoManage),
            "available":   d.convertToListValue(extensions.Available),
            "requested":   d.convertToListValue(extensions.Requested),
        },
    )

    return extensionsObj
}

func (d *databasesDataSource) connectionAttrTypes() map[string]attr.Type {
    return map[string]attr.Type{
        "database":            types.StringType,
        "host":                types.StringType,
        "password":            types.StringType,
        "port":                types.Int64Type,
        "username":            types.StringType,
        "external_ip_address": types.StringType,
        "internal_ip_address": types.StringType,
        "internal_host":       types.StringType,
    }
}

func (d *databasesDataSource) locationAttrTypes() map[string]attr.Type {
    return map[string]attr.Type{
        "code":        types.StringType,
        "country":     types.StringType,
        "latitude":    types.Float64Type,
        "longitude":   types.Float64Type,
        "name":        types.StringType,
        "region":      types.StringType,
        "region_code": types.StringType,
        "timezone":    types.StringType,
        "postal_code": types.StringType,
        "metro_code":  types.StringType,
        "city":        types.StringType,
    }
}

func (d *databasesDataSource) regionAttrTypes() map[string]attr.Type {
    return map[string]attr.Type{
        "active":             types.BoolType,
        "availability_zones": types.ListType{ElemType: types.StringType},
        "cloud":              types.StringType,
        "code":               types.StringType,
        "name":               types.StringType,
        "parent":             types.StringType,
    }
}

func (d *databasesDataSource) nodeExtensionsAttrTypes() map[string]attr.Type {
    return map[string]attr.Type{
        "errors":    types.MapType{ElemType: types.StringType},
        "installed": types.ListType{ElemType: types.StringType},
    }
}

func (d *databasesDataSource) mapNodesToModel(nodes []*models.DatabaseNode) types.Map {
    nodeMap := make(map[string]attr.Value)
    nodeAttrTypes := map[string]attr.Type{
        "name":       types.StringType,
        "connection": types.ObjectType{AttrTypes: d.connectionAttrTypes()},
        "location":   types.ObjectType{AttrTypes: d.locationAttrTypes()},
        "region":     types.ObjectType{AttrTypes: d.regionAttrTypes()},
        "extensions": types.ObjectType{AttrTypes: d.nodeExtensionsAttrTypes()},
    }

    for _, node := range nodes {
        if node.Name == nil {
            continue
        }
        nodeObj, _ := types.ObjectValue(
            nodeAttrTypes,
            map[string]attr.Value{
                "name":       types.StringPointerValue(node.Name),
                "connection": d.mapConnectionToModel(node.Connection),
                "location":   d.mapLocationToModel(node.Location),
                "region":     d.mapRegionToModel(node.Region),
                "extensions": d.mapNodeExtensionsToModel(node.Extensions),
            },
        )
        nodeMap[*node.Name] = nodeObj
    }

    return types.MapValueMust(types.ObjectType{AttrTypes: nodeAttrTypes}, nodeMap)
}

func (d *databasesDataSource) mapConnectionToModel(connection *models.Connection) types.Object {
    if connection == nil {
        return types.ObjectNull(d.connectionAttrTypes())
    }
    connectionObj, _ := types.ObjectValue(
        d.connectionAttrTypes(),
        map[string]attr.Value{
            "database":            types.StringPointerValue(connection.Database),
            "host":                types.StringValue(connection.Host),
            "password":            types.StringPointerValue(connection.Password),
            "port":                types.Int64PointerValue(connection.Port),
            "username":            types.StringPointerValue(connection.Username),
            "external_ip_address": types.StringValue(connection.ExternalIPAddress),
            "internal_ip_address": types.StringValue(connection.InternalIPAddress),
            "internal_host":       types.StringValue(connection.InternalHost),
        },
    )
    return connectionObj
}

func (d *databasesDataSource) mapLocationToModel(location *models.Location) types.Object {
    if location == nil {
        return types.ObjectNull(d.locationAttrTypes())
    }
    locationObj, _ := types.ObjectValue(
        d.locationAttrTypes(),
        map[string]attr.Value{
            "code":        types.StringValue(location.Code),
            "country":     types.StringValue(location.Country),
            "latitude":    types.Float64PointerValue(location.Latitude),
            "longitude":   types.Float64PointerValue(location.Longitude),
            "name":        types.StringValue(location.Name),
            "region":      types.StringValue(location.Region),
            "region_code": types.StringValue(location.RegionCode),
            "timezone":    types.StringValue(location.Timezone),
            "postal_code": types.StringValue(location.PostalCode),
            "metro_code":  types.StringValue(location.MetroCode),
            "city":        types.StringValue(location.City),
        },
    )
    return locationObj
}

func (d *databasesDataSource) mapRegionToModel(region *models.Region) types.Object {
    if region == nil {
        return types.ObjectNull(d.regionAttrTypes())
    }
    regionObj, _ := types.ObjectValue(
        d.regionAttrTypes(),
        map[string]attr.Value{
            "active":             types.BoolValue(region.Active),
            "availability_zones": types.ListValueMust(types.StringType, d.stringSliceToValueSlice(region.AvailabilityZones)),
            "cloud":              types.StringPointerValue(region.Cloud),
            "code":               types.StringPointerValue(region.Code),
            "name":               types.StringPointerValue(region.Name),
            "parent":             types.StringValue(region.Parent),
        },
    )
    return regionObj
}

func (d *databasesDataSource) mapNodeExtensionsToModel(extensions *models.DatabaseNodeExtensions) types.Object {
    if extensions == nil {
        return types.ObjectNull(d.nodeExtensionsAttrTypes())
    }
    errorsMap := make(map[string]attr.Value)
    for k, v := range extensions.Errors {
        errorsMap[k] = types.StringValue(v)
    }
    extensionsObj, _ := types.ObjectValue(
        d.nodeExtensionsAttrTypes(),
        map[string]attr.Value{
            "errors":    types.MapValueMust(types.StringType, errorsMap),
            "installed": types.ListValueMust(types.StringType, d.stringSliceToValueSlice(extensions.Installed)),
        },
    )
    return extensionsObj
}

func (d *databasesDataSource) mapRolesToModel(roles []*models.DatabaseRole) types.List {
    rolesList := []attr.Value{}
    roleAttrTypes := map[string]attr.Type{
        "name":             types.StringType,
        "bypass_rls":       types.BoolType,
        "connection_limit": types.Int64Type,
        "create_db":        types.BoolType,
        "create_role":      types.BoolType,
        "inherit":          types.BoolType,
        "login":            types.BoolType,
        "replication":      types.BoolType,
        "superuser":        types.BoolType,
    }

    for _, role := range roles {
        roleObj, _ := types.ObjectValue(
            roleAttrTypes,
            map[string]attr.Value{
                "name":             types.StringPointerValue(role.Name),
                "bypass_rls":       types.BoolPointerValue(role.BypassRls),
                "connection_limit": types.Int64PointerValue(role.ConnectionLimit),
                "create_db":        types.BoolPointerValue(role.CreateDb),
                "create_role":      types.BoolPointerValue(role.CreateRole),
                "inherit":          types.BoolPointerValue(role.Inherit),
                "login":            types.BoolPointerValue(role.Login),
                "replication":      types.BoolPointerValue(role.Replication),
                "superuser":        types.BoolPointerValue(role.Superuser),
            },
        )
        rolesList = append(rolesList, roleObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: roleAttrTypes}, rolesList)
}

// Helper function to convert []string to []attr.Value
func (d *databasesDataSource) stringSliceToValueSlice(slice []string) []attr.Value {
    valueSlice := make([]attr.Value, len(slice))
    for i, s := range slice {
        valueSlice[i] = types.StringValue(s)
    }
    return valueSlice
}