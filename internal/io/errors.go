package io

import (
	"errors"
	"fmt"
)

func ErrFailedToReadFile(path string) error {
	return fmt.Errorf("failed to read file '%s'", path)
}

var ErrEmptyStdin = errors.New("failed to read from stdin")

var ErrEmptyPath = errors.New("cannot read empty path")

var ErrNoFormatsProvided = errors.New("no formats provided")

func ErrUnsupportedFormat(format string) error {
	return fmt.Errorf("unsupported format '%s'", format)
}
