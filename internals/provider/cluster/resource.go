package cluster

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider/common"
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
				Required:    true,
				Description: "Node location of the cluster. Must be either 'public' or 'private'.",
				Validators: []validator.String{
					stringvalidator.OneOf("public", "private"),
				},
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
						"availability_zone": schema.StringAttribute{Optional: true, Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
						"volume_size":       schema.Int64Attribute{Optional: true, Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}},
						"volume_type":       schema.StringAttribute{Optional: true, Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
						"volume_iops":       schema.Int64Attribute{Optional: true, Computed: true, PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Required:    true,
							Description: "CIDR of the network",
						},
						"external": schema.BoolAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
							Description: "Whether the network is external",
						},
						"external_id": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Description: "External ID of the network",
						},
						"name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Description: "Name of the network",
						},
						"private_subnets": schema.ListAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							ElementType: types.StringType,
							Description: "List of private subnets",
						},
						"public_subnets": schema.ListAttribute{
							Required:    true,
							ElementType: types.StringType,
							Description: "List of public subnets",
						},
						"region": schema.StringAttribute{
							Required:    true,
							Description: "Region of the network",
						},
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
			"backup_store_ids": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of backup store IDs to associate with the cluster",
			},
			"resource_tags": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
				Description: "A map of tags to assign to the cluster",
			},
			"capacity": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func ValidateFirewallRuleName(name string) error {
	validChars := "abcdefghijklmnopqrstuvwxyz0123456789-"
	for _, char := range name {
		if !strings.ContainsRune(validChars, char) {
			return fmt.Errorf("firewall rule name can only contain lowercase letters, numbers, and hyphens, got: %s", name)
		}
	}

	if len(name) < 3 || len(name) > 63 {
		return fmt.Errorf("firewall rule name must be between 3 and 63 characters, got length: %d", len(name))
	}

	if name[0] < 'a' || name[0] > 'z' {
		return fmt.Errorf("firewall rule name must start with a lowercase letter, got: %s", name)
	}

	return nil
}

func ValidateClusterConfiguration(plan *clusterResourceModel) error {
	if err := validateFirewallRules(plan.FirewallRules); err != nil {
		return err
	}

	if err := validateNetworks(plan.Networks, plan.Nodes, plan.NodeLocation); err != nil {
		return err
	}

	return nil
}

func validateFirewallRules(rules []firewallRuleModel) error {
	for i, rule := range rules {
		if err := ValidateFirewallRuleName(rule.Name.ValueString()); err != nil {
			return fmt.Errorf("invalid firewall rule at index %d: %w", i, err)
		}

		port := rule.Port.ValueInt64()
		if port < 1 || port > 65535 {
			return fmt.Errorf("invalid port number %d at index %d: must be between 1 and 65535", port, i)
		}

		sources := common.ConvertTFListToStringSlice(rule.Sources)
		if len(sources) == 0 {
			return fmt.Errorf("firewall rule at index %d must have at least one source", i)
		}

		for _, source := range sources {
			if !isValidCIDR(source) {
				return fmt.Errorf("invalid CIDR format %s in firewall rule at index %d", source, i)
			}
		}
	}
	return nil
}

func validateNetworks(networks []networkModel, nodes []nodeModel, nodeLocation types.String) error {
	if nodeLocation.ValueString() == "private" {
		nodesPerRegion := make(map[string]int)
		for _, node := range nodes {
			nodesPerRegion[node.Region.ValueString()]++
		}

		for i, network := range networks {
			region := network.Region.ValueString()

			if network.PrivateSubnets.IsNull() || len(network.PrivateSubnets.Elements()) == 0 {
				return fmt.Errorf("private subnets are required at node_groups.aws[%d].private_subnets for private nodes", i)
			}

			nodeCount := nodesPerRegion[region]
			subnetCount := len(network.PrivateSubnets.Elements())
			if nodeCount > subnetCount {
				return fmt.Errorf("more nodes (%d) than available subnets (%d) in node_groups.aws[%d] for private nodes",
					nodeCount, subnetCount, i)
			}
		}
	}
	return nil
}

func isValidCIDR(cidr string) bool {
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return false
	}

	ipParts := strings.Split(parts[0], ".")
	if len(ipParts) != 4 {
		return false
	}

	for _, part := range ipParts {
		num := 0
		for _, digit := range part {
			if digit < '0' || digit > '9' {
				return false
			}
			num = num*10 + int(digit-'0')
		}
		if num < 0 || num > 255 {
			return false
		}
	}

	prefix, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	if prefix < 0 || prefix > 32 {
		return false
	}

	return true
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan clusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := ValidateClusterConfiguration(&plan); err != nil {
		resp.Diagnostics.AddError(
			"Invalid Cluster Configuration",
			err.Error(),
		)
		return
	}

	if err := validateRegions(plan.Nodes, plan.Regions, plan.Networks); err != nil {
		resp.Diagnostics.AddError(
			"Invalid Region Configuration",
			err.Error(),
		)
		return
	}

	sort.Slice(plan.Nodes, func(i, j int) bool {
		return plan.Nodes[i].Name.ValueString() < plan.Nodes[j].Name.ValueString()
	})

	regions, diags := convertToStringSlice(ctx, plan.Regions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	backupStoreIds, diags := convertToStringSlice(ctx, plan.BackupStoreIDs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceTags, diags := convertToStringMap(ctx, plan.ResourceTags)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createInput := &models.CreateClusterInput{
		Name:           plan.Name.ValueStringPointer(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		NodeLocation:   plan.NodeLocation.ValueStringPointer(),
		Nodes:          make([]*models.ClusterNodeSettings, 0),
		Networks:       make([]*models.ClusterNetworkSettings, 0),
		FirewallRules:  make([]*models.ClusterFirewallRuleSettings, 0),
		BackupStoreIds: backupStoreIds,
		ResourceTags:   resourceTags,
		Capacity:       plan.Capacity.ValueInt64(),
	}
	createInput.Regions = regions

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

		privateSubnets, diags := convertToStringSlice(ctx, network.PrivateSubnets)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Ensure deterministic ordering
		sort.Strings(publicSubnets)
		sort.Strings(privateSubnets)

		createInput.Networks = append(createInput.Networks, &models.ClusterNetworkSettings{
			Region:         network.Region.ValueStringPointer(),
			Cidr:           network.Cidr.ValueString(),
			PublicSubnets:  publicSubnets,
			PrivateSubnets: privateSubnets,
			Name:           network.Name.ValueString(),
			External:       network.External.ValueBool(),
			ExternalID:     network.ExternalID.ValueString(),
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
		if cluster != nil {
			mappedCluster := r.mapClusterToResourceModel(cluster)
			mappedCluster.ResourceTags = types.MapNull(types.StringType)
			mappedCluster.BackupStoreIDs = types.ListNull(types.StringType)
			mappedCluster.Regions = types.ListNull(types.StringType)

			diags = resp.State.Set(ctx, mappedCluster)
			resp.Diagnostics.Append(diags...)
		}
		resp.Diagnostics.Append(common.HandleProviderError(err, "cluster creation"))
		return
	}

	tflog.Debug(ctx, "Created cluster", map[string]interface{}{"cluster": cluster})

	plan.ID = types.StringValue(cluster.ID.String())
	plan.Status = types.StringPointerValue(cluster.Status)
	plan.CreatedAt = types.StringPointerValue(cluster.CreatedAt)

	plan.CloudAccountID = types.StringPointerValue(cluster.CloudAccount.ID)
	plan.NodeLocation = types.StringPointerValue(cluster.NodeLocation)
	plan.Capacity = types.Int64Value(cluster.Capacity)

	if cluster.SSHKeyID != "" {
		plan.SSHKeyID = types.StringValue(cluster.SSHKeyID)
	}

	// Update Nodes - ensure consistent ordering by sorting by name
	plan.Nodes = r.mapNodesToResourceModel(cluster.Nodes)

	networks := make([]networkModel, 0)
	for _, network := range cluster.Networks {
		// Sort subnets to ensure deterministic ordering in state
		pub := append([]string(nil), network.PublicSubnets...)
		priv := append([]string(nil), network.PrivateSubnets...)
		sort.Strings(pub)
		sort.Strings(priv)
		networks = append(networks, networkModel{
			Region: types.StringPointerValue(network.Region),
			Cidr:   types.StringValue(network.Cidr),
			PublicSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(pub))
				for i, subnet := range pub {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			PrivateSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(priv))
				for i, subnet := range priv {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			Name:       types.StringValue(network.Name),
			External:   types.BoolValue(network.External),
			ExternalID: types.StringValue(network.ExternalID),
		})
	}
	plan.Networks = networks

	// Update Firewall Rules
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
	plan.FirewallRules = firewallRules

	plan.Regions = types.ListValueMust(types.StringType, func() []attr.Value {
		regions := make([]attr.Value, len(createInput.Regions))
		for i, region := range createInput.Regions {
			regions[i] = types.StringValue(region)
		}
		return regions
	}())

	plan.BackupStoreIDs = types.ListValueMust(types.StringType, func() []attr.Value {
		backupStoreIDs := make([]attr.Value, len(createInput.BackupStoreIds))
		for i, backupStoreID := range createInput.BackupStoreIds {
			backupStoreIDs[i] = types.StringValue(backupStoreID)
		}
		return backupStoreIDs
	}())

	plan.ResourceTags = types.MapValueMust(types.StringType, func() map[string]attr.Value {
		resourceTags := make(map[string]attr.Value)
		for k, v := range createInput.ResourceTags {
			resourceTags[k] = types.StringValue(v)
		}
		return resourceTags
	}())

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
		diag := common.HandleProviderError(err, "reading cluster")
		if diag == nil {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.Append(diag)
		return
	}

	state.Name = types.StringPointerValue(cluster.Name)
	state.CloudAccountID = types.StringPointerValue(cluster.CloudAccount.ID)
	state.NodeLocation = types.StringPointerValue(cluster.NodeLocation)
	state.Status = types.StringPointerValue(cluster.Status)
	state.CreatedAt = types.StringPointerValue(cluster.CreatedAt)
	state.Capacity = types.Int64Value(cluster.Capacity)

	state.BackupStoreIDs = types.ListValueMust(types.StringType,
		func() []attr.Value {
			ids := make([]attr.Value, len(cluster.BackupStoreIds))
			for i, id := range cluster.BackupStoreIds {
				ids[i] = types.StringValue(id)
			}
			return ids
		}())

	state.ResourceTags = types.MapValueMust(types.StringType,
		func() map[string]attr.Value {
			tags := make(map[string]attr.Value)
			for k, v := range cluster.ResourceTags {
				tags[k] = types.StringValue(v)
			}
			return tags
		}())

	if cluster.SSHKeyID != "" {
		state.SSHKeyID = types.StringValue(cluster.SSHKeyID)
	}

	var currentRegions []string
	state.Regions.ElementsAs(ctx, &currentRegions, false)

	if !compareRegions(cluster.Regions, currentRegions) {
		state.Regions = types.ListValueMust(types.StringType, func() []attr.Value {
			regions := make([]attr.Value, len(cluster.Regions))
			for i, region := range cluster.Regions {
				regions[i] = types.StringValue(region)
			}
			return regions
		}())
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Nodes = r.mapNodesToResourceModel(cluster.Nodes)

	networks := make([]networkModel, 0)
	for _, network := range cluster.Networks {
		pub := append([]string(nil), network.PublicSubnets...)
		priv := append([]string(nil), network.PrivateSubnets...)
		sort.Strings(pub)
		sort.Strings(priv)
		networks = append(networks, networkModel{
			Region: types.StringPointerValue(network.Region),
			Cidr:   types.StringValue(network.Cidr),
			PublicSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(pub))
				for i, subnet := range pub {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			PrivateSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(priv))
				for i, subnet := range priv {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			Name:       types.StringValue(network.Name),
			External:   types.BoolValue(network.External),
			ExternalID: types.StringValue(network.ExternalID),
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
	var plan, state clusterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := validateRegions(plan.Nodes, plan.Regions, plan.Networks); err != nil {
		resp.Diagnostics.AddError(
			"Invalid Region Configuration",
			err.Error(),
		)
		return
	}

	regions, diags := convertToStringSlice(ctx, plan.Regions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	planBackupStoreIds, diags := convertToStringSlice(ctx, plan.BackupStoreIDs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	stateBackupStoreIds, diags := convertToStringSlice(ctx, state.BackupStoreIDs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceTags, diags := convertToStringMap(ctx, plan.ResourceTags)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateInput := &models.UpdateClusterInput{
		Nodes:    make([]*models.ClusterNodeSettings, 0),
		Networks: make([]*models.ClusterNetworkSettings, 0),
	}
	updateInput.Regions = regions
	if !compareStringSlices(planBackupStoreIds, stateBackupStoreIds) {
		updateInput.BackupStoreIds = planBackupStoreIds
	}
	updateInput.ResourceTags = resourceTags

	// Check if nodes have changed
	nodesChanged := !compareNodes(plan.Nodes, state.Nodes)

	// Add nodes only if they have changed
	if nodesChanged {
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
	}

	// Check if networks have changed
	networksChanged := !compareNetworks(plan.Networks, state.Networks)

	// Add Networks only if they have changed
	if networksChanged {
		for _, network := range plan.Networks {
			publicSubnets, diags := convertToStringSlice(ctx, network.PublicSubnets)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			privateSubnets, diags := convertToStringSlice(ctx, network.PrivateSubnets)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			// Ensure deterministic ordering
			sort.Strings(publicSubnets)
			sort.Strings(privateSubnets)

			networkSettings := &models.ClusterNetworkSettings{
				Region:         network.Region.ValueStringPointer(),
				Cidr:           network.Cidr.ValueString(),
				PublicSubnets:  publicSubnets,
				PrivateSubnets: privateSubnets,
				Name:           network.Name.ValueString(),
				External:       network.External.ValueBool(),
				ExternalID:     network.ExternalID.ValueString(),
			}

			updateInput.Networks = append(updateInput.Networks, networkSettings)
		}
	}

	// Check if firewall rules have changed
	firewallRulesChanged := !compareFirewallRules(plan.FirewallRules, state.FirewallRules)

	// Set FirewallRules only if they have changed
	if firewallRulesChanged {
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
	}

	cluster, err := r.client.UpdateCluster(ctx, strfmt.UUID(*plan.ID.ValueStringPointer()), updateInput)
	if err != nil {
		resp.Diagnostics.Append(common.HandleProviderError(err, "updating cluster"))
		return
	}

	updatedPlan := r.mapClusterToResourceModel(cluster)

	if len(updateInput.BackupStoreIds) > 0 {
		updatedPlan.BackupStoreIDs = types.ListValueMust(types.StringType,
			func() []attr.Value {
				values := make([]attr.Value, len(updateInput.BackupStoreIds))
				for i, id := range updateInput.BackupStoreIds {
					values[i] = types.StringValue(id)
				}
				return values
			}(),
		)
	} else {
		updatedPlan.BackupStoreIDs = state.BackupStoreIDs
	}
	updatedPlan.Regions = plan.Regions
	updatedPlan.Status = types.StringPointerValue(cluster.Status)
	updatedPlan.ResourceTags = plan.ResourceTags

	updatedPlan.Networks = make([]networkModel, 0)
	for _, network := range cluster.Networks {
		pub := append([]string(nil), network.PublicSubnets...)
		priv := append([]string(nil), network.PrivateSubnets...)
		sort.Strings(pub)
		sort.Strings(priv)
		updatedPlan.Networks = append(updatedPlan.Networks, networkModel{
			Region: types.StringPointerValue(network.Region),
			Cidr:   types.StringValue(network.Cidr),
			PublicSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(pub))
				for i, subnet := range pub {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			PrivateSubnets: types.ListValueMust(types.StringType, func() []attr.Value {
				subnets := make([]attr.Value, len(priv))
				for i, subnet := range priv {
					subnets[i] = types.StringValue(subnet)
				}
				return subnets
			}()),
			Name:       types.StringValue(network.Name),
			External:   types.BoolValue(network.External),
			ExternalID: types.StringValue(network.ExternalID),
		})
	}

	if !plan.SSHKeyID.IsNull() || !plan.SSHKeyID.IsUnknown() {
		updatedPlan.SSHKeyID = plan.SSHKeyID
	}

	diags = resp.State.Set(ctx, updatedPlan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) mapClusterToResourceModel(cluster *models.Cluster) clusterResourceModel {
	clusterResource := clusterResourceModel{
		ID:             types.StringValue(cluster.ID.String()),
		Name:           types.StringPointerValue(cluster.Name),
		CloudAccountID: types.StringPointerValue(cluster.CloudAccount.ID),
		NodeLocation:   types.StringPointerValue(cluster.NodeLocation),
		CreatedAt:      types.StringPointerValue(cluster.CreatedAt),
		// Regions:        r.mapRegionsToResourceModel(cluster.Regions),
		Nodes:         r.mapNodesToResourceModel(cluster.Nodes),
		FirewallRules: r.mapFirewallRulesToResourceModel(cluster.FirewallRules),
		Capacity:      types.Int64Value(cluster.Capacity),
	}

	return clusterResource
}

func (r *clusterResource) mapRegionsToResourceModel(regions []string) types.List {
	elements := make([]attr.Value, len(regions))
	for i, region := range regions {
		elements[i] = types.StringValue(region)
	}
	return types.ListValueMust(types.StringType, elements)
}

func (r *clusterResource) mapNodesToResourceModel(nodes []*models.ClusterNodeSettings) []nodeModel {
	var result []nodeModel
	for _, node := range nodes {
		result = append(result, nodeModel{
			Name:             types.StringValue(node.Name),
			Region:           types.StringPointerValue(node.Region),
			InstanceType:     types.StringValue(node.InstanceType),
			AvailabilityZone: types.StringValue(node.AvailabilityZone),
			VolumeSize:       types.Int64Value(node.VolumeSize),
			VolumeType:       types.StringValue(node.VolumeType),
			VolumeIops:       types.Int64Value(node.VolumeIops),
		})
	}

	// Sort nodes by name to ensure consistent ordering
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name.ValueString() < result[j].Name.ValueString()
	})

	return result
}

func (r *clusterResource) mapFirewallRulesToResourceModel(rules []*models.ClusterFirewallRuleSettings) []firewallRuleModel {
	var result []firewallRuleModel
	for _, rule := range rules {
		result = append(result, firewallRuleModel{
			Name:    types.StringValue(rule.Name),
			Port:    types.Int64PointerValue(rule.Port),
			Sources: types.ListValueMust(types.StringType, r.stringSliceToValueSlice(rule.Sources)),
		})
	}
	return result
}

func (r *clusterResource) stringSliceToValueSlice(slice []string) []attr.Value {
	valueSlice := make([]attr.Value, len(slice))
	for i, s := range slice {
		valueSlice[i] = types.StringValue(s)
	}
	return valueSlice
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
		resp.Diagnostics.Append(common.HandleProviderError(err, "cluster deletion"))
		return
	}
}

type clusterResourceModel struct {
	ID             types.String        `tfsdk:"id"`
	Name           types.String        `tfsdk:"name"`
	CloudAccountID types.String        `tfsdk:"cloud_account_id"`
	SSHKeyID       types.String        `tfsdk:"ssh_key_id"`
	Regions        types.List          `tfsdk:"regions"`
	NodeLocation   types.String        `tfsdk:"node_location"`
	Status         types.String        `tfsdk:"status"`
	CreatedAt      types.String        `tfsdk:"created_at"`
	Nodes          []nodeModel         `tfsdk:"nodes"`
	Networks       []networkModel      `tfsdk:"networks"`
	FirewallRules  []firewallRuleModel `tfsdk:"firewall_rules"`
	BackupStoreIDs types.List          `tfsdk:"backup_store_ids"`
	ResourceTags   types.Map           `tfsdk:"resource_tags"`
	Capacity       types.Int64         `tfsdk:"capacity"`
}

type nodeModel struct {
	Name             types.String `tfsdk:"name"`
	Region           types.String `tfsdk:"region"`
	InstanceType     types.String `tfsdk:"instance_type"`
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	VolumeSize       types.Int64  `tfsdk:"volume_size"`
	VolumeType       types.String `tfsdk:"volume_type"`
	VolumeIops       types.Int64  `tfsdk:"volume_iops"`
}

type networkModel struct {
	Cidr           types.String `tfsdk:"cidr"`
	External       types.Bool   `tfsdk:"external"`
	ExternalID     types.String `tfsdk:"external_id"`
	Name           types.String `tfsdk:"name"`
	PrivateSubnets types.List   `tfsdk:"private_subnets"`
	PublicSubnets  types.List   `tfsdk:"public_subnets"`
	Region         types.String `tfsdk:"region"`
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

func convertToStringMap(ctx context.Context, m types.Map) (map[string]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	result := make(map[string]string)

	if m.IsNull() || m.IsUnknown() {
		return nil, diags
	}

	for k, v := range m.Elements() {
		if str, ok := v.(types.String); ok {
			result[k] = str.ValueString()
		} else {
			diags.AddError(
				"Conversion Error",
				fmt.Sprintf("Expected string value for key %s, got: %T", k, v),
			)
		}
	}

	return result, diags
}

func compareRegions(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[string]bool)
	for _, v := range a {
		aMap[v] = true
	}
	for _, v := range b {
		if !aMap[v] {
			return false
		}
	}
	return true
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Strings(aCopy)
	sort.Strings(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

func compareNodes(planNodes, stateNodes []nodeModel) bool {
	if len(planNodes) != len(stateNodes) {
		return false
	}

	// Sort both slices by name for comparison
	planCopy := make([]nodeModel, len(planNodes))
	stateCopy := make([]nodeModel, len(stateNodes))
	copy(planCopy, planNodes)
	copy(stateCopy, stateNodes)

	sort.Slice(planCopy, func(i, j int) bool {
		return planCopy[i].Name.ValueString() < planCopy[j].Name.ValueString()
	})
	sort.Slice(stateCopy, func(i, j int) bool {
		return stateCopy[i].Name.ValueString() < stateCopy[j].Name.ValueString()
	})

	for i := range planCopy {
		if planCopy[i].Name.ValueString() != stateCopy[i].Name.ValueString() ||
			planCopy[i].Region.ValueString() != stateCopy[i].Region.ValueString() ||
			planCopy[i].AvailabilityZone.ValueString() != stateCopy[i].AvailabilityZone.ValueString() ||
			planCopy[i].InstanceType.ValueString() != stateCopy[i].InstanceType.ValueString() ||
			planCopy[i].VolumeSize.ValueInt64() != stateCopy[i].VolumeSize.ValueInt64() ||
			planCopy[i].VolumeType.ValueString() != stateCopy[i].VolumeType.ValueString() ||
			planCopy[i].VolumeIops.ValueInt64() != stateCopy[i].VolumeIops.ValueInt64() {
			return false
		}
	}
	return true
}

func compareNetworks(planNetworks, stateNetworks []networkModel) bool {
	if len(planNetworks) != len(stateNetworks) {
		return false
	}

	// Sort both slices by region for comparison
	planCopy := make([]networkModel, len(planNetworks))
	stateCopy := make([]networkModel, len(stateNetworks))
	copy(planCopy, planNetworks)
	copy(stateCopy, stateNetworks)

	sort.Slice(planCopy, func(i, j int) bool {
		return planCopy[i].Region.ValueString() < planCopy[j].Region.ValueString()
	})
	sort.Slice(stateCopy, func(i, j int) bool {
		return stateCopy[i].Region.ValueString() < stateCopy[j].Region.ValueString()
	})

	for i := range planCopy {
		if planCopy[i].Region.ValueString() != stateCopy[i].Region.ValueString() ||
			planCopy[i].Cidr.ValueString() != stateCopy[i].Cidr.ValueString() ||
			planCopy[i].Name.ValueString() != stateCopy[i].Name.ValueString() ||
			planCopy[i].External.ValueBool() != stateCopy[i].External.ValueBool() ||
			planCopy[i].ExternalID.ValueString() != stateCopy[i].ExternalID.ValueString() {
			return false
		}

		// Compare public subnets
		planPublicSubnets, _ := convertToStringSlice(context.Background(), planCopy[i].PublicSubnets)
		statePublicSubnets, _ := convertToStringSlice(context.Background(), stateCopy[i].PublicSubnets)
		if !compareStringSlices(planPublicSubnets, statePublicSubnets) {
			return false
		}

		// Compare private subnets
		planPrivateSubnets, _ := convertToStringSlice(context.Background(), planCopy[i].PrivateSubnets)
		statePrivateSubnets, _ := convertToStringSlice(context.Background(), stateCopy[i].PrivateSubnets)
		if !compareStringSlices(planPrivateSubnets, statePrivateSubnets) {
			return false
		}
	}
	return true
}

func compareFirewallRules(planRules, stateRules []firewallRuleModel) bool {
	if len(planRules) != len(stateRules) {
		return false
	}

	// Sort both slices by name for comparison
	planCopy := make([]firewallRuleModel, len(planRules))
	stateCopy := make([]firewallRuleModel, len(stateRules))
	copy(planCopy, planRules)
	copy(stateCopy, stateRules)

	sort.Slice(planCopy, func(i, j int) bool {
		return planCopy[i].Name.ValueString() < planCopy[j].Name.ValueString()
	})
	sort.Slice(stateCopy, func(i, j int) bool {
		return stateCopy[i].Name.ValueString() < stateCopy[j].Name.ValueString()
	})

	for i := range planCopy {
		if planCopy[i].Name.ValueString() != stateCopy[i].Name.ValueString() ||
			planCopy[i].Port.ValueInt64() != stateCopy[i].Port.ValueInt64() {
			return false
		}

		// Compare sources
		planSources, _ := convertToStringSlice(context.Background(), planCopy[i].Sources)
		stateSources, _ := convertToStringSlice(context.Background(), stateCopy[i].Sources)
		if !compareStringSlices(planSources, stateSources) {
			return false
		}
	}
	return true
}

func validateRegions(nodes []nodeModel, regions types.List, networks []networkModel) error {
	nodeRegions := make(map[string]bool)
	for _, node := range nodes {
		nodeRegions[node.Region.ValueString()] = true
	}

	networkRegions := make(map[string]bool)
	for _, network := range networks {
		networkRegions[network.Region.ValueString()] = true
	}

	regionsArray := make(map[string]bool)
	for _, region := range regions.Elements() {
		if str, ok := region.(types.String); ok {
			regionsArray[str.ValueString()] = true
		}
	}

	if len(nodeRegions) != len(networkRegions) || len(nodeRegions) != len(regionsArray) {
		return fmt.Errorf("regions mismatch: the number of unique regions in nodes, networks, and regions array must be the same")
	}

	for region := range nodeRegions {
		if !regionsArray[region] {
			return fmt.Errorf("regions mismatch: all regions specified in nodes must be present in the regions array")
		}
	}

	for region := range networkRegions {
		if !regionsArray[region] {
			return fmt.Errorf("regions mismatch: all regions specified in networks must be present in the regions array")
		}
	}

	return nil
}
