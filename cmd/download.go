/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/shion24hub/hrms/bybit"
	"github.com/spf13/cobra"
)

var symbol, begin, end, output string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "",
	Long:  "Download historical trading data for a given date.",
	RunE: func(cmd *cobra.Command, args []string) error {

		// validate input
		if len(symbol) == 0 {
			return errors.New("symbol is required")
		}
		if len(begin) == 0 {
			return errors.New("begin date is required")
		}
		if len(end) == 0 {
			return errors.New("end date is required")
		}

		b, err := time.Parse("20060102", begin)
		if err != nil {
			return errors.New("failed to parse begin date")
		}
		e, err := time.Parse("20060102", end)
		if err != nil {
			return errors.New("failed to parse end date")
		}

		dateRange := []time.Time{}
		for d := b; d.Before(e); d = d.AddDate(0, 0, 1) {
			dateRange = append(dateRange, d)
		}
		for _, date := range dateRange {
			url, err := bybit.MakeUrl(symbol, date)
			if err != nil {
				return errors.New("failed to make url")
			}
			_, err = bybit.DownloadTradingData(url)
			if err != nil {
				return errors.New("failed to download trading data")
			}
		}

		fmt.Println("download called!")
		fmt.Println("symbol:", symbol)
		fmt.Println("begin:", begin)
		fmt.Println("end:", end)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Symbol to download. (e.g. BTCUSD)")
	downloadCmd.Flags().StringVarP(&begin, "begin", "b", "", "Begin date. (e.g. 20240101)")
	downloadCmd.Flags().StringVarP(&end, "end", "e", "", "End date. (e.g. 20240103)")
	downloadCmd.Flags().StringVarP(&output, "output", "o", "./", "Output directory. (e.g. /path/to/output)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
