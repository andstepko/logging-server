package logan

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/logan/v3/fields"
)

var (
	ErrorKey = logrus.ErrorKey
	StackKey = "stack"
)

type Entry struct {
	entry *logrus.Entry
}

// WithRecover creates error from the `recoverData` if it isn't actually an error already
// and returns Entry with this error and its stack.
func (e *Entry) WithRecover(recoverData interface{}) *Entry {
	err := errors.FromPanic(recoverData)
	return e.WithStack(err).WithError(err)
}

func (e *Entry) WithError(err error) *Entry {
	errorFields := errors.GetFields(err)

	return &Entry{
		entry: e.WithFields(errorFields).entry.WithError(err),
	}
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	return e.WithFields(F{key: value})
}

func (e *Entry) WithFields(f F) *Entry {
	expanded := fields.Expand(f)
	return &Entry{e.entry.WithFields(logrus.Fields(expanded))}
}

func (e *Entry) WithStack(err error) *Entry {
	return e.WithField(StackKey, errors.GetStack(err))
}

// Debugf logs a message at the debug severity.
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

// Debug logs a message at the debug severity.
func (e *Entry) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

// Infof logs a message at the Info severity.
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Info logs a message at the Info severity.
func (e *Entry) Info(args ...interface{}) {
	e.entry.Info(args...)
}

// Warnf logs a message at the Warn severity.
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

// Warn logs a message at the Warn severity.
func (e *Entry) Warn(args ...interface{}) {
	e.entry.Warn(args...)
}

// Errorf logs a message at the Error severity.
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Error logs a message at the Error severity.
func (e *Entry) Error(args ...interface{}) {
	e.entry.Error(args...)
}

// Fatalf logs a message at the Error severity.
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

// Fatal logs a message at the Error severity.
func (e *Entry) Fatal(args ...interface{}) {
	e.entry.Fatal(args...)
}

// Panicf logs a message at the Panic severity.
func (e *Entry) Panicf(format string, args ...interface{}) {
	e.entry.Panicf(format, args...)
}

// Panic logs a message at the Panic severity.
func (e *Entry) Panic(args ...interface{}) {
	e.entry.Panic(args...)
}
