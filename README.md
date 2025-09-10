<img alt="pgEdge" src="https://pgedge-public-assets.s3.amazonaws.com/product/images/pgedge_mark.svg" width="100px">

# pgEdge Cloud Terraform Provider

The official Terraform provider for [pgEdge Cloud](https://www.pgedge.com/cloud), designed to simplify the management of pgEdge Cloud resources for both **Developer** and **Enterprise** edition.

- **Documentation:** [pgEdge Terraform Docs](https://registry.terraform.io/providers/pgEdge/pgedge/latest/docs)
- **Website:** [pgEdge](https://www.pgedge.com/)
- **Discuss:** [GitHub Issues](https://github.com/pgEdge/terraform-provider-pgedge/issues)

## Installation

To get started, declare the pgEdge provider in your Terraform configuration. Running `terraform init` will automatically download and install the provider from the [Terraform Registry](https://registry.terraform.io/providers/pgEdge/pgedge/latest):

```hcl
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}
```

## Environment Setup

Before using the provider, you must create an API Client in [pgEdge Cloud](https://app.pgedge.com) and configure the following environment variables:

```sh
export PGEDGE_CLIENT_ID="your-client-id"
export PGEDGE_CLIENT_SECRET="your-client-secret"
```

These credentials authenticate the Terraform provider with your pgEdge Cloud account.

## Usage

### Developer Edition Configuration

For Developer Edition, pgEdge offers access to manage databases. Here’s an example setup for Developer Edition::

```hcl
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}

provider "pgedge" {}

# Define a database
resource "pgedge_database" "defaultdb" {
  name       = "defaultdb"
  cluster_id = "f12239ddq-df9d-4ded-adqwead9-3e2bvhe6d6ee"

  options = [
    "rest:enabled",
    "install:northwind"
  ]
}
```

### Enterprise Edition Configuration

Enterprise Edition users can manage Cloud Accounts, SSH keys, Backup Stores, and Clusters. Here's an Enterprise Edition example that includes mechanisms to manage various aspects of these resources:

```hcl
terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
}

provider "pgedge" {
  base_url = "https://api.pgedge.com"
}

# SSH Key resource
resource "pgedge_ssh_key" "example" {
  name       = "example-key"
  public_key = "ssh-ed25519 AAAAC3NzaC1wes241mmT63i04t5fvvsdwqVG7DkyxvyXbYQNhKP/rSeLY user@example.com"
}

# Cloud Account resource
resource "pgedge_cloud_account" "example" {
  name        = "my-aws-account"
  type        = "aws"
  description = "My AWS Cloud Account"

  credentials = {
    role_arn = "arn:aws:iam::0123456789:role/pgedge-13fe3332c"
  }

  depends_on = [pgedge_ssh_key.example]
}

# Backup Store resource
resource "pgedge_backup_store" "test_store" {
  name             = "test-store"
  cloud_account_id = pgedge_cloud_account.example.id
  region           = "us-east-1"

  depends_on = [pgedge_cloud_account.example]
}

# Cluster resource with update support
resource "pgedge_cluster" "example" {
  name             = "example"
  cloud_account_id = pgedge_cloud_account.example.id
  regions          = ["us-west-2", "us-east-1", "eu-central-1"]
  ssh_key_id       = pgedge_ssh_key.example.id
  backup_store_ids = [pgedge_backup_store.test_store.id]
  node_location    = "public"

  nodes = [
    {
      name          = "n1"
      region        = "us-west-2"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    },
    {
      name          = "n2"
      region        = "us-east-1"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    },
    {
      name          = "n3"
      region        = "eu-central-1"
      instance_type = "r6g.medium"
      volume_size   = 100
      volume_type   = "gp2"
    }
  ]

  networks = [
    {
      region         = "us-west-2"
      cidr           = "10.1.0.0/16"
      public_subnets = ["10.1.0.0/24"]
    # private_subnets = ["10.1.1.0/24"]
    },
    {
      region         = "us-east-1"
      cidr           = "10.2.0.0/16"
      public_subnets = ["10.2.0.0/24"]
    # private_subnets = ["10.2.1.0/24"]
    },
    {
      region         = "eu-central-1"
      cidr           = "10.3.0.0/16"
      public_subnets = ["10.3.0.0/24"]
    # private_subnets = ["10.3.1.0/24"]
    }
  ]

  firewall_rules = [
    {
      name    = "postgres"
      port    = 5432
      sources = ["192.0.2.44/32"]
    },
  ]

  depends_on = [
    pgedge_cloud_account.example,
    pgedge_backup_store.test_store,
    pgedge_ssh_key.example
  ]
}
```

### Updating a Cluster

To update an existing cluster, such as adding or removing nodes, follow these steps:

- **Add or remove nodes**: Modify the `nodes` block.
- **Update regions and networks**: If you're adding or removing nodes, you must also update the corresponding `regions` and `networks` blocks for those nodes.
- **Manage backup stores**: Adjust the `backup_store_ids` array if necessary.

For example, when removing a node, ensure the corresponding region and network are also removed.

---

**Example: Removing a Node from the Cluster**

To remove a node, you must also update the corresponding `regions` and `networks`:

```hcl
nodes = [
  {
    name          = "n1"
    region        = "us-west-2"
    instance_type = "r6g.medium"
    volume_size   = 100
    volume_type   = "gp2"
  },
  {
    name          = "n3"
    region        = "eu-central-1"
    instance_type = "r6g.medium"
    volume_size   = 100
    volume_type   = "gp2"
  }
]

regions = ["us-west-2", "eu-central-1"]

networks = [
  {
    region         = "us-west-2"
    cidr           = "10.1.0.0/16"
    public_subnets = ["10.1.0.0/24"]
  },
  {
    region         = "eu-central-1"
    cidr           = "10.3.0.0/16"
    public_subnets = ["10.3.0.0/24"]
  }
]
```

This ensures that all associated regions, nodes, and networks are synchronized when making changes.

---

### Database Resource with Update Support

```hcl
resource "pgedge_database" "example_db" {
  name       = "exampledb"
  cluster_id = pgedge_cluster.example.id

  options = [
    "autoddl:enabled",
#   "install:northwind",
#   "rest:enabled",
#   "cloudwatch_metrics:enabled"
  ]

  extensions = {
    auto_manage = true
    requested = [
      "postgis",
      "vector"
    ]
  }

  nodes = {
    n1 = {
      name = "n1"
    },
    n2 = {
      name = "n2"
    },
    n3 = {
      name = "n3"
    }
  }

  backups = {
    provider = "pgbackrest"
    config = [
      {
        id        = "default"
        schedules = [
          {
            type            = "full"
            cron_expression = "0 6 * * ?"
            id              = "daily-full-backup"
          },
          {
            type            = "incr"
            cron_expression = "0 * * * ?"
            id              = "hourly-incr-backup"
          }
        ]
      }
    ]
  }

  depends_on = [pgedge_cluster.example]
}
```

### Updating a Database

- **Nodes**: Modify the `nodes` block to add or remove database nodes.
- **Options**: Update the `options` array to configure additional settings.
- **Extensions**: Modify the `extensions.requested` array to manage database extensions.

---

**Example Updates**

**Removing a node from the database:**

```hcl
nodes = {
  n1 = {
    name = "n1"
  },
  n3 = {
    name = "n3"
  }
}
```

**Adding a new backup store to a cluster:**

```hcl
backup_store_ids = [
  pgedge_backup_store.test_store.id,
  "new-backup-store-id"
]
```

For further details on provider configuration, refer to the [Terraform Provider Configuration guide](https://developer.hashicorp.com/terraform/language/providers/configuration).

---

## Contributing

We welcome and appreciate contributions from the community. Please review the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to get involved.

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.
