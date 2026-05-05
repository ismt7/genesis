package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These variables are injected at build time via -ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of genesis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("genesis v%s (commit: %s, built: %s)\n", version, commit, date)
	},
}

func init() {
rootCmd.AddCommand(versionCmd)
}