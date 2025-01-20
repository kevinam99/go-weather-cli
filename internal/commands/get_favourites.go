package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

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

	// Display first 5 cities
	currentIndex := min(5, len(favorites))
	for i := 0; i < currentIndex; i++ {
		fmt.Println(favorites[i])
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

		if key == keyboard.KeyArrowDown && currentIndex < len(favorites) {
			fmt.Println(favorites[currentIndex])
			currentIndex++
		}
	}
}

func readFavorites() ([]string, error) {
	file, err := os.Open("favorites.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open favorites file: %v", err)
	}
	defer file.Close()

	var favorites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		favorites = append(favorites, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return favorites, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}