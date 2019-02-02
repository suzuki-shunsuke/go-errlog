package errlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type (
	// Logger is a logger for error object.
	Logger struct {
		logger LogrusLogger
		msgs   []string
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
func NewLogger(logger LogrusLogger, msgs ...string) *Logger {
	if logger == nil {
		logger = logrus.New()
	}
	if msgs == nil {
		msgs = []string{}
	}
	return &Logger{logger: logger, msgs: msgs}
}

// With returns a new logger added given fields and messages.
func (logger *Logger) With(fields logrus.Fields, msgs ...string) *Logger {
	return &Logger{
		logger: logger.logger.WithFields(fields),
		msgs:   append(logger.msgs, msgs...),
	}
}

// Withf returns a new logger added given fields and message.
func (logger *Logger) Withf(fields logrus.Fields, msg string, a ...interface{}) *Logger {
	return &Logger{
		logger: logger.logger.WithFields(fields),
		msgs:   append(logger.msgs, fmt.Sprintf(msg, a...)),
	}
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

// Debug outputs debug log.
// If err is nil, do nothing.
func (logger *Logger) Debug(err error, fields logrus.Fields, msgs ...string) {
	logger.debug(Wrap(err, fields, append(msgs, logger.msgs...)...))
}

// Debugf outputs debug log.
// If err is nil, do nothing.
func (logger *Logger) Debugf(err error, fields logrus.Fields, msg string, a ...interface{}) {
	logger.Debug(err, fields, fmt.Sprintf(msg, a...))
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
func (logger *Logger) Error(err error, fields logrus.Fields, msgs ...string) {
	logger.err(Wrap(err, fields, append(msgs, logger.msgs...)...))
}

// Errorf outputs fatal log.
// If err is nil, do nothing.
func (logger *Logger) Errorf(err error, fields logrus.Fields, msg string, a ...interface{}) {
	logger.Error(err, fields, fmt.Sprintf(msg, a...))
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
func (logger *Logger) Fatal(err error, fields logrus.Fields, msgs ...string) {
	logger.fatal(Wrap(err, fields, append(msgs, logger.msgs...)...))
}

// Fatalf outputs fatal log.
// If err is nil, do nothing.
func (logger *Logger) Fatalf(err error, fields logrus.Fields, msg string, a ...interface{}) {
	logger.Fatal(err, fields, fmt.Sprintf(msg, a...))
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
func (logger *Logger) Info(err error, fields logrus.Fields, msgs ...string) {
	logger.info(Wrap(err, fields, append(msgs, logger.msgs...)...))
}

// Infof outputs info log.
// If err is nil, do nothing.
func (logger *Logger) Infof(err error, fields logrus.Fields, msg string, a ...interface{}) {
	logger.Info(err, fields, fmt.Sprintf(msg, a...))
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
func (logger *Logger) Warn(err error, fields logrus.Fields, msgs ...string) {
	logger.warn(Wrap(err, fields, append(msgs, logger.msgs...)...))
}

// Warnf outputs warn log.
// If err is nil, do nothing.
func (logger *Logger) Warnf(err error, fields logrus.Fields, msg string, a ...interface{}) {
	logger.Warn(err, fields, fmt.Sprintf(msg, a...))
}
