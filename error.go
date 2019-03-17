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
)

// Cause returns a base error.
func (e *Error) Cause() error {
	if e == nil {
		return nil
	}
	return e.err
}

// CheckField checks the field's value.
func (e *Error) CheckField(key string, f func(v interface{}) bool) bool {
	if e == nil {
		return false
	}
	v, ok := e.Fields()[key]
	if ok {
		return f(v)
	}
	return false
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

// GetField returns the field value.
// If error is nil or doesn't have the field,
// nil and false are returned.
func (e *Error) GetField(key string) (interface{}, bool) {
	if e == nil {
		return nil, false
	}
	v, ok := e.Fields()[key]
	return v, ok
}

// HasField returns whether error has the field.
func (e *Error) HasField(key string) bool {
	if e == nil {
		return false
	}
	_, ok := e.Fields()[key]
	return ok
}

// HasMsg returns whether error has the message.
func (e *Error) HasMsg(msg string) bool {
	if e == nil {
		return false
	}
	for _, m := range e.Msgs() {
		if m == msg {
			return true
		}
	}
	return false
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
