package plugin

import (
	"context"

	"github.com/bob-bot/pinballmap-go/pinballmap"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define the table schema
func tablePinballMapLocation() *plugin.Table {
	return &plugin.Table{
		Name:        "pinballmap_location",
		Description: "Locations of pinball machines.",
		List: &plugin.ListConfig{
			Hydrate: listPinballMapLocations,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "city", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: "int", Description: "The location ID."},
			{Name: "name", Type: "string", Description: "The location name."},
			{Name: "city", Type: "string", Description: "The city where the location is."},
		},
	}
}

// Fetch location data
func listPinballMapLocations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client := pinballmap.NewClient("")
	city := d.KeyColumnQualString("city")

	locations, err := client.GetLocations(city)
	if err != nil {
		return nil, err
	}

	for _, location := range locations {
		d.StreamListItem(ctx, location)
	}

	return nil, nil
}