package main

import (
	"os"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/cli/cmd"
)

func main() {
	// because of the use of os.Exit(), defers won't be run
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}
