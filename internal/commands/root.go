package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-weather-cli",
		Short: "A sample CLI tool to get weather information",
		Long:  "Receive weather for given the city and/or latitude and longitude from OpenWeatherMap API",
	}

	// Add subcommands
	rootCmd.AddCommand(getWeather())
	rootCmd.AddCommand(getFavourites())
	rootCmd.AddCommand(setFavourites())

	return rootCmd
}
