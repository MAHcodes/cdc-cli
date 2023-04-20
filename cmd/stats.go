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

	blitz := lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Blitz"),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Best"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.ChessBlitz.Best.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessBlitz.Best.Date)),
				ui.SprintInfo("Game", s.ChessBlitz.Best.Game))))),

		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Last"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.ChessBlitz.Last.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessBlitz.Last.Date)),
				ui.SprintInfo("Rd", strconv.Itoa(s.ChessBlitz.Last.Rd)))))),

		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Record"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Win", strconv.Itoa(s.ChessBlitz.Record.Win)),
				ui.SprintInfo("Draw", strconv.Itoa(s.ChessBlitz.Record.Draw)),
				ui.SprintInfo("Loss", strconv.Itoa(s.ChessBlitz.Record.Loss)))))))

	rapid := lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Rapid"),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Best"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.ChessRapid.Best.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessRapid.Best.Date)),
				ui.SprintInfo("Game", s.ChessRapid.Best.Game))))),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Last"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.ChessRapid.Last.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessRapid.Last.Date)),
				ui.SprintInfo("Rd", strconv.Itoa(s.ChessRapid.Last.Rd)))))),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Record"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Win", strconv.Itoa(s.ChessRapid.Record.Win)),
				ui.SprintInfo("Draw", strconv.Itoa(s.ChessRapid.Record.Draw)),
				ui.SprintInfo("Loss", strconv.Itoa(s.ChessRapid.Record.Loss)))))))

	daily := lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Daily"),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Last"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.ChessDaily.Last.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.ChessDaily.Last.Date)),
				ui.SprintInfo("Rd", strconv.Itoa(s.ChessDaily.Last.Rd)))))),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Record"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Win", strconv.Itoa(s.ChessDaily.Record.Win)),
				ui.SprintInfo("Draw", strconv.Itoa(s.ChessDaily.Record.Draw)),
				ui.SprintInfo("Loss", strconv.Itoa(s.ChessDaily.Record.Loss)))))))

	tactics := lipgloss.JoinVertical(lipgloss.Left,
		ui.SprintTag("Tactics"),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Highest"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.Tactics.Highest.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.Tactics.Highest.Date)))))),
		ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.SprintTag("Lowest"),
			ui.BoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
				ui.SprintInfo("Rating", strconv.Itoa(s.Tactics.Lowest.Rating)),
				ui.SprintInfo("Date", utils.FormatUnixTime(s.Tactics.Lowest.Date)))))))

	all := lipgloss.JoinVertical(lipgloss.Left,
		blitz,
		rapid,
		daily,
		tactics,
	)

	fmt.Println(all)
}
