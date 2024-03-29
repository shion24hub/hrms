/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var symbol, begin, end, output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hrms",
	Short: "'hrms' is a CLI tool for downloading historical trading data and generating candlestick data for the Bybit exchange.",
	Long:  `'hrms' is a CLI tool for downloading historical trading data and generating candlestick data for the Bybit exchange.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "", "symbol to download (e.g. BTCUSDT)")
	rootCmd.PersistentFlags().StringVarP(&begin, "begin", "b", "", "begin date to download (format: YYYYMMDD)")
	rootCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "end date to download (format: YYYYMMDD)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "./", "output directory (default is ./)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
