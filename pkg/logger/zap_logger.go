package logger

import (
	"go.uber.org/zap"

	"order-app/config"
)

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger(cfg *config.Log) (*ZapLogger, error) {
	lc := zap.NewDevelopmentConfig()
	lc.Encoding = "console"

	if cfg.Level == "prod" {
		lc = zap.NewProductionConfig()
	}

	l, err := lc.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		zap: l,
	}, nil
}

func (l *ZapLogger) Debug(msg interface{}) {
	l.zap.Debug(checkMessage(msg))
}

func (l *ZapLogger) Debugf(msg string, args ...interface{}) {
	l.zap.Sugar().Debugf(msg, args...)
}

func (l *ZapLogger) Info(msg interface{}) {
	l.zap.Info(checkMessage(msg))
}

func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	l.zap.Sugar().Infof(msg, args...)
}

func (l *ZapLogger) Warn(msg interface{}) {
	l.zap.Warn(checkMessage(msg))
}

func (l *ZapLogger) Warnf(msg string, args ...interface{}) {
	l.zap.Sugar().Warnf(msg, args...)
}

func (l *ZapLogger) Error(msg interface{}) {
	l.zap.Error(checkMessage(msg))
}

func (l *ZapLogger) Errorf(msg string, args ...interface{}) {
	l.zap.Sugar().Errorf(msg, args...)
}

func (l *ZapLogger) Fatal(msg interface{}) {
	l.zap.Fatal(checkMessage(msg))
}

func (l *ZapLogger) Fatalf(msg string, args ...interface{}) {
	l.zap.Sugar().Fatalf(msg, args...)
}

func (l *ZapLogger) Close() error {
	return l.zap.Sync()
}

func checkMessage(msg interface{}) string {
	switch msg.(type) {
	case string:
		return msg.(string)
	case error:
		return msg.(error).Error()
	default:
		return ""
	}
}
