terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
  required_version = ">= 1.1.0"
}

provider "pgedge" {
  auth_header = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA1NDM1MjE2LCJleHAiOjE3MDU1MjE2MTYsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.EFJGCabsi16TvH1EGyPrCjEBF20nnABopczflulmWMDT8IjhdzZLQgYbHVcKIFvncAX2cnVcCrKVmG4xEJzntSE6GoEH_9dBwA9wlJq-7lCDThNCKVkKzrctdUK_Z9tgdxFKMo_AADmpg2PI3T1DI_c-3czGypEoPT5bdDBaG0PUWUXwdxN8UxWArKRA6K9nXYJSdddVYv7K32IcAyqas3d9PWsaRejnyIg5JRv118h64nG1wPJ4_5YEDydAnLfosDz6XZVAGoLIddNiac5Tgip-V-46P4AW8wHPUthkntdhpAkV1Q2Tc_a4G6Su1E76WuyM7Jwq0wjeVryNyOHz0w"
  cluster_id  = "5e7478e5-4e68-464b-902d-747db528eccc"
}

data "pgedge_databases" "tech" {
  # name       = "techie"
  # cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
}

resource "pgedge_databases" "tech" {
  databases = {
    name       = "newDatabase100",
    cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
  } //, options = ["install:northwind"]
}


output "tech_databases" {
  value = data.pgedge_databases.tech
}
