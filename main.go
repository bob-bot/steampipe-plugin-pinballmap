package main

import (
	"github.com/bob-bot/steampipe-plugin-pinballmap/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: plugin.Plugin,
	})
}