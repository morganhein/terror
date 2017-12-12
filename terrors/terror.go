package terrors

import (
	"sync"

	"fmt"

	"reflect"

	. "github.com/morganhein/terror"
	"github.com/morganhein/terror/stack"
)

type terror struct {
	traces    map[string]*trace // map[function scope]Trace
	order     []string
	lock      *sync.RWMutex
	errorType interface{}
}

func initTerror(message string, depth int) *terror {
	t := &terror{
		lock:   &sync.RWMutex{},
		traces: make(map[string]*trace),
	}
	_, _ = t.insureTrace(message, depth+1)
	return t
}

//Error returns the Terror in string format
func (t *terror) Error() string {
	err := ""

	// iterate the traces in reverse order, LIFO
	for i := len(t.order) - 1; i >= 0; i-- {
		k := t.order[i]
		v := t.traces[k]
		err += fmt.Sprintf("--- %s\n", k)
		if t.errorType != nil {
			err += fmt.Sprintf("errType: %T\n", t.errorType)
		}
		if len(v.errMsg) > 0 {
			err += fmt.Sprintf("errMsg: %s\n", v.errMsg)
		}

		// iterate the fields added to the trace in FIFO order
		for _, f := range v.fields {
			err += fmt.Sprintf("%v: %s=%s\n", f.CallInfo.Line, f.Name, f.Msg)
		}

	}
	return err
}

//SetType sets the type for the overall Terror
func (t *terror) SetType(errorType interface{}) Terror {
	t.setType(errorType, 3)
	return t
}

func (t *terror) setType(errorType interface{}, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	_, _ = t.insureTrace("", depth+1)
	t.errorType = errorType
}

//WithField adds the key, value pair to this Terror
func (t *terror) WithField(key string, value interface{}) Terror {
	t.withField(key, value, 3)
	return t
}

func (t *terror) withField(key string, value interface{}, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	tr, call := t.insureTrace("", depth+1)
	tr.fields = append(tr.fields, &stack.Entry{
		Name:     key,
		Msg:      value,
		CallInfo: call,
	})
}

//WithError adds the passed error to this Terror
func (t *terror) WithError(err error) Terror {
	t.withError(err, 3)
	return t
}

func (t *terror) withError(err error, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	tr, call := t.insureTrace("", depth+1)
	tr.fields = append(tr.fields, &stack.Entry{
		Name:     "error",
		Msg:      err.Error(),
		CallInfo: call,
	})
}

//Trace returns all the Traces added to this Terror
func (t *terror) Trace() map[string]stack.Trace {
	ts := make(map[string]stack.Trace)
	for k, v := range t.traces {
		ts[k] = v
	}
	return ts
}

//ErrorType returns the string formatted name of this Terror's errorType.
//It will return a blank string if the errorType is null.
func (t *terror) ErrorType() string {
	if t.errorType == nil {
		return ""
	}
	return fmt.Sprintf("%T", t.errorType)
}

//IsType compares if the other object's type is the same as this Terror's errorType.
//If either are nil, this will return false.
func (t *terror) IsType(other interface{}) bool {
	//nil equals nothing, so it can't be equal to anything else
	if t.errorType == nil || other == nil {
		return false
	}
	type1 := reflect.TypeOf(t.errorType)
	type2 := reflect.TypeOf(other)
	return type1 == type2
}

func (t *terror) insureTrace(message string, depth int) (tr *trace, call *stack.Call) {
	if t.traces == nil {
		t.traces = make(map[string]*trace)
	}
	if t.order == nil {
		t.order = make([]string, 0)
	}
	call = caller(depth)
	// insure the trace exists in the map
	if _, ok := t.traces[call.Full]; !ok {
		t.traces[call.Full] = &trace{
			fields: make([]*stack.Entry, 0, 1),
			errMsg: message,
		}
		t.order = append(t.order, call.Full)
	}
	//set return with default value
	tr = t.traces[call.Full]
	// if the message is blank, don't try to over-write
	if message == "" {
		//fmt.Println(t)
		return
	}
	// set the message if it's different than what was there previously
	if val, ok := t.traces[call.Full]; ok && val.errMsg != message {
		val.errMsg = message
		t.traces[call.Full] = val
	}
	return
}
