// Package cli defines the command line interface for supply run.
package cli

import "github.com/spf13/cobra"

// New creates a new cli root command.
func New(version string) *cobra.Command {
	root := &cobra.Command{
		Version: version,
		Use:     "supplyrun [COMMAND]",
		Short:   "The Supply Run API web service",
	}

	root.AddCommand(startCmd(version))

	return root
}
