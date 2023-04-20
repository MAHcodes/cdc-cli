package cmd

import (
	"fmt"
	"strconv"

	"github.com/MAHcodes/cdc-cli/api"
	"github.com/MAHcodes/cdc-cli/api/types"
	"github.com/MAHcodes/cdc-cli/ui"
	"github.com/MAHcodes/cdc-cli/utils"
	"github.com/charmbracelet/lipgloss"
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

	blitzBest := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Best"),
		ui.SprintInfo("Rating", strconv.Itoa(s.ChessBlitz.Best.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessBlitz.Best.Date)),
		ui.SprintInfo("Game", s.ChessBlitz.Best.Game),
	))

	blitzLast := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Last"),
		ui.SprintInfo("Rating", strconv.Itoa(s.ChessBlitz.Last.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessBlitz.Last.Date)),
		ui.SprintInfo("Rd", strconv.Itoa(s.ChessBlitz.Last.Rd)),
	))

	blitzRecord := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Record"),
		ui.SprintInfo("Win", strconv.Itoa(s.ChessBlitz.Record.Win)),
		ui.SprintInfo("Draw", strconv.Itoa(s.ChessBlitz.Record.Draw)),
		ui.SprintInfo("Loss", strconv.Itoa(s.ChessBlitz.Record.Loss)),
	))

	blitzAll := lipgloss.JoinHorizontal(lipgloss.Left, blitzRecord, blitzLast, blitzBest)

	blitz := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Blitz"), blitzAll))

	rapidBest := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Best"),
		ui.SprintInfo("Rating", strconv.Itoa(s.ChessRapid.Best.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessRapid.Best.Date)),
		ui.SprintInfo("Game", s.ChessRapid.Best.Game),
	))

	rapidLast := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Last"),
		ui.SprintInfo("Rating", strconv.Itoa(s.ChessRapid.Last.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessRapid.Last.Date)),
		ui.SprintInfo("Rd", strconv.Itoa(s.ChessRapid.Last.Rd)),
	))

	rapidRecord := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Record"),
		ui.SprintInfo("Win", strconv.Itoa(s.ChessRapid.Record.Win)),
		ui.SprintInfo("Draw", strconv.Itoa(s.ChessRapid.Record.Draw)),
		ui.SprintInfo("Loss", strconv.Itoa(s.ChessRapid.Record.Loss)),
	))

	rapidAll := lipgloss.JoinHorizontal(lipgloss.Left, rapidRecord, rapidLast, rapidBest)

	rapid := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Rapid"), rapidAll))

	dailyLast := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Last"),
		ui.SprintInfo("Rating", strconv.Itoa(s.ChessDaily.Last.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessDaily.Last.Date)),
		ui.SprintInfo("Rd", strconv.Itoa(s.ChessDaily.Last.Rd)),
	))
	dailyRecord := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Record"),
		ui.SprintInfo("Win", strconv.Itoa(s.ChessDaily.Record.Win)),
		ui.SprintInfo("Draw", strconv.Itoa(s.ChessDaily.Record.Draw)),
		ui.SprintInfo("Loss", strconv.Itoa(s.ChessDaily.Record.Loss)),
	))

	dailyAll := lipgloss.JoinHorizontal(lipgloss.Left, dailyLast, dailyRecord)

	daily := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Daily"), dailyAll))

	tacticsHighest := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Highest"),
		ui.SprintInfo("Rating", strconv.Itoa(s.Tactics.Highest.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.Tactics.Highest.Date)),
	))
	tacticsLowest := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Lowest"),
		ui.SprintInfo("Rating", strconv.Itoa(s.Tactics.Lowest.Rating)),
		ui.SprintInfo("Date", utils.FormatUnixTime(s.Tactics.Lowest.Date)),
	))

	tacticsAll := lipgloss.JoinHorizontal(lipgloss.Left, tacticsHighest, tacticsLowest)

	tactics := ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Tactics"), tacticsAll))

  all := lipgloss.JoinVertical(lipgloss.Left, blitz, rapid, daily, tactics)

	fmt.Println(all)
}
