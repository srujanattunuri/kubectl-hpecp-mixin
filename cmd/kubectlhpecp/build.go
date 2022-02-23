package main

import (
	"github.com/spf13/cobra"
	"github.com/srujanattunuri/kubectl-hpecp-mixin/pkg/kubectlhpecp"
)

func buildBuildCommand(m *kubectlhpecp.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
