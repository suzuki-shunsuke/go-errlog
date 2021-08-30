package errlog

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	base struct {
		err    error
		msgs   []string
		fields logrus.Fields
	}

	causer interface {
		Cause() error
	}
)

func (e *base) getFields() logrus.Fields {
	if e == nil {
		return nil
	}
	return e.fields
}

func (e *base) checkField(key string, f func(v interface{}) bool) bool {
	fields := e.getFields()
	v, ok := fields[key]
	if ok {
		return f(v)
	}
	return false
}

func (e *base) getField(key string) (interface{}, bool) {
	fields := e.getFields()
	v, ok := fields[key]
	return v, ok
}

func (e *base) hasField(key string) bool {
	_, ok := e.getFields()[key]
	return ok
}

func (e *base) hasMsg(msg string) bool {
	if e == nil {
		return false
	}
	for _, m := range e.msgs {
		if m == msg {
			return true
		}
	}
	return false
}

// Cause returns a base error.
func (e *base) Cause() error {
	if e == nil {
		return nil
	}
	var err causer
	if errors.As(e.err, &err) {
		return err.Cause() //nolint:wrapcheck
	}
	return e.err
}

func (e *base) Unwrap() error {
	if e == nil || e.err == nil {
		return nil
	}
	return errors.Unwrap(e.err)
}

func join(msgs ...string) string {
	return strings.Join(msgs, " : ")
}

// Error returns a message represents error.
func (e *base) Error() string {
	if e == nil {
		return ""
	}
	return join(e.msgs...)
}
