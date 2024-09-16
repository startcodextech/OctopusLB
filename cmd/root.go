/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var globalUsage = `
Common actions for the Octopus Load Balancer:

- octopuslb status
- octopuslb start
- octopuslb stop
- octopuslb restart

By default, the default directories depend on the operating system. The defaults are listed below:

| Operating System    | Configuration Path                         | Data Path                      |
|---------------------|--------------------------------------------|--------------------------------|
| Linux               | $HOME/.config/octopuslb.json               | $home/.local/share/octopuslb   |
| MacOS               | $HOME/Library/Preferences/octopuslb.json   | $HOME/Library/octopuslb        |
| Windows             | %APPDATA%\octopuslb.json                   | %APPDATA%\octopuslb            |
`

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "octopuslb",
		Short:        "Octopus Load Balancer for managing load balancing, reverse proxy DHCP, BGP and more",
		Long:         globalUsage,
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()

	settings.AddFlags(flags)

	return cmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
