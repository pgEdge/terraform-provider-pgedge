package database

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	// "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
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

type backupsPlanModifier struct{}

func (m backupsPlanModifier) Description(_ context.Context) string {
	return "Prevents modifications to backup configuration while allowing computed value changes"
}

func (m backupsPlanModifier) MarkdownDescription(_ context.Context) string {
	return "Prevents modifications to backup configuration while allowing computed value changes"
}

func (m backupsPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
    if req.StateValue.IsNull() {
        return
    }

    if req.PlanValue.IsUnknown() {
        resp.PlanValue = req.StateValue
        return
    }

    planValue := req.PlanValue
    stateValue := req.StateValue
    
    planAttrs := planValue.Attributes()
    stateAttrs := stateValue.Attributes()

    planConfigs, ok := planAttrs["config"].(types.List)
    if !ok {
        return
    }
    stateConfigs, ok := stateAttrs["config"].(types.List)
    if !ok {
        return
    }

    hasBackupChanges := false

    for i, planConfig := range planConfigs.Elements() {
        if i >= len(stateConfigs.Elements()) {
            hasBackupChanges = true
            break
        }
        stateConfig := stateConfigs.Elements()[i]

        planConfigObj := planConfig.(types.Object)
        stateConfigObj := stateConfig.(types.Object)

        planRepos := planConfigObj.Attributes()["repositories"].(types.List)
        stateRepos := stateConfigObj.Attributes()["repositories"].(types.List)

        for j, planRepo := range planRepos.Elements() {
            if j >= len(stateRepos.Elements()) {
                hasBackupChanges = true
                break
            }
            stateRepo := stateRepos.Elements()[j]

            planRepoObj := planRepo.(types.Object)
            stateRepoObj := stateRepo.(types.Object)

            planBackupStoreID := planRepoObj.Attributes()["backup_store_id"].(types.String)
            stateBackupStoreID := stateRepoObj.Attributes()["backup_store_id"].(types.String)

            if !planBackupStoreID.IsNull() && !stateBackupStoreID.IsNull() &&
               planBackupStoreID.ValueString() != stateBackupStoreID.ValueString() {
                hasBackupChanges = true
            }

            if m.hasRetentionChanges(planRepoObj.Attributes(), stateRepoObj.Attributes()) {
                hasBackupChanges = true
            }
        }

        planSchedules, hasPS := planConfigObj.Attributes()["schedules"].(types.List)
        stateSchedules, hasSS := stateConfigObj.Attributes()["schedules"].(types.List)

        if hasPS && hasSS {
            planSchedsElements := planSchedules.Elements()
            stateSchedsElements := stateSchedules.Elements()

            if len(planSchedsElements) != len(stateSchedsElements) {
                hasBackupChanges = true
            } else {
                for k := range planSchedsElements {
                    planSched := planSchedsElements[k].(types.Object)
                    stateSched := stateSchedsElements[k].(types.Object)

                    if !m.schedulesEqual(planSched.Attributes(), stateSched.Attributes()) {
                        hasBackupChanges = true
                    }
                }
            }
        }
    }

    if hasBackupChanges {
        resp.Diagnostics.AddError(
            "Invalid Backup Configuration Modification",
            "Backup configuration cannot be modified after database creation. You must create a new database to change backup settings.",
        )
        return
    }

    resp.PlanValue = req.StateValue
}

func (m backupsPlanModifier) hasRetentionChanges(plan, state map[string]attr.Value) bool {
    if planRetention, ok := plan["retention_full"].(types.Int64); ok && !planRetention.IsNull() {
        if stateRetention, ok := state["retention_full"].(types.Int64); ok && !stateRetention.IsNull() {
            if planRetention.ValueInt64() != stateRetention.ValueInt64() {
                return true
            }
        }
    }

    if planType, ok := plan["retention_full_type"].(types.String); ok && !planType.IsNull() {
        if stateType, ok := state["retention_full_type"].(types.String); ok && !stateType.IsNull() {
            if planType.ValueString() != stateType.ValueString() {
                return true
            }
        }
    }

    return false
}

func (m backupsPlanModifier) schedulesEqual(plan, state map[string]attr.Value) bool {
    if !m.stringAttrEqual(plan["type"], state["type"]) {
        return false
    }

    if !m.stringAttrEqual(plan["cron_expression"], state["cron_expression"]) {
        return false
    }

    return true
}

func (m backupsPlanModifier) stringAttrEqual(plan, state attr.Value) bool {
    planStr, ok1 := plan.(types.String)
    stateStr, ok2 := state.(types.String)

    if !ok1 || !ok2 {
        return true
    }

    if planStr.IsNull() || stateStr.IsNull() {
        return true
    }

    return planStr.ValueString() == stateStr.ValueString()
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the database was created.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"pg_version": schema.StringAttribute{
				Description: "The PostgreSQL version of the database.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Computed: true,
			},
			// "storage_used": schema.Int64Attribute{
			// 	Description: "The amount of storage used by the database in bytes.",
			// 	PlanModifiers: []planmodifier.Int64{
			// 		int64planmodifier.RequiresReplaceIfConfigured(),
			// 	},
			// 	Computed:    true,
			// },
			"domain": schema.StringAttribute{
				Description: "The domain associated with the database.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"config_version": schema.StringAttribute{
				Description: "The configuration version of the database.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					configVersionPlanModifier{},
				},
			},
			"options": schema.ListAttribute{
				Description: "A list of options for the database.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"backups": schema.SingleNestedAttribute{
				Description: "Backup configuration for the database.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.Object{
					backupsPlanModifier{},
				},
				Attributes: map[string]schema.Attribute{
					"provider": schema.StringAttribute{
						Description: "The backup provider.",
						Computed:    true,
						Optional:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"config": schema.ListNestedAttribute{
						Description: "List of backup configurations.",
						Computed:    true,
						Optional:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Unique identifier for the backup config.",
									Computed:    true,
									Optional:    true,
								},
								"node_name": schema.StringAttribute{
									Description: "Name of the node.",
									Computed:    true,
									Optional:    true,
								},
								"repositories": schema.ListNestedAttribute{
									Description: "List of backup repositories.",
									Computed:    true,
									Optional:    true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.UseStateForUnknown(),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "Repository identifier.",
												Optional:    true,
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "Repository type (e.g., s3, gcs, azure).",
												Optional:    true,
												Computed:    true,
											},
											"backup_store_id": schema.StringAttribute{
												Description: "ID of the backup store to use. If specified, other fields will be " +
													"automatically populated.",
												Optional: true,
											},
											"s3_bucket": schema.StringAttribute{
												Description: "S3 bucket name for s3-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"s3_region": schema.StringAttribute{
												Description: "S3 region for s3-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"s3_endpoint": schema.StringAttribute{
												Description: "S3 endpoint for s3-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"gcs_bucket": schema.StringAttribute{
												Description: "GCS bucket name for gcs-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"gcs_endpoint": schema.StringAttribute{
												Description: "GCS endpoint for gcs-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"azure_account": schema.StringAttribute{
												Description: "Azure account for azure-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"azure_container": schema.StringAttribute{
												Description: "Azure container for azure-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"azure_endpoint": schema.StringAttribute{
												Description: "Azure endpoint for azure-type repositories.",
												Optional:    true,
												Computed:    true,
											},
											"base_path": schema.StringAttribute{
												Description: "Base path for the repository.",
												Optional:    true,
												Computed:    true,
											},
											"retention_full": schema.Int64Attribute{
												Description: "Retention period for full backups.",
												Optional:    true,
												Computed:    true,
											},
											"retention_full_type": schema.StringAttribute{
												Description: "Type of retention for full backups.",
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"schedules": schema.ListNestedAttribute{
									Description: "List of backup schedules.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "Unique identifier for the schedule.",
												Required:    true,
											},
											"type": schema.StringAttribute{
												Description: "Type of the schedule.",
												Required:    true,
											},
											"cron_expression": schema.StringAttribute{
												Description: "Cron expression for the schedule.",
												Required:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"components": schema.ListNestedAttribute{
				Description: "List of components in the database.",
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":           schema.StringAttribute{Computed: true},
						"name":         schema.StringAttribute{Computed: true},
						"version":      schema.StringAttribute{Computed: true},
						"release_date": schema.StringAttribute{Computed: true},
						"status":       schema.StringAttribute{Computed: true},
					},
				},
			},
			"extensions": schema.SingleNestedAttribute{
				Description: "Extensions configuration for the database.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"auto_manage": schema.BoolAttribute{Computed: true, Optional: true},
					"available": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"requested": schema.ListAttribute{Computed: true, Optional: true, ElementType: types.StringType, PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()}},
				},
			},
			"nodes": schema.MapNestedAttribute{
				Description: "Map of nodes in the database.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"connection": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"database":            schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"host":                schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"password":            schema.StringAttribute{Computed: true, Sensitive: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"port":                schema.Int64Attribute{Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}},
								"username":            schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"external_ip_address": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"internal_ip_address": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"internal_host":       schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
							},
						},
						"location": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"code":        schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"country":     schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"latitude":    schema.Float64Attribute{Computed: true, PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()}},
								"longitude":   schema.Float64Attribute{Computed: true, PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()}},
								"name":        schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"region":      schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"region_code": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"timezone":    schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"postal_code": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"metro_code":  schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"city":        schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
							},
						},
						"region": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"active":             schema.BoolAttribute{Computed: true, PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()}},
								"availability_zones": schema.ListAttribute{Computed: true, ElementType: types.StringType, PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()}},
								"cloud":              schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"code":               schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"name":               schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
								"parent":             schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
							},
						},
						"extensions": schema.SingleNestedAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.Object{
								conditionalUseStateForUnknownModifier{},
							},
							Attributes: map[string]schema.Attribute{
								"errors":    schema.MapAttribute{Computed: true, ElementType: types.StringType, PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()}},
								"installed": schema.ListAttribute{Computed: true, ElementType: types.StringType, PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()}},
							},
						},
					},
				},
			},
			"roles": schema.ListNestedAttribute{
				Description: "List of roles in the database.",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
					Attributes: map[string]schema.Attribute{
						"name":             schema.StringAttribute{Computed: true},
						"bypass_rls":       schema.BoolAttribute{Computed: true},
						"connection_limit": schema.Int64Attribute{Computed: true},
						"create_db":        schema.BoolAttribute{Computed: true},
						"create_role":      schema.BoolAttribute{Computed: true},
						"inherit":          schema.BoolAttribute{Computed: true},
						"login":            schema.BoolAttribute{Computed: true},
						"replication":      schema.BoolAttribute{Computed: true},
						"superuser":        schema.BoolAttribute{Computed: true},
					},
				},
			},
		},
	}
}

type configVersionPlanModifier struct{}

func (m configVersionPlanModifier) Description(_ context.Context) string {
	return "Prevents modifications to config_version after resource creation"
}

func (m configVersionPlanModifier) MarkdownDescription(_ context.Context) string {
	return "Prevents modifications to config_version after resource creation"
}

func (m configVersionPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}

	resp.PlanValue = req.StateValue
}

// Plan modifier for node extensions
type conditionalUseStateForUnknownModifier struct{}

func (m conditionalUseStateForUnknownModifier) Description(_ context.Context) string {
	return "Uses the prior state for nodeextensions if extensions hasn't been modified."
}

func (m conditionalUseStateForUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Uses the prior state for nodeextensions if extensions hasn't been modified."
}

func (m conditionalUseStateForUnknownModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.StateValue.IsNull() {
		return
	}

	if !req.PlanValue.IsUnknown() {
		return
	}

	var configRequested, stateInstalled []string

	if !req.Config.Raw.IsNull() {
		var configData map[string]tftypes.Value
		err := req.Config.Raw.As(&configData)
		if err == nil {
			if extVal, ok := configData["extensions"]; ok {
				var extMap map[string]tftypes.Value
				err = extVal.As(&extMap)
				if err == nil {
					if reqVal, ok := extMap["requested"]; ok {
						var requestedList []tftypes.Value
						err = reqVal.As(&requestedList)
						if err == nil {
							for _, v := range requestedList {
								var s string
								if err := v.As(&s); err == nil {
									configRequested = append(configRequested, s)
								}
							}
						} else {
							// TODO: log error
						}
					}
				}
			}
		}
	}

	if !req.StateValue.IsNull() {
		stateData := req.StateValue.Attributes()
		if installedVal, ok := stateData["installed"].(types.List); ok {
			stateInstalled = make([]string, 0, len(installedVal.Elements()))
			for _, elem := range installedVal.Elements() {
				if strVal, ok := elem.(types.String); ok {
					stateInstalled = append(stateInstalled, strVal.ValueString())
				}
			}
		}
	}

	if len(stateInstalled) == 0 {
		return
	}

	if compareStringSlices(configRequested, stateInstalled) {
		resp.PlanValue = req.StateValue
		return
	}

}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[string]int)

	for _, v := range a {
		countMap[v]++
	}

	for _, v := range b {
		countMap[v]--
		if countMap[v] == 0 {
			delete(countMap, v)
		}
	}

	return len(countMap) == 0
}

func New() planmodifier.Object {
	return &conditionalUseStateForUnknownModifier{}
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
	AutoManage types.Bool `tfsdk:"auto_manage"`
	Available  types.List `tfsdk:"available"`
	Requested  types.List `tfsdk:"requested"`
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan databaseResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createInput := &models.CreateDatabaseInput{
		Name:      plan.Name.ValueStringPointer(),
		ClusterID: strfmt.UUID(plan.ClusterID.ValueString()),
		Options:   convertToStringSlice(plan.Options),
	}

	if !plan.ConfigVersion.IsNull() && !plan.ConfigVersion.IsUnknown() {
		createInput.ConfigVersion = plan.ConfigVersion.ValueString()
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
			Requested:  common.ConvertTFListToStringSlice(extensionsData.Requested),
		}
	}

	// Handle Backups
	if !plan.Backups.IsNull() && !plan.Backups.IsUnknown() {
		var backupsData struct {
			Provider types.String `tfsdk:"provider"`
			Config   types.List   `tfsdk:"config"`
		}
		diags := plan.Backups.As(ctx, &backupsData, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		backups := &models.Backups{
			Provider: backupsData.Provider.ValueStringPointer(),
		}

		if !backupsData.Config.IsNull() && !backupsData.Config.IsUnknown() {
			var configData []struct {
				ID           types.String `tfsdk:"id"`
				NodeName     types.String `tfsdk:"node_name"`
				Repositories types.List   `tfsdk:"repositories"`
				Schedules    types.List   `tfsdk:"schedules"`
			}
			diags := backupsData.Config.ElementsAs(ctx, &configData, false)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			backups.Config = make([]*models.BackupConfig, 0, len(configData))

			for _, config := range configData {
				backupConfig := &models.BackupConfig{
					ID:           config.ID.ValueStringPointer(),
					NodeName:     *config.NodeName.ValueStringPointer(),
					Repositories: make([]*models.BackupRepository, 0),
					Schedules:    make([]*models.BackupSchedule, 0),
				}

				if !config.Repositories.IsNull() && !config.Repositories.IsUnknown() {
					var repositories []struct {
						ID                types.String `tfsdk:"id"`
						Type              types.String `tfsdk:"type"`
						BackupStoreID     types.String `tfsdk:"backup_store_id"`
						BasePath          types.String `tfsdk:"base_path"`
						S3Bucket          types.String `tfsdk:"s3_bucket"`
						S3Region          types.String `tfsdk:"s3_region"`
						S3Endpoint        types.String `tfsdk:"s3_endpoint"`
						GcsBucket         types.String `tfsdk:"gcs_bucket"`
						GcsEndpoint       types.String `tfsdk:"gcs_endpoint"`
						AzureAccount      types.String `tfsdk:"azure_account"`
						AzureContainer    types.String `tfsdk:"azure_container"`
						AzureEndpoint     types.String `tfsdk:"azure_endpoint"`
						RetentionFull     types.Int64  `tfsdk:"retention_full"`
						RetentionFullType types.String `tfsdk:"retention_full_type"`
					}
					diags := config.Repositories.ElementsAs(ctx, &repositories, false)
					resp.Diagnostics.Append(diags...)
					if resp.Diagnostics.HasError() {
						return
					}
					for _, repo := range repositories {
						backupConfig.Repositories = append(backupConfig.Repositories, &models.BackupRepository{
							ID:                *repo.ID.ValueStringPointer(),
							Type:              *repo.Type.ValueStringPointer(),
							BackupStoreID:     *repo.BackupStoreID.ValueStringPointer(),
							BasePath:          *repo.BasePath.ValueStringPointer(),
							S3Bucket:          *repo.S3Bucket.ValueStringPointer(),
							S3Region:          *repo.S3Region.ValueStringPointer(),
							S3Endpoint:        *repo.S3Endpoint.ValueStringPointer(),
							GcsBucket:         *repo.GcsBucket.ValueStringPointer(),
							GcsEndpoint:       *repo.GcsEndpoint.ValueStringPointer(),
							AzureAccount:      *repo.AzureAccount.ValueStringPointer(),
							AzureContainer:    *repo.AzureContainer.ValueStringPointer(),
							AzureEndpoint:     *repo.AzureEndpoint.ValueStringPointer(),
							RetentionFull:     *repo.RetentionFull.ValueInt64Pointer(),
							RetentionFullType: *repo.RetentionFullType.ValueStringPointer(),
						})
					}
				}

				if !config.Schedules.IsNull() && !config.Schedules.IsUnknown() {
					var schedules []struct {
						ID             types.String `tfsdk:"id"`
						Type           types.String `tfsdk:"type"`
						CronExpression types.String `tfsdk:"cron_expression"`
					}
					diags := config.Schedules.ElementsAs(ctx, &schedules, false)
					resp.Diagnostics.Append(diags...)
					if resp.Diagnostics.HasError() {
						return
					}

					for _, schedule := range schedules {
						backupConfig.Schedules = append(backupConfig.Schedules, &models.BackupSchedule{
							ID:             schedule.ID.ValueStringPointer(),
							Type:           schedule.Type.ValueStringPointer(),
							CronExpression: schedule.CronExpression.ValueStringPointer(),
						})
					}
				}

				backups.Config = append(backups.Config, backupConfig)
			}
		}

		createInput.Backups = backups
	}

	tflog.Debug(ctx, "Creating pgEdge database", map[string]interface{}{
		"create_input": createInput,
	})

	database, err := r.client.CreateDatabase(ctx, createInput)
	if err != nil {
		if database != nil {
			mappedDatabase := r.mapDatabaseToResourceModel(database)
			diags = resp.State.Set(ctx, mappedDatabase)
			resp.Diagnostics.Append(diags...)
		}
		resp.Diagnostics.Append(common.HandleProviderError(err, "database creation"))
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
		"create_output": plan,
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
		diag := common.HandleProviderError(err, "database retrieval")
		if diag == nil {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.Append(diag)
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

func (r *databaseResource) nodesEqual(planNodes, stateNodes types.Map) bool {
	if planNodes.IsNull() || stateNodes.IsNull() {
		return planNodes.IsNull() == stateNodes.IsNull()
	}

	planElements := planNodes.Elements()
	stateElements := stateNodes.Elements()

	if len(planElements) != len(stateElements) {
		return false
	}

	planNodeNames := make(map[string]bool)
	stateNodeNames := make(map[string]bool)

	for key, nodeElem := range planElements {
		if nodeObj, ok := nodeElem.(types.Object); ok {
			var node struct {
				Name types.String `tfsdk:"name"`
			}
			diags := nodeObj.As(context.Background(), &node, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				continue
			}
			planNodeNames[key] = true
			// Ensure the key matches the node name
			if key != node.Name.ValueString() {
				return false
			}
		}
	}

	for key, nodeElem := range stateElements {
		if nodeObj, ok := nodeElem.(types.Object); ok {
			var node struct {
				Name types.String `tfsdk:"name"`
			}
			diags := nodeObj.As(context.Background(), &node, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				continue
			}
			stateNodeNames[key] = true
			// Ensure the key matches the node name
			if key != node.Name.ValueString() {
				return false
			}
		}
	}

	for name := range planNodeNames {
		if !stateNodeNames[name] {
			return false
		}
	}

	for name := range stateNodeNames {
		if !planNodeNames[name] {
			return false
		}
	}

	return true
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state databaseResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Backups.Equal(state.Backups) {
		resp.Diagnostics.AddError(
			"Invalid Update Operation",
			"Backup configuration cannot be modified after database creation. You must create a new database to change backup settings.",
		)
		return
	}

	// Check how many fields are being updated
	updateCount := 0
	var updateField string

	if !plan.Options.Equal(state.Options) {
		updateCount++
		updateField = "options"
	}
	if !plan.Extensions.Equal(state.Extensions) {
		updateCount++
		updateField = "extensions"
	}
	if !r.nodesEqual(plan.Nodes, state.Nodes) {
		updateCount++
		updateField = "nodes"
	}

	// If more than one field is being updated, throw an error
	if updateCount > 1 {
		resp.Diagnostics.AddError(
			"Multiple Field Update Not Allowed",
			"Only one field (options, extensions, or nodes) can be updated at a time.",
		)
		return
	}

	// If no fields are being updated, return early
	// if updateCount == 0 {
	// 	return
	// }

	// Proceed with the update based on which field is being changed
	switch updateField {
	case "options":
		extensionsUpdateInput := &models.UpdateDatabaseInput{
			Options: convertToStringSlice(plan.Options),
		}

		tflog.Debug(ctx, "Updating pgEdge database options", map[string]interface{}{
			"update_input": extensionsUpdateInput,
		})

		updatedDatabase, err := r.client.UpdateDatabase(ctx, strfmt.UUID(plan.ID.ValueString()), extensionsUpdateInput)
		if err != nil {
			resp.Diagnostics.Append(common.HandleProviderError(err, "database options update"))
			return
		}
		plan = r.mapDatabaseToResourceModel(updatedDatabase)

	case "extensions":
		extensionsUpdateInput := &models.UpdateDatabaseInput{}

		if !plan.Extensions.IsNull() {
			var extensionsData extensionsModel
			diags := plan.Extensions.As(ctx, &extensionsData, basetypes.ObjectAsOptions{})
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			extensionsUpdateInput.Extensions = &models.Extensions{
				AutoManage: extensionsData.AutoManage.ValueBool(),
				Requested:  common.ConvertTFListToStringSlice(extensionsData.Requested),
			}
		}

		tflog.Debug(ctx, "Updating pgEdge database extensions", map[string]interface{}{
			"update_input": extensionsUpdateInput,
		})

		updatedDatabase, err := r.client.UpdateDatabase(ctx, strfmt.UUID(plan.ID.ValueString()), extensionsUpdateInput)
		if err != nil {
			resp.Diagnostics.Append(common.HandleProviderError(err, "database extensions update"))
			return
		}

		plan = r.mapDatabaseToResourceModel(updatedDatabase)

	case "nodes":
		nodesUpdateInput := &models.UpdateDatabaseInput{}

		nodes, diags := r.nodesToUpdateInput(ctx, plan.Nodes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		nodesUpdateInput.Nodes = nodes

		tflog.Debug(ctx, "Updating pgEdge database nodes", map[string]interface{}{
			"update_input": nodesUpdateInput,
		})

		updatedDatabase, err := r.client.UpdateDatabase(ctx, strfmt.UUID(plan.ID.ValueString()), nodesUpdateInput)
		if err != nil {
			resp.Diagnostics.Append(common.HandleProviderError(err, "database nodes update"))
			return
		}

		plan = r.mapDatabaseToResourceModel(updatedDatabase)
	}

	plan.ConfigVersion = state.ConfigVersion

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
		resp.Diagnostics.Append(common.HandleProviderError(err, "database deletion"))
		return
	}

	tflog.Info(ctx, "Deleted pgEdge database", map[string]interface{}{
		"database_id": state.ID.ValueString(),
	})
}

func (r *databaseResource) nodesToUpdateInput(ctx context.Context, nodesMap types.Map) ([]*models.DatabaseNode, diag.Diagnostics) {
	var nodes []*models.DatabaseNode
	var diags diag.Diagnostics

	for key, nodeElem := range nodesMap.Elements() {
		nodeObj, ok := nodeElem.(types.Object)
		if !ok {
			diags.AddError("Invalid node type", fmt.Sprintf("Expected types.Object, got %T", nodeElem))
			continue
		}

		var node struct {
			Name       types.String `tfsdk:"name"`
			Connection types.Object `tfsdk:"connection"`
			Location   types.Object `tfsdk:"location"`
			Region     types.Object `tfsdk:"region"`
			Extensions types.Object `tfsdk:"extensions"`
		}

		diags.Append(nodeObj.As(ctx, &node, basetypes.ObjectAsOptions{})...)
		if diags.HasError() {
			return nil, diags
		}

		// Ensure the key matches the node name
		if key != node.Name.ValueString() {
			diags.AddError("Node name mismatch", fmt.Sprintf("Node key '%s' does not match node name '%s'", key, node.Name.ValueString()))
			continue
		}

		dbNode := &models.DatabaseNode{
			Name:       node.Name.ValueStringPointer(),
			Connection: r.connectionFromObject(ctx, node.Connection),
			Location:   r.locationFromObject(ctx, node.Location),
			Region:     r.regionFromObject(ctx, node.Region),
			Extensions: r.extensionsFromObject(ctx, node.Extensions),
		}

		nodes = append(nodes, dbNode)
	}

	return nodes, diags
}

func (r *databaseResource) connectionFromObject(ctx context.Context, obj types.Object) *models.Connection {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var conn struct {
		Database          types.String `tfsdk:"database"`
		Host              types.String `tfsdk:"host"`
		Password          types.String `tfsdk:"password"`
		Port              types.Int64  `tfsdk:"port"`
		Username          types.String `tfsdk:"username"`
		ExternalIPAddress types.String `tfsdk:"external_ip_address"`
		InternalIPAddress types.String `tfsdk:"internal_ip_address"`
		InternalHost      types.String `tfsdk:"internal_host"`
	}

	obj.As(ctx, &conn, basetypes.ObjectAsOptions{})

	return &models.Connection{
		Database:          conn.Database.ValueStringPointer(),
		Host:              conn.Host.ValueString(),
		Password:          conn.Password.ValueStringPointer(),
		Port:              conn.Port.ValueInt64Pointer(),
		Username:          conn.Username.ValueStringPointer(),
		ExternalIPAddress: conn.ExternalIPAddress.ValueString(),
		InternalIPAddress: conn.InternalIPAddress.ValueString(),
		InternalHost:      conn.InternalHost.ValueString(),
	}
}

func (r *databaseResource) locationFromObject(ctx context.Context, obj types.Object) *models.Location {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var loc struct {
		Code       types.String  `tfsdk:"code"`
		Country    types.String  `tfsdk:"country"`
		Latitude   types.Float64 `tfsdk:"latitude"`
		Longitude  types.Float64 `tfsdk:"longitude"`
		Name       types.String  `tfsdk:"name"`
		Region     types.String  `tfsdk:"region"`
		RegionCode types.String  `tfsdk:"region_code"`
		Timezone   types.String  `tfsdk:"timezone"`
		PostalCode types.String  `tfsdk:"postal_code"`
		MetroCode  types.String  `tfsdk:"metro_code"`
		City       types.String  `tfsdk:"city"`
	}

	obj.As(ctx, &loc, basetypes.ObjectAsOptions{})

	return &models.Location{
		Code:       loc.Code.ValueString(),
		Country:    loc.Country.ValueString(),
		Latitude:   loc.Latitude.ValueFloat64Pointer(),
		Longitude:  loc.Longitude.ValueFloat64Pointer(),
		Name:       loc.Name.ValueString(),
		Region:     loc.Region.ValueString(),
		RegionCode: loc.RegionCode.ValueString(),
		Timezone:   loc.Timezone.ValueString(),
		PostalCode: loc.PostalCode.ValueString(),
		MetroCode:  loc.MetroCode.ValueString(),
		City:       loc.City.ValueString(),
	}
}

func (r *databaseResource) regionFromObject(ctx context.Context, obj types.Object) *models.Region {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var reg struct {
		Active            types.Bool   `tfsdk:"active"`
		AvailabilityZones types.List   `tfsdk:"availability_zones"`
		Cloud             types.String `tfsdk:"cloud"`
		Code              types.String `tfsdk:"code"`
		Name              types.String `tfsdk:"name"`
		Parent            types.String `tfsdk:"parent"`
	}

	obj.As(ctx, &reg, basetypes.ObjectAsOptions{})

	return &models.Region{
		Active:            reg.Active.ValueBool(),
		AvailabilityZones: common.ConvertTFListToStringSlice(reg.AvailabilityZones),
		Cloud:             reg.Cloud.ValueStringPointer(),
		Code:              reg.Code.ValueStringPointer(),
		Name:              reg.Name.ValueStringPointer(),
		Parent:            reg.Parent.ValueString(),
	}
}

func (r *databaseResource) extensionsFromObject(ctx context.Context, obj types.Object) *models.DatabaseNodeExtensions {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}

	var ext struct {
		Errors    types.Map  `tfsdk:"errors"`
		Installed types.List `tfsdk:"installed"`
	}

	obj.As(ctx, &ext, basetypes.ObjectAsOptions{})

	errors := make(map[string]string)
	for k, v := range ext.Errors.Elements() {
		if strVal, ok := v.(types.String); ok {
			errors[k] = strVal.ValueString()
		}
	}

	return &models.DatabaseNodeExtensions{
		Errors:    errors,
		Installed: common.ConvertTFListToStringSlice(ext.Installed),
	}
}

func (r *databaseResource) mapDatabaseToResourceModel(database *models.Database) databaseResourceModel {
	model := databaseResourceModel{
		ID:        types.StringValue(database.ID.String()),
		Name:      types.StringPointerValue(database.Name),
		ClusterID: types.StringValue(database.ClusterID.String()),
		Status:    types.StringPointerValue(database.Status),
		CreatedAt: types.StringPointerValue(database.CreatedAt),
		PgVersion: types.StringValue(database.PgVersion),
		// StorageUsed:   types.Int64Value(database.StorageUsed),
		Domain:        types.StringValue(database.Domain),
		ConfigVersion: types.StringValue(database.ConfigVersion),
		Options:       convertToListValue(database.Options),
		Backups:       r.mapBackupsToResourceModel(database.Backups),
		Components:    r.mapComponentsToResourceModel(database.Components),
		Extensions:    r.mapExtensionsToResourceModel(database.Extensions),
		Nodes:         r.mapNodesToResourceModel(database.Nodes),
		Roles:         r.mapRolesToResourceModel(database.Roles),
	}
	if r.backupsChanged(database.Backups, model.Backups) {
		model.Backups = r.mapBackupsToResourceModel(database.Backups)
	}

	return model
}

var backupConfigType = map[string]attr.Type{
	"id":           types.StringType,
	"node_name":    types.StringType,
	"repositories": types.ListType{ElemType: types.ObjectType{AttrTypes: backupRepositoryType}},
	"schedules":    types.ListType{ElemType: types.ObjectType{AttrTypes: backupScheduleType}},
}

var backupRepositoryType = map[string]attr.Type{
	"id":                  types.StringType,
	"type":                types.StringType,
	"backup_store_id":     types.StringType,
	"base_path":           types.StringType,
	"s3_bucket":           types.StringType,
	"s3_region":           types.StringType,
	"s3_endpoint":         types.StringType,
	"gcs_bucket":          types.StringType,
	"gcs_endpoint":        types.StringType,
	"azure_account":       types.StringType,
	"azure_container":     types.StringType,
	"azure_endpoint":      types.StringType,
	"retention_full":      types.Int64Type,
	"retention_full_type": types.StringType,
}

var backupScheduleType = map[string]attr.Type{
	"id":              types.StringType,
	"type":            types.StringType,
	"cron_expression": types.StringType,
}

func (r *databaseResource) mapBackupsToResourceModel(backups *models.Backups) types.Object {
	if backups == nil {
		return types.ObjectNull(map[string]attr.Type{
			"provider": types.StringType,
			"config":   types.ListType{ElemType: types.ObjectType{AttrTypes: backupConfigType}},
		})
	}

	configList := []attr.Value{}
	for _, config := range backups.Config {
		configObj, _ := types.ObjectValue(
			backupConfigType,
			map[string]attr.Value{
				"id":           types.StringPointerValue(config.ID),
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
			"config":   types.ListType{ElemType: types.ObjectType{AttrTypes: backupConfigType}},
		},
		map[string]attr.Value{
			"provider": types.StringPointerValue(backups.Provider),
			"config":   types.ListValueMust(types.ObjectType{AttrTypes: backupConfigType}, configList),
		},
	)

	return backupsObj
}

func (r *databaseResource) backupsChanged(apiBackups *models.Backups, modelBackups types.Object) bool {
	if apiBackups == nil {
		return false
	}

	if modelBackups.IsNull() || modelBackups.IsUnknown() {
		return true
	}

	// Compare relevant fields only
	var stateBackups struct {
		Provider string `tfsdk:"provider"`
		Config   []struct {
			ID           string `tfsdk:"id"`
			Repositories []struct {
				BackupStoreID string `tfsdk:"backup_store_id"`
			} `tfsdk:"repositories"`
		} `tfsdk:"config"`
	}

	modelBackups.As(context.Background(), &stateBackups, basetypes.ObjectAsOptions{})

	if stateBackups.Provider != *apiBackups.Provider {
		return true
	}

	configsMatch := len(stateBackups.Config) == len(apiBackups.Config)
	if !configsMatch {
		return true
	}

	for i, config := range stateBackups.Config {
		apiConfig := apiBackups.Config[i]
		if config.ID != *apiConfig.ID {
			return true
		}

		if len(config.Repositories) != len(apiConfig.Repositories) {
			return true
		}

		for j, repo := range config.Repositories {
			if repo.BackupStoreID != apiConfig.Repositories[j].BackupStoreID {
				return true
			}
		}
	}

	return false
}

func (r *databaseResource) mapComponentsToResourceModel(components []*models.DatabaseComponentVersion) types.List {
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

func (r *databaseResource) nodeAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       types.StringType,
		"connection": types.ObjectType{AttrTypes: r.connectionAttrTypes()},
		"location":   types.ObjectType{AttrTypes: r.locationAttrTypes()},
		"region":     types.ObjectType{AttrTypes: r.regionAttrTypes()},
		"extensions": types.ObjectType{AttrTypes: r.nodeExtensionsAttrTypes()},
	}
}

func (r *databaseResource) mapNodesToResourceModel(nodes []*models.DatabaseNode) types.Map {
	nodeMap := make(map[string]attr.Value)
	nodeAttrTypes := map[string]attr.Type{
		"name":       types.StringType,
		"connection": types.ObjectType{AttrTypes: r.connectionAttrTypes()},
		"location":   types.ObjectType{AttrTypes: r.locationAttrTypes()},
		"region":     types.ObjectType{AttrTypes: r.regionAttrTypes()},
		"extensions": types.ObjectType{AttrTypes: r.nodeExtensionsAttrTypes()},
	}

	for _, node := range nodes {
		if node.Name == nil {
			continue
		}

		regionObj := r.mapRegionToResourceModel(node.Region)
		if regionObj.IsNull() {
			regionObj, _ = types.ObjectValue(r.regionAttrTypes(), map[string]attr.Value{
				"active":             types.BoolNull(),
				"availability_zones": types.ListNull(types.StringType),
				"cloud":              types.StringNull(),
				"code":               types.StringNull(),
				"name":               types.StringNull(),
				"parent":             types.StringNull(),
			})
		}

		nodeObj, _ := types.ObjectValue(
			nodeAttrTypes,
			map[string]attr.Value{
				"name":       types.StringValue(*node.Name),
				"connection": r.mapConnectionToResourceModel(node.Connection),
				"location":   r.mapLocationToResourceModel(node.Location),
				"region":     regionObj,
				"extensions": r.mapNodeExtensionsToResourceModel(node.Extensions),
			},
		)
		nodeMap[*node.Name] = nodeObj
	}

	return types.MapValueMust(types.ObjectType{AttrTypes: nodeAttrTypes}, nodeMap)
}

func sortNodes(nodes []*models.DatabaseNode) []*models.DatabaseNode {
	sortedNodes := make([]*models.DatabaseNode, len(nodes))
	copy(sortedNodes, nodes)
	sort.Slice(sortedNodes, func(i, j int) bool {
		return *sortedNodes[i].Name < *sortedNodes[j].Name
	})
	return sortedNodes
}

func (r *databaseResource) mapRolesToResourceModel(roles []*models.DatabaseRole) types.List {
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

// Helper functions

func (r *databaseResource) mapBackupRepositoriesToResourceModel(repositories []*models.BackupRepository) types.List {
	repoList := []attr.Value{}
	for _, repo := range repositories {
		repoObj, _ := types.ObjectValue(
			backupRepositoryType,
			map[string]attr.Value{
				"id":                  types.StringValue(repo.ID),
				"type":                types.StringValue(repo.Type),
				"backup_store_id":     types.StringValue(repo.BackupStoreID),
				"base_path":           types.StringValue(repo.BasePath),
				"s3_bucket":           types.StringValue(repo.S3Bucket),
				"s3_region":           types.StringValue(repo.S3Region),
				"s3_endpoint":         types.StringValue(repo.S3Endpoint),
				"gcs_bucket":          types.StringValue(repo.GcsBucket),
				"gcs_endpoint":        types.StringValue(repo.GcsEndpoint),
				"azure_account":       types.StringValue(repo.AzureAccount),
				"azure_container":     types.StringValue(repo.AzureContainer),
				"azure_endpoint":      types.StringValue(repo.AzureEndpoint),
				"retention_full":      types.Int64Value(repo.RetentionFull),
				"retention_full_type": types.StringValue(repo.RetentionFullType),
			},
		)
		repoList = append(repoList, repoObj)
	}
	return types.ListValueMust(types.ObjectType{AttrTypes: backupRepositoryType}, repoList)
}

func (r *databaseResource) mapBackupSchedulesToResourceModel(schedules []*models.BackupSchedule) types.List {
	scheduleList := []attr.Value{}
	for _, schedule := range schedules {
		scheduleObj, _ := types.ObjectValue(
			backupScheduleType,
			map[string]attr.Value{
				"id":              types.StringPointerValue(schedule.ID),
				"type":            types.StringPointerValue(schedule.Type),
				"cron_expression": types.StringPointerValue(schedule.CronExpression),
			},
		)
		scheduleList = append(scheduleList, scheduleObj)
	}
	return types.ListValueMust(types.ObjectType{AttrTypes: backupScheduleType}, scheduleList)
}

func (r *databaseResource) connectionAttrTypes() map[string]attr.Type {
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

func (r *databaseResource) locationAttrTypes() map[string]attr.Type {
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

func (r *databaseResource) regionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"active":             types.BoolType,
		"availability_zones": types.ListType{ElemType: types.StringType},
		"cloud":              types.StringType,
		"code":               types.StringType,
		"name":               types.StringType,
		"parent":             types.StringType,
	}
}

func (r *databaseResource) nodeExtensionsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"errors":    types.MapType{ElemType: types.StringType},
		"installed": types.ListType{ElemType: types.StringType},
	}
}

func (r *databaseResource) mapConnectionToResourceModel(connection *models.Connection) types.Object {
	if connection == nil {
		return types.ObjectNull(r.connectionAttrTypes())
	}
	connectionObj, _ := types.ObjectValue(
		r.connectionAttrTypes(),
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
		return types.ObjectNull(r.locationAttrTypes())
	}
	locationObj, _ := types.ObjectValue(
		r.locationAttrTypes(),
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
		return types.ObjectNull(r.regionAttrTypes())
	}
	regionObj, _ := types.ObjectValue(
		r.regionAttrTypes(),
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
		return types.ObjectNull(r.nodeExtensionsAttrTypes())
	}
	errorsMap := make(map[string]attr.Value)
	for k, v := range extensions.Errors {
		errorsMap[k] = types.StringValue(v)
	}
	extensionsObj, _ := types.ObjectValue(
		r.nodeExtensionsAttrTypes(),
		map[string]attr.Value{
			"errors":    types.MapValueMust(types.StringType, errorsMap),
			"installed": types.ListValueMust(types.StringType, r.stringSliceToValueSlice(extensions.Installed)),
		},
	)
	return extensionsObj
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
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	ClusterID types.String `tfsdk:"cluster_id"`
	Status    types.String `tfsdk:"status"`
	CreatedAt types.String `tfsdk:"created_at"`
	PgVersion types.String `tfsdk:"pg_version"`
	// StorageUsed   types.Int64  `tfsdk:"storage_used"`
	Domain        types.String `tfsdk:"domain"`
	ConfigVersion types.String `tfsdk:"config_version"`
	Options       types.List   `tfsdk:"options"`
	Backups       types.Object `tfsdk:"backups"`
	Components    types.List   `tfsdk:"components"`
	Extensions    types.Object `tfsdk:"extensions"`
	Nodes         types.Map    `tfsdk:"nodes"`
	Roles         types.List   `tfsdk:"roles"`
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
