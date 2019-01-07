package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestCheckField(t *testing.T) {
	require.False(t, CheckField(nil, "foo", func(v interface{}) bool {
		return v == 1
	}))
	var e *Error
	require.False(t, CheckField(e, "foo", func(v interface{}) bool {
		return v == 1
	}))
	require.False(t, CheckField(fmt.Errorf("foo"), "foo", func(v interface{}) bool {
		return v == 1
	}))
}

func TestGetField(t *testing.T) {
	_, ok := GetField(nil, "foo")
	require.False(t, ok)
	_, ok = GetField(fmt.Errorf("foo"), "foo")
	require.False(t, ok)
	var e *Error
	_, ok = GetField(e, "foo")
	require.False(t, ok)
	v, ok := GetField(&Error{fields: logrus.Fields{"foo": "bar"}}, "foo")
	require.True(t, ok)
	require.Equal(t, "bar", v)
}

func TestHasField(t *testing.T) {
	require.False(t, HasField(nil, "foo"))
	require.False(t, HasField(fmt.Errorf("foo"), "foo"))
	var e *Error
	require.False(t, HasField(e, "foo"))
}

func TestHasMsg(t *testing.T) {
	require.False(t, HasMsg(nil, "foo"))
	require.True(t, HasMsg(fmt.Errorf("foo"), "foo"))
	var e *Error
	require.False(t, HasMsg(e, "foo"))
}

func TestNew(t *testing.T) {
	require.Equal(t, &Error{err: fmt.Errorf("foo"), msgs: []string{"foo"}}, New(nil, "foo"))
	require.Equal(
		t, &Error{err: fmt.Errorf("foo"), msgs: []string{"foo", "bar"}, fields: logrus.Fields{"program": "main"}},
		New(logrus.Fields{"program": "main"}, "foo", "bar"))
}

func TestNewf(t *testing.T) {
	require.Equal(
		t, &Error{err: fmt.Errorf("foo bar"), msgs: []string{"foo bar"}},
		Newf(nil, "foo %s", "bar"))
	require.Equal(
		t, &Error{
			err: fmt.Errorf("foo"), msgs: []string{"foo"},
			fields: logrus.Fields{"program": "main"}},
		Newf(logrus.Fields{"program": "main"}, "foo"))
}

func TestWrap(t *testing.T) {
	data := []struct {
		err       error
		fields    logrus.Fields
		msgs      []string
		expFields logrus.Fields
		expMsgs   []string
	}{{
		fmt.Errorf("foo"), nil, nil, logrus.Fields{}, []string{"foo"},
	}, {
		&Error{err: fmt.Errorf("foo"), msgs: []string{"foo"}},
		logrus.Fields{"foo": "bar"}, nil, logrus.Fields{"foo": "bar"}, []string{"foo"},
	}}
	for _, d := range data {
		err := Wrap(d.err, d.fields, d.msgs...)
		require.NotNil(t, err)
		if e, ok := err.(*Error); ok {
			require.Equal(t, d.expFields, e.Fields())
			require.Equal(t, d.expMsgs, e.Msgs())
		}
	}
	require.Nil(t, Wrap(nil, nil, "foo"))
	require.Nil(t, Wrap(Wrap(nil, nil, "bar"), nil, "foo"))
}

func TestWrapf(t *testing.T) {
	e := Wrapf(fmt.Errorf("foo"), nil, "failed to %s", "get users")
	e2 := Wrap(fmt.Errorf("foo"), nil, fmt.Sprintf("failed to %s", "get users"))
	require.Equal(t, e, e2)
	require.Nil(t, Wrapf(nil, nil, "failed to %s", "foo"))
}
