resource "pgedge_cluster" "example" {
  name             = ""
  cloud_account_id = ""
  firewall = [
    {
      type    = ""
      port    = 0
      sources = [""]
    }
  ]
  node_groups = {
    aws = [
      {
        region             = ""
        instance_type      = ""
        cidr               = ""
        node_location      = ""
        public_subnets     = [""]
        private_subnets    = [""]
        volume_size        = 0
        volume_type        = ""
        volume_iops        = 0
        availability_zones = [""]
        nodes = [
          {
            display_name = ""
            is_active    = true
            ip_address   = ""
          }
        ]
      },
    ],
    azure = [
      {
        region             = ""
        instance_type      = ""
        cidr               = ""
        node_location      = ""
        public_subnets     = [""]
        private_subnets    = [""]
        volume_size        = 0
        volume_type        = ""
        volume_iops        = 0
        availability_zones = [""]
        nodes = [
          {
            display_name = ""
            is_active    = true
            ip_address   = ""
          }
        ]
      },
    ],
    google = [
      {
        region             = ""
        instance_type      = ""
        cidr               = ""
        node_location      = ""
        public_subnets     = [""]
        private_subnets    = [""]
        volume_size        = 0
        volume_type        = ""
        volume_iops        = 0
        availability_zones = [""]
        nodes = [
          {
            display_name = ""
            is_active    = true
            ip_address   = ""
          }
        ]
      },
    ]
  }
}
