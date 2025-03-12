package cmd

import (
	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [productName] [quantity]",
	Short: "Add a product to the stock",
	Long:  `Add a product to the stock. For example: stock-cli add -p "product name" -q 10`,
	Run: func(cmd *cobra.Command, args []string) {
		sm := cmd.Context().Value("stockManager").(*stock.StockManager)

		productName, _ := cmd.Flags().GetString("productName")
		quantity, _ := cmd.Flags().GetInt("quantity")

		p := stock.Product{
			Name:     productName,
			Quantity: quantity,
		}

		if err := sm.Add(p); err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Printf("Product %s added successfully", productName)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("productName", "p", "", "Name of the product")
	addCmd.Flags().IntP("quantity", "q", 0, "Quantity of the product")
	addCmd.MarkFlagRequired("productName")
	addCmd.MarkFlagRequired("quantity")
}
