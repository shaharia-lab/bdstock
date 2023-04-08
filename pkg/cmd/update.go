// Package cmd provides commands
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"bd-stock-market/pkg/stock"

	"github.com/spf13/cobra"
)

// NewUpdateCommand build "display" name
func NewUpdateCommand() *cobra.Command {
	var filename string

	cmd := cobra.Command{
		Use:   "update",
		Short: "Update stock price information for companies",
		Long:  "Update stock price information for companies",
		RunE: func(cmd *cobra.Command, args []string) error {
			dseEndpoint, envFound := os.LookupEnv("DSE_ENDPOINT")
			if !envFound {
				dseEndpoint = "https://www.dsebd.org"
			}

			st := stock.NewStock(dseEndpoint, filename != "")
			stockData, err := st.GetData()
			if err != nil {
				return err
			}

			jsonData, err := json.MarshalIndent(stockData, " ", "   ")
			if err != nil {
				return fmt.Errorf("failed to marshal json. erro: %w", err)
			}

			if filename != "" {
				err = ioutil.WriteFile(filename, jsonData, 0644)
				if err != nil {
					return fmt.Errorf("failed to write data to file. error: %w", err)
				}
			} else {
				fmt.Println(string(jsonData))
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&filename, "file", "f", "", "output filename")

	return &cmd
}
