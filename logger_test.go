package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogrusLogger(t *testing.T) {
	// test *logrus.Logger and *logrus.Entry implement LogrusLogger
	NewLogger(&logrus.Logger{})
	NewLogger(&logrus.Entry{})
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger(nil)
	require.NotNil(t, logger.logger)
}

func TestLoggerWith(t *testing.T) {
	logger := NewLogger(nil)
	logger = logger.With(logrus.Fields{"foo": "bar"})
	logger.Info(fmt.Errorf("hello"), nil)
}

func TestLogger_debug(t *testing.T) {
	logger := NewLogger(nil)
	logger.debug(nil)
	logger.debug(fmt.Errorf("hello"))
	logger.debug(New(nil, "bar"))
	var e *Error
	logger.debug(e)
}

func TestLoggerDebug(t *testing.T) {
	logger := NewLogger(nil)
	logger.Debug(fmt.Errorf("invalid user name"), nil)
}

func TestLoggerDebugf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Debugf(fmt.Errorf("invalid user name"), nil, "hello %s", "bob")
}

func TestLogger_err(t *testing.T) {
	logger := NewLogger(nil)
	logger.err(nil)
	logger.err(fmt.Errorf("hello"))
	logger.err(New(nil, "bar"))
	var e *Error
	logger.err(e)
}

func TestLoggerError(t *testing.T) {
	logger := NewLogger(nil)
	logger.Error(fmt.Errorf("invalid user name"), nil)
}

func TestLoggerErrorf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Errorf(fmt.Errorf("invalid user name"), nil, "hello %s", "bob")
}

func TestLogger_fatal(t *testing.T) {
	logger := NewLogger(nil)
	logger.fatal(nil)
	var e *Error
	logger.fatal(e)
}

func TestLoggerFatal(t *testing.T) {
	logger := NewLogger(nil)
	logger.Fatal(nil, nil)
}

func TestLoggerFatalf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Fatalf(nil, nil, "hello %s", "bob")
}

func TestLogger_info(t *testing.T) {
	logger := NewLogger(nil)
	logger.info(nil)
	logger.info(fmt.Errorf("hello"))
	logger.info(New(nil, "bar"))
	var e *Error
	logger.info(e)
}

func TestLoggerInfo(t *testing.T) {
	logger := NewLogger(nil)
	logger.Info(fmt.Errorf("invalid user name"), nil)
}

func TestLoggerInfof(t *testing.T) {
	logger := NewLogger(nil)
	logger.Infof(fmt.Errorf("invalid user name"), nil, "hello %s", "bob")
}

func TestLogger_warn(t *testing.T) {
	logger := NewLogger(nil)
	logger.warn(nil)
	logger.warn(fmt.Errorf("hello"))
	logger.warn(New(nil, "bar"))
	var e *Error
	logger.warn(e)
}

func TestLoggerWarn(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warn(fmt.Errorf("invalid user name"), nil)
}

func TestLoggerWarnf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warnf(fmt.Errorf("invalid user name"), nil, "hello %s", "bob")
}
