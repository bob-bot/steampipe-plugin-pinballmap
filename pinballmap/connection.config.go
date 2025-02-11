package plugin

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// ConnectionConfig defines the connection configuration for the plugin.
func ConnectionConfig() *plugin.ConnectionConfigSchema {
	return &plugin.ConnectionConfigSchema{
		NewInstance: func() (interface{}, error) { return nil, nil },
	}
}
