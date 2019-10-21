package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/terraform-providers/terraform-provider-dome9/dome9"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dome9.Provider,
	})
}
