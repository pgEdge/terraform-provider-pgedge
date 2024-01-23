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
    name       = "test15"
    cloud_account_id = "5984a9ec-7786-4ad9-9739-bbdf386eafec"
}


output "tech_clusters" {
  value = data.pgedge_clusters.tech
}
