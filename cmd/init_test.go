package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFromTemplate(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	var tests = []struct {
		template    string
		projectName string
		author      string
		email       string
		expected    string
	}{
		{
			template:    "Hello {{.ProjectName}} by {{.Author}} ({{.Email}})",
			projectName: "racoon-restaurant",
			author:      "Linda Belcher",
			email:       "linda.belcher@example.com",
			expected:    "Hello racoon-restaurant by Linda Belcher (linda.belcher@example.com)",
		},
		{
			template:    "Project: {{.ProjectName}}",
			projectName: "",
			author:      "",
			email:       "",
			expected:    "Project: ",
		},
		{
			template:    "{{.ProjectName}} - {{.Author}}",
			projectName: "community-garden-plot",
			author:      "Bob Belcher",
			email:       "bob.belcher@test.com",
			expected:    "community-garden-plot - Bob Belcher",
		},
	}

	for _, test := range tests {
		// Create template file
		templatePath := filepath.Join(tempDir, "template.txt")
		os.WriteFile(templatePath, []byte(test.template), 0644)

		// Test the function
		outputPath := filepath.Join(tempDir, "output.txt")
		createFromTemplate(templatePath, outputPath, test.projectName, test.author, test.email)

		// Check result
		content, _ := os.ReadFile(outputPath)
		if string(content) != test.expected {
			t.Errorf("createFromTemplate() = %v, want %v", string(content), test.expected)
		}
	}
}

func TestCopyFile(t *testing.T) {
	tempDir := t.TempDir()

	var tests = []struct {
		content  string
		expected string
	}{
		{
			content:  "This is a test file",
			expected: "This is a test file",
		},
		{
			content:  "",
			expected: "",
		},
		{
			content:  "\nRed 1\nBlue 2\nGreen 3",
			expected: "\nRed 1\nBlue 2\nGreen 3",
		},
	}

	for _, test := range tests {
		// Create source file
		srcPath := filepath.Join(tempDir, "source.txt")
		os.WriteFile(srcPath, []byte(test.content), 0644)

		// Test the copyFile function
		dstPath := filepath.Join(tempDir, "destination.txt")
		copyFile(srcPath, dstPath)

		// Check result
		content, _ := os.ReadFile(dstPath)
		if string(content) != test.expected {
			t.Errorf("copyFile() = Actual: %v\n,Expected: %v", string(content), test.expected)
		}
	}
}

func benchmarkCreateFromTemplate(projectName, author, email string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		createFromTemplate("templates/main.go.template", "temp_output.txt", projectName, author, email)
	}
}

func BenchmarkCreateFromTemplateSimple(b *testing.B) {
	benchmarkCreateFromTemplate("kuchi-kopi-project", "Louise Belcher", "louise.belcher@example.com", b)
}
func BenchmarkCreateFromTemplateComplex(b *testing.B) {
	benchmarkCreateFromTemplate("long-project-name-chungking-express-by-wong-kar-wai", "王菲", "faye.wong@example.com", b)
}
func BenchmarkCreateFromTemplateEmpty(b *testing.B) { benchmarkCreateFromTemplate("", "", "", b) }
