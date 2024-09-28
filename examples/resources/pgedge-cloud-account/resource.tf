resource "pgedge_cloud_account" "example" {
  name        = "example-aws-account"
  type        = "aws"
  description = "Example AWS Cloud Account"

  credentials = {
    role_arn = "arn:aws:iam::123456789012:role/pgedge-example-role"
  }
}