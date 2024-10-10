package cloudaccount

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
)

var (
	_ resource.Resource              = &cloudAccountResource{}
	_ resource.ResourceWithConfigure = &cloudAccountResource{}
)

func NewCloudAccountResource() resource.Resource {
	return &cloudAccountResource{}
}

type cloudAccountResource struct {
	client *pgEdge.Client
}

func (r *cloudAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_account"
}

func (r *cloudAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type CloudAccountResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Description types.String `tfsdk:"description"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	Credentials types.Map    `tfsdk:"credentials"`
}

func (r *cloudAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"credentials": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *cloudAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CloudAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	credentials := make(map[string]interface{})
	for k, v := range plan.Credentials.Elements() {
		credentials[k] = v.(types.String).ValueString()
	}

	createInput := &models.CreateCloudAccountInput{
		Name:        *plan.Name.ValueStringPointer(),
		Type:        plan.Type.ValueStringPointer(),
		Description: *plan.Description.ValueStringPointer(),
		Credentials: credentials,
	}

	account, err := r.client.CreateCloudAccount(ctx, createInput)
	if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "cloud account creation"))
        return
    }

	plan.ID = types.StringValue(account.ID.String())
	plan.CreatedAt = types.StringValue(*account.CreatedAt)
	plan.UpdatedAt = types.StringValue(*account.UpdatedAt)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cloudAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CloudAccountResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	account, err := r.client.GetCloudAccount(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		diag := common.HandleProviderError(err, "reading cloud account")
        if diag == nil {
            resp.State.RemoveResource(ctx)
            return
        }
        resp.Diagnostics.Append(diag)
        return
    }

	state.Name = types.StringValue(*account.Name)
	state.Type = types.StringValue(*account.Type)
	state.Description = types.StringValue(account.Description)
	state.CreatedAt = types.StringValue(*account.CreatedAt)
	state.UpdatedAt = types.StringValue(*account.UpdatedAt)

	// Note: We don't update the credentials here as they are not returned by the API for security reasons

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cloudAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// As per the current API, there's no update operation for cloud accounts
	// If this changes in the future, implement the update logic here
	resp.Diagnostics.AddError(
		"Error Updating Cloud Account",
		"Cloud Account update is not supported by the API",
	)
}

func (r *cloudAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CloudAccountResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCloudAccount(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "cloud account deletion"))
        return
    }
}