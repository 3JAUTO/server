package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

var (
	l           *zap.Logger
	s           *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

func init() {
	if err := Init("stdout", "debug"); err != nil {
		panic(err)
	}
}

// Init initialize logger with output path and log level
func Init(output string, level string) error {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.DisableCaller = true
	zapCfg.DisableStacktrace = true
	atomicLevel = zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(level)); err != nil {
		return err
	}
	zapCfg.Level = atomicLevel
	zapCfg.Encoding = "console"
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.EncoderConfig.TimeKey = "ts"
	zapCfg.OutputPaths = []string{output}
	logger, err := zapCfg.Build()
	if err != nil {
		return err
	}
	l = logger
	s = l.Sugar()
	return nil
}

// Debug ...
func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	s.Debugf(format, args...)
}

// Debugw ...
func Debugw(msg string, keysAndValues ...interface{}) {
	s.Debugw(msg, keysAndValues...)
}

// Error ...
func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	s.Errorf(format, args...)
}

// Errorw ...
func Errorw(msg string, keysAndValues ...interface{}) {
	s.Errorw(msg, keysAndValues...)
}

// Info ...
func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	s.Infof(format, args...)
}

// Infow ...
func Infow(msg string, keysAndValues ...interface{}) {
	s.Infow(msg, keysAndValues...)
}

// Warn ...
func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	s.Warnf(format, args...)
}

// Warnw ...
func Warnw(msg string, keysAndValues ...interface{}) {
	s.Warnw(msg, keysAndValues...)
}

// Panic ...
func Panic(msg string, fields ...zap.Field) {
	l.Panic(msg, fields...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	s.Panicf(format, args...)
}

// Panicw ...
func Panicw(msg string, keysAndValues ...interface{}) {
	s.Panicw(msg, keysAndValues...)
}

// Fatal ...
func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	s.Fatalf(format, args...)
}

// Fatalw ...
func Fatalw(msg string, keysAndValues ...interface{}) {
	s.Fatalw(msg, keysAndValues...)
}

// dbLogger Database logger writer proxy
type dbLogger struct{}

// Printf impl gorm.logger.Writer
func (l *dbLogger) Printf(msg string, args ...interface{}) {
	all := make([]interface{}, len(args)+1)
	all = append(all, msg)
	all = append(all, args...)
	s.Debug(all...)
}

// GetDBLogger ...
func GetDBLogger() logger.Writer {
	return &dbLogger{}
}

// L ...
func L() *zap.Logger {
	return l
}

// S ...
func S() *zap.SugaredLogger {
	return s
}
