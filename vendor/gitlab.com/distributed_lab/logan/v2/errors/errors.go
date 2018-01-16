package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// FromPanic extracts the err from the result of a recover() call.
// DEPRECATED
func FromPanic(rec interface{}) error {
	err, ok := rec.(error)
	if !ok {
		err = fmt.Errorf("%s", rec)
	}

	return err
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
// DEPRECATED
func New(msg string) error {
	return errors.New(msg)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
// DEPRECATED
func Wrap(base error, msg string) error {
	return errors.Wrap(base, msg)
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
// DEPRECATED
func Cause(err error) error {
	return errors.Cause(err)
}

// GetFields returns the underlying fields of the error and its nested cause-errors, if possible.
// An error value has fields if it (or any of its nested cause) implements the following interface:
//
//     type fieldsProvider interface {
//            GetFields() F
//     }
//
// If the error and all of its nested causes do not implement GetFields, empty fields map will
// be returned.
// DEPRECATED
func GetFields(err error) F {
	type fieldsProvider interface {
		GetFields() F
	}

	type causer interface {
		Cause() error
	}

	f := F{}
	for err != nil {
		fError, ok := err.(fieldsProvider)
		if ok {
			f = f.AddFields(fError.GetFields())
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return f
}

// DEPRECATED
type withFields struct {
	error
	F
}

// DEPRECATED
func (w *withFields) Error() string {
	return w.error.Error()
}

// DEPRECATED
func (w *withFields) GetFields() F {
	return w.F
}

// DEPRECATED
func (w *withFields) Cause() error {
	return w.error
}
