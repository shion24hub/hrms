/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"net/url"
	"os"
	"time"

	"github.com/shion24hub/hrms/bybit"
	"github.com/spf13/cobra"
)

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
		if len(output) == 0 {
			return errors.New("output directory is required")
		}

		b, err := time.Parse("20060102", begin)
		if err != nil {
			return errors.New("failed to parse begin date")
		}
		e, err := time.Parse("20060102", end)
		if err != nil {
			return errors.New("failed to parse end date")
		}

		// MAIN PROCESS

		dateRange := []time.Time{}
		for d := b; d.Before(e); d = d.AddDate(0, 0, 1) {
			dateRange = append(dateRange, d)
		}
		for _, date := range dateRange {

			// make url
			ep, err := bybit.MakeUrl(symbol, date)
			if err != nil {
				return errors.New("failed to make url")
			}

			// fetch trading data
			btrs, err := bybit.FetchTradingData(ep)
			if err != nil {
				return errors.New("failed to download trading data")
			}
			data := [][]string{}
			for _, btr := range btrs {
				data = append(data, []string{btr.Timestamp.Format("2006-01-02 15:04:05"), btr.Symbol, btr.Side, btr.Size, btr.Price})
			}

			// create file
			fileName := symbol + date.Format("2006-01-02") + ".csv"
			outputPath, err := url.JoinPath(output, fileName)
			if err != nil {
				return errors.New("failed to join path")
			}
			of, err := os.Create(outputPath)
			if err != nil {
				return errors.New("failed to create file")
			}

			// write to file
			w := csv.NewWriter(of)
			w.WriteAll(data)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
