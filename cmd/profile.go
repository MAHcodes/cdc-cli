package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/MAHcodes/cdc-cli/api"
	"github.com/MAHcodes/cdc-cli/ui"
	"github.com/spf13/cobra"
)

func getProfileEndpoint(username string) (endpoint string) {
	return fmt.Sprintf("https://api.chess.com/pub/player/%s", username)
}


type Profile struct {
	ID         string `json:"@id"`
	Country    string `json:"country"`
	Followers  int    `json:"followers"`
	IsStreamer bool   `json:"is_streamer"`
	Joined     int    `json:"joined"`
	LastOnline int    `json:"last_online"`
	League     string `json:"league"`
	PlayerID   int    `json:"player_id"`
	Status     string `json:"status"`
	URL        string `json:"url"`
	Username   string `json:"username"`
	Verified   bool   `json:"verified"`
}

func init() {
	rootCmd.AddCommand(usernameCmd)
}

var usernameCmd = &cobra.Command{
	Use:   "profile",
	Short: "Player Profile",
	Long:  `Get additional details about a player in a game.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run:   runProfileCmd,
}

func runProfileCmd(cmd *cobra.Command, args []string) {
	username := args[0]
	endpoint := getProfileEndpoint(username)
	profileJson, err := api.FetchEndpoint(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	var p Profile
	if err := json.Unmarshal(profileJson, &p); err != nil {
		fmt.Printf("failed to unmarshal JSON: %v\n", err)
		return
	}

  lastOnline := time.Unix(int64(p.LastOnline), 0)

  country, err := fetchCountry(p.Country)
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.PrintInfo("Username", p.Username)
	ui.PrintInfo("Followers", strconv.Itoa(p.Followers))
	ui.PrintInfo("League", p.League)
	ui.PrintInfo("Country", country.Name)
	ui.PrintInfo("LastOnline", lastOnline.String())
	ui.PrintInfo("Status", p.Status)
	ui.PrintInfo("URL", p.URL)
}
