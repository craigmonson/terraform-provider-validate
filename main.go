package main

import (
	"github.com/craigmonson/terraform-provider-validate/validate"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: validate.Provider,
	})
}
