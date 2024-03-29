---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pgedge_clusters Data Source - terraform-provider-pgedge"
subcategory: ""
description: |-
  Interface with the pgEdge service API for clusters.
---

# pgedge_clusters (Data Source)

Interface with the pgEdge service API for clusters.

## Example Usage

```hcl
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}

data "pgedge_clusters" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `clusters` (Attributes List) (see [below for nested schema](#nestedatt--clusters))

<a id="nestedatt--clusters"></a>
### Nested Schema for `clusters`

Required:

- `name` (String) Name of the cluster

Optional:

- `node_groups` (Attributes) (see [below for nested schema](#nestedatt--clusters--node_groups))

Read-Only:

- `cloud_account_id` (String) Cloud account ID of the cluster
- `created_at` (String) Created at of the cluster
- `firewall` (Attributes List) (see [below for nested schema](#nestedatt--clusters--firewall))
- `id` (String) ID of the cluster
- `status` (String) Status of the cluster

<a id="nestedatt--clusters--node_groups"></a>
### Nested Schema for `clusters.node_groups`

Optional:

- `aws` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--aws))
- `azure` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--azure))
- `google` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--google))

<a id="nestedatt--clusters--node_groups--aws"></a>
### Nested Schema for `clusters.node_groups.aws`

Optional:

- `nodes` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--aws--nodes))

Read-Only:

- `availability_zones` (List of String) Availability zones of the AWS node group
- `cidr` (String) CIDR of the AWS node group
- `instance_type` (String) Instance type of the AWS node group
- `node_location` (String) Node location of the AWS node group
- `private_subnets` (List of String)
- `public_subnets` (List of String)
- `region` (String) Region of the AWS node group
- `volume_iops` (Number) Volume IOPS of the AWS node group
- `volume_size` (Number) Volume size of the AWS node group
- `volume_type` (String) Volume type of the AWS node group

<a id="nestedatt--clusters--node_groups--aws--nodes"></a>
### Nested Schema for `clusters.node_groups.aws.volume_type`

Read-Only:

- `display_name` (String) Display name of the node
- `ip_address` (String) IP address of the node
- `is_active` (Boolean) Is the node active



<a id="nestedatt--clusters--node_groups--azure"></a>
### Nested Schema for `clusters.node_groups.azure`

Optional:

- `nodes` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--azure--nodes))

Read-Only:

- `availability_zones` (List of String) Availability zones of the AWS node group
- `cidr` (String) CIDR of the AWS node group
- `instance_type` (String) Instance type of the AWS node group
- `node_location` (String) Node location of the AWS node group
- `private_subnets` (List of String)
- `public_subnets` (List of String)
- `region` (String) Region of the AWS node group
- `volume_iops` (Number) Volume IOPS of the AWS node group
- `volume_size` (Number) Volume size of the AWS node group
- `volume_type` (String) Volume type of the AWS node group

<a id="nestedatt--clusters--node_groups--azure--nodes"></a>
### Nested Schema for `clusters.node_groups.azure.volume_type`

Read-Only:

- `display_name` (String) Display name of the node
- `ip_address` (String) IP address of the node
- `is_active` (Boolean) Is the node active



<a id="nestedatt--clusters--node_groups--google"></a>
### Nested Schema for `clusters.node_groups.google`

Optional:

- `nodes` (Attributes List) (see [below for nested schema](#nestedatt--clusters--node_groups--google--nodes))

Read-Only:

- `availability_zones` (List of String) Availability zones of the AWS node group
- `cidr` (String) CIDR of the AWS node group
- `instance_type` (String) Instance type of the AWS node group
- `node_location` (String) Node location of the AWS node group
- `private_subnets` (List of String)
- `public_subnets` (List of String)
- `region` (String) Region of the AWS node group
- `volume_iops` (Number) Volume IOPS of the AWS node group
- `volume_size` (Number) Volume size of the AWS node group
- `volume_type` (String) Volume type of the AWS node group

<a id="nestedatt--clusters--node_groups--google--nodes"></a>
### Nested Schema for `clusters.node_groups.google.volume_type`

Read-Only:

- `display_name` (String) Display name of the node
- `ip_address` (String) IP address of the node
- `is_active` (Boolean) Is the node active




<a id="nestedatt--clusters--firewall"></a>
### Nested Schema for `clusters.firewall`

Read-Only:

- `port` (Number) Port for the firewall rule
- `sources` (List of String) Sources for the firewall rule
- `type` (String) Type of the firewall rule
