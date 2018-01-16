package logan

import (
	"fmt"
	"github.com/pkg/errors"
)

// DEPRECATED
// FromPanic extracts the err from the result of a recover() call.
func FromPanic(rec interface{}) error {
	err, ok := rec.(error)
	if !ok {
		err = fmt.Errorf("%s", rec)
	}

	return err
}

// DEPRECATED
type FieldedErrorI interface {
	// DEPRECATED
	Error() string
	// DEPRECATED
	Fields() F
	// DEPRECATED
	WithField(key string, value interface{}) FieldedErrorI
	// DEPRECATED
	WithFields(fields F) FieldedErrorI
}

// DEPRECATED
type Stackable interface {
	Stack() []byte
}

// DEPRECATED
// If base is nil, Wrap returns nil.
func Wrap(base error, msg string) FieldedErrorI {
	if base == nil {
		return nil
	}

	fieldedError, ok := base.(*FError)
	if !ok {
		fieldedError = &FError{
			err:    base,
			fields: F{},
		}
	}

	fieldedError.err = errors.Wrap(fieldedError.err, msg)
	return fieldedError
}

// DEPRECATED
func NewError(msg string) FieldedErrorI {
	return &FError{
		err:    errors.New(msg),
		fields: F{},
	}
}

// DEPRECATED
type FError struct {
	err    error
	fields F
}

// DEPRECATED
func (e *FError) Error() string {
	return e.err.Error()
}

// DEPRECATED
func (e *FError) Fields() F {
	return e.fields
}

// DEPRECATED
// WithField returns the same instance
func (e *FError) WithField(key string, value interface{}) FieldedErrorI {
	if e == nil {
		return nil
	}

	fieldedEntity, ok := value.(FieldedEntityI)

	if ok {
		return e.WithFields(obtainFields(key, fieldedEntity))
	}

	// It's just a plain field.
	e.fields[key] = value
	return e
}

// DEPRECATED
// WithFields returns the same instance
func (e *FError) WithFields(fields F) FieldedErrorI {
	if e == nil {
		return nil
	}

	for key, value := range fields {
		e.fields[key] = value
	}
	return e
}

// DEPRECATED
func (e *FError) Cause() error {
	return errors.Cause(e.err)
}
