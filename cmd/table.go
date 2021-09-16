package cmd

import (
	"fmt"
	"log"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ringGameTableCmd = &cobra.Command{
	Use:   "table",
	Short: "Ring game tables",
	Long:  "This command performs various ring game operations on the Poker Mavens server software.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ring game table called.")
		x, err := ringgame.Tables(viper.GetString("Url"), viper.GetString("Password"))
		if err != nil {
			log.Fatal(err)
		}
		for i := range x {
			fmt.Println(x[i].Game)
		}
		//var t pokerapimodels.Table = pokerapimodels.Table{}
		//t.Game = "new game"
	},
}

func init() {
	rootCmd.AddCommand(ringGameTableCmd)
}
