/*
Package stackerr provides a richer error object which includes the stack trace for the goroutine that
generated the error. It also is fully compatible with the existing standard error, by allowing
to wrap existing error and implementing the error interface.
*/
package stackerr

import (
	"errors"
	"runtime"
)

const (
	BufferSize = 2048
)

// Error represents an error and is consisted of a simple string message and the stack of the goroutine
// at the moment the error was constructed.
type Err struct {
	error
	Stack []byte
}

// New will create a new Err object with the given message and the current goroutine stack.
// This replaces the standard errors.New function.
func New(msg string) error {
	var buf [BufferSize]byte

	return &Err{
		error: errors.New(msg),
		Stack: buf[:runtime.Stack(buf[:], false)],
	}
}

// Wrap allows you to wrap an existing standard error object. It's important to know that the stack
// will be from the moment this function is called rather than when the error was first created, but
// it allows to a certain level of traceability.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Err); ok {
		return e
	}

	var buf [BufferSize]byte

	return &Err{
		error: err,
		Stack: buf[:runtime.Stack(buf[:], false)],
	}
}

func Stack(err error) string {
	if e, ok := err.(*Err); ok {
		return string(e.Stack)
	} else {
		return ""
	}
}

// NewStack creates a copy of the error with the Stack updated to the current goroutine stack.
func (e Err) NewStack() error {
	var buf [BufferSize]byte

	return &Err{
		error: e.error,
		Stack: buf[:runtime.Stack(buf[:], false)],
	}
}
