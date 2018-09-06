package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-netlify/netlify"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netlify.Provider})
}
