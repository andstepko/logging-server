package logan

import (
	goerrors "github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/logan/v2/errors"
	"gitlab.com/distributed_lab/logan/v2/fields"
)

// F type is for logging fields, used to pass to `WithFields`.
// DEPRECATED
type F map[string]interface{}

var (
	// DEPRECATED
	ErrorKey = logrus.ErrorKey
	// DEPRECATED
	StackKey = "stack"
)

// DEPRECATED
type Entry struct {
	*logrus.Entry
}

// WithRecover creates `go-errors.Error` error from the `recoverData`
// and returns Entry with this error and its stack.
// DEPRECATED
func (e *Entry) WithRecover(recoverData interface{}) *Entry {
	err := errors.FromPanic(recoverData)
	goErr := goerrors.Wrap(err, 4)
	return e.WithStack(goErr).WithError(goErr)
}

// DEPRECATED
func (e *Entry) WithError(err error) *Entry {
	type fieldsProvider interface {
		GetFields() errors.F
	}

	fieldedError, ok := err.(fieldsProvider)
	if ok {
		e = &Entry{
			e.Entry.WithFields(logrus.Fields(fieldedError.GetFields())),
		}
	}

	return &Entry{
		Entry: e.Entry.WithError(err),
	}
}

// DEPRECATED
func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(fields.Obtain(key, value))
}

// DEPRECATED
func (e *Entry) WithFields(fields F) *Entry {
	return &Entry{e.Entry.WithFields(logrus.Fields(fields))}
}

// DEPRECATED
func (e *Entry) WithStack(err error) *Entry {
	return e.WithField(StackKey, getStack(err))
}

// Debugf logs a message at the debug severity.
// DEPRECATED
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.Entry.Debugf(format, args...)
}

// Debug logs a message at the debug severity.
// DEPRECATED
func (e *Entry) Debug(args ...interface{}) {
	e.Entry.Debug(args...)
}

// Infof logs a message at the Info severity.
// DEPRECATED
func (e *Entry) Infof(format string, args ...interface{}) {
	e.Entry.Infof(format, args...)
}

// Info logs a message at the Info severity.
// DEPRECATED
func (e *Entry) Info(args ...interface{}) {
	e.Entry.Info(args...)
}

// Warnf logs a message at the Warn severity.
// DEPRECATED
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.Entry.Warnf(format, args...)
}

// Warn logs a message at the Warn severity.
// DEPRECATED
func (e *Entry) Warn(args ...interface{}) {
	e.Entry.Warn(args...)
}

// Errorf logs a message at the Error severity.
// DEPRECATED
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.Entry.Errorf(format, args...)
}

// Error logs a message at the Error severity.
// DEPRECATED
func (e *Entry) Error(args ...interface{}) {
	e.Entry.Error(args...)
}

// Panicf logs a message at the Panic severity.
// DEPRECATED
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.Entry.Panicf(format, args...)
}

// Panic logs a message at the Panic severity.
// DEPRECATED
func (e *Entry) Panic(args ...interface{}) {
	e.Entry.Panic(args...)
}

// getStack returns the stack, as a string, if one can be extracted from `err`.
// If provided `err` is NOT of go-errors.Error type - returns "unknown".
// DEPRECATED
func getStack(err error) string {
	type stackProvider interface {
		Stack() []byte
	}

	if s, ok := err.(stackProvider); ok {
		return string(s.Stack())
	}

	return "unknown"
}
