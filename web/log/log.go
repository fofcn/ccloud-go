package log

import (
	"ccloud/web/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	consoleencoder := getConsoleEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
		zapcore.NewCore(consoleencoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	logger := zap.New(core)
	Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.GetInstance().LoggingConfig.File,
		MaxSize:    config.GetInstance().LoggingConfig.MaxSize,
		MaxAge:     config.GetInstance().LoggingConfig.MaxAge,
		MaxBackups: config.GetInstance().LoggingConfig.MaxBackups,
		LocalTime:  config.GetInstance().LoggingConfig.LocalTime,
		Compress:   config.GetInstance().LoggingConfig.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}
