package terrors

import (
	"testing"

	"fmt"

	"errors"

	. "github.com/morganhein/terror"
	"github.com/stretchr/testify/assert"
)

func TestInsureTrace(t *testing.T) {
	ter := initTerror("", 2)
	ter.insureTrace("new trace", 2)
	assert.Len(t, ter.traces, 1)
	pkg, fn, _ := caller(1)
	assert.NotNil(t, ter.traces[pkg+"."+fn])
	assert.Equal(t, "new trace", ter.traces[pkg+"."+fn].message)
}

func TestInitTerror(t *testing.T) {
	ter := initTerror("", 2)
	assert.Len(t, ter.traces, 1)
	pkg, fn, _ := caller(1)
	assert.NotNil(t, ter.traces[pkg+"."+fn])
	assert.Equal(t, "", ter.traces[pkg+"."+fn].message)
}

func TestTerrorToError(t *testing.T) {
	ter := New("new terror").
		WithError(errors.New("new error in a terror")).
		WithField("key", "value").
		WithField("hello", "world").
		SetType("terror error")
	testTerrorToErrorSecondary(ter)
	fmt.Println(ter)
}

func testTerrorToErrorSecondary(ter Terror) {
	ter.WithError(errors.New("deeper terror error"))
}
