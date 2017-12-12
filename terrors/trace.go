package terrors

import (
	"fmt"

	"github.com/morganhein/terror/stack"
)

type trace struct {
	errMsg string
	fields []*stack.Entry // map[key]value
}

//Error returns this Trace in string format
func (t *trace) Error() string {
	err := ""
	if len(t.errMsg) > 0 {
		err += fmt.Sprintf("errMsg: %s\n", t.errMsg)
	}
	if len(t.fields) > 0 {
		for _, v := range t.fields {
			err += fmt.Sprintf("%s: %s\n", v.Name, v.Msg)
		}
	}
	return err
}

//Message returns the error errMsg associated to this Trace.
func (t *trace) Message() string {
	return t.errMsg
}

//Fields returns all the key, value pairs associated to this Trace.
func (t *trace) Fields() []*stack.Entry {
	return t.fields
}
