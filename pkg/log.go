package pkg

import (
	"log"
	"time"

	"github.com/knight0zh/demo_config/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var AccessLogger *zap.Logger
var Logger *zap.Logger

func InitErrorLogger() (err error) {
	writeSyncer := getLogWriter("logs/error", time.Hour*24*3, 2)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte("debug"))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	Logger = zap.New(core, zap.AddCaller())
	return
}

func InitAccessLogger(cfg *config.LogCfg) (err error) {
	writeSyncer := getLogWriter("logs/access", time.Hour*24*3, 3)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	AccessLogger = zap.New(core, zap.AddCaller())
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, rotationTime time.Duration, maxBackups uint) zapcore.WriteSyncer {
	rotateLogger, err := rotatelogs.New(
		filename+"-%Y%m%d.log",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxBackups),
		rotatelogs.WithLinkName(filename+".log"),
	)
	if err != nil {
		log.Fatalf("failed to create rotatelogs: %s", err)
	}
	return zapcore.AddSync(rotateLogger)
}
