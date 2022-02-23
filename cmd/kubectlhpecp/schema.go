package main

import (
	"github.com/spf13/cobra"
	"github.com/srujanattunuri/kubectl-hpecp-mixin/pkg/kubectlhpecp"
)

func buildSchemaCommand(m *kubectlhpecp.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintSchema()
		},
	}
	return cmd
}
