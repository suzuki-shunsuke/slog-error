# slog-error

[![Go Reference](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/slog-error.svg)](https://pkg.go.dev/github.com/suzuki-shunsuke/slog-error)

Go library to embed args into error for [slog](https://pkg.go.dev/log/slog)

## Usage

This library provides only two APIs.

```go
// WithError gets args from err and returns a new logger with err and args.
func WithError(logger *slog.Logger, err error) *slog.Logger
// With returns an error with args.
func With(err error, args ...any) error
```

```go
package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/suzuki-shunsuke/slog-error/slogerr"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	if err := core(); err != nil {
		// time=2023-12-29T22:30:06.992+09:00 level=ERROR msg="command failed" name=mike error="user is not found"
		slogerr.WithError(logger, err).Error("command failed")
	}
}

func core() error {
	if err := getUser(); err != nil {
		return fmt.Errorf("get a user: %w", err)
	}
	return nil
}

func getUser() error {
	return slogerr.With(
		errors.New("user is not found"),
		"name", "mike")
}
```

## LICENSE

[MIT](LICENSE)
