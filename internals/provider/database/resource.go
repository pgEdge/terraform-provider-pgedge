package database

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	// "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

var (
	_ resource.Resource                = &databaseResource{}
	_ resource.ResourceWithConfigure   = &databaseResource{}
	_ resource.ResourceWithImportState = &databaseResource{}
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

func (r *databaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Description: "Manages a pgEdge database.",
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
                Description: "Unique identifier for the database.",
                Computed:    true,
                PlanModifiers: []planmodifier.String{
                    stringplanmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the database.",
                Required:    true,
            },
            "cluster_id": schema.StringAttribute{
                Description: "The ID of the cluster this database belongs to.",
                Required:    true,
            },
            "status": schema.StringAttribute{
                Description: "The current status of the database.",
                Computed:    true,
            },
            "created_at": schema.StringAttribute{
                Description: "The timestamp when the database was created.",
                Computed:    true,
            },
            "updated_at": schema.StringAttribute{
                Description: "The timestamp when the database was last updated.",
                Computed:    true,
            },
            "pg_version": schema.StringAttribute{
                Description: "The PostgreSQL version of the database.",
                Computed:    true,
            },
            "storage_used": schema.Int64Attribute{
                Description: "The amount of storage used by the database in bytes.",
                Computed:    true,
            },
            "domain": schema.StringAttribute{
                Description: "The domain associated with the database.",
                Computed:    true,
            },
            "config_version": schema.StringAttribute{
                Description: "The configuration version of the database.",
                Optional:    true,
                Computed:    true,
            },
            "options": schema.ListAttribute{
                Description: "A list of options for the database.",
                ElementType: types.StringType,
                Optional:    true,
            },
            // "backups": schema.SingleNestedAttribute{
            //     Description: "Backup configuration for the database.",
            //     Computed:    true,
			// 	Optional:    true,
            //     Attributes: map[string]schema.Attribute{
            //         "provider": schema.StringAttribute{
            //             Description: "The backup provider.",
            //             Computed:    true,
			// 			Optional:    true,
            //         },
            //         "config": schema.ListNestedAttribute{
            //             Description: "List of backup configurations.",
            //             Computed:    true,
			// 			Optional:    true,
            //             NestedObject: schema.NestedAttributeObject{
            //                 Attributes: map[string]schema.Attribute{
            //                     "id": schema.StringAttribute{
            //                         Description: "Unique identifier for the backup config.",
            //                         Computed:    true,
            //                     },
            //                     "node_name": schema.StringAttribute{
            //                         Description: "Name of the node.",
            //                         Computed:    true,
			// 						Optional:    true,
            //                     },
            //                     "repositories": schema.ListNestedAttribute{
            //                         Description: "List of backup repositories.",
            //                         Computed:    true,
            //                         NestedObject: schema.NestedAttributeObject{
            //                             Attributes: map[string]schema.Attribute{
            //                                 "id": schema.StringAttribute{
            //                                     Description: "Unique identifier for the repository.",
            //                                     Computed:    true,
            //                                 },
            //                                 "type": schema.StringAttribute{
            //                                     Description: "Type of the repository.",
            //                                     Computed:    true,
            //                                 },
            //                                 "backup_store_id": schema.StringAttribute{
            //                                     Description: "ID of the backup store.",
            //                                     Computed:    true,
            //                                 },
            //                                 "base_path": schema.StringAttribute{
            //                                     Description: "Base path for the repository.",
            //                                     Computed:    true,
            //                                 },
            //                                 "s3_bucket": schema.StringAttribute{
            //                                     Description: "S3 bucket name.",
            //                                     Computed:    true,
            //                                 },
            //                                 "s3_region": schema.StringAttribute{
            //                                     Description: "S3 region.",
            //                                     Computed:    true,
            //                                 },
            //                                 "s3_endpoint": schema.StringAttribute{
            //                                     Description: "S3 endpoint.",
            //                                     Computed:    true,
            //                                 },
            //                                 "gcs_bucket": schema.StringAttribute{
            //                                     Description: "GCS bucket name.",
            //                                     Computed:    true,
            //                                 },
            //                                 "gcs_endpoint": schema.StringAttribute{
            //                                     Description: "GCS endpoint.",
            //                                     Computed:    true,
            //                                 },
            //                                 "azure_account": schema.StringAttribute{
            //                                     Description: "Azure account.",
            //                                     Computed:    true,
            //                                 },
            //                                 "azure_container": schema.StringAttribute{
            //                                     Description: "Azure container.",
            //                                     Computed:    true,
            //                                 },
            //                                 "azure_endpoint": schema.StringAttribute{
            //                                     Description: "Azure endpoint.",
            //                                     Computed:    true,
            //                                 },
            //                                 "retention_full": schema.Int64Attribute{
            //                                     Description: "Retention period for full backups.",
            //                                     Computed:    true,
            //                                 },
            //                                 "retention_full_type": schema.StringAttribute{
            //                                     Description: "Type of retention for full backups.",
            //                                     Computed:    true,
            //                                 },
            //                             },
            //                         },
            //                     },
            //                     "schedules": schema.ListNestedAttribute{
            //                         Description: "List of backup schedules.",
            //                         Computed:    true,
			// 						Optional:    true,
            //                         NestedObject: schema.NestedAttributeObject{
            //                             Attributes: map[string]schema.Attribute{
            //                                 "id": schema.StringAttribute{
            //                                     Description: "Unique identifier for the schedule.",
            //                                     Computed:    true,
            //                                 },
            //                                 "type": schema.StringAttribute{
            //                                     Description: "Type of the schedule.",
            //                                     Computed:    true,
			// 									Optional:    true,
            //                                 },
            //                                 "cron_expression": schema.StringAttribute{
            //                                     Description: "Cron expression for the schedule.",
            //                                     Computed:    true,
			// 									Optional:    true,
            //                                 },
            //                             },
            //                         },
            //                     },
            //                 },
            //             },
            //         },
            //     },
            // },
            // "components": schema.ListNestedAttribute{
            //     Description: "List of components in the database.",
            //     Computed:    true,
            //     NestedObject: schema.NestedAttributeObject{
            //         Attributes: map[string]schema.Attribute{
            //             "id":           schema.StringAttribute{Computed: true},
            //             "name":         schema.StringAttribute{Computed: true},
            //             "version":      schema.StringAttribute{Computed: true},
            //             "release_date": schema.StringAttribute{Computed: true},
            //             "status":       schema.StringAttribute{Computed: true},
            //         },
            //     },
            // },
            "extensions": schema.SingleNestedAttribute{
                Description: "Extensions configuration for the database.",
                Computed:    true,
				Optional:    true,
                Attributes: map[string]schema.Attribute{
                    "auto_manage": schema.BoolAttribute{Computed: true, Optional: true},
                    "available":   schema.ListAttribute{Computed: true, ElementType: types.StringType},
                    "requested":   schema.ListAttribute{Computed: true, Optional: true, ElementType: types.StringType},
                },
            },
            // "nodes": schema.ListNestedAttribute{
            //     Description: "List of nodes in the database.",
            //     Computed:    true,
            //     NestedObject: schema.NestedAttributeObject{
            //         Attributes: map[string]schema.Attribute{
            //             "name": schema.StringAttribute{Computed: true},
            //             "connection": schema.SingleNestedAttribute{
            //                 Computed: true,
            //                 Attributes: map[string]schema.Attribute{
            //                     "database":            schema.StringAttribute{Computed: true},
            //                     "host":                schema.StringAttribute{Computed: true},
            //                     "password":            schema.StringAttribute{Computed: true, Sensitive: true},
            //                     "port":                schema.Int64Attribute{Computed: true},
            //                     "username":            schema.StringAttribute{Computed: true},
            //                     "external_ip_address": schema.StringAttribute{Computed: true},
            //                     "internal_ip_address": schema.StringAttribute{Computed: true},
            //                     "internal_host":       schema.StringAttribute{Computed: true},
            //                 },
            //             },
            //             "location": schema.SingleNestedAttribute{
            //                 Computed: true,
            //                 Attributes: map[string]schema.Attribute{
            //                     "code":        schema.StringAttribute{Computed: true},
            //                     "country":     schema.StringAttribute{Computed: true},
            //                     "latitude":    schema.Float64Attribute{Computed: true},
            //                     "longitude":   schema.Float64Attribute{Computed: true},
            //                     "name":        schema.StringAttribute{Computed: true},
            //                     "region":      schema.StringAttribute{Computed: true},
            //                     "region_code": schema.StringAttribute{Computed: true},
            //                     "timezone":    schema.StringAttribute{Computed: true},
            //                     "postal_code": schema.StringAttribute{Computed: true},
            //                     "metro_code":  schema.StringAttribute{Computed: true},
            //                     "city":        schema.StringAttribute{Computed: true},
            //                 },
            //             },
            //             "region": schema.SingleNestedAttribute{
            //                 Computed: true,
            //                 Attributes: map[string]schema.Attribute{
            //                     "active":             schema.BoolAttribute{Computed: true},
            //                     "availability_zones": schema.ListAttribute{Computed: true, ElementType: types.StringType},
            //                     "cloud":              schema.StringAttribute{Computed: true},
            //                     "code":               schema.StringAttribute{Computed: true},
            //                     "name":               schema.StringAttribute{Computed: true},
            //                     "parent":             schema.StringAttribute{Computed: true},
            //                 },
            //             },
            //             "extensions": schema.SingleNestedAttribute{
            //                 Computed: true,
            //                 Attributes: map[string]schema.Attribute{
            //                     "errors":    schema.MapAttribute{Computed: true, ElementType: types.StringType},
            //                     "installed": schema.ListAttribute{Computed: true, ElementType: types.StringType},
            //                 },
            //             },
            //         },
            //     },
            // },
            // "roles": schema.ListNestedAttribute{
            //     Description: "List of roles in the database.",
            //     Computed:    true,
            //     NestedObject: schema.NestedAttributeObject{
            //         Attributes: map[string]schema.Attribute{
            //             "name":             schema.StringAttribute{Computed: true},
            //             "bypass_rls":       schema.BoolAttribute{Computed: true},
            //             "connection_limit": schema.Int64Attribute{Computed: true},
            //             "create_db":        schema.BoolAttribute{Computed: true},
            //             "create_role":      schema.BoolAttribute{Computed: true},
            //             "inherit":          schema.BoolAttribute{Computed: true},
            //             "login":            schema.BoolAttribute{Computed: true},
            //             "replication":      schema.BoolAttribute{Computed: true},
            //             "superuser":        schema.BoolAttribute{Computed: true},
            //         },
            //     },
            // },
            // "tables": schema.ListNestedAttribute{
            //     Description: "List of tables in the database.",
            //     Computed:    true,
            //     NestedObject: schema.NestedAttributeObject{
            //         Attributes: map[string]schema.Attribute{
            //             "name":             schema.StringAttribute{Computed: true},
            //             "schema":           schema.StringAttribute{Computed: true},
            //             "primary_key":      schema.ListAttribute{Computed: true, ElementType: types.StringType},
            //             "replication_sets": schema.ListAttribute{Computed: true, ElementType: types.StringType},
            //             "columns": schema.ListNestedAttribute{
            //                 Computed: true,
            //                 NestedObject: schema.NestedAttributeObject{
            //                     Attributes: map[string]schema.Attribute{
            //                         "name":             schema.StringAttribute{Computed: true},
            //                         "data_type":        schema.StringAttribute{Computed: true},
            //                         "default":          schema.StringAttribute{Computed: true},
            //                         "is_nullable":      schema.BoolAttribute{Computed: true},
            //                         "is_primary_key":   schema.BoolAttribute{Computed: true},
            //                         "ordinal_position": schema.Int64Attribute{Computed: true},
            //                     },
            //                 },
            //             },
            //             "status": schema.ListNestedAttribute{
            //                 Computed: true,
            //                 NestedObject: schema.NestedAttributeObject{
            //                     Attributes: map[string]schema.Attribute{
            //                         "aligned":     schema.BoolAttribute{Computed: true},
            //                         "node_name":   schema.StringAttribute{Computed: true},
            //                         "present":     schema.BoolAttribute{Computed: true},
            //                         "replicating": schema.BoolAttribute{Computed: true},
            //                     },
            //                 },
            //             },
            //         },
            //     },
            // },
        },
    }
}


func (r *databaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pgEdge.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *pgEdge.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

type extensionsModel struct {
    AutoManage types.Bool   `tfsdk:"auto_manage"`
    Available  types.List   `tfsdk:"available"`
    Requested  types.List   `tfsdk:"requested"`
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan databaseResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    createInput := &models.CreateDatabaseInput{
        Name:           plan.Name.ValueStringPointer(),
        ClusterID:      strfmt.UUID(plan.ClusterID.ValueString()),
        ConfigVersion:  plan.ConfigVersion.ValueString(),
        Options:        convertToStringSlice(plan.Options),
    }

	 // Handle Extensions
     if !plan.Extensions.IsNull() && !plan.Extensions.IsUnknown() {
        var extensionsData extensionsModel
        diags := plan.Extensions.As(ctx, &extensionsData, basetypes.ObjectAsOptions{})
        resp.Diagnostics.Append(diags...)
        if resp.Diagnostics.HasError() {
            return
        }
        
        createInput.Extensions = &models.Extensions{
            AutoManage: extensionsData.AutoManage.ValueBool(),
            Requested:  convertTFListToStringSlice(extensionsData.Requested),
        }
    }

    // Handle Backups
    // if !plan.Backups.IsNull() {
    //     var backupsData struct {
    //         Provider string `tfsdk:"provider"`
    //         Config   []struct {
    //             NodeName  string `tfsdk:"node_name"`
    //             Schedules []struct {
    //                 Type            string `tfsdk:"type"`
    //                 CronExpression string `tfsdk:"cron_expression"`
    //             } `tfsdk:"schedules"`
    //         } `tfsdk:"config"`
    //     }
    //     diags := plan.Backups.As(ctx, &backupsData, basetypes.ObjectAsOptions{})
    //     resp.Diagnostics.Append(diags...)
    //     if resp.Diagnostics.HasError() {
    //         return
    //     }

    //     createInput.Backups = &models.Backups{
    //         Provider: &backupsData.Provider,
    //         Config:   make([]*models.BackupConfig, len(backupsData.Config)),
    //     }

    //     for i, config := range backupsData.Config {
    //         createInput.Backups.Config[i] = &models.BackupConfig{
    //             NodeName:  config.NodeName,
    //             Schedules: make([]*models.BackupSchedule, len(config.Schedules)),
    //         }
    //         for j, schedule := range config.Schedules {
    //             createInput.Backups.Config[i].Schedules[j] = &models.BackupSchedule{
    //                 Type:           &schedule.Type,
    //                 CronExpression: &schedule.CronExpression,
    //             }
    //         }
    //     }
    // }

    tflog.Debug(ctx, "Creating pgEdge database", map[string]interface{}{
        "create_input": createInput,
    })

    database, err := r.client.CreateDatabase(ctx, createInput)
    if err != nil {
        resp.Diagnostics.AddError(
            "Error creating database",
            "Could not create database, unexpected error: "+err.Error(),
        )
        return
    }

    // Map response body to schema and populate Computed attribute values
    plan = r.mapDatabaseToResourceModel(database)

    // Set state to fully populated data
    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    tflog.Info(ctx, "Created pgEdge database", map[string]interface{}{
        "database_id": plan.ID.ValueString(),
    })
}

func (r *databaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    var state databaseResourceModel
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

    // Map response body to schema and populate Computed attribute values
    state = r.mapDatabaseToResourceModel(database)

    // Set refreshed state
    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    var plan databaseResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    updateInput := &models.UpdateDatabaseInput{
        Options: convertToStringSlice(plan.Options),
    }

     // Handle Extensions
     if !plan.Extensions.IsNull() {
        var extensionsData extensionsModel
        diags := plan.Extensions.As(ctx, &extensionsData, basetypes.ObjectAsOptions{})
        resp.Diagnostics.Append(diags...)
        if resp.Diagnostics.HasError() {
            return
        }
        
        updateInput.Extensions = &models.Extensions{
            AutoManage: extensionsData.AutoManage.ValueBool(),
            Requested:  convertTFListToStringSlice(extensionsData.Requested),
        }
    }

	 // Handle Extensions
	//  if !plan.Extensions.IsNull() {
    //     var extensionsData struct {
    //         AutoManage bool     `tfsdk:"auto_manage"`
    //         Requested  []string `tfsdk:"requested"`
    //     }
    //     diags := plan.Extensions.As(ctx, &extensionsData, basetypes.ObjectAsOptions{})
    //     resp.Diagnostics.Append(diags...)
    //     if resp.Diagnostics.HasError() {
    //         return
    //     }
    //     updateInput.Extensions = &models.Extensions{
    //         AutoManage: extensionsData.AutoManage,
    //         Requested:  extensionsData.Requested,
    //     }
    // }

    tflog.Debug(ctx, "Updating pgEdge database", map[string]interface{}{
        "update_input": updateInput,
    })

    database, err := r.client.UpdateDatabase(ctx, strfmt.UUID(plan.ID.ValueString()), updateInput)
    if err != nil {
        resp.Diagnostics.AddError(
            "Error Updating pgEdge Database",
            "Could not update database, unexpected error: "+err.Error(),
        )
        return
    }

    // Map response body to schema and populate Computed attribute values
    plan = r.mapDatabaseToResourceModel(database)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    tflog.Info(ctx, "Updated pgEdge database", map[string]interface{}{
        "database_id": plan.ID.ValueString(),
    })
}

func (r *databaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var state databaseResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    tflog.Info(ctx, "Deleting pgEdge database", map[string]interface{}{
        "database_id": state.ID.ValueString(),
    })

    err := r.client.DeleteDatabase(ctx, strfmt.UUID(state.ID.ValueString()))
    if err != nil {
        resp.Diagnostics.AddError(
            "Error Deleting pgEdge Database",
            "Could not delete database, unexpected error: "+err.Error(),
        )
        return
    }

    tflog.Info(ctx, "Deleted pgEdge database", map[string]interface{}{
        "database_id": state.ID.ValueString(),
    })
}

func (r *databaseResource) mapDatabaseToResourceModel(database *models.Database) databaseResourceModel {
    return databaseResourceModel{
        ID:             types.StringValue(database.ID.String()),
        Name:           types.StringPointerValue(database.Name),
        ClusterID:      types.StringValue(database.ClusterID.String()),
        Status:         types.StringPointerValue(database.Status),
        CreatedAt:      types.StringPointerValue(database.CreatedAt),
        UpdatedAt:      types.StringPointerValue(database.UpdatedAt),
        PgVersion:      types.StringValue(database.PgVersion),
        StorageUsed:    types.Int64Value(database.StorageUsed),
        Domain:         types.StringValue(database.Domain),
        ConfigVersion:  types.StringValue(database.ConfigVersion),
        Options:        convertToListValue(database.Options),
        // Backups:        r.mapBackupsToResourceModel(database.Backups),
        // Components:     r.mapComponentsToResourceModel(database.Components),
        Extensions:     r.mapExtensionsToResourceModel(database.Extensions),
        // Nodes:          r.mapNodesToResourceModel(database.Nodes),
        // Roles:          r.mapRolesToResourceModel(database.Roles),
        // Tables:         r.mapTablesToResourceModel(database.Tables),
    }
}

func (r *databaseResource) mapBackupsToResourceModel(backups *models.Backups) types.Object {
    if backups == nil {
        return types.ObjectNull(map[string]attr.Type{
            "provider": types.StringType,
            "config":   types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
        })
    }

    configList := []attr.Value{}
    for _, config := range backups.Config {
        configObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "id":           types.StringType,
                "node_name":    types.StringType,
                "repositories": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
                "schedules":    types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
            },
            map[string]attr.Value{
                "id":           types.StringValue(config.ID.String()),
                "node_name":    types.StringValue(config.NodeName),
                "repositories": r.mapBackupRepositoriesToResourceModel(config.Repositories),
                "schedules":    r.mapBackupSchedulesToResourceModel(config.Schedules),
            },
        )
        configList = append(configList, configObj)
    }

    backupsObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "provider": types.StringType,
            "config":   types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
        },
        map[string]attr.Value{
            "provider": types.StringPointerValue(backups.Provider),
            "config":   types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, configList),
        },
    )

    return backupsObj
}

func (r *databaseResource) mapComponentsToResourceModel(components []*models.DatabaseComponentVersion) types.List {
    componentsList := []attr.Value{}
    for _, component := range components {
        componentObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "id":           types.StringType,
                "name":         types.StringType,
                "version":      types.StringType,
                "release_date": types.StringType,
                "status":       types.StringType,
            },
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

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, componentsList)
}

func (r *databaseResource) mapExtensionsToResourceModel(extensions *models.Extensions) types.Object {
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
            "available":   types.ListValueMust(types.StringType, r.stringSliceToValueSlice(extensions.Available)),
            "requested":   types.ListValueMust(types.StringType, r.stringSliceToValueSlice(extensions.Requested)),
        },
    )

    return extensionsObj
}

func (r *databaseResource) mapNodesToResourceModel(nodes []*models.DatabaseNode) types.List {
    nodesList := []attr.Value{}
    for _, node := range nodes {
        nodeObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "name":       types.StringType,
                "connection": types.ObjectType{AttrTypes: map[string]attr.Type{}},
                "location":   types.ObjectType{AttrTypes: map[string]attr.Type{}},
                "region":     types.ObjectType{AttrTypes: map[string]attr.Type{}},
                "extensions": types.ObjectType{AttrTypes: map[string]attr.Type{}},
            },
            map[string]attr.Value{
                "name":       types.StringPointerValue(node.Name),
                "connection": r.mapConnectionToResourceModel(node.Connection),
                "location":   r.mapLocationToResourceModel(node.Location),
                "region":     r.mapRegionToResourceModel(node.Region),
                "extensions": r.mapNodeExtensionsToResourceModel(node.Extensions),
            },
        )
        nodesList = append(nodesList, nodeObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, nodesList)
}

func (r *databaseResource) mapRolesToResourceModel(roles []*models.DatabaseRole) types.List {
    rolesList := []attr.Value{}
    for _, role := range roles {
        roleObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "name":             types.StringType,
                "bypass_rls":       types.BoolType,
                "connection_limit": types.Int64Type,
                "create_db":        types.BoolType,
                "create_role":      types.BoolType,
                "inherit":          types.BoolType,
                "login":            types.BoolType,
                "replication":      types.BoolType,
                "superuser":        types.BoolType,
            },
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

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, rolesList)
}

func (r *databaseResource) mapTablesToResourceModel(tables []*models.DatabaseTable) types.List {
    tablesList := []attr.Value{}
    for _, table := range tables {
        tableObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "name":             types.StringType,
                "schema":           types.StringType,
                "primary_key":      types.ListType{ElemType: types.StringType},
                "replication_sets": types.ListType{ElemType: types.StringType},
                "columns":          types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
                "status":           types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{}}},
            },
            map[string]attr.Value{
                "name":             types.StringPointerValue(table.Name),
                "schema":           types.StringPointerValue(table.Schema),
                "primary_key":      types.ListValueMust(types.StringType, r.stringSliceToValueSlice(table.PrimaryKey)),
                "replication_sets": types.ListValueMust(types.StringType, r.stringSliceToValueSlice(table.ReplicationSets)),
                "columns":          r.mapColumnsToResourceModel(table.Columns),
                "status":           r.mapTableStatusToResourceModel(table.Status),
            },
        )
        tablesList = append(tablesList, tableObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, tablesList)
}

// Helper functions

func (r *databaseResource) mapBackupRepositoriesToResourceModel(repositories []*models.BackupRepository) types.List {
    repoList := []attr.Value{}
    for _, repo := range repositories {
        repoObj, _ := types.ObjectValue(
            map[string]attr.Type{
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
            },
            map[string]attr.Value{
                "id":                   types.StringValue(repo.ID.String()),
                "type":                 types.StringValue(repo.Type),
                "backup_store_id":      types.StringValue(repo.BackupStoreID.String()),
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
    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, repoList)
}

func (r *databaseResource) mapBackupSchedulesToResourceModel(schedules []*models.BackupSchedule) types.List {
    scheduleList := []attr.Value{}
    for _, schedule := range schedules {
        scheduleObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "id":               types.StringType,
                "type":             types.StringType,
                "cron_expression":  types.StringType,
            },
            map[string]attr.Value{
                "id":               types.StringValue(schedule.ID.String()),
                "type":             types.StringPointerValue(schedule.Type),
                "cron_expression":  types.StringPointerValue(schedule.CronExpression),
            },
        )
        scheduleList = append(scheduleList, scheduleObj)
    }
    return types.ListValueMust(
        types.ObjectType{
            AttrTypes: map[string]attr.Type{
                "id":               types.StringType,
                "type":             types.StringType,
                "cron_expression":  types.StringType,
            },
        },
        scheduleList,
    )
}

func (r *databaseResource) mapConnectionToResourceModel(connection *models.Connection) types.Object {
    if connection == nil {
        return types.ObjectNull(map[string]attr.Type{})
    }
    connectionObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "database":            types.StringType,
            "host":                types.StringType,
            "password":            types.StringType,
            "port":                types.Int64Type,
            "username":            types.StringType,
            "external_ip_address": types.StringType,
            "internal_ip_address": types.StringType,
            "internal_host":       types.StringType,
        },
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

func (r *databaseResource) mapLocationToResourceModel(location *models.Location) types.Object {
    if location == nil {
        return types.ObjectNull(map[string]attr.Type{})
    }
    locationObj, _ := types.ObjectValue(
        map[string]attr.Type{
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
        },
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

func (r *databaseResource) mapRegionToResourceModel(region *models.Region) types.Object {
    if region == nil {
        return types.ObjectNull(map[string]attr.Type{
            "active":             types.BoolType,
            "availability_zones": types.ListType{ElemType: types.StringType},
            "cloud":              types.StringType,
            "code":               types.StringType,
            "name":               types.StringType,
            "parent":             types.StringType,
        })
    }

    regionObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "active":             types.BoolType,
            "availability_zones": types.ListType{ElemType: types.StringType},
            "cloud":              types.StringType,
            "code":               types.StringType,
            "name":               types.StringType,
            "parent":             types.StringType,
        },
        map[string]attr.Value{
            "active":             types.BoolValue(region.Active),
            "availability_zones": types.ListValueMust(types.StringType, r.stringSliceToValueSlice(region.AvailabilityZones)),
            "cloud":              types.StringPointerValue(region.Cloud),
            "code":               types.StringPointerValue(region.Code),
            "name":               types.StringPointerValue(region.Name),
            "parent":             types.StringValue(region.Parent),
        },
    )

    return regionObj
}

func (r *databaseResource) mapNodeExtensionsToResourceModel(extensions *models.DatabaseNodeExtensions) types.Object {
    if extensions == nil {
        return types.ObjectNull(map[string]attr.Type{
            "errors":    types.MapType{ElemType: types.StringType},
            "installed": types.ListType{ElemType: types.StringType},
        })
    }

    errorsMap := make(map[string]attr.Value)
    for k, v := range extensions.Errors {
        errorsMap[k] = types.StringValue(v)
    }

    extensionsObj, _ := types.ObjectValue(
        map[string]attr.Type{
            "errors":    types.MapType{ElemType: types.StringType},
            "installed": types.ListType{ElemType: types.StringType},
        },
        map[string]attr.Value{
            "errors":    types.MapValueMust(types.StringType, errorsMap),
            "installed": types.ListValueMust(types.StringType, r.stringSliceToValueSlice(extensions.Installed)),
        },
    )

    return extensionsObj
}

func (r *databaseResource) mapColumnsToResourceModel(columns []*models.DatabaseColumn) types.List {
    columnsList := []attr.Value{}
    for _, column := range columns {
        columnObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "name":             types.StringType,
                "data_type":        types.StringType,
                "default":          types.StringType,
                "is_nullable":      types.BoolType,
                "is_primary_key":   types.BoolType,
                "ordinal_position": types.Int64Type,
            },
            map[string]attr.Value{
                "name":             types.StringPointerValue(column.Name),
                "data_type":        types.StringPointerValue(column.DataType),
                "default":          types.StringPointerValue(column.Default),
                "is_nullable":      types.BoolPointerValue(column.IsNullable),
                "is_primary_key":   types.BoolPointerValue(column.IsPrimaryKey),
                "ordinal_position": types.Int64PointerValue(column.OrdinalPosition),
            },
        )
        columnsList = append(columnsList, columnObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, columnsList)
}

func (r *databaseResource) mapTableStatusToResourceModel(status []*models.DatabaseTableStatus) types.List {
    statusList := []attr.Value{}
    for _, s := range status {
        statusObj, _ := types.ObjectValue(
            map[string]attr.Type{
                "aligned":     types.BoolType,
                "node_name":   types.StringType,
                "present":     types.BoolType,
                "replicating": types.BoolType,
            },
            map[string]attr.Value{
                "aligned":     types.BoolPointerValue(s.Aligned),
                "node_name":   types.StringPointerValue(s.NodeName),
                "present":     types.BoolPointerValue(s.Present),
                "replicating": types.BoolPointerValue(s.Replicating),
            },
        )
        statusList = append(statusList, statusObj)
    }

    return types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{}}, statusList)
}

func (r *databaseResource) stringSliceToValueSlice(slice []string) []attr.Value {
    valueSlice := make([]attr.Value, len(slice))
    for i, s := range slice {
        valueSlice[i] = types.StringValue(s)
    }
    return valueSlice
}



func (r *databaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type databaseResourceModel struct {
    ID             types.String `tfsdk:"id"`
    Name           types.String `tfsdk:"name"`
    ClusterID      types.String `tfsdk:"cluster_id"`
    Status         types.String `tfsdk:"status"`
    CreatedAt      types.String `tfsdk:"created_at"`
    UpdatedAt      types.String `tfsdk:"updated_at"`
    PgVersion      types.String `tfsdk:"pg_version"`
    StorageUsed    types.Int64  `tfsdk:"storage_used"`
    Domain         types.String `tfsdk:"domain"`
    ConfigVersion  types.String `tfsdk:"config_version"`
    Options        types.List   `tfsdk:"options"`
    // Backups        types.Object `tfsdk:"backups"`
    // Components     types.List   `tfsdk:"components"`
    Extensions     types.Object `tfsdk:"extensions"`
    // Nodes          types.List   `tfsdk:"nodes"`
    // Roles          types.List   `tfsdk:"roles"`
    // Tables         types.List   `tfsdk:"tables"`
}

func convertToStringSlice(list types.List) []string {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}
	var result []string
	for _, elem := range list.Elements() {
		if str, ok := elem.(types.String); ok {
			result = append(result, str.ValueString())
		}
	}
	return result
}

func convertToListValue(slice []string) types.List {
	elements := make([]attr.Value, len(slice))
	for i, s := range slice {
		elements[i] = types.StringValue(s)
	}
	return types.ListValueMust(types.StringType, elements)
}

// Helper function to convert types.List to []string
func convertTFListToStringSlice(list types.List) []string {
    if list.IsNull() || list.IsUnknown() {
        return nil
    }
    var result []string
    for _, elem := range list.Elements() {
        if str, ok := elem.(types.String); ok {
            result = append(result, str.ValueString())
        }
    }
    return result
}