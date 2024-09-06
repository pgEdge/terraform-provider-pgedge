package cluster

import (
	"context"
	"fmt"
	"regexp"

	// "time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

func (r *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *clusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *clusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "ID of the cluster",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-z0-9]+$`),
						"must contain only lowercase alphanumeric characters",
					),
				},
				Description: "Name of the cluster",
			},
			"cloud_account_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the target cloud account",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "Creation time of the cluster",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Status of the cluster",
			},
			"ssh_key_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the SSH key to add to the cluster nodes",
			},
			"regions": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"node_location": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network location for nodes (public or private)",
			},
			"firewall_rules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Required:    true,
							Description: "Port whose traffic is allowed",
						},
						"sources": schema.ListAttribute{
							Required:    true,
							ElementType: types.StringType,
							Description: "CIDRs and/or IP addresses allowed",
						},
					},
				},
			},
			"nodes": ClusterNodeAttribute,
			"networks": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Name of the network",
						},
						"region": schema.StringAttribute{
							Required:    true,
							Description: "Region of the network",
						},
						"external": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Is the network externally defined",
						},
						"external_id": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "ID of the network, if externally defined",
						},
						"cidr": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "CIDR range for the network",
						},
						"public_subnets": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
						"private_subnets": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
		Description: "Interface with the pgEdge service API for clusters.",
	}
}

func fireWallRulesClusterReq(ctx context.Context, resp *resource.CreateResponse, firewallRuleReq []basetypes.ObjectValue) (*resource.CreateResponse, []*models.FirewallRule) {
	var (
		firewallRules    []*models.FirewallRule
		firewallRuleType FirewallRule
		sources          []string
	)
	for _, firewallRule := range firewallRuleReq {
		diags := firewallRule.As(ctx, &firewallRuleType, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		diags = firewallRuleType.Sources.ElementsAs(ctx, &sources, false)
		resp.Diagnostics.Append(diags...)
		firewallRules = append(firewallRules, &models.FirewallRule{
			Port:    firewallRuleType.Port.ValueInt64(),
			Sources: sources,
		})
	}
	return resp, firewallRules
}

func networksAccountClusterReq(ctx context.Context, resp *resource.CreateResponse, networksReq []basetypes.ObjectValue) (*resource.CreateResponse, []*models.Network) {
	var networks []*models.Network
	for _, network := range networksReq {
		var networkType ClusterNetworks
		diags := network.As(ctx, &networkType, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		networks = append(networks, &models.Network{
			Name:   networkType.Name.ValueString(),
			Region: networkType.Region.ValueString(),
			Cidr:   networkType.Cidr.ValueString(),
			PublicSubnets: func() []string {
				if networkType.PublicSubnets.IsUnknown() {
					return nil
				}
				var publicSubnets []string
				diags = networkType.PublicSubnets.ElementsAs(ctx, &publicSubnets, false)
				resp.Diagnostics.Append(diags...)
				return publicSubnets
			}(),
			PrivateSubnets: func() []string {
				if networkType.PrivateSubnets.IsUnknown() {
					return nil
				}
				var privateSubnets []string
				diags = networkType.PrivateSubnets.ElementsAs(ctx, &privateSubnets, false)
				resp.Diagnostics.Append(diags...)
				return privateSubnets
			}(),
			External:   networkType.External.ValueBool(),
			ExternalID: networkType.ExternalId.ValueString(),
		})
	}
	return resp, networks
}

func nodesClusterReq(ctx context.Context, resp *resource.CreateResponse, nodesReq []basetypes.ObjectValue) (*resource.CreateResponse, []*models.ClusterNode) {
	var nodes []*models.ClusterNode
	for _, node := range nodesReq {
		var nodeType ClusterNode
		diags := node.As(ctx, &nodeType, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		nodes = append(nodes, &models.ClusterNode{
			Name:   nodeType.Name.ValueString(),
			Region: nodeType.Region.ValueString(),
			AvailabilityZone: func() string {
				if nodeType.AvailabilityZone.IsUnknown() {
					return ""
				}
				return nodeType.AvailabilityZone.ValueString()
			}(),
			Options: func() []string {
				if nodeType.Options.IsUnknown() {
					return nil
				}
				var options []string
				diags = nodeType.Options.ElementsAs(ctx, &options, false)
				resp.Diagnostics.Append(diags...)
				return options
			}(),
			InstanceType: nodeType.InstanceType.ValueString(),
			VolumeSize:   nodeType.VolumeSize.ValueInt64(),
			VolumeIops:   nodeType.VolumeIOPS.ValueInt64(),
			VolumeType:   nodeType.VolumeType.ValueString(),
		})
	}
	return resp, nodes
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClusterDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterCreationRequest := &models.ClusterCreationRequest{
		Name:           plan.Name.ValueString(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		NodeLocation:   plan.NodeLocation.ValueString(),
		SSHKeyID:       plan.SSHKeyID.ValueString(),
	}

	if plan.Firewall != nil {
		resp, firewallRules := fireWallRulesClusterReq(ctx, resp, plan.Firewall)
		if resp.Diagnostics.HasError() {
			return
		}
		clusterCreationRequest.FirewallRules = firewallRules
	}

	if plan.Networks != nil {
		resp, networks := networksAccountClusterReq(ctx, resp, plan.Networks)
		if resp.Diagnostics.HasError() {
			return
		}
		clusterCreationRequest.Networks = networks
	}

	if plan.Nodes != nil {
		resp, nodes := nodesClusterReq(ctx, resp, plan.Nodes)
		if resp.Diagnostics.HasError() {
			return
		}
		clusterCreationRequest.Nodes = nodes
	}

	if !plan.Regions.IsUnknown() {
		var regions []string
		diags = plan.Regions.ElementsAs(ctx, &regions, false)
		resp.Diagnostics.Append(diags...)
		clusterCreationRequest.Regions = regions
	}

	createdCluster, err := r.client.CreateCluster(ctx, clusterCreationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error creating pgEdge Cluster", "Could not create Cluster, unexpected error: "+err.Error())
		return
	}

	plan.ID = types.StringValue(createdCluster.ID.String())
	plan.Name = types.StringValue(createdCluster.Name)
	plan.CloudAccountID = types.StringValue(createdCluster.CloudAccount.ID)
	plan.CreatedAt = types.StringValue(createdCluster.CreatedAt.String())
	plan.Status = types.StringValue(createdCluster.Status)
	plan.NodeLocation = types.StringValue(createdCluster.NodeLocation)
	plan.SSHKeyID = types.StringValue(createdCluster.SSHKeyID)
	plan.Regions = func() types.List {
		var regions []attr.Value
		for _, region := range createdCluster.Regions {
			regions = append(regions, types.StringValue(region))
		}
		regionsList, _ := types.ListValue(types.StringType, regions)
		return regionsList
	}()

	var firewallResp []types.Object
	for _, firewall := range createdCluster.FirewallRules {
		var firewallSources []attr.Value
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}
		firewallSourceList, diags := types.ListValue(types.StringType, firewallSources)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallElements := map[string]attr.Value{
			"port":    firewallPort,
			"sources": firewallSourceList,
		}
		firewallObjectValue, diags := types.ObjectValue(FireWallType, firewallElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallResp = append(firewallResp, firewallObjectValue)
	}

	if len(createdCluster.FirewallRules) > 0 {
		plan.Firewall = firewallResp
	}

	var networkResp []types.Object
	for _, network := range createdCluster.Networks {
		networkElements := map[string]attr.Value{
			"region": types.StringValue(network.Region),
			"cidr":   types.StringValue(network.Cidr),
			"public_subnets": func() types.List {
				var publicSubnets []attr.Value
				for _, publicSubnet := range network.PublicSubnets {
					publicSubnets = append(publicSubnets, types.StringValue(publicSubnet))
				}
				publicSubnetList, _ := types.ListValue(types.StringType, publicSubnets)
				if publicSubnetList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return publicSubnetList
			}(),
			"private_subnets": func() types.List {
				var privateSubnets []attr.Value
				for _, privateSubnet := range network.PrivateSubnets {
					privateSubnets = append(privateSubnets, types.StringValue(privateSubnet))
				}
				privateSubnetList, _ := types.ListValue(types.StringType, privateSubnets)
				if privateSubnetList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return privateSubnetList
			}(),
			"name":        types.StringValue(network.Name),
			"external":    types.BoolValue(network.External),
			"external_id": types.StringValue(network.ExternalID),
		}
		networkObjectValue, diags := types.ObjectValue(NetworksType, networkElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		networkResp = append(networkResp, networkObjectValue)
	}

	if len(createdCluster.Networks) > 0 {
		plan.Networks = networkResp
	}

	var nodeResp []types.Object
	for _, node := range createdCluster.Nodes {
		nodeElements := map[string]attr.Value{
			"name":              types.StringValue(node.Name),
			"region":            types.StringValue(node.Region),
			"availability_zone": types.StringValue(node.AvailabilityZone),
			"options": func() types.List {
				var options []attr.Value
				for _, option := range node.Options {
					options = append(options, types.StringValue(option))
				}
				optionsList, _ := types.ListValue(types.StringType, options)
				if optionsList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return optionsList
			}(),
			"volume_size":   types.Int64Value(node.VolumeSize),
			"volume_iops":   types.Int64Value(node.VolumeIops),
			"volume_type":   types.StringValue(node.VolumeType),
			"instance_type": types.StringValue(node.InstanceType),
		}
		nodeObjectValue, diags := types.ObjectValue(ClusterNodeTypes, nodeElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		nodeResp = append(nodeResp, nodeObjectValue)
	}

	if len(createdCluster.Nodes) > 0 {
		plan.Nodes = nodeResp
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ClusterDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cluster, err := r.client.GetCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading pgEdge Cluster",
			"Could not read Cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state.ID = types.StringValue(cluster.ID.String())
	state.Name = types.StringValue(cluster.Name)
	state.CloudAccountID = types.StringValue(cluster.CloudAccount.ID.String())
	state.CreatedAt = types.StringValue(cluster.CreatedAt.String())
	state.Status = types.StringValue(cluster.Status)
	state.NodeLocation = types.StringValue(cluster.NodeLocation)
	state.SSHKeyID = types.StringValue(cluster.SSHKeyID)
	state.Regions = func() types.List {
		var regions []attr.Value
		for _, region := range cluster.Regions {
			regions = append(regions, types.StringValue(region))
		}
		regionsList, _ := types.ListValue(types.StringType, regions)
		return regionsList
	}()

	var firewallResp []types.Object
	for _, firewall := range cluster.FirewallRules {
		var firewallSources []attr.Value
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}
		firewallSourcesList, diags := types.ListValue(types.StringType, firewallSources)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallElements := map[string]attr.Value{
			"port":    firewallPort,
			"sources": firewallSourcesList,
		}
		firewallObjectValue, diags := types.ObjectValue(FireWallType, firewallElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallResp = append(firewallResp, firewallObjectValue)
	}

	if len(cluster.FirewallRules) > 0 {
		state.Firewall = firewallResp
	}

	var networkResp []types.Object
	for _, network := range cluster.Networks {
		networkElements := map[string]attr.Value{
			"name":     types.StringValue(network.Name),
			"external": types.BoolValue(network.External),
			"cidr":     types.StringValue(network.Cidr),
			"region":   types.StringValue(network.Region),
			"public_subnets": func() types.List {
				var publicSubnets []attr.Value
				for _, publicSubnet := range network.PublicSubnets {
					publicSubnets = append(publicSubnets, types.StringValue(publicSubnet))
				}
				publicSubnetsList, _ := types.ListValue(types.StringType, publicSubnets)
				return publicSubnetsList
			}(),
			"private_subnets": func() types.List {
				var privateSubnets []attr.Value
				for _, privateSubnet := range network.PrivateSubnets {
					privateSubnets = append(privateSubnets, types.StringValue(privateSubnet))
				}
				privateSubnetsList, _ := types.ListValue(types.StringType, privateSubnets)
				return privateSubnetsList
			}(),
			"external_id": types.StringValue(network.ExternalID),
		}
		networkObjectValue, diags := types.ObjectValue(NetworksType, networkElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		networkResp = append(networkResp, networkObjectValue)
	}

	if len(cluster.Networks) > 0 {
		state.Networks = networkResp
	}

	var nodeResp []types.Object
	for _, node := range cluster.Nodes {
		nodeElements := map[string]attr.Value{
			"name":              types.StringValue(node.Name),
			"volume_type":       types.StringValue(node.VolumeType),
			"instance_type":     types.StringValue(node.InstanceType),
			"availability_zone": types.StringValue(node.AvailabilityZone),
			"volume_size":       types.Int64Value(node.VolumeSize),
			"volume_iops":       types.Int64Value(node.VolumeIops),
			"options": func() types.List {
				var options []attr.Value
				for _, option := range node.Options {
					options = append(options, types.StringValue(option))
				}
				optionsList, _ := types.ListValue(types.StringType, options)
				return optionsList
			}(),
			"region": types.StringValue(node.Region),
		}
		nodeObjectValue, diags := types.ObjectValue(ClusterNodeTypes, nodeElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		nodeResp = append(nodeResp, nodeObjectValue)
	}

	if len(cluster.Nodes) > 0 {
		state.Nodes = nodeResp
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ClusterDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ClusterDetails
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newResp := &resource.CreateResponse{}

	*newResp = resource.CreateResponse(*resp)

	newResp, firewallRules := fireWallRulesClusterReq(ctx, newResp, plan.Firewall)
	if newResp.Diagnostics.HasError() {
		return
	}

	newResp, networks := networksAccountClusterReq(ctx, newResp, plan.Networks)
	if newResp.Diagnostics.HasError() {
		return
	}

	newResp, nodes := nodesClusterReq(ctx, newResp, plan.Nodes)
	if newResp.Diagnostics.HasError() {
		return
	}

	*resp = resource.UpdateResponse(*newResp)

	clusterUpdateRequest := &models.ClusterUpdateRequest{
		FirewallRules: firewallRules,
		Networks:      networks,
		Nodes:         nodes,
		Regions: func() []string {
			var regions []string
			for _, region := range plan.Regions.Elements() {
				regions = append(regions, region.String())
			}
			return regions
		}(),
		SSHKeyID: plan.SSHKeyID.ValueString(),
	}

	updatedCluster, err := r.client.UpdateCluster(ctx, strfmt.UUID(state.ID.ValueString()), clusterUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error updating pgEdge Cluster", "Could not update Cluster, unexpected error: "+err.Error())
		return
	}

	plan.ID = types.StringValue(updatedCluster.ID.String())
	plan.Name = types.StringValue(updatedCluster.Name)
	plan.CloudAccountID = types.StringValue(updatedCluster.CloudAccount.ID.String())
	plan.CreatedAt = types.StringValue(updatedCluster.CreatedAt.String())
	plan.Status = types.StringValue(updatedCluster.Status)
	plan.NodeLocation = types.StringValue(updatedCluster.NodeLocation)
	plan.SSHKeyID = types.StringValue(updatedCluster.SSHKeyID)
	plan.Regions = func() types.List {
		var regions []attr.Value
		for _, region := range updatedCluster.Regions {
			regions = append(regions, types.StringValue(region))
		}
		regionsList, _ := types.ListValue(types.StringType, regions)

		return regionsList
	}()

	var firewallResp []types.Object
	for _, firewall := range updatedCluster.FirewallRules {
		var firewallSources []attr.Value
		firewallPort := types.Int64Value(firewall.Port)
		for _, source := range firewall.Sources {
			firewallSources = append(firewallSources, types.StringValue(source))
		}
		firewallSourceList, diags := types.ListValue(types.StringType, firewallSources)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallElements := map[string]attr.Value{
			"port":    firewallPort,
			"sources": firewallSourceList,
		}
		firewallObjectValue, diags := types.ObjectValue(FireWallType, firewallElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		firewallResp = append(firewallResp, firewallObjectValue)
	}

	if len(updatedCluster.FirewallRules) > 0 {
		plan.Firewall = firewallResp
	}

	var networkResp []types.Object
	for _, network := range updatedCluster.Networks {
		networkElements := map[string]attr.Value{
			"region": types.StringValue(network.Region),
			"cidr":   types.StringValue(network.Cidr),
			"public_subnets": func() types.List {
				var publicSubnets []attr.Value
				for _, publicSubnet := range network.PublicSubnets {
					publicSubnets = append(publicSubnets, types.StringValue(publicSubnet))
				}

				publicSubnetList, _ := types.ListValue(types.StringType, publicSubnets)
				return publicSubnetList
			}(),
			"private_subnets": func() types.List {
				var privateSubnets []attr.Value
				for _, privateSubnet := range network.PrivateSubnets {
					privateSubnets = append(privateSubnets, types.StringValue(privateSubnet))
				}

				privateSubnetList, _ := types.ListValue(types.StringType, privateSubnets)
				if privateSubnetList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return privateSubnetList
			}(),

			"name":        types.StringValue(network.Name),
			"external":    types.BoolValue(network.External),
			"external_id": types.StringValue(network.ExternalID),
		}
		networkObjectValue, diags := types.ObjectValue(NetworksType, networkElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		networkResp = append(networkResp, networkObjectValue)
	}

	if len(updatedCluster.Networks) > 0 {
		plan.Networks = networkResp
	}

	var nodeResp []types.Object
	for _, node := range updatedCluster.Nodes {
		nodeElements := map[string]attr.Value{
			"name":              types.StringValue(node.Name),
			"region":            types.StringValue(node.Region),
			"availability_zone": types.StringValue(node.AvailabilityZone),
			"options": func() types.List {
				var options []attr.Value
				for _, option := range node.Options {
					options = append(options, types.StringValue(option))
				}
				optionsList, _ := types.ListValue(types.StringType, options)
				if optionsList.IsNull() {
					return types.ListNull(types.StringType)
				}
				return optionsList
			}(),
			"volume_size":   types.Int64Value(node.VolumeSize),
			"volume_iops":   types.Int64Value(node.VolumeIops),
			"volume_type":   types.StringValue(node.VolumeType),
			"instance_type": types.StringValue(node.InstanceType),
		}
		nodeObjectValue, diags := types.ObjectValue(ClusterNodeTypes, nodeElements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		nodeResp = append(nodeResp, nodeObjectValue)
	}

	if len(updatedCluster.Nodes) > 0 {
		plan.Nodes = nodeResp
	}

	// plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ClusterDetails
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCluster(ctx, strfmt.UUID(state.ID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting pgEdge Database",
			"Could not delete Database, unexpected error: "+err.Error(),
		)
		return
	}
}
