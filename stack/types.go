package stack

//Trace is stack trace containing an error message, and a set of key value
//pairs associated to the Trace
type Trace interface {
	error

	Message() string
	Fields() []*Entry
}

type Entry struct {
	Name     string
	Msg      interface{}
	CallInfo *Call
}

type Call struct {
	Full string
	Pkg  string
	Fn   string
	Line int
}
