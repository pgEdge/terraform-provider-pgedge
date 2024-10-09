package cloudaccount

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
	_ datasource.DataSource              = &cloudAccountsDataSource{}
	_ datasource.DataSourceWithConfigure = &cloudAccountsDataSource{}
)

func NewCloudAccountsDataSource() datasource.DataSource {
	return &cloudAccountsDataSource{}
}

type cloudAccountsDataSource struct {
	client *pgEdge.Client
}

func (d *cloudAccountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_accounts"
}

func (d *cloudAccountsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

type CloudAccountsDataSourceModel struct {
	CloudAccounts []CloudAccountDetails `tfsdk:"cloud_accounts"`
}

type CloudAccountDetails struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	Description types.String `tfsdk:"description"`
	Properties  types.Map    `tfsdk:"properties"`
}

func (d *cloudAccountsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cloud_accounts": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "ID of the cloud account",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the cloud account",
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the cloud account (e.g., AWS, Azure, GCP)",
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: "Creation time of the cloud account",
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: "Last update time of the cloud account",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "Description of the cloud account",
						},
						"properties": schema.MapAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: "Additional properties of the cloud account",
						},
					},
				},
			},
		},
		Description: "Data source for pgEdge cloud accounts.",
	}
}

func (d *cloudAccountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state CloudAccountsDataSourceModel

	cloudAccounts, err := d.client.GetCloudAccounts(ctx)
	if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "reading cloud accounts"))
        return
    }

	for _, account := range cloudAccounts {
		properties := make(map[string]attr.Value)
		// Check if Properties is a map[string]interface{}
		if props, ok := account.Properties.(map[string]interface{}); ok {
			for k, v := range props {
				// Convert each value to a string
				strValue := fmt.Sprintf("%v", v)
				properties[k] = types.StringValue(strValue)
			}
		}
		propertiesMap, diags := types.MapValue(types.StringType, properties)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		accountDetails := CloudAccountDetails{
			ID:          types.StringValue(account.ID.String()),
			Name:        types.StringValue(*account.Name),
			Type:        types.StringValue(*account.Type),
			CreatedAt:   types.StringValue(*account.CreatedAt),
			UpdatedAt:   types.StringValue(*account.UpdatedAt),
			Description: types.StringValue(account.Description),
			Properties:  propertiesMap,
		}

		state.CloudAccounts = append(state.CloudAccounts, accountDetails)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}