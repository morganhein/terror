package terrors

import (
	"fmt"
	"reflect"
)

type trace struct {
	errorType interface{}
	message   string
	fields    map[string]interface{} // map[key]value
}

//Error returns this Trace in string format
func (t *trace) Error() string {
	err := ""
	if t.errorType != nil {
		err += fmt.Sprintf("type: %s\n", reflect.TypeOf(t.errorType))
	}
	if len(t.message) > 0 {
		err += fmt.Sprintf("message: %s\n", t.message)
	}
	if len(t.fields) > 0 {
		for k, v := range t.fields {
			err += fmt.Sprintf("%s: %s\n", k, v)
		}
	}
	return err
}

//Message returns the error message associated to this Trace.
func (t *trace) Message() string {
	return t.message
}

//Fields returns all the key, value pairs associated to this Trace.
func (t *trace) Fields() map[string]interface{} {
	return t.fields
}
