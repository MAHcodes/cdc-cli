package cmd

import (
	"fmt"

	"github.com/MAHcodes/cdc-cli/api"
	"github.com/MAHcodes/cdc-cli/api/types"
	"github.com/spf13/cobra"
)

func getStatsEndpoint(username string) (endpoint string) {
	return fmt.Sprintf("https://api.chess.com/pub/player/%s/stats", username)
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get ratings, win/loss, and other stats about a player's game play, tactics, lessons and Puzzle Rush score.",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run:   runStatsCmd,
}

func runStatsCmd(cmd *cobra.Command, args []string) {
	username := args[0]
	endpoint := getStatsEndpoint(username)

	var stats types.Stats
	s, err := api.FetchJSON(endpoint, &stats)
	if err != nil {
		fmt.Println(err)
		return
	}

  fmt.Printf("%#v", s)
}
