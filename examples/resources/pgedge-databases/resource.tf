resource "pgedge_database" "example" {
  database = {
    name = "", // (Required) The name of the database to create.
  }            //, options = ["install:northwind"]
}