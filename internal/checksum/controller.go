package checksum

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Controller interface {
	Checksum(cmd *cobra.Command, args []string)
}

type ControllerImpl struct {
	svc Service
}

func NewController() *ControllerImpl {
	return &ControllerImpl{
		svc: NewService(),
	}
}

func (c *ControllerImpl) Checksum(cmd *cobra.Command, args []string) {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	// only return md5 checksum if user requests
	isMd5, err := cmd.Flags().GetBool("md5")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	if isMd5 {
		checksum, err := c.svc.MD5Checksum(filePath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %w", err))
			return
		}
		fmt.Println(checksum)
		return
	}

	// only return sha1 checksum if user requests
	isSha1, err := cmd.Flags().GetBool("sha1")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	if isSha1 {
		checksum, err := c.svc.SHA1Checksum(filePath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %w", err))
			return
		}
		fmt.Println(checksum)
		return
	}

	// only return sha256 checksum if user requests
	isSha256, err := cmd.Flags().GetBool("sha256")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	if isSha256 {
		checksum, err := c.svc.SHA256Checksum(filePath)
		if err != nil {
			fmt.Println(fmt.Errorf("error: %w", err))
			return
		}
		fmt.Println(checksum)
		return
	}

	// throw error if user didn't request any checksum algorithm
	fmt.Println("error: Please specify at least one checksum algorithm (--md5, --sha1, or --sha256)")
}
