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
  # client_id     = "CIzx5xcvt9MFRYVIoFl7Bz9Kl8ryNSdh"
  # client_secret = "XqRDtkdyyVKNjjT-NiDXdP-ovAJMEmTqKlbMD89WonZhRLyQocKA11rddxw85H8r"
}

data "pgedge_clusters" "tech" {
}

# resource "pgedge_cluster" "tech" {
#   cluster = {
#     name       = "newDatabase101",
#     cluster_id = "a961031f-5865-43ce-bef1-24461ac16f1a"
#     options    = ["install:northwind"]
#   } //, options = ["install:northwind"]
# }


# output "tech_clusters" {
#   value = data.pgedge_clusters.tech
# }
