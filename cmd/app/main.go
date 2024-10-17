package main

import (
	"fmt"
	"os"

	"github.com/Sotatek-HuyDao/golang-practice-101/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
