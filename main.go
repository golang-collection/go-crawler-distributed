package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-crawler-distributed/router"
	"go-crawler-distributed/router/middleware"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-03 19:02
* @Description:
**/

var port string
var cfg = pflag.StringP("config", "c", "", "config file path")

func main() {
	pflag.Parse()

	//init config
	if err := watchConfig.Init(*cfg); err != nil {
		unifiedLog.GetLogger().Error("init config error", zap.Error(err))
		panic(err)
	}

	port = viper.GetString("addr")

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			unifiedLog.GetLogger().Error("The router has no response, or it might took too long to start up.", zap.Error(err))
		}else{
			unifiedLog.GetLogger().Info("The router has been deployed successfully.")
		}
	}()

	unifiedLog.GetLogger().Info("Start to listening the incoming requests on http address" + port)
	unifiedLog.GetLogger().Info(http.ListenAndServe(port, g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		unifiedLog.GetLogger().Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
