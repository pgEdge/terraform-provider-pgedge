terraform {
  required_providers {
    pgedge = {
      source = "pgEdge/pgedge"
    }
  }
  required_version = ">= 1.1.0"
}

provider "pgedge" {
}

data "pgedge_databases" "tech" {
}

resource "pgedge_database" "tech" {
  name       = "newDatabase2002"
  cluster_id = ""
  options    = ["install:northwind"]
}


output "tech_databases" {
  value = data.pgedge_databases.tech
}
