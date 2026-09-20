package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// appVersion and buildTimestamp are overridden at build time via ldflags:
//   -X main.appVersion={{.Version}} -X main.buildTimestamp={{.Date}}
// see .goreleaser.yaml.
var (
	appVersion     = "n/a (dev build)"
	buildTimestamp = "undefined"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print version and exit",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf(
			"%s v%s / built at %s\n",
			os.Args[0], appVersion, buildTimestamp,
		)
		return nil
	},
}
