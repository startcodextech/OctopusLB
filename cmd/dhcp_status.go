package cmd

import (
	"github.com/spf13/cobra"
	"github.com/startcodextech/octopuslb/cmd/require"
)

func newDhcpStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "status of DHCP server",
		Args:  require.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
