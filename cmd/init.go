package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Long:  "init generates a minimal Go project (go.mod, main.go, README, LICENSE, .gitignore).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Initializing project: %s\n", projectName)

		// Create project directory
		if err := os.Mkdir(projectName, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		// Create files from templates
		if err := createFromTemplate("templates/go.mod.template",
			filepath.Join(projectName, "go.mod"), projectName); err != nil {
			fmt.Printf("Error creating go.mod: %v\n", err)
			return
		}

		if err := createFromTemplate("templates/main.go.template",
			filepath.Join(projectName, "main.go"), projectName); err != nil {
			fmt.Printf("Error creating main.go: %v\n", err)
			return
		}

		if err := createFromTemplate("templates/README.md.template",
			filepath.Join(projectName, "README.md"), projectName); err != nil {
			fmt.Printf("Error creating README.md: %v\n", err)
			return
		}

		// Copy static template files
		if err := copyFile("templates/LICENSE", filepath.Join(projectName, "LICENSE")); err != nil {
			fmt.Printf("Error copying LICENSE: %v\n", err)
			return
		}

		if err := copyFile("templates/.gitignore", filepath.Join(projectName, ".gitignore")); err != nil {
			fmt.Printf("Error copying .gitignore: %v\n", err)
			return
		}

		fmt.Printf("Project %s created successfully!\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createFromTemplate(templatePath, outputPath, projectName string) error {
	template, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	content := strings.ReplaceAll(string(template), "{{.ProjectName}}", projectName)

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, sourceFile, 0644); err != nil {
		return err
	}

	return nil
}
