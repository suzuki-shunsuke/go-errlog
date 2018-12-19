// Package errlog is a library for logging error with logrus more easily.
package errlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Wrap returns an error added fields and msgs.
// If err is nil, returns nil.
func Wrap(err error, fields logrus.Fields, msgs ...string) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		if e == nil {
			return nil
		}
		e.msgs = append(e.msgs, msgs...)
		if e.fields == nil {
			e.fields = logrus.Fields{}
		}
		for k, v := range fields {
			e.fields[k] = v
		}
		return e
	}
	return &Error{err: err, msgs: msgs, fields: fields}
}

// Wrapf is a shordhand of combination of Wrap and fmt.Sprintf .
// If err is nil, returns nil.
func Wrapf(err error, fields logrus.Fields, msg string, a ...interface{}) *Error {
	return Wrap(err, fields, fmt.Sprintf(msg, a...))
}

// New is a shorthand of combination of Wrap and fmt.Errorf .
func New(fields logrus.Fields, msg string, msgs ...string) *Error {
	return &Error{err: fmt.Errorf(msg), msgs: msgs, fields: fields}
}

// Newf is a shorthand of combination of New and fmt.Sprintf .
func Newf(fields logrus.Fields, msg string, a ...interface{}) *Error {
	return &Error{err: fmt.Errorf(msg, a...), fields: fields}
}
