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
		s := len(se.attrs)
		args := make([]any, s+1)
		for i, attr := range se.attrs {
			args[i] = attr
		}
		args[s] = slog.String("error", se.Error())
		return logger.With(args...)
	}
	return logger.With(slog.String("error", se.Error()))
}

func WithAttrs(err error, attrs ...slog.Attr) error {
	if err == nil {
		return nil
	}
	se := &sError{}
	if errors.As(err, &se) {
		se.attrs = append(se.attrs, attrs...)
		return se
	}
	return &sError{
		err:   err,
		attrs: attrs,
	}
}

type sError struct {
	err   error
	attrs []slog.Attr
}

func (e *sError) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

func (e *sError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}
