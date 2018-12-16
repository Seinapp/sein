package lgr

// Printer is an interface used to write messages in the right format, at
// the right place
type Printer interface {
	// PrintError prints an error on stderr
	PrintError(e error)
	// PrintSuccess prints a success message on stdout
	PrintSuccess(msg string)
	// PrintWarning prints a warning message on stderr
	PrintWarning(msg string)
	// PrintInfo prints a info message on stderr
	PrintInfo(msg string)
	// Type returns the type of the printer
	Type() Type
	// UseLogFile will create a log file in the provided directory
	UseLogFile(logDirectory string)
	// CloseLogFile will close a log file if there was one
	CloseLogFile() error
}
