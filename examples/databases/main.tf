terraform {
  required_providers {
    pgedge = {
      source = "pgedge.com/tech/pgedge"
    }
  }
}

provider "pgedge" {
  auth_header = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA1MjM5MTgwLCJleHAiOjE3MDUzMjU1ODAsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.E1fo83Qnez44Q4UzX3_WOrUBy5d5uSIDK5Jp22R4u-ymsBHW_iBOyI4cuks7BuWd8EkNDDdqx_OIFZfmSjM--PcYV93F_Nle50pJTqLo2mC3vKyLaOD2PNq7fXoEJvgp4rG__VS0bUSHQL-buGEIeViTh7rJiFn9kkdHvbqV5ysehx51dkN8AbmMTtVWU_ITskMa-KqqzOcIMtDByUSLsstBK8F1yEAOL9YNWpccHINxVNGS60GvP3ycBXpqIuR0pOE9MJAB14TuuBSO3UcdsO0jCOYNnK7Sl8aZFZFq_9CyP_E_SOdgc-ofaI_LvMqJGVV4iyF-s-kXe7i3HhEoxQ"
  cluster_id  = "5e7478e5-4e68-464b-902d-747db528eccc"
}

data "pgedge_databases" "tech" {
  # name       = "techie"
  # cluster_id = "5e7478e5-4e68-464b-902d-747db528eccc"
}

output "tech_databases" {
  value = data.pgedge_databases.tech
}
