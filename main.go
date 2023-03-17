package main

import (
	"fmt"
	"github.com/spf13/cobra"
	appCmd "gohub/app/cmd"
	appMake "gohub/app/cmd/make"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.InitConfig(appCmd.Env)

			bootstrap.SetupLogger()

			bootstrap.SetupDB()

			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		appCmd.CmdServe,
		appCmd.CmdKey,
		appCmd.CmdPlay,
		appMake.CmdMake,
		appMake.CmdMakeModel,
		appMake.CmdMakeAPIController,
		appMake.CmdMakeRequest,
	)

	appCmd.RegisterDefaultCmd(rootCmd, appCmd.CmdServe)

	appCmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
