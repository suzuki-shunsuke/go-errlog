// Package errlog is a library for logging error with logrus more easily.
package errlog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// CheckField checks the field's value.
func CheckField(err error, key string, f func(v interface{}) bool) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.CheckField(key, f)
	}
	return false
}

// GetField returns the field value.
// If error is nil or doesn't have the field,
// nil and false are returned.
func GetField(err error, key string) (interface{}, bool) {
	if err == nil {
		return nil, false
	}
	if e, ok := err.(*Error); ok {
		return e.GetField(key)
	}
	return nil, false
}

// HasField returns whether error has the field.
func HasField(err error, key string) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.HasField(key)
	}
	return false
}

// HasMsg returns whether error has the message.
// If err is nil, returns false.
// If err isn't Error, returns err.Error() == msg .
func HasMsg(err error, msg string) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.HasMsg(msg)
	}
	return err.Error() == msg
}

// New is a shorthand of combination of Wrap and fmt.Errorf .
func New(fields logrus.Fields, msg string, msgs ...string) *Error {
	return &Error{
		err: fmt.Errorf(msg), msgs: append([]string{msg}, msgs...), fields: fields}
}

// Newf is a shorthand of combination of New and fmt.Sprintf .
func Newf(fields logrus.Fields, msg string, a ...interface{}) *Error {
	s := fmt.Sprintf(msg, a...)
	return &Error{err: fmt.Errorf(s), msgs: []string{s}, fields: fields}
}

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
		ret := &Error{
			err:    e.err,
			msgs:   append(e.msgs, msgs...),
			fields: e.fields,
		}
		if ret.fields == nil {
			ret.fields = logrus.Fields{}
		}
		for k, v := range fields {
			ret.fields[k] = v
		}
		return ret
	}
	return &Error{
		err: err, msgs: append([]string{err.Error()}, msgs...), fields: fields}
}

// Wrapf is a shordhand of combination of Wrap and fmt.Sprintf .
// If err is nil, returns nil.
func Wrapf(err error, fields logrus.Fields, msg string, a ...interface{}) *Error {
	return Wrap(err, fields, fmt.Sprintf(msg, a...))
}
