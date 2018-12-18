package errlog

import (
	"fmt"
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
	return e.err
}

// Error returns a message represents error.
func (e *Error) Error() string {
	msg := strings.Join(e.msgs, " : ")
	if len(e.msgs) == 0 {
		return e.err.Error()
	}
	return fmt.Sprintf("%s : %s", e.err, msg)
}

// Fields returns structured data of error.
func (e *Error) Fields() logrus.Fields {
	return e.fields
}

// Msgs returns messages.
func (e Error) Msgs() []string {
	return e.msgs
}
