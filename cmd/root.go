package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var username string

var rootCmd = &cobra.Command{
	Use:   "cdc-cli",
	Short: "Unofficial Chess.com CLI",
	Long: `cdc-cli is an unofficial chess.com command-line tool written in Go,
allowing you to easily access the chess.com directly from your terminal.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// player username flag
	// rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "player username on chess.com")
	// viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	// rootCmd.MarkPersistentFlagRequired("username")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cdc-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cdc-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
