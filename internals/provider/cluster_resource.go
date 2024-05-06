package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	pgEdge "github.com/pgEdge/terraform-provider-pgedge/client"
	"github.com/pgEdge/terraform-provider-pgedge/models"
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
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cluster",
			},
			"cloud_account_id": schema.StringAttribute{
				Computed:    true,
				Description: "Cloud account ID of the cluster",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "Created at of the cluster",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Status of the cluster",
			},
			"ssh_key_id": schema.StringAttribute{
				Computed:    true,
				Description: "SSH key ID of the cluster",
			},
			// "resource_tags": schema.MapAttribute{
			// 	ElementType: types.StringType,
			// 	Computed:    true,
			// },
			"regions": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"node_location": schema.StringAttribute{
				Computed:    true,
				Description: "Node location of the cluster",
			},
			"cloud_account": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Display name of the node",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "IP address of the node",
						},
						"Type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the node",
						},
					},
				},
			},
			"firewall_rules": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the firewall rule",
						},
						"port": schema.Int64Attribute{
							Computed:    true,
							Description: "Port for the firewall rule",
						},
						"sources": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Sources for the firewall rule",
						},
					},
				},
			},
			"nodes": ClusterNodeDataSourceType,
			"networks": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the network",
						},
						"region": schema.StringAttribute{
							Computed:    true,
							Description: "Region of the network",
						},
						"external": schema.BoolAttribute{
							Computed:    true,
							Description: "Is the network external",
						},
						"external_id": schema.StringAttribute{
							Computed:    true,
							Description: "External ID of the network",
						},
						"cidr": schema.StringAttribute{
							Computed:    true,
							Description: "CIDR of the AWS node group",
						},
						"public_subnets": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"private_subnets": schema.ListAttribute{
							ElementType: types.StringType,
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
			Name:    firewallRuleType.Name.ValueString(),
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
				var publicSubnets []string
				diags = networkType.PublicSubnets.ElementsAs(ctx, &publicSubnets, false)
				resp.Diagnostics.Append(diags...)
				return publicSubnets
			}(),
			PrivateSubnets: func() []string {
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
			Name: nodeType.Name.ValueString(),
		})
	}
	return resp, nodes
}

func cloudAccountClusterReq(ctx context.Context, resp *resource.CreateResponse, cloudAccountReq basetypes.ObjectValue) (*resource.CreateResponse, *models.ClusterCreationRequestCloudAccount) {
	var cloudAccountType CloudAccount
	diags := cloudAccountReq.As(ctx, &cloudAccountType, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	return resp, &models.ClusterCreationRequestCloudAccount{
		ID:   strfmt.UUID(cloudAccountType.ID.ValueString()),
		Name: cloudAccountType.Name.ValueString(),
		Type: cloudAccountType.Type.ValueString(),
	}
}
func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClusterDetails
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp, firewallRules := fireWallRulesClusterReq(ctx, resp, plan.Firewall)
	if resp.Diagnostics.HasError() {
		return
	}

	resp, cloudAccount := cloudAccountClusterReq(ctx, resp, plan.CloudAccount)
	if resp.Diagnostics.HasError() {
		return
	}

	resp, networks := networksAccountClusterReq(ctx, resp, plan.Networks)
	if resp.Diagnostics.HasError() {
		return
	}

	resp, nodes := nodesClusterReq(ctx, resp, plan.Nodes)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterCreationRequest := &models.ClusterCreationRequest{
		Name:           plan.Name.ValueString(),
		CloudAccountID: plan.CloudAccountID.ValueString(),
		FirewallRules:  firewallRules,
		CloudAccount:   cloudAccount,
		Networks:       networks,
		Nodes:          nodes,
		NodeLocation:   plan.NodeLocation.ValueString(),
		Regions: func() []string {
			var regions []string
			for _, region := range plan.Regions.Elements() {
				regions = append(regions, region.String())
			}
			return regions
		}(),
		// ResourceTags: ,
		SSHKeyID: plan.SSHKeyID.ValueString(),
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
	plan.CloudAccount = func() types.Object {
		cloudAccountElements := map[string]attr.Value{
			"id":   types.StringValue(createdCluster.CloudAccount.ID),
			"name": types.StringValue(createdCluster.CloudAccount.Name),
			"type": types.StringValue(createdCluster.CloudAccount.Type),
		}

		cloudAccountObjectValue, _ := types.ObjectValue(CloudAccountType, cloudAccountElements)
		resp.Diagnostics.Append(diags...)
		return cloudAccountObjectValue

	}()
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
		firewallName := types.StringValue(firewall.Name)
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
			"name":    firewallName,
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
				return publicSubnetList
			}(),
			"private_subnets": func() types.List {
				var privateSubnets []attr.Value
				for _, privateSubnet := range network.PrivateSubnets {
					privateSubnets = append(privateSubnets, types.StringValue(privateSubnet))
				}

				privateSubnetList, _ := types.ListValue(types.StringType, privateSubnets)
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
			"image":             types.StringValue(node.Image),
			"availability_zone": types.StringValue(node.AvailabilityZone),
			"options": func() types.List {
				var options []attr.Value
				for _, option := range node.Options {
					options = append(options, types.StringValue(option))
				}
				optionsList, _ := types.ListValue(types.StringType, options)
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
	state.CloudAccount = func() types.Object {
		cloudAccountElements := map[string]attr.Value{
			"id":   types.StringValue(cluster.CloudAccount.ID.String()),
			"name": types.StringValue(cluster.CloudAccount.Name),
			"type": types.StringValue(cluster.CloudAccount.Type),
		}

		cloudAccountObjectValue, _ := types.ObjectValue(CloudAccountType, cloudAccountElements)
		resp.Diagnostics.Append(diags...)
		return cloudAccountObjectValue

	}()
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
		firewallName := types.StringValue(firewall.Name)
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
			"name":    firewallName,
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
			"image":  types.StringValue(node.Image),
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
		// ResourceTags: ,
		SSHKeyID: plan.SSHKeyID.ValueString(),
	}

	updatedCluster, err := r.client.UpdateCluster(ctx, strfmt.UUID(plan.ID.ValueString()), clusterUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error updating pgEdge Cluster", "Could not update Cluster, unexpected error: "+err.Error())
		return
	}

	plan.ID = types.StringValue(updatedCluster.ID.String())
	plan.Name = types.StringValue(updatedCluster.Name)
	plan.CloudAccountID = types.StringValue(updatedCluster.CloudAccount.ID.String())
	plan.CreatedAt = types.StringValue(updatedCluster.CreatedAt.String())
	plan.Status = types.StringValue(updatedCluster.Status)
	plan.CloudAccount = func() types.Object {
		cloudAccountElements := map[string]attr.Value{
			"id":   types.StringValue(updatedCluster.CloudAccount.ID.String()),
			"name": types.StringValue(updatedCluster.CloudAccount.Name),
			"type": types.StringValue(updatedCluster.CloudAccount.Type),
		}

		cloudAccountObjectValue, _ := types.ObjectValue(CloudAccountType, cloudAccountElements)
		resp.Diagnostics.Append(diags...)
		return cloudAccountObjectValue

	}()
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
		firewallName := types.StringValue(firewall.Name)
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
			"name":    firewallName,
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
			"image":             types.StringValue(node.Image),
			"availability_zone": types.StringValue(node.AvailabilityZone),
			"options": func() types.List {
				var options []attr.Value
				for _, option := range node.Options {
					options = append(options, types.StringValue(option))
				}
				optionsList, _ := types.ListValue(types.StringType, options)
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

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

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
