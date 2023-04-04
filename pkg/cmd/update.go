// Package cmd provides commands
package cmd

import (
	"bd-stock-market/pkg/stock"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// NewUpdateCommand build "display" name
func NewUpdateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "update",
		Short: "Update stock price information for companies",
		Long:  "Update stock price information for companies",
		RunE: func(cmd *cobra.Command, args []string) error {
			dseEndpoint, envFound := os.LookupEnv("DSE_ENDPOINT")
			if !envFound {
				dseEndpoint = "https://www.dsebd.org"
			}

			st := stock.NewStock(dseEndpoint, false)
			stockData, err := st.GetData()
			if err != nil {
				return err
			}

			for _, info := range stockData {
				fmt.Printf("\n\n==%s==\nLast Trading Price: %s\nClosing Price: %s\nLast Update: %s\nDay's Range: %s\nWeeks Moving Range: %s\nOpening Price: %s\nDays Volume: %s\nAdjusted Opening: %s\nDays Trade: %s\nYesterday Closing: %s\nMarket Capitalization: %s\n",
					info.StockCode, info.LastTradingPrice, info.ClosingPrice, info.LastUpdate, info.DaysRange, info.WeeksMovingRange, info.OpeningPrice, info.DaysVolume, info.AdjustedOpening, info.DaysTrade, info.YesterdayClosing, info.MarketCapitalization)
			}

			return nil
		},
	}

	return &cmd
}
