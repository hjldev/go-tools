package logger

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func NewLogger(logName string) {
	writeSyncer := getLogWriter(logName)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	defer logger.Sync()
}

func WriteStr(content string) {
	fmt.Println(content)
	logger.Info(content)
}

func WriteObj(o interface{}) {
	data, _ := json.Marshal(o)
	fmt.Println(string(data))
	logger.Info(string(data))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logName string) zapcore.WriteSyncer {
	if logName == "" {
		logName = "./record.log"
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logName, // 日志输出文件
		MaxSize:    1,       // 日志最大保存1M
		MaxBackups: 5,       // 就日志保留5个备份
		MaxAge:     30,      // 最多保留30个日志 和MaxBackups参数配置1个就可以
		Compress:   false,   // 自导打 gzip包 默认false
	}
	return zapcore.AddSync(lumberJackLogger)
}
