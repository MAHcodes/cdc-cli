package cmd

import (
	"fmt"
	"strconv"

	"github.com/MAHcodes/cdc-cli/api"
	"github.com/MAHcodes/cdc-cli/api/types"
	"github.com/MAHcodes/cdc-cli/ui"
	"github.com/MAHcodes/cdc-cli/utils"
	"github.com/spf13/cobra"
)

func getProfileEndpoint(username string) (endpoint string) {
	return fmt.Sprintf("https://api.chess.com/pub/player/%s", username)
}

func init() {
	rootCmd.AddCommand(usernameCmd)
}

var usernameCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get additional details about a player in a game.",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run:   runProfileCmd,
}

func runProfileCmd(cmd *cobra.Command, args []string) {
	username := args[0]
	endpoint := getProfileEndpoint(username)

	var profile types.Profile
	p, err := api.FetchJSON(endpoint, &profile)
	if err != nil {
		fmt.Println(err)
		return
	}

	country, err := fetchCountry(p.Country)
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.PrintInfo("Username", p.Username)
	ui.PrintInfo("Player ID", strconv.Itoa(p.PlayerID))
	ui.PrintInfo("Followers", strconv.Itoa(p.Followers))
	ui.PrintInfo("League", p.League)
	ui.PrintInfo("Country", country.Name)
	ui.PrintInfo("Status", p.Status)
	ui.PrintInfo("Joined", utils.FormatUnixTime(p.Joined))
	ui.PrintInfo("Last Online", utils.FormatUnixTime(p.LastOnline))
	ui.PrintInfo("URL", p.URL)

	if p.Verified {
		ui.PrintTag("verified")
	}
	if p.IsStreamer {
		ui.PrintTag("streamer")
	}
	fmt.Println()
}
