package cmd

import (
	"context"
	"os"

	"github.com/ocobobyte/stock-cli/stock"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stock",
	Short: "Stock is a CLI tool to manage your stock",
	Long:  `Stock is a CLI tool to manage your stock. It allows you to add, update, and list products in your stock.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(sm *stock.StockManager) {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Store the StockManager instance in the command context or a global config variable.
		cmd.SetContext(context.WithValue(cmd.Context(), "stockManager", sm))
		return nil
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
