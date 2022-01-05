/**
 * File: logger.go
 * Author: Sarnava Mukherjee
 * Contact: (sarnavamukherjee20@gmail.com)
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"os"
	"strings"
)

var serviceName string

//CreateLogger with specific log level
func CreateLogger(logLevel, applicationName string) {

	serviceName = applicationName
	level := getLogLevel(logLevel)
	alevel := zap.NewAtomicLevelAt(level)

	var config zap.Config
	if level == zapcore.DebugLevel {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.DisableCaller = false
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.FunctionKey = "func"
	config.DisableStacktrace = true
	config.DisableCaller = true
	config.Level = alevel
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Encoding = "json"

	logger, err := config.Build()
	zap.ReplaceGlobals(logger)

	if err != nil {
		panic(err.Error())
	}
}

func getLogLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// GetLogger get native not sugared logger
func GetLogger() *zap.Logger {
	return zap.L()
}

// ObserveLogging constructs a logger through the zap/zaptest/observer framework
// so that logs will be accessible in tests.
func ObserveLogging(level zapcore.Level) *observer.ObservedLogs {
	observedLogger, logs := observer.New(level)
	logger := zap.New(observedLogger)
	zap.ReplaceGlobals(logger)
	return logs
}

// Debug logs a debug message with the given fields
func Debug(msg string, fields ...interface{}) {
	zap.S().Debugf(msg, fields...)
}

// Info logs a debug message with the given fields
func Info(msg string, fields ...interface{}) {
	zap.S().Infof(msg, fields...)
}

// Warn logs a debug message with the given fields
func Warn(msg string, fields ...interface{}) {
	zap.S().Warnf(msg, fields...)
}

// Error logs a debug message with the given fields
func Error(msg string, fields ...interface{}) {
	zap.S().Errorf(msg, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(msg string, fields ...interface{}) {
	zap.S().Fatalf(msg, fields...)
	zap.S().Sync()
	os.Exit(1)
}
