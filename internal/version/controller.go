package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Controller interface {
	Version(cmd *cobra.Command, args []string)
}

type ControllerImpl struct {
	svc Service
}

func NewController() *ControllerImpl {
	return &ControllerImpl{
		svc: NewService(),
	}
}

func (c *ControllerImpl) Version(cmd *cobra.Command, args []string) {
	version, err := c.svc.Version()
	if err != nil {
		fmt.Println(fmt.Errorf("error when get version: %w", err))
	}
	fmt.Println(version)
}
