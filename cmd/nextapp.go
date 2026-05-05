package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nextappCmd = &cobra.Command{
	Use:   "nextapp [project-name]",
	Short: "Create a new Next.js app using create-next-app",
	Long:  `Create a new Next.js project in the current directory using 'npx create-next-app@latest'.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		npxArgs := []string{"create-next-app@latest"}
		if len(args) == 1 {
			npxArgs = append(npxArgs, args[0])
		}

		c := exec.Command("npx", npxArgs...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return fmt.Errorf("failed to run create-next-app: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nextappCmd)
}