package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/lgr"
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/server"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func startSuccessResponse(port int) {
	msg := fmt.Sprintf("server running on port: %d", port)

	if cliLogger.Type() == lgr.TypeJSON {
		msg = fmt.Sprintf(`{"port":"%d"}`, port)
	}
	cliLogger.PrintSuccess(msg)
}

func startCmd(env *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "starts the gRPC server",
	}

	// flags
	port := cmd.Flags().IntP("port", "p", 51556, "port to use")
	asDaemon := cmd.Flags().BoolP("daemon", "d", false, "start the server as a daemon")
	attach := cmd.Flags().BoolP("attach-if-exists", "a", false, "start the server as a daemon")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return start(*env, *port, *asDaemon, *attach)
	}
	return cmd
}

func start(env string, port int, asDaemon, attach bool) error {
	isRunning, _, err := server.RunningServer(env, cliLogger)
	if err != nil {
		return err
	}

	// TODO(melvin): Return nil if asDaemon=true and attach=false
	//if srv.IsRunning() && asDaemon {
	if isRunning {
		if asDaemon {
			startSuccessResponse(port)
			return nil
		}
		return errors.New("server already running")
	}

	srv, err := server.New(env, cliLogger)
	if err != nil {
		return errors.Wrap(err, "Could not create a new server")
	}
	if asDaemon {
		return srv.Daemonize()
	}

	// once we return we need to make sure we remove the file
	// this line might be redundant in some cases, but it's fine
	defer server.RemoveProcessInfo(env)

	// we setup a channel
	done := make(chan os.Signal, 1)

	var RunErr error
	go func() {
		RunErr = srv.Run(port)
		done <- os.Interrupt
	}()

	startSuccessResponse(port)

	// Because we cannot write on the log file from multiple process at
	// the same time, we need to close it for now since this process
	// will last for an unknown amount of time
	cliLogger.CloseLogFile()

	signal.Notify(done, server.ExitSignals...)
	<-done

	// now that we're closing we can re-enable the logs
	cliLogger.UseLogFile(srv.AppDataDir)

	// if RunErr is not nil it means we stopped because we could not run
	// the server
	if RunErr != nil {
		return RunErr
	}
	return srv.Stop()
}
