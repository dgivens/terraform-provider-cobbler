package main

import (
	"github.com/dgivens/terraform-provider-cobbler/cobbler"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cobbler.Provider})
}
