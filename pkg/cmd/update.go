// Package cmd provides commands
package cmd

import (
	"bd-stock-market/pkg/stock"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// NewUpdateCommand build "display" name
func NewUpdateCommand() *cobra.Command {
	var outputFormat string
	var outputFile string
	var collectionBatchSize int
	var verbose bool

	cmd := cobra.Command{
		Use:   "update",
		Short: "Update stock price information for companies",
		Long:  "Update stock price information for companies",
		RunE: func(cmd *cobra.Command, args []string) error {
			st := stock.NewStock("https://www.dsebd.org", false)

			stockData, err := st.GetData(collectionBatchSize)
			if err != nil {
				return err
			}

			if outputFormat == "json" {
				jsonData, err := json.MarshalIndent(stockData, "", "  ")
				if err != nil {
					return fmt.Errorf("error encoding JSON")
				}

				if outputFile != "" {
					err = ioutil.WriteFile(outputFile, jsonData, 0644)
					if err != nil {
						return fmt.Errorf("error writing file: %w", err)
					}

					fmt.Println("data has been saved")
				}
			}

			for _, info := range stockData {
				fmt.Printf("\n\n==%s==\nLast Trading Price: %s\nClosing Price: %s\nLast Update: %s\nDay's Range: %s\nWeeks Moving Range: %s\nOpening Price: %s\nDays Volume: %s\nAdjusted Opening: %s\nDays Trade: %s\nYesterday Closing: %s\nMarket Capitalization: %s\n",
					info.StockCode, info.LastTradingPrice, info.ClosingPrice, info.LastUpdate, info.DaysRange, info.WeeksMovingRange, info.OpeningPrice, info.DaysVolume, info.AdjustedOpening, info.DaysTrade, info.YesterdayClosing, info.MarketCapitalization)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&outputFormat, "format", "c", "", "Display format. eg: json, table")
	cmd.PersistentFlags().StringVarP(&outputFile, "file", "f", "", "File path to save the output")
	cmd.PersistentFlags().IntVarP(&collectionBatchSize, "batch-size", "b", 10, "Batch size to fetch data concurrently")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	return &cmd
}
