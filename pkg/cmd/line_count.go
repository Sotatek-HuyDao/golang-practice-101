package cmd

import (
	"github.com/Sotatek-HuyDao/golang-practice-101/internal/line_count"
	"github.com/spf13/cobra"
)

func NewLineCountCmd() *cobra.Command {
	lineCountCtr := line_count.NewController()
	cmd := &cobra.Command{
		Use:   "linecount",
		Short: "Print line count of file",
		Run:   lineCountCtr.LineCount,
	}
	cmd.Flags().StringP("file", "f", "", "the input file")
	cmd.MarkFlagRequired("file")
	return cmd
}
