package slogerr

import (
	"errors"
	"log/slog"
)

func WithError(logger *slog.Logger, err error) *slog.Logger {
	if err == nil {
		return logger
	}
	se := &sError{}
	if errors.As(err, &se) {
		return logger.With(se.args...).With("error", err.Error())
	}
	return logger.With("error", err.Error())
}

func With(err error, args ...any) error {
	if err == nil {
		return nil
	}
	se := &sError{}
	if errors.As(err, &se) {
		se.args = append(se.args, args...)
		return se
	}
	return &sError{
		err:  err,
		args: args,
	}
}

type sError struct {
	err  error
	args []any
}

// Error implements the error interface.
// It returns the underlying error's message, or an empty string if e is nil.
func (e *sError) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

// Unwrap returns the underlying error.
// It implements the Unwrap method for Go's error wrapping conventions,
// allowing errors.Is and errors.As to work correctly with wrapped errors.
// Returns nil if e is nil.
func (e *sError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}
