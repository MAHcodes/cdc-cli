package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of chess-cli",
	Long:  `All software has versions. This is chess-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cdc-cli v0.1 -- HEAD")
	},
}
