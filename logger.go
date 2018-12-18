package errlog

import (
	"github.com/sirupsen/logrus"
)

type (
	// Logger is a logger for error object.
	Logger struct {
		logger LogrusLogger
	}

	// LogrusLogger is an interface for logrus.Logger and logrus.Entry .
	LogrusLogger interface {
		Debug(args ...interface{})
		Debugf(format string, args ...interface{})
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
		Fatal(args ...interface{})
		Fatalf(format string, args ...interface{})
		Info(args ...interface{})
		Infof(format string, args ...interface{})
		Warn(args ...interface{})
		Warnf(format string, args ...interface{})
		WithField(key string, value interface{}) *logrus.Entry
		WithFields(fields logrus.Fields) *logrus.Entry
	}
)

// NewLogger returns a logger.
func NewLogger(logger LogrusLogger) *Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return &Logger{logger: logger}
}

// WithField returns a new logger added given field.
func (logger *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{logger: logger.logger.WithField(key, value)}
}

// WithFields returns a new logger added given fields.
func (logger *Logger) WithFields(fields logrus.Fields) *Logger {
	return &Logger{logger: logger.logger.WithFields(fields)}
}

// Debug outputs debug log.
func (logger *Logger) Debug(err error) {
	if e, ok := err.(Error); ok {
		logger.logger.WithFields(e.Fields()).Debug(e)
		return
	}
	logger.logger.Debug(err)
}

// Error outputs error log.
func (logger *Logger) Error(err error) {
	if e, ok := err.(Error); ok {
		logger.logger.WithFields(e.Fields()).Error(e)
		return
	}
	logger.logger.Error(err)
}

// Fatal outputs fatal log.
func (logger *Logger) Fatal(err error) {
	if e, ok := err.(Error); ok {
		logger.logger.WithFields(e.Fields()).Fatal(e)
		return
	}
	logger.logger.Fatal(err)
}

// Info outputs info log.
func (logger *Logger) Info(err error) {
	if e, ok := err.(Error); ok {
		logger.logger.WithFields(e.Fields()).Info(e)
		return
	}
	logger.logger.Info(err)
}

// Warn outputs warn log.
func (logger *Logger) Warn(err error) {
	if e, ok := err.(Error); ok {
		logger.logger.WithFields(e.Fields()).Warn(e)
		return
	}
	logger.logger.Warn(err)
}
