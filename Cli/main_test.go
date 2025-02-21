package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func TestDevAssist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dev Assist CLI Suite")
}

var _ = Describe("CLI Commands", func() {
	Context("Root Command", func() {
		It("should have correct command name", func() {
			Expect(rootCmd.Use).To(Equal("dev-assist"))
		})

		It("should have a description", func() {
			Expect(rootCmd.Long).NotTo(BeEmpty())
		})
	})

	Context("Summarize Command", func() {
		var cmd *cobra.Command

		BeforeEach(func() {
			cmd = summarizeCmd()
		})

		It("should have correct command name", func() {
			Expect(cmd.Use).To(Equal("summarize [path]"))
		})

		It("should require exactly one argument", func() {
			Expect(cmd.Args).NotTo(BeNil())
		})
	})

	Context("Generate README Command", func() {
		var cmd *cobra.Command

		BeforeEach(func() {
			cmd = generateReadmeCmd()
		})

		It("should have correct command name", func() {
			Expect(cmd.Use).To(Equal("generate-readme [path]"))
		})

		It("should require exactly one argument", func() {
			Expect(cmd.Args).NotTo(BeNil())
		})
	})

	Context("Generate Tests Command", func() {
		var cmd *cobra.Command

		BeforeEach(func() {
			cmd = generateTestsCmd()
		})

		It("should have correct command name", func() {
			Expect(cmd.Use).To(Equal("generate-tests [path]"))
		})

		It("should require exactly one argument", func() {
			Expect(cmd.Args).NotTo(BeNil())
		})
	})

	Context("Suggest CICD Command", func() {
		var cmd *cobra.Command

		BeforeEach(func() {
			cmd = suggestCicdCmd()
		})

		It("should have correct command name", func() {
			Expect(cmd.Use).To(Equal("suggest-cicd [path]"))
		})

		It("should require exactly one argument", func() {
			Expect(cmd.Args).NotTo(BeNil())
		})
	})
})
