resource "pgedge_ssh_key" "example" {
  name       = "example-ssh-key"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIExamplePublicKeyContentHere user@example.com"
}