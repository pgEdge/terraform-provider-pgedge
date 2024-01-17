terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
  required_version = ">= 1.1.0"
}

provider "pgedge" {
  auth_header = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA1NDgyNTg3LCJleHAiOjE3MDU1Njg5ODcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.gCqx0w21DTaEOjoPkLR1hOlA2I6rSLs9JeQO6IgvizELL_S29IYAmXz8GbjlTT5PamW1-IFhUNtoaj9zg8qeQ_ib5qK5W1w5F5FwEu4jvX4VXFwS1p1DnqYQC2V2EmMdHdGu30W-UZ3KTD9PNkqNARuQt-rvlUjIeQx7shOLtgwC_KanQE4qWp-sXtfAo2k5VvJC7dsHtVHF__ivruht9MuUcoPof4srTbCKsLxkUwkts3KwG8zB3xPAEruqyGotXxJqxvbawmZrkSrr56SE9J5VrWv7D_TMstEB7pRMYtZb9LhGySawA5QjJdTytV-yTlBRrqzZObC-CglXBqkxvw"
  cluster_id  = "5e7478e5-4e68-464b-902d-747db528eccc"
}

data "pgedge_databases" "tech" {
  # name       = "techie"
  # cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
}

resource "pgedge_databases" "tech" {
  databases = {
    name       = "newDatabase101",
    cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
  } //, options = ["install:northwind"]
}


output "tech_databases" {
  value = data.pgedge_databases.tech
}
