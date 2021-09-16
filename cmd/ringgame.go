package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ringGameCmd = &cobra.Command{
	Use:   "ringgame",
	Short: "Ring game operations",
	Long:  "This command performs various ring game operations on the Poker Mavens server software.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ring game called.")
	},
}

func init() {
	rootCmd.AddCommand(ringGameCmd)
}
