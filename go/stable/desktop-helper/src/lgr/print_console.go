package lgr

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// we make sure ConsolePrinter implements Printer
var _ Printer = (*ConsolePrinter)(nil)

// NewConsolePrinter returns a printer that outputs data readable
// in a console
func NewConsolePrinter() *ConsolePrinter {
	return &ConsolePrinter{}
}

// ConsolePrinter is a printer that outputs data as JSON
type ConsolePrinter struct {
	fileWritter io.WriteCloser
}

// UseLogFile will create a log file in the provided directory
func (p *ConsolePrinter) UseLogFile(logDirectory string) {
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
func (p *ConsolePrinter) CloseLogFile() error {
	if p.fileWritter != nil {
		return p.fileWritter.Close()
	}
	return nil
}

// Type returns TypeConsole
func (p *ConsolePrinter) Type() Type {
	return TypeConsole
}

// PrintError prints an error on stderr
func (p *ConsolePrinter) PrintError(e error) {
	p.print("Error", e.Error(), os.Stderr)
}

// PrintWarning prints a warning message on stderr
func (p *ConsolePrinter) PrintWarning(msg string) {
	p.print("Warning", msg, os.Stderr)
}

// PrintInfo prints an info message on stderr
func (p *ConsolePrinter) PrintInfo(msg string) {
	p.print("Info", msg, os.Stderr)
}

// PrintSuccess prints a success message on stdout
func (p *ConsolePrinter) PrintSuccess(msg string) {
	p.print("", msg, os.Stdout)
}

// print prints a message on stderr
func (p *ConsolePrinter) print(tag, msg string, out *os.File) {
	// Since the tag is optional, if we have one we want
	// Tag: message_to_log
	// otherwise we want
	// mesage_to_log
	fullTag := ""
	if tag != "" {
		fullTag = tag + ": "
	}

	data := fmt.Sprintf("%s%s\n", fullTag, msg)
	out.WriteString(data)

	fileData := []byte(fmt.Sprintf("[%s] %s", time.Now().UTC().String(), data))
	p.fileWritter.Write(fileData)
}
