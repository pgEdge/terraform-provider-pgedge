resource "pgedge_database" "example" {
  name       = "example-db"
  cluster_id = "asdqw-fdrt334-w123-1gsde"

  options = [
    "install:northwind",
    "rest:enabled",
    "autoddl:enabled"
  ]

  extensions = {
    auto_manage = true
    requested   = ["postgis"]
  }

  nodes = {
    n1 = {
      name = "n1"
    },
    n2 = {
      name = "n2"
    }
  }

  backups = {
    provider = "pgbackrest"
    config = [
      {
        id        = "default"
        node_name = "n1"
        schedules = [
          {
            type            = "full"
            cron_expression = "0 1 * * *"
            id              = "daily-full-backup"
          }
        ]
      }
    ]
  }
}