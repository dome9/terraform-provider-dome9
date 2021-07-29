package main

import (
	"github.com/dome9/terraform-provider-dome9/dome9"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dome9.Provider,
	})
}
