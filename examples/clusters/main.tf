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
    name       = "test1"
    cloud_account_id = ""
}


output "tech_clusters" {
  value = data.pgedge_clusters.tech
}
