package sshkey

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
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
    _ resource.Resource                = &sshKeyResource{}
    _ resource.ResourceWithConfigure   = &sshKeyResource{}
    _ resource.ResourceWithImportState = &sshKeyResource{}
)

func NewSSHKeyResource() resource.Resource {
    return &sshKeyResource{}
}

type sshKeyResource struct {
    client *pgEdge.Client
}

func (r *sshKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_ssh_key"
}

func (r *sshKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Description: "Manages a pgEdge SSH key.",
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
                Description: "Unique identifier for the SSH key.",
                Computed:    true,
                PlanModifiers: []planmodifier.String{
                    stringplanmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the SSH key.",
                Required:    true,
            },
            "public_key": schema.StringAttribute{
                Description: "The public key.",
                Required:    true,
            },
            "created_at": schema.StringAttribute{
                Description: "The timestamp when the SSH key was created.",
                Computed:    true,
            },
        },
    }
}

func (r *sshKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type sshKeyResourceModel struct {
    ID        types.String `tfsdk:"id"`
    Name      types.String `tfsdk:"name"`
    PublicKey types.String `tfsdk:"public_key"`
    CreatedAt types.String `tfsdk:"created_at"`
}

func (r *sshKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan sshKeyResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    createInput := &models.CreateSSHKeyInput{
        Name:      plan.Name.ValueStringPointer(),
        PublicKey: plan.PublicKey.ValueStringPointer(),
    }

    sshKey, err := r.client.CreateSSHKey(ctx, createInput)
    if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "ssh key creation"))
        return
    }

    plan.ID = types.StringValue(sshKey.ID.String())
    plan.CreatedAt = types.StringValue(*sshKey.CreatedAt)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *sshKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    var state sshKeyResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    sshKey, err := r.client.GetSSHKey(ctx, strfmt.UUID(state.ID.ValueString()))
    if err != nil {
		diag := common.HandleProviderError(err, "ssh key retrieval")
        if diag == nil {
            resp.State.RemoveResource(ctx)
            return
        }
        resp.Diagnostics.Append(diag)
        return
    }

    state.Name = types.StringValue(*sshKey.Name)
    state.PublicKey = types.StringValue(*sshKey.PublicKey)
    state.CreatedAt = types.StringValue(*sshKey.CreatedAt)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *sshKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    // SSH keys cannot be updated in this API, so we'll return an error
    resp.Diagnostics.AddError(
        "Error Updating SSH Key",
        "SSH Key update is not supported by the API",
    )
}

func (r *sshKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var state sshKeyResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    err := r.client.DeleteSSHKey(ctx, strfmt.UUID(state.ID.ValueString()))
    if err != nil {
        resp.Diagnostics.Append(common.HandleProviderError(err, "ssh key deletion"))
        return
    }
}

func (r *sshKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}