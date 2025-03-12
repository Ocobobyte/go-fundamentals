package cmd

import (
	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [productName] [quantity]",
	Short: "Update a product in the stock",
	Long:  `Update a product in the stock. For example: stock-cli update -p "product name" -q 10`,
	Run: func(cmd *cobra.Command, args []string) {
		sm := cmd.Context().Value("stockManager").(*stock.StockManager)

		productName, _ := cmd.Flags().GetString("productName")
		quantity, _ := cmd.Flags().GetInt("quantity")

		p := stock.Product{
			Name:     productName,
			Quantity: quantity,
		}

		if err := sm.Update(p); err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Printf("Product %s updated successfully", productName)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("productName", "p", "", "Name of the product")
	updateCmd.Flags().IntP("quantity", "q", 0, "Quantity of the product")
	updateCmd.MarkFlagRequired("productName")
	updateCmd.MarkFlagRequired("quantity")
}
