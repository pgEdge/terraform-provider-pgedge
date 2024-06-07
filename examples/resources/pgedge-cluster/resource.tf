
resource "pgedge_cluster" "main" {
  name             = "example"
  cloud_account_id = "b8959307-be7e-4f6c-b29e-f753dbc39e4e"
  ssh_key_id       = "b2ffbbd5-91b2-43f8-ae7f-6e47fcc71044"
  regions          = ["us-west-1", "us-west-2"]
  networks = [
    {
      region         = "us-west-1"
      cidr           = "10.1.0.0/16"
      public_subnets = ["10.1.1.0/24"]
    },
    {
      region         = "us-west-2"
      cidr           = "10.2.0.0/16"
      public_subnets = ["10.2.1.0/24"]
    }
  ]
  firewall_rules = [
    {
      port    = 5432
      sources = ["0.0.0.0/0"]
    },
    {
      port    = 22
      sources = ["0.0.0.0/0"]
    }
  ]
  nodes = [
    {
      name              = "n1"
      region            = "us-west-1"
      availability_zone = "us-west-1a"
      instance_type     = "t4g.medium"
      volume_size       = 20
      volume_type       = "gp2"
    },
    {
      name              = "n2"
      region            = "us-west-2"
      availability_zone = "us-west-2a"
      instance_type     = "t4g.medium"
      volume_size       = 20
      volume_type       = "gp2"
    }
  ]
}
