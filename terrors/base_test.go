package terrors

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/morganhein/terror"
	"github.com/stretchr/testify/assert"
)

func TestCaller(t *testing.T) {
	c := caller(1)
	assert.Equal(t, "github.com/morganhein/terror/terrors.TestCaller", c.Full)
	assert.Equal(t, "TestCaller", c.Fn)
	assert.NotEmpty(t, c.Line)
}

func TestCallerOnEmbeddedFunc(t *testing.T) {
	testTerrorToErrorSecondary := func(ter Terror) {
		ter.WithError(errors.New("deeper terror error"))
	}

	ter := New("new terror")
	testTerrorToErrorSecondary(ter)

	//todo: assertions. srsly.
}

func TestNew(t *testing.T) {
	c := caller(1)
	te := New("new terror")
	x := fmt.Sprintf("%s.%s", c.Pkg, c.Fn)

	assert.NotNil(t, te)
	assert.Len(t, te.Trace(), 1)
	ts := te.Trace()
	assert.NotNil(t, ts[x])
	assert.Equal(t, "new terror", ts[x].Message())
}

func TestWithError(t *testing.T) {
	c := caller(1)
	te := WithError(errors.New("new error"))
	x := fmt.Sprintf("%s.%s", c.Pkg, c.Fn)

	fields := te.Trace()[x].Fields()
	assert.Equal(t, "new error", fields[0].Msg)
}
