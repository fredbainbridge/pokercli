package cmd

import (
	"log"

	"github.com/fredbainbridge/pokermavensclient/ringgame"
	"github.com/fredbainbridge/pokermavensclient/tournament"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var wait bool
var all bool

const WAITCMD string = "wait"
const WAITDESC string = "Wait to take offline until current hand is finished."
const ALLCMD string = "all"
const ALLDESC string = "Set all games to offline."

var offlineTournamentCmd = &cobra.Command{
	Use:   "offline",
	Short: "Set a tournament game offline.",
	Long:  "This command will set one or moe tournaments offline.",
	Run: func(cmd *cobra.Command, args []string) {
		var tableNames []string
		if all {
			if len(args) != 0 {
				log.Fatal("Cannot specify \"All\" and game name(s)")
			}
			tables, err := tournament.Tables(viper.GetString("Url"), viper.GetString("Password"))
			for i := range tables {
				tableNames = append(tableNames, tables[i].Name)
			}
			if err != nil {
				log.Fatal(err)
			}
		} else {
			for i := range args {
				tableNames = append(tableNames, args[i])
			}
		}
		for i := range tableNames {
			err := tournament.Offline(tableNames[i], getWait(wait), viper.GetString("Url"), viper.GetString("Password"))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

var offlineRinggameCmd = &cobra.Command{
	Use:   "offline",
	Short: "Set a ring game offline.",
	Long:  "This command will set one or more ring games offline.",
	Run: func(cmd *cobra.Command, args []string) {
		var tableNames []string
		if all {
			if len(args) != 0 {
				log.Fatal("Cannot specify \"All\" and game name(s)")
			}
			tables, err := ringgame.Tables(viper.GetString("Url"), viper.GetString("Password"))
			for i := range tables {
				tableNames = append(tableNames, tables[i].Name)
			}
			if err != nil {
				log.Fatal(err)
			}
		} else {
			for i := range args {
				tableNames = append(tableNames, args[i])
			}
		}
		for i := range tableNames {
			err := ringgame.Offline(tableNames[i], getWait(wait), viper.GetString("Url"), viper.GetString("Password"))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func getWait(waitFlag bool) string {
	if wait {
		return "No"
	}
	return "Yes"
}

func setFlags(cmds ...*cobra.Command) {
	for i := range cmds {
		cmds[i].Flags().BoolVarP(&wait, WAITCMD, "w", false, WAITDESC)
		cmds[i].Flags().BoolVarP(&all, ALLCMD, "a", false, ALLDESC)
	}
}
func init() {
	setFlags(offlineRinggameCmd, offlineTournamentCmd)
	tournamentCmd.AddCommand(offlineTournamentCmd)
	ringGameCmd.AddCommand((offlineRinggameCmd))
}
