package terrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaller(t *testing.T) {
	pkg, fn, line := caller(1)
	assert.Equal(t, "github.com/morganhein/terror/terrors", pkg)
	assert.Equal(t, "TestCaller", fn)
	assert.NotEmpty(t, line)
	fmt.Println("")
}

func TestNew(t *testing.T) {
	te := New("new terror")
	assert.NotNil(t, te)
	assert.Len(t, te.Trace(), 1)
	ts := te.Trace()
	pkg, fn, _ := caller(1)
	assert.NotNil(t, ts[pkg+"."+fn])
	assert.Equal(t, "new terror", te.Trace()["error"].Message())
}

func TestWithError(t *testing.T) {
	te := WithError(errors.New("new error"))
	fields := te.Trace()["error"].Fields()
	val, ok := fields["error"].(error)
	assert.True(t, ok)
	assert.Equal(t, "new error", val.Error())
}
