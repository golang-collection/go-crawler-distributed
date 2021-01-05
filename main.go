package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/cache"
	"go-crawler-distributed/pkg/db"
	"go-crawler-distributed/pkg/elastic"
	"go-crawler-distributed/pkg/idGenerator"
	"go-crawler-distributed/pkg/logger"
	"go-crawler-distributed/pkg/mq"
	"go-crawler-distributed/pkg/setting"
	"go-crawler-distributed/pkg/tracer"

	"gopkg.in/natefinch/lumberjack.v2"
)

/**
* @Author: super
* @Date: 2021-01-05 14:25
* @Description:
**/
var (
	config    string
	isVersion bool
)

func init() {
	//读取命令行参数
	err := setupFlag()
	if err != nil {
		log.Printf("init.setupFlag err: %v\n", err)
	}
	//初始化配置
	err = setupSetting()
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Printf("init setupLogger err: %v\n", err)
	}
	//初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Printf("init setupDBEngine err: %v\n", err)
	}
	//初始化redis
	err = setupCacheEngine()
	if err != nil {
		log.Printf("init setupCacheEngine err: %v\n", err)
	}
	//初始化RabbitMQ
	err = setupRabbitMQEngine()
	if err != nil {
		log.Printf("init setupRabbitMQEngine err: %v\n", err)
	}
	//初始化elastic
	err = setupElasticEngine()
	if err != nil {
		log.Printf("init setupElasticEngine err: %v\n", err)
	}
	//初始化追踪
	err = setupTracer()
	if err != nil {
		log.Printf("init.setupTracer err: %v\n", err)
	}
	//初始化ID生成器
	err = idGenerator.InitSnowflake()
	if err != nil {
		log.Printf("init.snowflak err: %v\n", err)
	}
}

func main() {

}

func setupFlag() error {
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Elastic", &global.ElasticSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupCacheEngine() error {
	var err error
	global.RedisEngine, err = cache.NewRedisEngine(global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupRabbitMQEngine() error {
	var err error
	global.RabbitMQEngine, err = mq.NewRabbitMQEngine(global.RabbitMQSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupElasticEngine() error {
	var err error
	global.ElasticEngine, err = elastic.NewElasticEngine(global.ElasticSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	log.Println("log file name ", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("bedtimeStory", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
