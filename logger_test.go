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

func TestLogger_WithField(t *testing.T) {
	logger := NewLogger(nil)
	logger = logger.WithField("foo", "bar")
	logger.Info(fmt.Errorf("hello"))
}

func TestLogger_WithFields(t *testing.T) {
	logger := NewLogger(nil)
	logger = logger.WithFields(logrus.Fields{"foo": "bar"})
	logger.Info(fmt.Errorf("hello"))
}

func TestLogger_Logger(t *testing.T) {
	logger := NewLogger(nil)
	require.NotNil(t, logger.Logger())
}

func TestLogger_debug(t *testing.T) {
	logger := NewLogger(nil)
	logger.debug(nil)
	logger.debug(fmt.Errorf("hello"))
	logger.debug(New(nil, "bar"))
	var e *base
	logger.debug(e)
}

func TestLogger_Sdebug(t *testing.T) {
	logger := NewLogger(nil)
	logger.Sdebug("invalid user name")
}

func TestLogger_Sdebugf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Sdebugf("hello %s", "bob")
}

func TestLogger_Swarn(t *testing.T) {
	logger := NewLogger(nil)
	logger.Swarn("invalid user name")
}

func TestLogger_Swarnf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Swarnf("hello %s", "bob")
}

func TestLogger_Serror(t *testing.T) {
	logger := NewLogger(nil)
	logger.Serror("invalid user name")
}

func TestLogger_Serrorf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Serrorf("hello %s", "bob")
}

func TestLogger_Sinfo(t *testing.T) {
	logger := NewLogger(nil)
	logger.Sinfo("invalid user name")
}

func TestLogger_Sinfof(t *testing.T) {
	logger := NewLogger(nil)
	logger.Sinfof("hello %s", "bob")
}

func TestLogger_Debug(t *testing.T) {
	logger := NewLogger(nil)
	logger.Debug(fmt.Errorf("invalid user name"))
}

func TestLogger_Debugf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Debugf(fmt.Errorf("invalid user name"), "hello %s", "bob")
}

func TestLogger_err(t *testing.T) {
	logger := NewLogger(nil)
	logger.err(nil)
	logger.err(fmt.Errorf("hello"))
	logger.err(New(nil, "bar"))
	var e *base
	logger.err(e)
}

func TestLogger_Error(t *testing.T) {
	logger := NewLogger(nil)
	logger.Error(fmt.Errorf("invalid user name"))
}

func TestLogger_Errorf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Errorf(fmt.Errorf("invalid user name"), "hello %s", "bob")
}

func TestLogger_fatal(t *testing.T) {
	logger := NewLogger(nil)
	logger.fatal(nil)
	var e *base
	logger.fatal(e)
}

func TestLogger_Fatal(t *testing.T) {
	logger := NewLogger(nil)
	logger.Fatal(nil)
}

func TestLogger_Fatalf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Fatalf(nil, "hello %s", "bob")
}

func TestLogger_info(t *testing.T) {
	logger := NewLogger(nil)
	logger.info(nil)
	logger.info(fmt.Errorf("hello"))
	logger.info(New(nil, "bar"))
	var e *base
	logger.info(e)
}

func TestLogger_Info(t *testing.T) {
	logger := NewLogger(nil)
	logger.Info(fmt.Errorf("invalid user name"))
}

func TestLogger_Infof(t *testing.T) {
	logger := NewLogger(nil)
	logger.Infof(fmt.Errorf("invalid user name"), "hello %s", "bob")
}

func TestLogger_warn(t *testing.T) {
	logger := NewLogger(nil)
	logger.warn(nil)
	logger.warn(fmt.Errorf("hello"))
	logger.warn(New(nil, "bar"))
	var e *base
	logger.warn(e)
}

func TestLogger_Warn(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warn(fmt.Errorf("invalid user name"))
}

func TestLogger_Warnf(t *testing.T) {
	logger := NewLogger(nil)
	logger.Warnf(fmt.Errorf("invalid user name"), "hello %s", "bob")
}
