package cmd

import (
	"fmt"

	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// lowstockCmd represents the lowstock command
var lowstockCmd = &cobra.Command{
	Use:   "lowstock",
	Short: "List all low stock products",
	Long:  `List all low stock products. For example: stock-cli lowstock`,
	Run: func(cmd *cobra.Command, args []string) {
		sm := cmd.Context().Value("stockManager").(*stock.StockManager)

		lowStockProducts := sm.LowStock()

		if len(lowStockProducts) == 0 {
			fmt.Println("No low stock products")
			return
		}

		for _, p := range lowStockProducts {
			fmt.Printf("Name: %s, Quantity: %d\n", p.Name, p.Quantity)
		}
	},
}

func init() {
	rootCmd.AddCommand(lowstockCmd)
}
