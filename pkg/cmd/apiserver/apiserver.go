// Package apiserver provides a command line interface for managing the API server.
package apiserver

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/minuk-dev/minuk-boilerplate/pkg/cmd/apiserver/start"
)

// Options contains the options for the API server command.
type Options struct{}

// NewCommand creates a new command for the API server.
func NewCommand(opts Options) *cobra.Command {
	//exhaustruct:ignore
	cmd := &cobra.Command{
		Use:  "apiserver",
		Long: "apiserver is a command line tool to manage the API server.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := opts.Prepare(cmd, args)
			if err != nil {
				return fmt.Errorf("failed to prepare: %w", err)
			}

			err = opts.Run(cmd, args)
			if err != nil {
				return fmt.Errorf("failed to run: %w", err)
			}

			return nil
		},
	}

	cmd.AddCommand(start.NewCommand(start.Options{}))

	return cmd
}

// Prepare prepares the command options.
func (o *Options) Prepare(*cobra.Command, []string) error {
	// Here you can add any preparation logic you need
	// For example, you can parse flags or set up logging
	return nil
}

// Run runs the command.
func (o *Options) Run(cmd *cobra.Command, _ []string) error {
	// Here you can add the main logic for your API server
	// For example, you can start an HTTP server or connect to a database
	err := cmd.Help()
	if err != nil {
		return fmt.Errorf("failed to display help: %w", err)
	}

	return nil
}
