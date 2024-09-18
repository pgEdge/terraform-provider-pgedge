package common

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/pgEdge/terraform-provider-pgedge/internals/provider"
)

// TestAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"pgedge": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// ProviderConfig is a shared configuration to combine with the actual
// test configuration so the PgEdge client is properly configured.
// It is also possible to use the PGEDGE_ environment variables instead,
// leaving this constant empty.
const ProviderConfig = `
provider "pgedge" {}
`