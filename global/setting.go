package global

import (
	"go-crawler-distributed/pkg/logger"
	"go-crawler-distributed/pkg/setting"
)

/**
* @Author: super
* @Date: 2020-09-18 08:32
* @Description: 全局配置包括：服务，数据库，Email，JWT和日志
**/

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	RabbitMQSetting *setting.RabbitMQSettingS
	ElasticSetting  *setting.ElasticSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	ConsulSetting   *setting.ConsulSettingS
	MongoDBSetting  *setting.MongoDBSettingS
	EtcdSetting     *setting.EtcdSettingS
	TracerSetting   *setting.TracerSettingS
)
