package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  "Manage configuration settings for the gostart tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updated configuration settings")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
