package cmd

import "github.com/spf13/cobra"

var mcpserverCmd = &cobra.Command{
	Use:   "mcpserver [project-name]",
	Short: "Create a new MCP server using @modelcontextprotocol/create-server",
	Long:  `Create a new MCP server project in the current directory using 'npx @modelcontextprotocol/create-server'.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		npxArgs := []string{"@modelcontextprotocol/create-server"}
		if len(args) == 1 {
			npxArgs = append(npxArgs, args[0])
		}

		return runAttachedCommand("npx", npxArgs, "@modelcontextprotocol/create-server")
	},
}

func init() {
	rootCmd.AddCommand(mcpserverCmd)
}
