package cmd

import (
	"github.com/Sotatek-HuyDao/golang-practice-101/internal/checksum"
	"github.com/spf13/cobra"
)

func NewChecksumCmd() *cobra.Command {
	checksumCtr := checksum.NewController()
	cmd := &cobra.Command{
		Use:   "checksum",
		Short: "Print checksum of file",
		Run:   checksumCtr.Checksum,
	}
	cmd.Flags().StringP("file", "f", "", "the input file")
	cmd.Flags().Bool("md5", false, "calculate MD5 checksum")
	cmd.Flags().Bool("sha1", false, "calculate SHA1 checksum")
	cmd.Flags().Bool("sha256", false, "calculate SHA256 checksum")
	cmd.MarkFlagRequired("file")
	return cmd
}
