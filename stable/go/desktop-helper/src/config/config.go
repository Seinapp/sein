package config

import "fmt"

const (
	// HelperServerLogRelPath is a relative path to the log file containing
	// all the server logs
	HelperServerLogRelPath = "helper-server.log"

	// HelperCLILogRelPath is a relative path to the log file containing
	// all the cli logs
	HelperCLILogRelPath = "helper-cli.log"

	// AppName contains the production name of the app
	AppName = "Sein"

	// EnvProduction contains the name of the production env
	EnvProduction = "production"
)

func AppNameForEnv(env string) string {
	if env == "" || env == EnvProduction {
		return AppName
	}
	return fmt.Sprintf("%s-%s", AppName, env)
}
