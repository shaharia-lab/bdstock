// Package cmd provides commands
package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd build root command
func NewRootCmd(version string) *cobra.Command {
	cmd := cobra.Command{
		Version: version,
		Short:   "Get Bangladeshi stock market information",
		Long:    "Get the stock price information from Bangladesh Stock market",
	}

	cmd.AddCommand(NewUpdateCommand())

	return &cmd
}
