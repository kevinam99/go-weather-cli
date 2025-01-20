package main

import (
	"go-weather-cli/internal/commands"
	"log"
)

func main() {
	rootCmd := commands.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
