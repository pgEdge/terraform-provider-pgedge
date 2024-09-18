package cluster

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

var (
	_ resource.Resource              = &clusterResource{}
	_ resource.ResourceWithConfigure = &clusterResource{}
)

func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

type clusterResource struct {
	client *pgEdge.Client
}

func (r *clusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *clusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"cloud_account_id": schema.StringAttribute{
				Required: true,
			},
			"ssh_key_id": schema.StringAttribute{
				Optional: true,
			},
			"regions": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"node_location": schema.StringAttribute{
				Required: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"nodes": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":              schema.StringAttribute{Required: true},
						"region":            schema.StringAttribute{Required: true},
						"instance_type":     schema.StringAttribute{Required: true},
						"availability_zone": schema.StringAttribute{Optional: true},
						"volume_size":       schema.Int64Attribute{Optional: true},
						"volume_type":       schema.StringAttribute{Optional: true},
						"volume_iops":       schema.Int64Attribute{Optional: true},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region":         schema.StringAttribute{Required: true},
						"cidr":           schema.StringAttribute{Required: true},
						"public_subnets": schema.ListAttribute{Required: true, ElementType: types.StringType},
					},
				},
			},
			"firewall_rules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":    schema.StringAttribute{Required: true},
						"port":    schema.Int64Attribute{Required: true},
						"sources": schema.ListAttribute{Required: true, ElementType: types.StringType},
					},
				},
			},
		},
	}
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan clusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	regions, diags := convertToStringSlice(ctx, plan.Regions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sortedRegions := sortRegions(regions)


	createInput := &models.CreateClusterInput{
		Name:           plan.Name.ValueStringPointer(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		NodeLocation:   plan.NodeLocation.ValueStringPointer(),
		Regions:        sortedRegions,
		Nodes:          make([]*models.ClusterNodeSettings, 0),
		Networks:       make([]*models.ClusterNetworkSettings, 0),
		FirewallRules:  make([]*models.ClusterFirewallRuleSettings, 0),
	}

	if !plan.SSHKeyID.IsNull() {
		createInput.SSHKeyID = plan.SSHKeyID.ValueString()
	}

	for _, node := range plan.Nodes {
		createInput.Nodes = append(createInput.Nodes, &models.ClusterNodeSettings{
			Name:             node.Name.ValueString(),
			Region:           node.Region.ValueStringPointer(),
			InstanceType:     node.InstanceType.ValueString(),
			AvailabilityZone: node.AvailabilityZone.ValueString(),
			VolumeSize:       node.VolumeSize.ValueInt64(),
			VolumeType:       node.VolumeType.ValueString(),
			VolumeIops:       node.VolumeIops.ValueInt64(),
		})
	}

	for _, network := range plan.Networks {
		publicSubnets, diags := convertToStringSlice(ctx, network.PublicSubnets)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		createInput.Networks = append(createInput.Networks, &models.ClusterNetworkSettings{
			Region:        network.Region.ValueStringPointer(),
			Cidr:          network.Cidr.ValueString(),
			PublicSubnets: publicSubnets,
		})
	}

	for _, rule := range plan.FirewallRules {
		sources, diags := convertToStringSlice(ctx, rule.Sources)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}


		createInput.FirewallRules = append(createInput.FirewallRules, &models.ClusterFirewallRuleSettings{
			Name:    rule.Name.ValueString(),
			Port:    rule.Port.ValueInt64Pointer(),
			Sources: sources,
		})
	}

	tflog.Debug(ctx, "Creating cluster", map[string]interface{}{"create_input": createInput})

	cluster, err := r.client.CreateCluster(ctx, createInput)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating cluster",
			"Could not create cluster, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Created cluster", map[string]interface{}{"cluster": cluster})

	plan.ID = types.StringValue(cluster.ID.String())
	plan.Status = types.StringPointerValue(cluster.Status)
	plan.CreatedAt = types.StringPointerValue(cluster.CreatedAt)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state clusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cluster, err := r.client.GetCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading cluster",
			"Could not read cluster ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Name = types.StringPointerValue(cluster.Name)
	state.CloudAccountID = types.StringPointerValue(cluster.CloudAccount.ID)
	state.NodeLocation = types.StringPointerValue(cluster.NodeLocation)
	state.Status = types.StringPointerValue(cluster.Status)
	state.CreatedAt = types.StringPointerValue(cluster.CreatedAt)

	if cluster.SSHKeyID != "" {
		state.SSHKeyID = types.StringValue(cluster.SSHKeyID)
	}

	sortedRegions := sortRegions(cluster.Regions)

	state.Regions = types.ListValueMust(types.StringType, func() []attr.Value {
		regions := make([]attr.Value, len(sortedRegions))
		for i, region := range sortedRegions {
			regions[i] = types.StringValue(region)
		}
		return regions
	}())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nodes := make([]nodeModel, 0)
	for _, node := range cluster.Nodes {
		nodes = append(nodes, nodeModel{
			Name:             types.StringValue(node.Name),
			Region:           types.StringValue(*node.Region),
			InstanceType:     types.StringValue(node.InstanceType),
			AvailabilityZone: types.StringValue(node.AvailabilityZone),
			VolumeSize:       types.Int64Value(node.VolumeSize),
			VolumeType:       types.StringValue(node.VolumeType),
			VolumeIops:       types.Int64Value(node.VolumeIops),
		})
	}
	state.Nodes = nodes

	// Handle Networks
	networks := make([]networkModel, 0)
	for _, network := range cluster.Networks {
		networks = append(networks, networkModel{
			Region: types.StringValue(*network.Region),
			Cidr:   types.StringValue(network.Cidr),
			PublicSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(network.PublicSubnets))
				for i, subnet := range network.PublicSubnets {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
		})
	}
	state.Networks = networks

	// Handle Firewall Rules
	firewallRules := make([]firewallRuleModel, 0)
	for _, rule := range cluster.FirewallRules {
		firewallRules = append(firewallRules, firewallRuleModel{
			Name: types.StringValue(rule.Name),
			Port: types.Int64Value(*rule.Port),
			Sources: types.ListValueMust(types.StringType, func() []attr.Value {
				sources := make([]attr.Value, len(rule.Sources))
				for i, source := range rule.Sources {
					sources[i] = types.StringValue(source)
				}
				return sources
			}()),
		})
	}
	state.FirewallRules = firewallRules

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan clusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	regions, diags := convertToStringSlice(ctx, plan.Regions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sortedRegions := sortRegions(regions)

	updateInput := &models.UpdateClusterInput{
		SSHKeyID: plan.SSHKeyID.ValueString(),
		Regions:  sortedRegions,
		Nodes:    make([]*models.ClusterNodeSettings, 0),
	}

	// Add nodes
	for _, node := range plan.Nodes {
        nodeSettings := &models.ClusterNodeSettings{
            Name:             node.Name.ValueString(),
            Region:           node.Region.ValueStringPointer(),
            AvailabilityZone: node.AvailabilityZone.ValueString(),
            InstanceType:     node.InstanceType.ValueString(),
            VolumeSize:       node.VolumeSize.ValueInt64(),
            VolumeType:       node.VolumeType.ValueString(),
            VolumeIops:       node.VolumeIops.ValueInt64(),
        }

        updateInput.Nodes = append(updateInput.Nodes, nodeSettings)
    }

	
  // Add Networks
  for _, network := range plan.Networks {
	publicSubnets, diags := convertToStringSlice(ctx, network.PublicSubnets)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networkSettings := &models.ClusterNetworkSettings{
		Region:        network.Region.ValueStringPointer(),
		Cidr:          network.Cidr.ValueString(),
		PublicSubnets: publicSubnets,
	}

	updateInput.Networks = append(updateInput.Networks, networkSettings)
}

    // Set FirewallRules
    for _, rule := range plan.FirewallRules {
        sources, diags := convertToStringSlice(ctx, rule.Sources)
        resp.Diagnostics.Append(diags...)
        if resp.Diagnostics.HasError() {
            return
        }

        updateInput.FirewallRules = append(updateInput.FirewallRules, &models.ClusterFirewallRuleSettings{
            Name:    rule.Name.ValueString(),
            Port:    rule.Port.ValueInt64Pointer(),
            Sources: sources,
        })
    }

	cluster, err := r.client.UpdateCluster(ctx, strfmt.UUID(*plan.ID.ValueStringPointer()), updateInput)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating cluster",
			"Could not update cluster, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Status = types.StringValue(*cluster.Status)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state clusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting cluster",
			"Could not delete cluster, unexpected error: "+err.Error(),
		)
		return
	}
}

type clusterResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	CloudAccountID types.String `tfsdk:"cloud_account_id"`
	SSHKeyID       types.String `tfsdk:"ssh_key_id"`
	Regions        types.List   `tfsdk:"regions"`
	NodeLocation   types.String `tfsdk:"node_location"`
	Status         types.String `tfsdk:"status"`
	CreatedAt      types.String `tfsdk:"created_at"`
	Nodes          []nodeModel  `tfsdk:"nodes"`
	Networks       []networkModel `tfsdk:"networks"`
	FirewallRules  []firewallRuleModel `tfsdk:"firewall_rules"`
}

type nodeModel struct {
    Name             types.String  `tfsdk:"name"`
    Region           types.String  `tfsdk:"region"`
    InstanceType     types.String  `tfsdk:"instance_type"`
    AvailabilityZone types.String  `tfsdk:"availability_zone"`
    VolumeSize       types.Int64   `tfsdk:"volume_size"`
    VolumeType       types.String  `tfsdk:"volume_type"`
    VolumeIops       types.Int64   `tfsdk:"volume_iops"`
}

type networkModel struct {
	Region        types.String `tfsdk:"region"`
	Cidr          types.String `tfsdk:"cidr"`
	PublicSubnets types.List   `tfsdk:"public_subnets"`
}

type firewallRuleModel struct {
	Name    types.String `tfsdk:"name"`
	Port    types.Int64  `tfsdk:"port"`
	Sources types.List   `tfsdk:"sources"`
}

func convertToStringSlice(ctx context.Context, list types.List) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var result []string

	for _, item := range list.Elements() {
		if str, ok := item.(types.String); ok {
			result = append(result, str.ValueString())
		} else {
			diags.AddError(
				"Conversion Error",
				fmt.Sprintf("Expected string value, got: %T", item),
			)
		}
	}

	return result, diags
}


func sortRegions(regions []string) []string {
    sort.Strings(regions)
    return regions
}