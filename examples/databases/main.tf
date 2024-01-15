terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
}

provider "pgedge" {
  auth_header = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA1MzIyNzAxLCJleHAiOjE3MDU0MDkxMDEsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.l33c6lqgHPQ0vg7RH-UGnevTD1eDzq2EyhVNif5dWY0F0mWHqHnraBk-i0o1q_mcLKsg-fTPLuT1xr8jTkp95orVIe5j8Hbe0CSH5M1dbormSc5wU2mO-7qkMAUSouiSBL9s_1_MpOW6Qwi_4oEdQ5S92uGmT-0qVwauxKK1BJNbvRn1mk64QroeXwP4OeQ7om0T9FHELfNuR0KH4mMHVWgLRbbB41oTCSxc59O0ZpF14KN2gi0Pz8vdRiluGTitdlN-RBf2iOnoO8SRtH6u-Okp2o6dV8-GiHXYzPKYaB2ZCDnMlJyczZI3tbLVT8SHCrayeXwho8b32HeT6XwQEA"
  cluster_id  = "5e7478e5-4e68-464b-902d-747db528eccc"
}

data "pgedge_databases" "tech" {
  # name       = "techie"
  # cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
}

output "tech_databases" {
  value = data.pgedge_databases.tech
}
