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

resource "pgedge_cluster" "tech" {
    name       = "test15"
    cloud_account_id = "5984a9ec-7786-4ad9-9739-bbdf386eafec"
     aws = {
     key_pair = "asdasd"
     role_arn = "asdasd"
     tags     = {"test":"test"}
       }
}


output "tech_clusters" {
  value = data.pgedge_clusters.tech
}
