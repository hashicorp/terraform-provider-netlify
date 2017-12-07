package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mitchellh/terraform-provider-netlify/netlify"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netlify.Provider})
}
