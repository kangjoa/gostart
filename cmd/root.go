package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gostart",
	Short: "A CLI tool to generate Go project boilerplate",
	Long:  `A CLI tool that instantly generates standardized, lightweight Go project boilerplate.`,
}

func init() {
	viper.SetDefault("author", "Your Name")
	viper.SetDefault("email", "your.email@example.com")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
