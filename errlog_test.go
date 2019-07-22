package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestCheckField(t *testing.T) {
	f := func(v interface{}) bool {
		return v == 1
	}
	var e *Error
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "Error is nil",
			err:   e,
		},
		{
			title: "not Error",
			err:   fmt.Errorf("foo"),
		},
		{
			title: "Error doesn't have the key",
			err:   &Error{},
		},
		{
			title: "the function returns false",
			err:   &Error{fields: logrus.Fields{"foo": 0}},
		},
		{
			title: "the function returns true",
			err:   &Error{fields: logrus.Fields{"foo": 1}},
			exp:   true,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			if d.exp {
				require.True(t, CheckField(d.err, "foo", f))
				return
			}
			require.False(t, CheckField(d.err, "foo", f))
		})
	}
}

func TestGetField(t *testing.T) {
	var e *Error
	data := []struct {
		title string
		err   error
		expV  interface{}
		expB  bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "err is a normal error",
			err:   fmt.Errorf("foo"),
		},
		{
			title: "err is an Error but nil",
			err:   e,
		},
		{
			title: "err is an Error",
			err:   &Error{fields: logrus.Fields{"foo": "bar"}},
			expV:  "bar",
			expB:  true,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			v, ok := GetField(d.err, "foo")
			if !d.expB {
				require.False(t, ok)
				return
			}
			require.True(t, ok)
			require.Equal(t, d.expV, v)
		})
	}
}

func TestHasField(t *testing.T) {
	var e *Error
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "err is not Error",
			err:   fmt.Errorf("foo"),
		},
		{
			title: "err is an Error but nil",
			err:   e,
		},
		{
			title: "err is an Error",
			err:   &Error{fields: logrus.Fields{"foo": "bar"}},
			exp:   true,
		},
		{
			title: "err is an Error but doesn't have the field",
			err:   &Error{},
			exp:   false,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			if d.exp {
				require.True(t, HasField(d.err, "foo"))
				return
			}
			require.False(t, HasField(d.err, "foo"))
		})
	}
}

func TestHasMsg(t *testing.T) {
	var e *Error
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "err isn't an Error but has the message",
			err:   fmt.Errorf("foo"),
			exp:   true,
		},
		{
			title: "err isn't Error",
			err:   fmt.Errorf("bar"),
		},
		{
			title: "err is an Error but nil",
			err:   e,
		},
		{
			title: "err is an Error but doesn't have the message",
			err:   &Error{},
		},
		{
			title: "err is an Error and has the message",
			err:   &Error{msgs: []string{"foo"}},
			exp:   true,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			if d.exp {
				require.True(t, HasMsg(d.err, "foo"))
				return
			}
			require.False(t, HasMsg(d.err, "foo"))
		})
	}
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
	var e *Error
	data := []struct {
		title  string
		err    error
		fields logrus.Fields
		msgs   []string
		exp    error
	}{{
		title: "err is nil",
	}, {
		title: "err is Error but nil",
		err:   e,
	}, {
		title: "err is not Error",
		err:   fmt.Errorf("foo"),
		msgs:  []string{"bar"},
		fields: logrus.Fields{
			"foo": "bar",
		},
		exp: &Error{
			err:  fmt.Errorf("foo"),
			msgs: []string{"foo", "bar"},
			fields: logrus.Fields{
				"foo": "bar",
			},
		},
	}, {
		title: "err is an Error",
		err: &Error{
			err: fmt.Errorf("foo"),
			fields: logrus.Fields{
				"foo": "bar",
				"zoo": "1",
			},
			msgs: []string{"foo"},
		},
		fields: logrus.Fields{
			"foo": "goo",
		},
		msgs: []string{"zoo"},
		exp: &Error{
			err:  fmt.Errorf("foo"),
			msgs: []string{"foo", "zoo"},
			fields: logrus.Fields{
				"foo": "goo",
				"zoo": "1",
			},
		},
	}}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			require.Equal(t, d.exp, Wrap(d.err, d.fields, d.msgs...))
		})
	}
}

func TestWrapf(t *testing.T) {
	e := Wrapf(fmt.Errorf("foo"), nil, "failed to %s", "get users")
	e2 := Wrap(fmt.Errorf("foo"), nil, fmt.Sprintf("failed to %s", "get users"))
	require.Equal(t, e, e2)
	require.Nil(t, Wrapf(nil, nil, "failed to %s", "foo"))
}
