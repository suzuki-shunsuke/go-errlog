package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestError_Cause(t *testing.T) {
	e := fmt.Errorf("foo")
	err := Error{err: e}
	require.Equal(t, e, err.Cause())
	var e2 *Error
	require.Nil(t, e2.Cause())
	e3 := Error{err: &err}
	require.Equal(t, e, e3.Cause())
}

func TestError_Error(t *testing.T) {
	err := Wrap(fmt.Errorf("foo"), nil)
	require.Equal(t, "foo", err.Error())
	e := Wrap(err, nil, "bar")
	require.Equal(t, "foo : bar", e.Error())
	var e2 *Error
	require.Equal(t, "", e2.Error())
}

func TestError_Fields(t *testing.T) {
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

func TestError_Msgs(t *testing.T) {
	msgs := []string{"foo", "bar"}
	err := &Error{err: fmt.Errorf("hello"), msgs: msgs}
	require.Equal(t, msgs, err.Msgs())
	err = nil
	require.Equal(t, []string{}, err.Msgs())
	err = &Error{err: fmt.Errorf("hello")}
	require.Equal(t, []string{}, err.Msgs())
}
