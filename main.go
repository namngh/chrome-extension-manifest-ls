package main

import (
	"fmt"

	"github.com/namngh/chrome-extension-manifest-ls/langserver"
	"github.com/spf13/cobra"
	serverpkg "github.com/tliron/glsp/server"
	"github.com/tliron/kutil/util"
)

const toolName = "chrome-extension-manifest-language-server"

var protocol string
var verbose int

var command = &cobra.Command{
	Use:   toolName,
	Short: "Start the Chrome extension manifest language server",
	Run: func(cmd *cobra.Command, args []string) {
		err := Run(protocol)
		util.FailOnError(err)
	},
}

func init() {
	command.Flags().StringVarP(&protocol, "protocol", "p", "stdio", "protocol (\"stdio\")")
	command.Flags().CountVarP(&verbose, "verbose", "v", "verbosity level")
}

func Run(protocol string) error {
	server := serverpkg.NewServer(&langserver.Handler, toolName, verbose > 0)

	switch protocol {
	case "stdio":
		return server.RunStdio()

	default:
		return fmt.Errorf("unsupperted protocol: %s", protocol)
	}
}
