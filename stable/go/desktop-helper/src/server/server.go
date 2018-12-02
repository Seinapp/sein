// Package server contains functions and structs to manager the gRPC server
package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/lgr"
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/services/utils"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunningServer checks if a server is already running, and returns it.
// TODO(melvin): Rename to RunningServer()
func RunningServer(env string, cliLogger lgr.Printer) (bool, *Server, error) {
	// init a new server for the specific env
	srv, err := New(env, cliLogger)
	if err != nil {
		return false, nil, errors.Wrap(err, "could not get the server data")
	}

	// populate the server with the data from the running server (if there's one)
	srv.ProcessInfo, err = ReadProcessInfo(env)
	if err != nil && !os.IsNotExist(err) {
		return false, nil, errors.Wrap(err, "could not read the process info")
	}
	if os.IsNotExist(err) {
		return false, nil, nil
	}

	// now let's make sure the server is running
	var isRunning bool
	isRunning, srv.process = isProcessRunning(srv.ProcessInfo.PID)
	if !isRunning {
		srv.CLILogger.PrintInfo("state file found but process is not running")
		RemoveProcessInfo(env)
		return false, nil, nil
	}

	// At that point we're 99.9% sure everything is fine, but maybe the server
	// has a bug and cannot accept incoming connections.
	// Let's ping it to make sure it's working
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", srv.ProcessInfo.Port), grpc.WithInsecure())
	if err != nil {
		// we kill the process, so we can start a new server
		if err = srv.process.Kill(); err != nil {
			return false, nil, errors.Wrap(err, "could not connect nor kill the running server")
		}
		return false, nil, err
	}
	defer conn.Close()

	client := utils.NewUtilsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx, &utils.PingRequest{})
	// If we get any errors we can assume the server is not working properly
	if err != nil {
		srv.CLILogger.PrintError(errors.Wrap(err, "the server could not be pinged, killing it"))
		// we kill the process, so we can start a new server
		if err = srv.process.Kill(); err != nil {
			return false, nil, errors.Wrap(err, "could not ping nor kill the running server")
		}
		RemoveProcessInfo(env)
		return false, nil, nil
	}
	return true, srv, nil
}

// New creates a new server, fails if a server already exists for this env
func New(env string, printer lgr.Printer) (*Server, error) {
	appName := config.AppNameForEnv(env)
	appDataDir, err := config.AppDataDir(appName)
	if err != nil {
		return nil, errors.Wrap(err, "could not find the app data directory")
	}

	return &Server{
		Env:        env,
		AppDataDir: appDataDir,
		CLILogger:  printer,
	}, nil
}

// Server represents the configuration of the app
type Server struct {
	Env         string
	ProcessInfo *ProcessInfo
	AppDataDir  string
	process     *os.Process
	listener    net.Listener
	// ServerLogger is a logger only set for the process running the server
	// and is *only* logging server related data:
	// incoming connection, messages, errors occuring in a message
	// implementation, etc.
	ServerLogger *zap.Logger
	// CLILogger is an interface that prints data on stdout and stderr only
	// It's meant to send data back to the process that ran the command
	// running this code to let them know of warnings, infos, errors, etc.
	// occuring when starting, stoping, and attaching to a server
	// ServerLogger & CLILogger are separated to prevent issues when writting
	// on a log file concurrently from multiple processes
	CLILogger lgr.Printer
}

// IsDeamon returns false if the current process is the one running the
// server, true if the server is running on another process
func (s *Server) IsDeamon() bool {
	return s.ProcessInfo.PID != os.Getpid()
}

// FIXME(melvin): thoughts:
// Always run in daemon mode?
// - In regular mode and when attaching, we redirect stdout and stderr
//   of the daemon process to stdout and stderr of the current process
// - Handle keyboard event to detach (ctrl-j + d)
//    - could be a 2nd goroutine, careful not leaking it
//    - https://tutorialedge.net/golang/reading-console-input-golang/
//    - buffio with channel?
// - pass signals around

// Run runs the gRPC server if it's not already running
func (s *Server) Run(port int) error {
	s.ServerLogger = newLogger(s.AppDataDir, s.CLILogger)
	defer s.ServerLogger.Sync()

	s.ProcessInfo = &ProcessInfo{
		Port: port,
		PID:  os.Getpid(),
	}

	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return errors.Wrapf(err, "could not listen on the provided port")
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.UnaryLoggerInterceptor(s.ServerLogger),
			s.UnaryPanicInterceptor(),
		)),
	)

	utils.RegisterUtilsServer(grpcServer, utils.NewServer(s.ServerLogger))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Now that everything is ready we can create the info file
	// for other app to use
	if err := s.WriteProcessInfo(); err != nil {
		return errors.Wrap(err, "could not write the info file")
	}

	if err := grpcServer.Serve(s.listener); err != nil {
		return errors.Wrap(err, "could not accept incoming connections")
	}

	return nil
}

// Attach re-attaches a daemon
func (s *Server) Attach() error {
	// TODO(melvin): implement
	return nil
}

// Daemonize detaches a daemon
func (s *Server) Daemonize() error {
	// TODO(melvin): implement
	// os.startProcess
	return nil
}

// Stop stops the gRPC server
func (s *Server) Stop() error {
	if !s.IsDeamon() {
		if s.listener == nil {
			return nil
		}
		// Stop current process
		s.listener.Close()
		return RemoveProcessInfo(s.Env)
	}

	return s.stopDaemon()
}

// ShouldBeRunning returns true if the server should be running (it has
// been running)
func (s *Server) ShouldBeRunning() bool {
	return s.ProcessInfo != nil && s.ProcessInfo.PID > 0
}

// IsRunning will send a ping to the server to see if the server
// really is running
func (s *Server) IsRunning() bool {
	// TODO(melvin): implement
	return false
}
