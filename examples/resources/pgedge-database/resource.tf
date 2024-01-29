resource "pgedge_database" "example" {
  database = {
    name       = "",  // (Required) The name of the database to create.
    cluster_id = ""   // (Required) The ID of the cluster to create the database in.
    options    = [""] // (Optional) The options to apply to the database.
  }

}
