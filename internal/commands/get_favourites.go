package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getFavourites() *cobra.Command {
	return &cobra.Command{
		Use:   "get-favourites",
		Short: "Executes get-favourites",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			fmt.Println("get-favourites executed!")
		},
	}
}
