terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
}

provider "pgedge" {}

data "pgEdge_databases" "example" {}
