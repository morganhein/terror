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
	lock      *sync.RWMutex
	errorType interface{}
}

func initTerror(message string, depth int) *terror {
	t := &terror{
		lock:   &sync.RWMutex{},
		traces: make(map[string]*trace),
	}
	_ = t.insureTrace(message, depth+1)
	return t
}

//Error returns the Terror in string format
func (t *terror) Error() string {
	err := ""
	for k, v := range t.traces {
		err += fmt.Sprintf("--- %s\n%s", k, v.Error())
	}
	return err
}

//SetType sets the type for the overall Terror
func (t *terror) SetType(errorType interface{}) Terror {
	t.setType(errorType, 4)
	return t
}

func (t *terror) setType(errorType interface{}, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	_ = t.insureTrace("", depth)
	t.errorType = errorType
}

//WithField adds the key, value pair to this Terror
func (t *terror) WithField(key string, value interface{}) Terror {
	t.withField(key, value, 4)
	return t
}

func (t *terror) withField(key string, value interface{}, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	tr := t.insureTrace("", depth)
	tr.fields[key] = value
}

//WithError adds the passed error to this Terror
func (t *terror) WithError(err error) Terror {
	t.withError(err, 4)
	return t
}

func (t *terror) withError(err error, depth int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	tr := t.insureTrace("", depth)
	tr.fields["error"] = err
}

//WithTrace adds a stack trace containing of package/function/line number to this Terror
func (t *terror) WithTrace(message string) Terror {
	t.lock.Lock()
	defer t.lock.Unlock()
	_ = t.insureTrace(message, 3)
	return t
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

func (t *terror) insureTrace(message string, depth int) (tr *trace) {
	if t.traces == nil {
		t.traces = make(map[string]*trace)
	}
	pkg, fn, _ := caller(depth)
	// insure the trace exists in the map
	if _, ok := t.traces[pkg+"."+fn]; !ok {
		t.traces[pkg+"."+fn] = &trace{
			fields:  make(map[string]interface{}),
			message: message,
		}
	}
	tr = t.traces[pkg+"."+fn]
	// if the message is blank, don't try to over-write
	if message == "" {
		return
	}
	// set the message if it's different than what was there previously
	if val, ok := t.traces[pkg+"."+fn]; ok && val.message != message {
		val.message = message
		t.traces[pkg+"."+fn] = val
	}
	return
}
