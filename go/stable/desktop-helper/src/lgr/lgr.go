// Package lgr contains helper to simplify printing messages
// in the right format
package lgr

// Type represents the type of the printer
type Type int

// List of all printer types
const (
	TypeJSON Type = iota
	TypeConsole
)
