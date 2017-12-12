package terrors

import (
	"testing"

	"errors"

	"fmt"

	. "github.com/morganhein/terror"
	"github.com/stretchr/testify/assert"
)

func TestInsureTrace(t *testing.T) {
	ter := initTerror("", 2)
	ter.insureTrace("new trace", 2)
	assert.Len(t, ter.traces, 1)
	c := caller(1)
	assert.NotNil(t, ter.traces[c.Full])
	assert.Equal(t, "new trace", ter.traces[c.Full].errMsg)
}

func TestInitTerror(t *testing.T) {
	c := caller(1)
	ter := initTerror("", 2)
	x := fmt.Sprintf("%s.%s", c.Pkg, c.Fn)

	assert.Len(t, ter.traces, 1)
	assert.NotNil(t, ter.traces[x])
	assert.Equal(t, "", ter.traces[x].errMsg)

	c = caller(1)
	ter = initTerror("error message", 2)
	x = fmt.Sprintf("%s.%s", c.Pkg, c.Fn)

	assert.Len(t, ter.traces, 1)
	assert.NotNil(t, ter.traces[x])
	assert.Equal(t, "error message", ter.traces[x].errMsg)
}

func TestTerrorToError(t *testing.T) {
	type (
		Error1 int
		Error2 int
		Error3 int
	)

	testTerrorToErrorSecondary := func(ter Terror) {
		ter.WithError(errors.New("deeper terror error"))
	}

	ter := New("new terror").
		WithError(errors.New("new error in a terror")).
		WithField("key", "value").
		WithField("hello", "world").
		SetType(new(Error1))
	testTerrorToErrorSecondary(ter)

	ter.WithField("later", "terror trace")
	output := fmt.Sprintf("%s", ter)

	assert.Contains(t, output, "new terror")
	assert.Contains(t, output, "new error in a terror")
	assert.Contains(t, output, "value")
	assert.Contains(t, output, "terrors.TestTerrorToError")
	assert.Contains(t, output, "deeper terror error")
	assert.Contains(t, output, "later")
	assert.Contains(t, output, "terror trace")
	//fmt.Println(output)
}
