package backupstore

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
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

var (
    _ resource.Resource                = &backupStoreResource{}
    _ resource.ResourceWithConfigure   = &backupStoreResource{}
    _ resource.ResourceWithImportState = &backupStoreResource{}
)

func NewBackupStoreResource() resource.Resource {
    return &backupStoreResource{}
}

type backupStoreResource struct {
    client *pgEdge.Client
}

type backupStoreResourceModel struct {
    ID               types.String `tfsdk:"id"`
    CloudAccountID   types.String `tfsdk:"cloud_account_id"`
    CloudAccountType types.String `tfsdk:"cloud_account_type"`
    CreatedAt        types.String `tfsdk:"created_at"`
    UpdatedAt        types.String `tfsdk:"updated_at"`
    Status           types.String `tfsdk:"status"`
    Name             types.String `tfsdk:"name"`
    Properties       types.Map    `tfsdk:"properties"`
    ClusterIDs       types.List   `tfsdk:"cluster_ids"`
    Region           types.String `tfsdk:"region"`
}

func (r *backupStoreResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_backup_store"
}

func (r *backupStoreResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
                Computed: true,
                PlanModifiers: []planmodifier.String{
                    stringplanmodifier.UseStateForUnknown(),
                },
            },
            "cloud_account_id": schema.StringAttribute{
                Required: true,
            },
            "cloud_account_type": schema.StringAttribute{
                Computed: true,
            },
            "created_at": schema.StringAttribute{
                Computed: true,
            },
            "updated_at": schema.StringAttribute{
                Computed: true,
            },
            "status": schema.StringAttribute{
                Computed: true,
            },
            "name": schema.StringAttribute{
                Required: true,
            },
            "properties": schema.MapAttribute{
                Computed:    true,
                ElementType: types.StringType,
            },
            "cluster_ids": schema.ListAttribute{
                Computed:    true,
                ElementType: types.StringType,
            },
            "region": schema.StringAttribute{
                Required: true,
            },
        },
    }
}

func (r *backupStoreResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *backupStoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan backupStoreResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    input := &models.CreateBackupStoreInput{
        Name:           plan.Name.ValueStringPointer(),
        CloudAccountID: plan.CloudAccountID.ValueStringPointer(),
    }

    if !plan.Region.IsNull() {
        input.Region = plan.Region.ValueString()
    }

    backupStore, err := r.client.CreateBackupStore(ctx, input)
    if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "backup store creation"))
        return
    }

    plan.ID = types.StringValue(backupStore.ID.String())
    plan.CloudAccountType = types.StringPointerValue(backupStore.CloudAccountType)
    plan.CreatedAt = types.StringPointerValue(backupStore.CreatedAt)
    plan.UpdatedAt = types.StringPointerValue(backupStore.UpdatedAt)
    plan.Status = types.StringPointerValue(backupStore.Status)

    // Handle Properties
    plan.Properties = r.convertPropertiesToMap(backupStore.Properties)

    // Handle ClusterIds
    clusterIDs := make([]attr.Value, len(backupStore.ClusterIds))
    for i, id := range backupStore.ClusterIds {
        clusterIDs[i] = types.StringValue(id)
    }
    plan.ClusterIDs = types.ListValueMust(types.StringType, clusterIDs)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *backupStoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    var state backupStoreResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    backupStore, err := r.client.GetBackupStore(ctx, strfmt.UUID(state.ID.ValueString()))
    if err != nil {
        diag := common.HandleProviderError(err, "reading backup store")
        if diag == nil {
            resp.State.RemoveResource(ctx)
            return
        }
        resp.Diagnostics.Append(diag)
        return
    }

    state.CloudAccountID = types.StringPointerValue(backupStore.CloudAccountID)
    state.CloudAccountType = types.StringPointerValue(backupStore.CloudAccountType)
    state.CreatedAt = types.StringPointerValue(backupStore.CreatedAt)
    state.UpdatedAt = types.StringPointerValue(backupStore.UpdatedAt)
    state.Status = types.StringPointerValue(backupStore.Status)
    state.Name = types.StringPointerValue(backupStore.Name)

    // Handle Properties
    state.Properties = r.convertPropertiesToMap(backupStore.Properties)

    // Handle ClusterIds
    clusterIDs := make([]attr.Value, len(backupStore.ClusterIds))
    for i, id := range backupStore.ClusterIds {
        clusterIDs[i] = types.StringValue(id)
    }
    state.ClusterIDs = types.ListValueMust(types.StringType, clusterIDs)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *backupStoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    resp.Diagnostics.AddError(
        "Error Updating Backup Store",
        "Backup Store update is not supported by the API",
    )
}

func (r *backupStoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var state backupStoreResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    err := r.client.DeleteBackupStore(ctx, strfmt.UUID(state.ID.ValueString()))
    if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "backup store deletion"))
        return
    }
}

func (r *backupStoreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *backupStoreResource) convertPropertiesToMap(properties interface{}) types.Map {
    propertiesMap := make(map[string]attr.Value)
    
    if props, ok := properties.(map[string]interface{}); ok {
        for k, v := range props {
            propertiesMap[k] = types.StringValue(fmt.Sprintf("%v", v))
        }
    }
    
    return types.MapValueMust(types.StringType, propertiesMap)
}