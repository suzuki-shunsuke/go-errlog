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

// Fields returns structured data of error.
func (e *Error) Fields() logrus.Fields {
	if e == nil {
		return logrus.Fields{}
	}
	if e.fields == nil {
		e.fields = logrus.Fields{}
	}
	return e.fields
}

// Msgs returns messages.
func (e *Error) Msgs() []string {
	if e == nil {
		return []string{}
	}
	if e.msgs == nil {
		e.msgs = []string{}
	}
	return e.msgs
}
