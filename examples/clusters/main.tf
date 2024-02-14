terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
  # required_version = ">= 1.1.0"
}

provider "pgedge" {
}

data "pgedge_clusters" "tech" {
}

resource "pgedge_cluster" "tech" {
  name             = "testing10712"
  cloud_account_id = ""
  firewall = [
    {
      type    = "postgres"
      port    = 5432
      sources = ["0.0.0.0/0"]
    }
  ]
   node_groups = {
    aws = [
      {
        region        = "us-west-2"
        instance_type = "t4g.small",
        # nodes = [],
        #  availability_zones = [
  #         "us-west-2a",
        # ]
      },
    ]
  }
  # node_groups = {
  #   aws = [
  #     {
  #       region        = "us-west-2"
  #       instance_type = "t4g.small"
  #       availability_zones = [
  #         "us-west-2a",
  #       ]
  #       nodes = [
  #         {
  #           display_name = "n1"
  #           is_active    = true
  #         }
  #       ]
  #     },
  #   ]
  # }
}


output "tech_clusters" {
  value = data.pgedge_clusters.tech
}
