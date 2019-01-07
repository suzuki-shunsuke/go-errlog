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

func TestLoggerWithField(t *testing.T) {
	logger := NewLogger(nil)
	logger = logger.WithField("foo", "bar")
	logger.Info(fmt.Errorf("hello"))
}

func TestLoggerWithFields(t *testing.T) {
	logger := NewLogger(nil)
	logger = logger.WithFields(logrus.Fields{"foo": "bar"})
	logger.Info(fmt.Errorf("hello"))
}

func TestLoggerDebug(t *testing.T) {
	logger := NewLogger(nil)
	logger.Debug(nil)
	logger.Debug(fmt.Errorf("hello"))
	logger.Debug(New(nil, "bar"))
	var e *Error
	logger.Debug(e)
}

func TestLoggerError(t *testing.T) {
	logger := NewLogger(nil)
	logger.Error(nil)
	logger.Error(fmt.Errorf("hello"))
	logger.Error(New(nil, "bar"))
	var e *Error
	logger.Error(e)
}

func TestLoggerFatal(t *testing.T) {
	logger := NewLogger(nil)
	logger.Fatal(nil)
	var e *Error
	logger.Fatal(e)
}

func TestLoggerInfo(t *testing.T) {
	logger := NewLogger(nil)
	logger.Info(nil)
	logger.Info(fmt.Errorf("hello"))
	logger.Info(New(nil, "bar"))
	var e *Error
	logger.Info(e)
}

func TestLoggerWarn(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warn(nil)
	logger.Warn(fmt.Errorf("hello"))
	logger.Warn(New(nil, "bar"))
	var e *Error
	logger.Warn(e)
}
