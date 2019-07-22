package errlog

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	// Error is a structured error.
	Error struct {
		err    error
		msgs   []string
		fields logrus.Fields
	}

	causer interface {
		Cause() error
	}
)

// Cause returns a base error.
func (e *Error) Cause() error {
	if e == nil {
		return nil
	}
	if err, ok := e.err.(causer); ok {
		return err.Cause()
	}
	return e.err
}

func join(msgs ...string) string {
	return strings.Join(msgs, " : ")
}

// Error returns a message represents error.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return join(e.msgs...)
}
