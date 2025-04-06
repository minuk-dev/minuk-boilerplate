// Package start provides the command to start the API server.
package start

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver"
)

// Options contains the options for the start command.
// It is a struct that holds the command line flags & internal state.
type Options struct {
	addr string
	db   string
}

// NewCommand creates a new start command.
func NewCommand(opts Options) *cobra.Command {
	//exhaustruct:ignore
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the API server",
		Long:  "Start the API server with the specified options.",
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

	cmd.Flags().StringVar(&opts.addr, "addr", ":8080", "Address to bind the server to")
	cmd.Flags().StringVar(&opts.db, "db", "sqlite:test.db", "Database file to use")

	return cmd
}

// Prepare prepares the command options.
func (o *Options) Prepare(*cobra.Command, []string) error {
	return nil
}

// Run runs the command.
func (o *Options) Run(*cobra.Command, []string) error {
	apiserver := apiserver.New(
		apiserver.Settings{
			Addr: o.addr,
			DB:   o.db,
		},
	)

	apiserver.Run()

	return nil
}
