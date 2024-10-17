package cmd

import (
	"github.com/Sotatek-HuyDao/golang-practice-101/internal/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	verController := version.NewController()
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version info",
		Run:   verController.Version,
	}
}
