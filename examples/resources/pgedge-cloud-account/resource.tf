resource "pgedge_cloud_account" "aws_account" {
  name        = "my-aws-account"
  type        = "aws"
  description = "My AWS Cloud Account"

  credentials = {
    role_arn = "arn:aws:iam::212312312439:role/rolearn"
  }
}