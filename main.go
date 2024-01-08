package main

import (
	"fmt"
    "golearning/app/cmd"
    "golearning/bootstrap"
	"golearning/pkg/config"
	btsConfig "golearning/config"
	"golearning/pkg/console"
	"os"
	"golearning/app/cmd/make"
	

    "github.com/spf13/cobra"
)

func init(){
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:    "Golearning",
		Short:  "A simple forum project",
		Long:   `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
	        bootstrap.SetupRedis()
			bootstrap.SetupCache()
		},
	}
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
		cmd.CmdCacheClear,
		cmd.CmdCacheForget,
	)

	cmd.RegisterGlobalFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil{
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}