terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
  # required_version = ">= 1.1.0"
}

provider "pgedge" {
  base_url = "https://devapi.pgedge.com"
}

data "pgedge_clusters" "tech" {
}

resource "pgedge_cluster" "tech" {
  name             = "test124"
  cloud_account_id = ""
  firewall = [
    {
      type    = "https"
      port    = 5432
      sources = ["0.0.0.0/0"]
    }
  ]
  node_groups = {
    aws = [
      {
        region        = "us-east-1"
        instance_type = "t4g.small"
        availability_zones = [
          "us-east-1a",
        ]
        nodes = [
          {
            display_name = "Node1"
            is_active    = true
          }
        ]
      },
    ]
  }
}


output "tech_clusters" {
  value = data.pgedge_clusters.tech
}
