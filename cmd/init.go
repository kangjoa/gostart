package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Long:  "init generates a minimal Go project (go.mod, main.go, README, LICENSE, .gitignore).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Initializing project: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
