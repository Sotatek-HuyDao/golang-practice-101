package line_count

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Controller interface {
	LineCount(cmd *cobra.Command, args []string)
}

type ControllerImpl struct {
	svc Service
}

func NewController() *ControllerImpl {
	return &ControllerImpl{
		svc: NewService(),
	}
}

func (c *ControllerImpl) LineCount(cmd *cobra.Command, args []string) {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	lineCount, err := c.svc.LineCount(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		return
	}
	fmt.Println(lineCount)
}
