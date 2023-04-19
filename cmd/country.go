package cmd

import (
	"fmt"
	"strings"

	"github.com/MAHcodes/cdc-cli/api"
	"github.com/MAHcodes/cdc-cli/api/types"
	"github.com/MAHcodes/cdc-cli/ui"
	"github.com/spf13/cobra"
)

func getCountryEndponit(code string) (endpoint string) {
	return fmt.Sprintf("https://api.chess.com/pub/country/%s", code)
}

func init() {
	rootCmd.AddCommand(countryCmd)
}

var countryCmd = &cobra.Command{
	Use:   "country",
	Short: "Get additional details about a country.",
	Long:  `All country-based URLs use the country's 2-character ISO 3166 code (capitalized) to specify which country you want data for.
https://api.chess.com/pub/country/{iso}
https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2

Chess.com supports players and clubs identifying with some regions that are not recognized countries in this ISO list. These countries are identified with codes within the "user-assigned code" ranges. When possible, we tried to use codes that are commonly used in other applications. This is a list of these codes at Chess.com:

    XA : "Canary Islands"
    XB : "Basque Country"
    XC : "Catalonia"
    XE : "England"
    XG : "Galicia"
    XK : "Kosovo"
    XP : "Palestine"
    XS : "Scotland"
    XW : "Wales"
    XX : "International"

The presence or absence of any region on this list does not reflect any political opinion by Chess.com. We're just here to play chess.`,
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		countryCode := strings.ToUpper(args[0])
    countryEndpoint := getCountryEndponit(countryCode)
		country, err := fetchCountry(countryEndpoint)
		if err != nil {
			fmt.Println(err)
		}
    ui.PrintInfo("ID", country.ID)
    ui.PrintInfo("Code", country.Code)
    ui.PrintInfo("Name", country.Name)
	},
}

func fetchCountry(countryEndpoint string) (c types.Country, err error) {
	country, err := api.FetchJSON(countryEndpoint, &c)
	if err != nil {
		fmt.Println(err)
		return c, err
	}
	return *country, nil
}
