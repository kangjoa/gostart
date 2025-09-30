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

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.go-new-project")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
