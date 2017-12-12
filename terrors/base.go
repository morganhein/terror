package terrors

import (
	"runtime"
	"strings"

	. "github.com/morganhein/terror"
	"github.com/morganhein/terror/stack"
)

//New creates a new Terror with the passed error message.
func New(message string) Terror {
	t := initTerror(message, 3)
	return t
}

//WithError creates a new Terror with the passed error.
func WithError(err error) Terror {
	t := initTerror("", 3)
	t.withError(err, 3)
	return t
}

//WithField creates a new Terror with the passed key, value pair.
func WithField(key string, value interface{}) Terror {
	t := initTerror("", 3)
	t.withField(key, value, 3)
	return t
}

//SetType creates a new Terror with the passed errorType.
func SetType(errorType interface{}) Terror {
	t := initTerror("", 3)
	t.setType(errorType, 3)
	return t
}

func caller(depth int) (c *stack.Call) {
	if pc, _, l, ok := runtime.Caller(depth); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			pkg, fn := splitPkgFn(f.Name())
			c = &stack.Call{
				Full: f.Name(),
				Pkg:  pkg,
				Fn:   fn,
				Line: l,
			}
			return
		}
	}
	return
}

func splitPkgFn(location string) (pkg, fn string) {
	i := strings.LastIndex(location, ".")
	if i == -1 {
		return
	}
	pkg = location[0:i]
	fn = location[i+1:]
	return
}
