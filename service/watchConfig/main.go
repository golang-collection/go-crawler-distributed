package watchConfig

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go-crawler-distributed/config"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-20 21:06
* @Description:
**/

var logger = unifiedLog.GetLogger()

func init() {
	err := viper.AddRemoteProvider("consul", config.ConsulURL, config.ConsulConfigPath)
	if err != nil {
		logger.Error("read config",zap.Error(err))
		return
	}
	viper.SetConfigType("json") // Need to explicitly set this to json
	if err := viper.ReadRemoteConfig(); err != nil {
		logger.Error("read config",zap.Error(err))
		return
	}
}

func GetMysqlUrl() (string, error) {
	mysqlHost := viper.GetString("mysql.host")
	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlDBName := viper.GetString("mysql.db_name")
	mysqlURL := mysqlUser + ":" + mysqlPassword + "@(" + mysqlHost + ")/" + mysqlDBName
	return mysqlURL, nil
}

func GetRedisUrl() (string, error) {
	redisURL := viper.GetString("redis.host")
	return redisURL, nil
}

func GetRabbitMQUrl() (string, error) {
	mqHost := viper.GetString("rabbitmq.host")
	mqUser := viper.GetString("rabbitmq.user")
	mqPassword := viper.GetString("rabbitmq.password")
	mqURL := "amqp://" + mqUser + ":" + mqPassword + "@" + mqHost + "/"
	return mqURL, nil
}

func GetElasticUrl() (string, error) {
	elasticURL := viper.GetString("elastic.url")
	return elasticURL, nil
}

func GetElasticIndex() (string, error) {
	elasticURL := viper.GetString("elastic.index")
	return elasticURL, nil
}
