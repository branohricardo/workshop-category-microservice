package logger

import "go.uber.org/zap"

// Logger Define the Logger proxy
type Logger struct {
	internalLogger *zap.SugaredLogger
}

var Log Logger

//New : instantiate a new logger
func New() {
	zapLogger, _ := zap.NewProduction()
	Log = Logger{
		internalLogger: zapLogger.Sugar(),
	}
}

//Info Log debug information
func (l Logger) Info(message string) {
	l.internalLogger.Info(message)
}

//Error Log error message and custom parameters
func (l Logger) Error(err error, args ...interface{}) {
	l.internalLogger.Errorf("Error: %s | Info: %s", err.Error(), args)
}

//Warning Log warning message
func (l Logger) Warning(message string) {
	l.internalLogger.Warn(message)
}

//Fatal Log fatal error message with arguments
func (l Logger) Fatal(err error, args ...interface{}) {
	l.internalLogger.Fatalf("Fatal Error: %s | Info: %s", err.Error(), args)
}
