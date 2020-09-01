package watchConfig

import (
	"github.com/spf13/viper"
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
	viper.SetConfigFile("config/config.json") //文件名
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		logger.Error("viper read config error", zap.Error(err))
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
	mqURL := "amqp://" + mqUser + ":" + mqPassword + "@"+mqHost+"/"
	return mqURL, nil
}

func GetElasticUrl() (string, error){
	elasticURL := viper.GetString("elastic.url")
	return elasticURL, nil
}