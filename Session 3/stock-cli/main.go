package main

import (
	"fmt"
	"os"

	"github.com/ocobobyte/stock-cli/cmd"
	"github.com/ocobobyte/stock-cli/stock"
)

func main() {
	sm, err := stock.NewStockManager("stock.json")
	if err != nil {
		fmt.Printf("Error initializing stock manager: %v\n", err)
		os.Exit(1)
	}
	cmd.Execute(sm)
}
