package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Inventory client is properly configured.
	providerConfig = `
	  provider "pgedge" {
		base_url = "https://devapi.pgedge.com"
		cluster_id  = "5e7478e5-4e68-464b-902d-747db528eccc"
		client_id = "CIzx5xcvt9MFRYVIoFl7Bz9Kl8ryNSdh"
		client_secret = "XqRDtkdyyVKNjjT-NiDXdP-ovAJMEmTqKlbMD89WonZhRLyQocKA11rddxw85H8r"
	  }	  
`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"pgedge": providerserver.NewProtocol6WithError(New("test")()),
	  }
)
