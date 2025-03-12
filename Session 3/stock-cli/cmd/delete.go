package cmd

import (
	"fmt"

	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [productName]",
	Short: "Delete a product from the stock",
	Long:  `Delete a product from the stock. For example: stock-cli delete -p "product name"`,
	Run: func(cmd *cobra.Command, args []string) {
		sm := cmd.Context().Value("stockManager").(*stock.StockManager)
		productName, _ := cmd.Flags().GetString("productName")

		if err := sm.Delete(productName); err != nil {
			cmd.PrintErr(err)
			return
		}

		fmt.Printf("Product %s deleted successfully", productName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("productName", "p", "", "Name of the product")
	deleteCmd.MarkFlagRequired("productName")
}
