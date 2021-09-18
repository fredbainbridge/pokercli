package cmd

import (
	"fmt"
	"log"

	"github.com/fredbainbridge/pokermavensclient/tournament"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tournamentCmd = &cobra.Command{
	Use:   "tournament",
	Short: "Ring game operations",
	Long:  "This command performs various ring game operations on the Poker Mavens server software.",
	Run: func(cmd *cobra.Command, args []string) {
		x, err := tournament.Tables(viper.GetString("Url"), viper.GetString("Password"))
		if err != nil {
			log.Fatal(err)
		}
		for i := range x {
			fmt.Println(x[i].Name + " " + x[i].Game)
		}
	},
}

func init() {
	rootCmd.AddCommand(tournamentCmd)
}
