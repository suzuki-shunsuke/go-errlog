package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

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
	logger.Debug(fmt.Errorf("hello"))
	logger.Debug(New(nil, "bar"))
}

func TestLoggerError(t *testing.T) {
	logger := NewLogger(nil)
	logger.Error(fmt.Errorf("hello"))
	logger.Error(New(nil, "bar"))
}

func TestLoggerInfo(t *testing.T) {
	logger := NewLogger(nil)
	logger.Info(fmt.Errorf("hello"))
	logger.Info(New(nil, "bar"))
}

func TestLoggerWarn(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warn(fmt.Errorf("hello"))
	logger.Warn(New(nil, "bar"))
}
