package cmd

import (
	"errors"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/server"
	"github.com/spf13/cobra"
)

func stopCmd(env *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stops the gRPC server",
	}

	// flags
	port := cmd.Flags().Int("port", 51556, "port to use")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return stop(*env, *port)
	}
	return cmd
}

func stop(env string, port int) error {
	isRunning, srv, err := server.RunningServer(env, cliLogger)
	if err != nil {
		return err
	}

	if !isRunning {
		return errors.New("no server running")
	}

	return srv.Stop()
}
