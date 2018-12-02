package lgr

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// we make sure JSONPrinter implements Printer
var _ Printer = (*JSONPrinter)(nil)

// NewJSONPrinter returns a printer that outputs data as JSON
func NewJSONPrinter() *JSONPrinter {
	return &JSONPrinter{}
}

// JSONPrinter is a printer that outputs data as JSON
type JSONPrinter struct {
	fileWritter io.WriteCloser
}

// UseLogFile will create a log file in the provided directory
func (p *JSONPrinter) UseLogFile(logDirectory string) {
	if p.fileWritter != nil {
		p.fileWritter.Close()
	}
	p.fileWritter = &lumberjack.Logger{
		Filename:   filepath.Join(logDirectory, config.HelperCLILogRelPath),
		MaxSize:    20,
		MaxBackups: 3,
		MaxAge:     30,
	}
}

// CloseLogFile will close a log file if there was one
func (p *JSONPrinter) CloseLogFile() error {
	if p.fileWritter != nil {
		return p.fileWritter.Close()
	}
	return nil
}

// Type returns TypeJSON
func (p *JSONPrinter) Type() Type {
	return TypeJSON
}

// PrintError prints an error as JSON on stderr
func (p *JSONPrinter) PrintError(e error) {
	p.print(os.Stderr, "error", e.Error(), struct {
		Error string    `json:"error"`
		At    time.Time `json:"at"`
	}{Error: e.Error(), At: time.Now().UTC()})
}

// PrintWarning prints a warning message as JSON on stderr
func (p *JSONPrinter) PrintWarning(msg string) {
	p.print(os.Stderr, "warning", msg, struct {
		Warning string    `json:"warning"`
		At      time.Time `json:"at"`
	}{Warning: msg, At: time.Now().UTC()})
}

// PrintInfo prints an info message as JSON on stderr
func (p *JSONPrinter) PrintInfo(msg string) {
	p.print(os.Stderr, "info", msg, struct {
		Info string    `json:"info"`
		At   time.Time `json:"at"`
	}{Info: msg, At: time.Now().UTC()})
}

// PrintSuccess prints a success message on stdout
func (p *JSONPrinter) PrintSuccess(msg string) {
	// let's make sure the string it's not already a json string
	if msg[0] == '{' || msg[0] == '[' {
		os.Stdout.WriteString(msg)
		p.fileWritter.Write([]byte(msg))
		return
	}

	p.print(os.Stdout, "success", msg, struct {
		Success string    `json:"success"`
		At      time.Time `json:"at"`
	}{Success: msg, At: time.Now().UTC()})
}

// print prints a message as JSON on stderr
func (p *JSONPrinter) print(out *os.File, tag, originalData string, object interface{}) {
	var msg string

	resp, err := json.Marshal(object)
	if err != nil {
		// if we could not encode the thing (which should never happen)
		// we can try to hardcode the string (which may result in an invalid
		// json)
		msg = fmt.Sprintf(`{"%s":"%s", "at":"%s"}`,
			tag,
			strings.Replace(originalData, `"`, `\\"`, -1),
			time.Now().UTC().String(),
		)
	} else {
		msg = string(resp)
	}

	msg += "\n"
	out.WriteString(msg)
	p.fileWritter.Write([]byte(msg))
}
