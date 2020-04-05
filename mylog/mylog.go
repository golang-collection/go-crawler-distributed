package mylog

import "go.uber.org/zap"

var logger *zap.Logger

func init(){
	var err error
	logger, err = zap.NewProduction()
	if err != nil{
		panic(err)
	}
}

func LogError(msg string, err error){
	logger.Error(msg, zap.String("err", err.Error()))
}

func LogInfo(msg string, info string){
	logger.Info(msg, zap.String("info", info))
}