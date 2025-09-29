package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gostart",
	Short: "A CLI tool to generate Go project boilerplate",
	Long:  `A CLI tool that instantly generates standardized, lightweight Go project boilerplate.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
