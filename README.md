<img alt="pgEdge" src="https://pgedge-public-assets.s3.amazonaws.com/product/images/pgedge_mark.svg" width="100px">

# pgEdge Terraform Provider

The official Terraform provider for [pgEdge](https://www.pgedge.com/).

- **Documentation:** https://registry.terraform.io/providers/pgEdge/pgedge/latest/docs
- **Website**: https://www.pgedge.com/
- **Discuss**: https://github.com/pgEdge/terraform-provider-pgedge/issues

## Installation

Declare the provider in your configuration and `terraform init` will automatically fetch and install the provider for you from the [Terraform Registry](https://registry.terraform.io/providers/pgEdge/pgedge/latest):

```hcl
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
    required_version = ">= 1.1.0"
}
```

## Usage

[Create an API Client in pgEdge Cloud](https://dev.pgedge.com), and set the Client ID and Client Secret as environment variables:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
```

Then, you can use the provider in your configuration:

```hcl
provider "pgedge" {
  base_url = "https://api.pgedge.com" #(Optional)
}

data "pgedge_clusters" "cloud" {
}

data "pgedge_databases" "cloud" {
}

# Create a cluster
resource "pgedge_cluster" "cloud" {
  # ...
}

# Create a database
resource "pgedge_database" "cloud" {
  # ...
}
```

For more information on configuring providers in general, see the [Provider Configuration documentation](https://developer.hashicorp.com/terraform/language/providers/configuration).
