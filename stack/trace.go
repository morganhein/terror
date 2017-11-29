package stack

//Trace is stack trace containing an error message, and a set of key value
//pairs associated to the Trace
type Trace interface {
	error

	Message() string
	Fields() map[string]interface{}
}
