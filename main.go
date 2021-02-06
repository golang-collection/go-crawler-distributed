package main

import (
	"context"
	"errors"
	"flag"
	"go-crawler-distributed/global"
	"go-crawler-distributed/initConf"
	"go-crawler-distributed/internal/routers"
	"log"
	"net/http"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-21 20:37
* @Description:
**/
var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	}
	initConf.Init(config)
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := pingServer(); err != nil {
			global.Logger.Errorf(context.Background(), "The server has no response, or it might took too long to start up.")
		}
		global.Logger.Info(context.Background(), "The server has been deployed successfully.")
	}()

	global.Logger.Infof(context.Background(), "Start to listening the incoming requests on http address :%s", global.ServerSetting.HttpPort)
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Fatalf(context.Background(), "start listen server err: %v", err)
	}
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(":" + global.ServerSetting.HttpPort + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		global.Logger.Info(context.Background(), "Waiting for the server, retry in 1 second.")
	}
	return errors.New("cannot connect to the server")
}
