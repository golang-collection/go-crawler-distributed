package unifiedLog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-21 09:12
* @Description: 方便后期进行统一配置
**/

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *zap.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init(){
	hook := &lumberjack.Logger{
		Filename:   getLogFileFullPath(), // 日志文件路径
		MaxSize:    500,            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,              // 日志文件最多保存多少个备份
		MaxAge:     28,             // 文件最多保存多少天
		Compress:   true,           // 是否压缩
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),                       // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), // 打印到控制台和文件
		zap.InfoLevel, // 日志级别
	)

	logger = zap.New(core, zap.AddCaller())
}

func GetLogger() *zap.Logger {
	return logger
}