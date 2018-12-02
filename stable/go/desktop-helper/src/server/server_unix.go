package server

import (
	"os"

	"golang.org/x/sys/unix"
)

// ExitSignals contains all the signals that should cause the program
// to exit when received.
// SIGKILL and SIGSTOP cannot be caught since they are handled by the
// kernel
// os.Interupt = unix.SIGINT
var ExitSignals = []os.Signal{unix.SIGINT, unix.SIGTERM, unix.SIGHUP, unix.SIGQUIT, unix.SIGILL, unix.SIGTRAP, unix.SIGABRT, unix.SIGSYS, unix.SIGEMT}

// isProcessRunning returns a boolean as well as a process object if
func isProcessRunning(pid int) (bool, *os.Process) {
	process, err := os.FindProcess(pid)
	processExists := err != nil
	// Unfortunately, on Unix systems this always returns true. We need to send
	// a fake signal to know for sure.
	if err == nil {
		processExists = process.Signal(unix.Signal(0)) == nil
	}
	if !processExists {
		return false, nil
	}
	return true, process
}

// stopDaemon stops the server ran as daemon
func (s *Server) stopDaemon() error {
	// process should not be nil, but just in case...
	if s.process == nil {
		var err error
		s.process, err = os.FindProcess(s.ProcessInfo.PID)
		if err != nil {
			// FindProcess never fails on unix systems, even if the process doesn't
			// exists, so that's just in case they update it, one day
			return nil
		}
	}
	// We don't return the response of Signal, because we have no idea if the
	// process is actually running. The server might have been stopped already
	s.process.Signal(os.Interrupt)
	return nil
}
