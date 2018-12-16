package cmd

import (
	"github.com/pkg/errors"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/lgr"
)

var cliLogger lgr.Printer

func setLogger(env string, outputJSON bool) error {
	cliLogger = lgr.NewConsolePrinter()
	if outputJSON {
		cliLogger = lgr.NewJSONPrinter()
		cliLogger.CloseLogFile()
	}

	appDataDir, err := config.AppDataDirForEnv(env)
	if err != nil {
		return errors.Wrap(err, "could not find the app data directory")
	}
	cliLogger.UseLogFile(appDataDir)
	return nil
}
