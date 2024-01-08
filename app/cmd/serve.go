package cmd

import (
	"golearning/bootstrap"
	"golearning/pkg/config"
	"golearning/pkg/console"
	"golearning/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:       "serve",
	Short:     "Start web serve",
	Run:       runWeb,
	Args:      cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string){
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	bootstrap.SetupRoute(router)
	err := router.Run(":"+ config.Get("app.port"))
	if err != nil{
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
		
	}
}