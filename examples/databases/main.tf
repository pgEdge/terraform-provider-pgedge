terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
  required_version = ">= 1.1.0"
}

provider "pgedge" {
  base_url      = "https://devapi.pgedge.com"
  # client_id     = "CIzx5xcvt9MFRYVIoFl7Bz9Kl8ryNSdh"
  # client_secret = "XqRDtkdyyVKNjjT-NiDXdP-ovAJMEmTqKlbMD89WonZhRLyQocKA11rddxw85H8r"
}

data "pgedge_databases" "tech" {
}

resource "pgedge_database" "tech" {
  database = {
    name = "newDatabase101",
    cluster_id    = "5e7478e5-4e68-464b-902d-747db528eccc"
  } //, options = ["install:northwind"]
}


output "tech_databases" {
  value = data.pgedge_databases.tech
}
