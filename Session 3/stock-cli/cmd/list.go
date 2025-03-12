package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [limit]",
	Short: "List all products in the stock",
	Long:  `List all products in the stock. For example: stock-cli list -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		sm := cmd.Context().Value("stockManager").(*stock.StockManager)

		limit, _ := cmd.Flags().GetInt("limit")
		products := sm.List(limit)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Product", "Quantity"})
		for i, product := range products {
			t.AppendRow(table.Row{i + 1, product.Name, product.Quantity})
		}
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntP("limit", "l", 10, "Product to list")
}
