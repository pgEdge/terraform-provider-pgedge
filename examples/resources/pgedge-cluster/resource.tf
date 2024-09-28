resource "pgedge_cluster" "example" {
  name             = "example-cluster"
  cloud_account_id = "b8959307-asdxqwe-4f6c-b29e-fjtyrvv554"
  regions          = ["us-west-2", "us-east-1"]
  node_location    = "public"

  ssh_key_id = "b2ffbbd5-qweasw11-43f8-ae7f-6e47fcc71044"

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
    }
  ]

  networks = [
    {
      region         = "us-west-2"
      cidr           = "10.1.0.0/16"
      public_subnets = ["10.1.1.0/24"]
      # private_subnets = ["10.1.2.0/24"]
    },
    {
      region         = "us-east-1"
      cidr           = "10.2.0.0/16"
      public_subnets = ["10.2.1.0/24"]
      # private_subnets = ["10.2.2.0/24"]
    }
  ]

  backup_store_ids = [
    "b8959307-dtqwd1-4f6c-b29e-f753dbc39e4e",
    "b8959307-dfgw2-4f6c-b29e-f753dbc39e4e"
  ]

  firewall_rules = [
    {
      name    = "postgres"
      port    = 5432
      sources = ["103.213.321.452/32"]
    }
  ]
}