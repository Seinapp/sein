package server

import (
	"os"
)

// ExitSignals contains all the signals that should cause the program
// to exit when received.
// os.Kill cannot be caught since they it's handled by the
// kernel
// os.Interupt is only triggered by ^C (Control-C) or ^BREAK (Control-Break)
var ExitSignals = []os.Signal{os.Interrupt}

func isProcessRunning(pid int) (bool, *os.Process) {
	process, err := os.FindProcess(pid)
	return err != nil, process
}

// stopDaemon stops the server ran as daemon
func (s *Server) stopDaemon() error {
	// process should not be nil, but just in case...
	if s.process == nil {
		var err error
		s.process, err = os.FindProcess(s.ProcessInfo.PID)
		if err != nil {
			return nil
		}
	}

	// For Windows, we can "receive" os.Interrupt but not send it.
	// It's only triggered by ^C (Control-C) or ^BREAK (Control-Break)
	// see https://golang.org/pkg/os/signal/#hdr-Windows
	// This means we need to kill the server the hard way...
	return s.process.Kill()
}
