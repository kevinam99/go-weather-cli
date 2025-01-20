package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func getWeather() *cobra.Command {
	return &cobra.Command{
		Use:   "get-weather",
		Short: "Executes get-weather",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get-weather executed!")
		},
	}
}
