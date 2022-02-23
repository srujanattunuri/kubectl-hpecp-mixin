package main

import (
	"github.com/spf13/cobra"
	"github.com/srujanreddya/kubectl-hpecp-mixin/pkg/kubectl-hpecp"
)

func buildSchemaCommand(m *kubectl-hpecp.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintSchema()
		},
	}
	return cmd
}
