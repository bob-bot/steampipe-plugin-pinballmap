package plugin

import (
	"context"

	"github.com/bob-bot/pinballmap-go/pinballmap"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define the table schema
func tablePinballMapMachine() *plugin.Table {
	return &plugin.Table{
		Name:        "pinballmap_machine",
		Description: "Available pinball machines.",
		List: &plugin.ListConfig{
			Hydrate: listPinballMapMachines,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: "int", Description: "The machine ID."},
			{Name: "name", Type: "string", Description: "The machine name."},
		},
	}
}

// Fetch machine data
func listPinballMapMachines(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client := pinballmap.NewClient("")
	machines, err := client.GetMachines()
	if err != nil {
		return nil, err
	}

	for _, machine := range machines {
		d.StreamListItem(ctx, machine)
	}

	return nil, nil
}