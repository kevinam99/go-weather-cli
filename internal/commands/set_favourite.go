package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func setFavourite() *cobra.Command {
	return &cobra.Command{
		Use:   "set-favourite [city] [latitude] [longitude]",
		Short: "Add a city to favourite with its coordinates",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			city := args[0]
			var lat, lon float64
			fmt.Sscanf(args[1], "%f", &lat)
			fmt.Sscanf(args[2], "%f", &lon)

			if err := saveFavouriteCity(city, lat, lon); err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Added %s to favourites!\n", city)
		},
	}
}

func saveFavouriteCity(city string, lat, lon float64) error {
	locations, err := readFavorites()
	if err != nil {
		return fmt.Errorf("error reading favourites: %v", err)
	}

	// Check if city already exists
	for _, loc := range locations {
		if loc.City == city {
			return fmt.Errorf("city %s already exists in favourites", city)
		}
	}

	// Add new location
	locations = append(locations, Location{
		City:      city,
		Latitude:  lat,
		Longitude: lon,
	})

	// Save back to file
	data, err := json.MarshalIndent(locations, "", "    ")
	if err != nil {
		return fmt.Errorf("error encoding favourites: %v", err)
	}

	if err := os.WriteFile("favourites.json", data, 0644); err != nil {
		return fmt.Errorf("error saving favourites: %v", err)
	}

	return nil
}
