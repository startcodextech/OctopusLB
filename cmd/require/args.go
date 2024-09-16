package require

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.Errorf(
			"%q accepts no arguments\n\nUsage: %s",
			cmd.CommandPath(),
			cmd.UseLine(),
		)
	}
	return nil
}
