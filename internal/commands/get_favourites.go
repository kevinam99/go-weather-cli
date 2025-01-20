package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

type Location struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getFavourites() *cobra.Command {
	return &cobra.Command{
		Use:   "get-favourites",
		Short: "Displays favourite cities and their coordinates",
		Run: func(cmd *cobra.Command, args []string) {
			if err := displayFavourites(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		},
	}
}

func displayFavourites() error {
	// Read all favorites into memory
	favorites, err := readFavorites()
	if err != nil {
		return fmt.Errorf("failed to read favorites: %v", err)
	}

	if len(favorites) == 0 {
		return fmt.Errorf("no favorites found")
	}

	// Format locations for display
	displayStrings := make([]string, len(favorites))
	for i, loc := range favorites {
		displayStrings[i] = fmt.Sprintf("%s: %.4f° N, %.4f° E", loc.City, loc.Latitude, loc.Longitude)
	}

	// Display first 5 cities
	currentIndex := min(5, len(displayStrings))
	for i := 0; i < currentIndex; i++ {
		fmt.Println(displayStrings[i])
	}

	// Setup keyboard events
	if err := keyboard.Open(); err != nil {
		return fmt.Errorf("failed to open keyboard: %v", err)
	}
	defer keyboard.Close()

	fmt.Println("\nPress down arrow for more, 'q' to quit")

	// Handle keyboard input
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return fmt.Errorf("error reading keyboard: %v", err)
		}

		if char == 'q' {
			return nil
		}

		if key == keyboard.KeyArrowDown && currentIndex < len(displayStrings) {
			fmt.Println(displayStrings[currentIndex])
			currentIndex++
		}
	}
}

func readFavorites() ([]Location, error) {
	data, err := os.ReadFile("favourites.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Location{}, nil
		}
		return nil, fmt.Errorf("failed to open favourites file: %v", err)
	}

	var locations []Location
	if err := json.Unmarshal(data, &locations); err != nil {
		return nil, fmt.Errorf("error parsing favourites: %v", err)
	}

	return locations, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
