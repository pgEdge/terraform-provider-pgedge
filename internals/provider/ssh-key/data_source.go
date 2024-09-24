package sshkey

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
)

var (
    _ datasource.DataSource              = &sshKeysDataSource{}
    _ datasource.DataSourceWithConfigure = &sshKeysDataSource{}
)

func NewSSHKeysDataSource() datasource.DataSource {
    return &sshKeysDataSource{}
}

type sshKeysDataSource struct {
    client *pgEdge.Client
}

func (d *sshKeysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_ssh_keys"
}

func (d *sshKeysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "ssh_keys": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.StringAttribute{
                            Computed:    true,
                            Description: "ID of the SSH key",
                        },
                        "name": schema.StringAttribute{
                            Computed:    true,
                            Description: "Name of the SSH key",
                        },
                        "public_key": schema.StringAttribute{
                            Computed:    true,
                            Description: "Public key",
                        },
                        "created_at": schema.StringAttribute{
                            Computed:    true,
                            Description: "Creation time of the SSH key",
                        },
                    },
                },
            },
        },
        Description: "Data source for pgEdge SSH keys.",
    }
}

type SSHKeysDataSourceModel struct {
    SSHKeys []SSHKeyDetails `tfsdk:"ssh_keys"`
}

type SSHKeyDetails struct {
    ID        types.String `tfsdk:"id"`
    Name      types.String `tfsdk:"name"`
    PublicKey types.String `tfsdk:"public_key"`
    CreatedAt types.String `tfsdk:"created_at"`
}

func (d *sshKeysDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *sshKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    var state SSHKeysDataSourceModel

    sshKeys, err := d.client.GetSSHKeys(ctx)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Read pgEdge SSH Keys",
            err.Error(),
        )
        return
    }

    for _, sshKey := range sshKeys {
        sshKeyDetails := SSHKeyDetails{
            ID:        types.StringValue(sshKey.ID.String()),
            Name:      types.StringValue(*sshKey.Name),
            PublicKey: types.StringValue(*sshKey.PublicKey),
            CreatedAt: types.StringValue(*sshKey.CreatedAt),
        }

        state.SSHKeys = append(state.SSHKeys, sshKeyDetails)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}