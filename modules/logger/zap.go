package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	log *zap.Logger
}

func newDevelopmentConfig() zap.Config {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapConfig.Sampling = nil
	zapConfig.Development = true
	zapConfig.Encoding = "json"
	return zapConfig
}

func NewZapLogger(debug bool) *ZapLogger {
	var zapConfig zap.Config

	if debug {
		zapConfig = newDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 关闭采样以采集所有日志
	zapConfig.Sampling = nil
	logger, _ := zapConfig.Build(zap.AddStacktrace(zapcore.WarnLevel))
	return &ZapLogger{
		log: logger,
	}
}

func (z *ZapLogger) convertField(fields []LogField) []zap.Field {
	if len(fields) == 0 {
		return nil
	}
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Val)
	}
	return zapFields
}

func (z *ZapLogger) Debug(msg string, fields ...LogField) {
	z.log.Debug(msg, z.convertField(fields)...)
}

func (z *ZapLogger) Info(msg string, fields ...LogField) {
	z.log.Info(msg, z.convertField(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...LogField) {
	z.log.Warn(msg, z.convertField(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...LogField) {
	z.log.Error(msg, z.convertField(fields)...)
}
