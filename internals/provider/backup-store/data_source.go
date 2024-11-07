package backupstore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

var (
    _ datasource.DataSource              = &backupStoresDataSource{}
    _ datasource.DataSourceWithConfigure = &backupStoresDataSource{}
)

func NewBackupStoresDataSource() datasource.DataSource {
    return &backupStoresDataSource{}
}

type backupStoresDataSource struct {
    client *pgEdge.Client
}

func (d *backupStoresDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_backup_stores"
}

func (d *backupStoresDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type BackupStoresDataSourceModel struct {
    BackupStores []BackupStoreModel `tfsdk:"backup_stores"`
}

type BackupStoreModel struct {
    ID               types.String `tfsdk:"id"`
    CloudAccountID   types.String `tfsdk:"cloud_account_id"`
    CloudAccountType types.String `tfsdk:"cloud_account_type"`
    CreatedAt        types.String `tfsdk:"created_at"`
    Status           types.String `tfsdk:"status"`
    Name             types.String `tfsdk:"name"`
    Properties       types.Map    `tfsdk:"properties"`
    ClusterIDs       types.List   `tfsdk:"cluster_ids"`
}

func (d *backupStoresDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "backup_stores": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.StringAttribute{
                            Computed: true,
                        },
                        "cloud_account_id": schema.StringAttribute{
                            Computed: true,
                        },
                        "cloud_account_type": schema.StringAttribute{
                            Computed: true,
                        },
                        "created_at": schema.StringAttribute{
                            Computed: true,
                        },
                        "status": schema.StringAttribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "properties": schema.MapAttribute{
                            Computed:    true,
                            ElementType: types.StringType,
                        },
                        "cluster_ids": schema.ListAttribute{
                            Computed:    true,
                            ElementType: types.StringType,
                        },
                    },
                },
            },
        },
    }
}

func (d *backupStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    var state BackupStoresDataSourceModel

    backupStores, err := d.client.GetBackupStores(ctx, nil, nil, nil, nil, nil)
    if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "reading backup stores"))
        return
    }

    for _, store := range backupStores {
        storeModel := BackupStoreModel{
            ID:               types.StringValue(store.ID.String()),
            CloudAccountID:   types.StringPointerValue(store.CloudAccountID),
            CloudAccountType: types.StringPointerValue(store.CloudAccountType),
            CreatedAt:        types.StringPointerValue(store.CreatedAt),
            Status:           types.StringPointerValue(store.Status),
            Name:             types.StringPointerValue(store.Name),
            Properties:       d.convertPropertiesToMap(store.Properties),
            ClusterIDs:       types.ListValueMust(types.StringType, d.convertToValueSlice(store.ClusterIds)),
        }

        state.BackupStores = append(state.BackupStores, storeModel)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (d *backupStoresDataSource) convertPropertiesToMap(properties interface{}) types.Map {
    propertiesMap := make(map[string]attr.Value)
    
    if props, ok := properties.(map[string]interface{}); ok {
        for k, v := range props {
            propertiesMap[k] = types.StringValue(fmt.Sprintf("%v", v))
        }
    }
    
    return types.MapValueMust(types.StringType, propertiesMap)
}

func (d *backupStoresDataSource) convertToValueSlice(slice []string) []attr.Value {
    valueSlice := make([]attr.Value, len(slice))
    for i, v := range slice {
        valueSlice[i] = types.StringValue(v)
    }
    return valueSlice
}
