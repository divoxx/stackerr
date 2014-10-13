package stackerr

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestErrorWithStack(t *testing.T) {
	err := New("a beautiful error message")

	if err.Message != "a beautiful error message" {
		t.Error("unexpected error message:", err.Message)
	}

	if !strings.Contains(string(err.Stack), "github.com/divoxx/stackerr.TestErrorWithStack") {
		t.Errorf("unexpected error stack to be correct")
	}
}

func TestStdErrorCompatibility(t *testing.T) {
	err := New("a beautiful error message")

	if err, ok := interface{}(err).(error); !ok {
		t.Error("expected stackerr.Error to implement standard error interface")
	} else {
		if err.Error() != "a beautiful error message" {
			t.Error("expected error message to match 'stackerr.Error'.Message")
		}
	}
}

func TestWrapStdError(t *testing.T) {
	err := Wrap(errors.New("standard error message"))

	if err.Message != "standard error message" {
		t.Error("expected wrapped standard error to keep same message")
	}

	if !strings.Contains(string(err.Stack), "github.com/divoxx/stackerr.TestWrapStdError") {
		t.Errorf("unexpected wrapped error stack to be correct")
	}
}

func TestWrapNilError(t *testing.T) {
	err := Wrap(nil)

	if err != nil {
		t.Error("expected wrapped nil error to be nil")
	}
}

func TestWrapStackerr(t *testing.T) {
	err := New("an error")
	werr := Wrap(err)

	if !bytes.Equal(err.Stack, werr.Stack) {
		t.Error("expected re-wrapping an stackerr to not chang stack")
	}
}
