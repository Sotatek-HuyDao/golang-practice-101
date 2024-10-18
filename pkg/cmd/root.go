package cmd

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "futil",
		Short: "File Utility",
		Long:  `A simple file utility for line counting and checksum calculation.`,
	}
	versionCmd := NewVersionCmd()
	rootCmd.AddCommand(versionCmd)
	lineCountCmd := NewLineCountCmd()
	rootCmd.AddCommand(lineCountCmd)
	checksumCmd := NewChecksumCmd()
	rootCmd.AddCommand(checksumCmd)
	return rootCmd
}
