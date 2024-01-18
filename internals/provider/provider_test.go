package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerConfig = `
	  provider "pgedge" {
		base_url = "https://devapi.pgedge.com"
	  }	  
`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"pgedge": providerserver.NewProtocol6WithError(New("test")()),
	  }
)
