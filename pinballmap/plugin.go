package plugin

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Create the plugin definition
func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-pinballmap",
		ConnectionConfigSchema: ConnectionConfig(),
		TableMap: map[string]*plugin.Table{
			"pinballmap_location": tablePinballMapLocation(),
			"pinballmap_machine":  tablePinballMapMachine(),
			"pinballmap_region":   tablePinballMapRegion(),
		},
	}
}
