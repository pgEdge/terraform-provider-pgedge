resource "pgedge_databases" "example" {
  databases = {
    name = "", // (Required) The name of the database to create.
  }            //, options = ["install:northwind"]
}