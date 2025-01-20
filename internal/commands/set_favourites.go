package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func setFavourites() *cobra.Command {
	return &cobra.Command{
		Use:   "set-favourites",
		Short: "Executes set-favourites",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("set-favourites executed!")
		},
	}
}
