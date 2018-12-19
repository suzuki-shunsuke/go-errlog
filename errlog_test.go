package errlog

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	data := []struct {
		err       error
		fields    logrus.Fields
		msgs      []string
		expFields logrus.Fields
		expMsgs   []string
	}{{
		fmt.Errorf("foo"), nil, nil, logrus.Fields{}, []string{},
	}, {
		&Error{err: fmt.Errorf("foo")}, logrus.Fields{"foo": "bar"}, nil, logrus.Fields{"foo": "bar"}, []string{},
	}}
	for _, d := range data {
		e := Wrap(d.err, d.fields, d.msgs...)
		require.NotNil(t, e)
		require.Equal(t, d.expFields, e.Fields())
		require.Equal(t, d.expMsgs, e.Msgs())
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

func TestNew(t *testing.T) {
	require.Equal(t, &Error{err: fmt.Errorf("foo")}, New(nil, "foo"))
	require.Equal(
		t, &Error{err: fmt.Errorf("foo"), msgs: []string{"bar"}, fields: logrus.Fields{"program": "main"}},
		New(logrus.Fields{"program": "main"}, "foo", "bar"))
}

func TestNewf(t *testing.T) {
	require.Equal(t, &Error{err: fmt.Errorf("foo bar")}, Newf(nil, "foo %s", "bar"))
	require.Equal(
		t, &Error{err: fmt.Errorf("foo"), msgs: nil, fields: logrus.Fields{"program": "main"}},
		Newf(logrus.Fields{"program": "main"}, "foo"))
}
