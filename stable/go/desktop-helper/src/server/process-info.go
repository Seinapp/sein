package server

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	"github.com/pkg/errors"
)

const (
	// ProcessInfoRelPath contains the relative path (from the Sein's app
	// data dir) to the file that contains the server's process info
	ProcessInfoRelPath = "helper-process.json"
)

// ProcessInfo contains process information about a running gRPC server
type ProcessInfo struct {
	Port int `json:"port"`
	PID  int `json:"pid"`
}

// ReadProcessInfo returns the the process info for the server running on the
// current Env. If the file doesn't exist an error matching os.IsNotExist
// is returned
func ReadProcessInfo(env string) (*ProcessInfo, error) {
	appDataDir, err := config.AppDataDirForEnv(env)
	if err != nil {
		return nil, errors.Wrap(err, "could not find the app data directory")
	}

	processinfoPath := filepath.Join(appDataDir, ProcessInfoRelPath)
	file, err := os.Open(processinfoPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "could not open the process info file")
	}

	// for IsNotExist, we just pass the error down so it can be checked by
	// the caller
	if os.IsNotExist(err) {
		return nil, err
	}
	defer file.Close()

	processinfo := &ProcessInfo{}
	if err = json.NewDecoder(file).Decode(processinfo); err != nil {
		return nil, errors.Wrap(err, "could not decode the file content")
	}
	return processinfo, nil
}

// RemoveProcessInfo removes any process info file attached to this env
func RemoveProcessInfo(env string) error {
	appDataDir, err := config.AppDataDirForEnv(env)
	if err != nil {
		return errors.Wrap(err, "could not find the app data directory")
	}

	processinfoPath := filepath.Join(appDataDir, ProcessInfoRelPath)
	err = os.Remove(processinfoPath)
	if !os.IsNotExist(err) {
		return errors.Wrap(err, "could not remove the process info file")
	}
	return nil
}

// WriteProcessInfo write the process info in a file
func (s *Server) WriteProcessInfo() error {
	if s.ProcessInfo == nil {
		return errors.New("no data to write")
	}

	processinfoPath := filepath.Join(s.AppDataDir, ProcessInfoRelPath)
	infoFile, err := os.Create(processinfoPath)
	if err != nil {
		return errors.Wrap(err, "could not open the file for writting")
	}
	defer infoFile.Close()

	if err = json.NewEncoder(infoFile).Encode(s.ProcessInfo); err != nil {
		return errors.Wrap(err, "could not encode the file content")
	}

	return nil
}
