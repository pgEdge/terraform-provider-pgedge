package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
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
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"pgedge": providerserver.NewProtocol6WithError(New("test")()),
	  }
)
