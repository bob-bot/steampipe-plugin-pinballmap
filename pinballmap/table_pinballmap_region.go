package plugin

import (
	"context"

	"github.com/bob-bot/pinballmap-go/pinballmap"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define the table schema
func tablePinballMapRegion() *plugin.Table {
	return &plugin.Table{
		Name:        "pinballmap_region",
		Description: "Available pinball regions.",
		List: &plugin.ListConfig{
			Hydrate: listPinballMapRegions,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: "int", Description: "The region ID."},
			{Name: "name", Type: "string", Description: "The region name."},
			{Name: "slug", Type: "string", Description: "Region slug."},
		},
	}
}

// Fetch region data
func listPinballMapRegions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client := pinballmap.NewClient("")
	regions, err := client.GetRegions()
	if err != nil {
		return nil, err
	}

	for _, region := range regions {
		d.StreamListItem(ctx, region)
	}

	return nil, nil
}