package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/samjegal/terraform-provider-fincloud/fincloud"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fincloud.Provider})
}
