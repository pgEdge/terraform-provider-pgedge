
resource "pgedge_database" "example" {
  name       = "exampledb"
  cluster_id = pgedge_cluster.main.id
}
