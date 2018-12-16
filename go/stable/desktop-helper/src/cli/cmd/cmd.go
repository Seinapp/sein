package cmd

import (
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	"github.com/spf13/cobra"
)

func rootCmd() (*cobra.Command, *string, *bool) {
	cmd := &cobra.Command{
		Use:           "sein-helper",
		Short:         "sein-helper is a gRPC server use to deal with Sein data",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	outputJSON := cmd.PersistentFlags().Bool("json", false, "output json")
	env := cmd.PersistentFlags().String("env", config.EnvProduction, "env of the server")

	// pre hook wont be called if an issue handled by cobra occurs
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return setLogger(*env, *outputJSON)
	}

	cmd.AddCommand(startCmd(env))
	cmd.AddCommand(stopCmd(env))
	return cmd, env, outputJSON
}

// Execute runs and executes the cli command
func Execute() error {
	cmd, env, outputJSON := rootCmd()

	defer func() {
		if cliLogger != nil {
			cliLogger.CloseLogFile()
		}
	}()

	if err := cmd.Execute(); err != nil {
		if cliLogger == nil {
			setLogger(*env, *outputJSON)
		}
		cliLogger.PrintError(err)
		return err
	}

	return nil
}
