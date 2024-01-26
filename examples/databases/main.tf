terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
  required_version = ">= 1.1.0"
}

provider "pgedge" {
  base_url = "https://devapi.pgedge.com"
}

data "pgedge_databases" "tech" {
}

resource "pgedge_database" "tech" {
    name       = "newDatabase1013"
    cluster_id = ""
    options    = ["install:northwind"]
}


output "tech_databases" {
  value = data.pgedge_databases.tech
}
