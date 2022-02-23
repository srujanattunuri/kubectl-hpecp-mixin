package main

import (
	"github.com/spf13/cobra"
	"github.com/srujanattunuri/kubectl-hpecp-mixin/pkg/kubectlhpecp"
)

func buildUninstallCommand(m *kubectlhpecp.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
