resource "pgedge_database" "example" {
  name       = "example-db"
  cluster_id = "asdqw-fdrt334-w123-1gsde"

  options = [
    "autoddl:enabled",
    #   "install:northwind",
    #   "rest:enabled",
    #   "cloudwatch_metrics:enabled"
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
        id = "default"
        schedules = [
          {
            type            = "full"
            cron_expression = "0 6 * * ?"
            id              = "daily-full-backup"
          },
          {
            type            = "incr"
            cron_expression = "0 * * * ?"
            id              = "hourly-incr-backup"
          }
        ]
      }
    ]
  }
}