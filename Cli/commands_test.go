package main

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("Command Execution", func() {
	var (
		cmd    *cobra.Command
		output *bytes.Buffer
	)

	BeforeEach(func() {
		output = new(bytes.Buffer)
		rootCmd.SetOut(output)
	})

	Context("when executing summarize command", func() {
		BeforeEach(func() {
			cmd = summarizeCmd()
		})

		It("should fail with no arguments", func() {
			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
		})

		It("should accept a valid path argument", func() {
			cmd.SetArgs([]string{"."})
			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())
			Expect(output.String()).To(ContainSubstring("Summarizing repository"))
		})
	})

	Context("when executing generate-readme command", func() {
		BeforeEach(func() {
			cmd = generateReadmeCmd()
		})

		It("should fail with no arguments", func() {
			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
		})

		It("should accept a valid path argument", func() {
			cmd.SetArgs([]string{"."})
			err := cmd.Execute()
			Expect(err).NotTo(HaveOccurred())
			Expect(output.String()).To(ContainSubstring("Generating README.md"))
		})
	})
}) 