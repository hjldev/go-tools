package main

import (
	"log"

	"github.com/hjldev/go-tools/cmd"
	"github.com/spf13/cobra"
)

var (
	version string = "v1.0.0"

	rootCmd = &cobra.Command{
		Use:     "tools",
		Short:   "tools.",
		Long:    `tools.`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(cmd.ProjectNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
