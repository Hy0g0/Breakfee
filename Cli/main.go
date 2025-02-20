package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-assist",
	Short: "AI-powered developer assistant CLI",
	Long: `dev-assist is a CLI tool that helps developers with various tasks using AI:
- Summarize repositories
- Generate README files
- Generate unit tests
- Suggest CI/CD pipeline configurations`,
}

var (
	ollamaURL string
	chromaURL string
)

func init() {
	// Add commands
	rootCmd.AddCommand(summarizeCmd())
	rootCmd.AddCommand(generateReadmeCmd())
	rootCmd.AddCommand(generateTestsCmd())
	rootCmd.AddCommand(suggestCicdCmd())

	// Add global flags
	rootCmd.PersistentFlags().StringVar(&ollamaURL, "ollama-url", "http://ollama:11434", "Ollama server URL")
	rootCmd.PersistentFlags().StringVar(&chromaURL, "chroma-url", "http://chromadb:8000", "ChromaDB server URL")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func summarizeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "summarize [path]",
		Short: "Summarize a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			fmt.Printf("Summarizing repository at: %s\n", path)
			// TODO: Implement repository summarization
			return nil
		},
	}
}

func generateReadmeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate-readme [path]",
		Short: "Generate a README.md file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			fmt.Printf("Generating README.md for repository at: %s\n", path)
			// TODO: Implement README generation
			return nil
		},
	}
}

func generateTestsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate-tests [path]",
		Short: "Generate unit tests",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			fmt.Printf("Generating unit tests for repository at: %s\n", path)
			// TODO: Implement test generation
			return nil
		},
	}
}

func suggestCicdCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "suggest-cicd [path]",
		Short: "Suggest CI/CD pipeline configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			fmt.Printf("Suggesting CI/CD pipeline for repository at: %s\n", path)
			// TODO: Implement CI/CD pipeline suggestion
			return nil
		},
	}
}
