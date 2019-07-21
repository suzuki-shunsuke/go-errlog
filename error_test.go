package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestErrorCause(t *testing.T) {
	msg := "foo"
	err := Error{err: fmt.Errorf(msg)}
	e := err.Cause()
	require.NotNil(t, e)
	require.Equal(t, msg, e.Error())
	var e2 *Error
	require.Nil(t, e2.Cause())
}

func TestErrorError(t *testing.T) {
	err := Wrap(fmt.Errorf("foo"), nil)
	require.Equal(t, "foo", err.Error())
	e := Wrap(err, nil, "bar")
	require.Equal(t, "foo : bar", e.Error())
	var e2 *Error
	require.Equal(t, "", e2.Error())
}

func TestErrorFields(t *testing.T) {
	data := []struct {
		fields logrus.Fields
		exp    logrus.Fields
	}{{
		nil, logrus.Fields{},
	}, {
		logrus.Fields{"foo": "bar"}, logrus.Fields{"foo": "bar"},
	}}
	for _, d := range data {
		err := Error{fields: d.fields}
		require.Equal(t, d.exp, err.Fields())
	}
	var e2 *Error
	require.Equal(t, logrus.Fields{}, e2.Fields())
}

func TestErrorMsgs(t *testing.T) {
	msgs := []string{"foo", "bar"}
	err := &Error{err: fmt.Errorf("hello"), msgs: msgs}
	require.Equal(t, msgs, err.Msgs())
	err = nil
	require.Equal(t, []string{}, err.Msgs())
	err = &Error{err: fmt.Errorf("hello")}
	require.Equal(t, []string{}, err.Msgs())
}
