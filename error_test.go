package errlog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_base_Cause(t *testing.T) {
	e := fmt.Errorf("foo")
	err := base{err: e}
	require.Equal(t, e, err.Cause())
	var e2 *base
	require.Nil(t, e2.Cause())
	e3 := base{err: &err}
	require.Equal(t, e, e3.Cause())
}

func Test_base_Error(t *testing.T) {
	err := Wrap(fmt.Errorf("foo"), nil)
	require.Equal(t, "foo", err.Error())
	e := Wrap(err, nil, "bar")
	require.Equal(t, "foo : bar", e.Error())
	var e2 *base
	require.Equal(t, "", e2.Error())
}
