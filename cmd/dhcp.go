/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/startcodextech/octopuslb/cmd/require"
)

var dhcp = `
This command consists of multiple subcommands to manage DHCP server.

It can be used to start, stop, restart, configure and check the status of the DHCP server.
`

func newDhcpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dhcp status [ARGS]",
		Short: "status of DHCP server",
		Long:  dhcp,
		Args:  require.NoArgs,
	}

	cmd.AddCommand(newDhcpStatusCmd())

	return cmd
}
