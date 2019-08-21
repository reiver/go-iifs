package iifs

import (
	"testing"
)

func TestInternalInternalErrorAsError(t *testing.T) {

	var err error = internalInternalError{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalInternalErrorAsInternalError(t *testing.T) {

	var complainer InternalError = internalInternalError{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
