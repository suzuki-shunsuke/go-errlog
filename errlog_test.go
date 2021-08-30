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
	var e *base
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "base is nil",
			err:   e,
		},
		{
			title: "not base",
			err:   fmt.Errorf("foo"),
		},
		{
			title: "base doesn't have the key",
			err:   &base{},
		},
		{
			title: "the function returns false",
			err:   &base{fields: logrus.Fields{"foo": 0}},
		},
		{
			title: "the function returns true",
			err:   &base{fields: logrus.Fields{"foo": 1}},
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
	var e *base
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
			title: "err is a base but nil",
			err:   e,
		},
		{
			title: "err is a base",
			err:   &base{fields: logrus.Fields{"foo": "bar"}},
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
	var e *base
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "err is not base",
			err:   fmt.Errorf("foo"),
		},
		{
			title: "err is a base but nil",
			err:   e,
		},
		{
			title: "err is an base",
			err:   &base{fields: logrus.Fields{"foo": "bar"}},
			exp:   true,
		},
		{
			title: "err is a base but doesn't have the field",
			err:   &base{},
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
	var e *base
	data := []struct {
		title string
		err   error
		exp   bool
	}{
		{
			title: "err is nil",
		},
		{
			title: "err isn't a base but has the message",
			err:   fmt.Errorf("foo"),
			exp:   true,
		},
		{
			title: "err isn't base",
			err:   fmt.Errorf("bar"),
		},
		{
			title: "err is a base but nil",
			err:   e,
		},
		{
			title: "err is a base but doesn't have the message",
			err:   &base{},
		},
		{
			title: "err is a base and has the message",
			err:   &base{msgs: []string{"foo"}},
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
	require.Equal(t, &base{err: fmt.Errorf("foo"), msgs: []string{"foo"}}, New(nil, "foo"))
	require.Equal(
		t, &base{err: fmt.Errorf("foo"), msgs: []string{"foo", "bar"}, fields: logrus.Fields{"program": "main"}},
		New(logrus.Fields{"program": "main"}, "foo", "bar"))
}

func TestNewf(t *testing.T) {
	require.Equal(
		t, &base{err: fmt.Errorf("foo bar"), msgs: []string{"foo bar"}},
		Newf(nil, "foo %s", "bar"))
	require.Equal(
		t, &base{
			err: fmt.Errorf("foo"), msgs: []string{"foo"},
			fields: logrus.Fields{"program": "main"},
		},
		Newf(logrus.Fields{"program": "main"}, "foo"))
}

func TestWrap(t *testing.T) {
	var e *base
	data := []struct {
		title  string
		err    error
		fields logrus.Fields
		msgs   []string
		exp    error
	}{{
		title: "err is nil",
	}, {
		title: "err is base but nil",
		err:   e,
	}, {
		title: "err is not base",
		err:   fmt.Errorf("foo"),
		msgs:  []string{"bar"},
		fields: logrus.Fields{
			"foo": "bar",
		},
		exp: &base{
			err:  fmt.Errorf("foo"),
			msgs: []string{"foo", "bar"},
			fields: logrus.Fields{
				"foo": "bar",
			},
		},
	}, {
		title: "err is a base",
		err: &base{
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
		exp: &base{
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
