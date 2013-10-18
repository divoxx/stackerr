package stackerr

import (
	"errors"
	"strings"
	"testing"
)

func TestErrorWithStack(t *testing.T) {
	err := New("an beautiful error message")

	if err.Message != "An beautiful error message" {
		t.Error("unexpected error message:", err.Message)
	}

	if !strings.Contains(string(err.Stack), "github.com/divoxx/stackerr.TestErrorWithStack") {
		t.Errorf("unexpected error stack to be correct")
	}
}

func TestStdErrorCompatibility(t *testing.T) {
	err := New("an beautiful error message")

	if err, ok := interface{}(err).(error); !ok {
		t.Error("expected stackerr.Error to implement standard error interface")
	} else {
		if err.Error() != "an beautiful error message" {
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
