package iifs

import (
	"fmt"
)

type InternalError interface {
	error
	InternalError() string
}

func errInternalError(msg string) error {
	var e InternalError = &internalInternalError{
		msg: msg,
	}

	return e
}

func errInternalErrorf(format string, a ...interface{}) error {
	return errInternalError(fmt.Sprintf(format, a...))
}

type internalInternalError struct {
	msg string
}

func (receiver internalInternalError) Error() string {
	return fmt.Sprintf("iifs: Internal Error: %s", receiver.msg)
}

func (receiver internalInternalError) InternalError() string {
	return receiver.msg
}
