// Package main
package main

import (
	"bd-stock-market/pkg/cmd"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	if err := cmd.NewRootCmd("1.0").ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
