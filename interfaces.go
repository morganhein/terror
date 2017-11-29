package terror

import "github.com/morganhein/terror/stack"

//Terror is the main exported interface for interacting with Terrors.
type Terror interface {
	error

	SetType(errorType interface{}) Terror
	WithField(key string, value interface{}) Terror
	WithError(error) Terror
	WithTrace(message string) Terror

	IsType(other interface{}) bool
	ErrorType() string
	Trace() map[string]stack.Trace
}
