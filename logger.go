package errlog

import (
	"fmt"

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

// WithField returns a new logger added given fields and messages.
func (logger *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		logger: logger.logger.WithField(key, value),
	}
}

// WithFields returns a new logger added given fields and messages.
func (logger *Logger) WithFields(fields logrus.Fields) *Logger {
	return &Logger{
		logger: logger.logger.WithFields(fields),
	}
}

// Logger returns a logger.
func (logger *Logger) Logger() LogrusLogger {
	return logger.logger
}

func (logger *Logger) debug(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return
		}
		logger.logger.WithFields(e.Fields()).Debug(e)
	}
	logger.logger.Debug(err)
}

// Sdebug outputs debug log.
func (logger *Logger) Sdebug(msgs ...string) {
	logger.logger.Debug(join(msgs...))
}

// Sdebugf outputs debug log.
func (logger *Logger) Sdebugf(msg string, a ...interface{}) {
	logger.logger.Debugf(msg, a...)
}

// Sfatal outputs fatal log.
func (logger *Logger) Sfatal(msgs ...string) {
	logger.logger.Fatal(join(msgs...))
}

// Sfatalf outputs fatal log.
func (logger *Logger) Sfatalf(msg string, a ...interface{}) {
	logger.logger.Fatalf(msg, a...)
}

// Swarn outputs warn log.
func (logger *Logger) Swarn(msgs ...string) {
	logger.logger.Warn(join(msgs...))
}

// Swarnf outputs warn log.
func (logger *Logger) Swarnf(msg string, a ...interface{}) {
	logger.logger.Warnf(msg, a...)
}

// Sinfo outputs info log.
func (logger *Logger) Sinfo(msgs ...string) {
	logger.logger.Info(join(msgs...))
}

// Sinfof outputs info log.
func (logger *Logger) Sinfof(msg string, a ...interface{}) {
	logger.logger.Infof(msg, a...)
}

// Serror outputs error log.
func (logger *Logger) Serror(msgs ...string) {
	logger.logger.Error(join(msgs...))
}

// Serrorf outputs error log.
func (logger *Logger) Serrorf(msg string, a ...interface{}) {
	logger.logger.Errorf(msg, a...)
}

// Debug outputs debug log.
// If err is nil, do nothing.
func (logger *Logger) Debug(err error, msgs ...string) {
	logger.debug(Wrap(err, nil, msgs...))
}

// Debugf outputs debug log.
// If err is nil, do nothing.
func (logger *Logger) Debugf(err error, msg string, a ...interface{}) {
	logger.Debug(err, fmt.Sprintf(msg, a...))
}

func (logger *Logger) err(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return
		}
		logger.logger.WithFields(e.Fields()).Error(e)
	}
	logger.logger.Error(err)
}

// Error outputs error log.
// If err is nil, do nothing.
func (logger *Logger) Error(err error, msgs ...string) {
	logger.err(Wrap(err, nil, msgs...))
}

// Errorf outputs fatal log.
// If err is nil, do nothing.
func (logger *Logger) Errorf(err error, msg string, a ...interface{}) {
	logger.Error(err, fmt.Sprintf(msg, a...))
}

func (logger *Logger) fatal(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return
		}
		logger.logger.WithFields(e.Fields()).Fatal(e)
	}
	logger.logger.Fatal(err)
}

// Fatal outputs fatal log.
// If err is nil, do nothing.
func (logger *Logger) Fatal(err error, msgs ...string) {
	logger.fatal(Wrap(err, nil, msgs...))
}

// Fatalf outputs fatal log.
// If err is nil, do nothing.
func (logger *Logger) Fatalf(err error, msg string, a ...interface{}) {
	logger.Fatal(err, fmt.Sprintf(msg, a...))
}

func (logger *Logger) info(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return
		}
		logger.logger.WithFields(e.Fields()).Info(e)
		return
	}
	logger.logger.Info(err)
}

// Info outputs info log.
// If err is nil, do nothing.
func (logger *Logger) Info(err error, msgs ...string) {
	logger.info(Wrap(err, nil, msgs...))
}

// Infof outputs info log.
// If err is nil, do nothing.
func (logger *Logger) Infof(err error, msg string, a ...interface{}) {
	logger.Info(err, fmt.Sprintf(msg, a...))
}

func (logger *Logger) warn(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return
		}
		logger.logger.WithFields(e.Fields()).Warn(e)
		return
	}
	logger.logger.Warn(err)
}

// Warn outputs warn log.
// If err is nil, do nothing.
func (logger *Logger) Warn(err error, msgs ...string) {
	logger.warn(Wrap(err, nil, msgs...))
}

// Warnf outputs warn log.
// If err is nil, do nothing.
func (logger *Logger) Warnf(err error, msg string, a ...interface{}) {
	logger.Warn(err, fmt.Sprintf(msg, a...))
}
