package io

import (
	"errors"
	"fmt"
)

func ErrFailedToReadFile(path string) error {
	return fmt.Errorf("failed to read file '%s'", path)
}

var ErrIoFailure = errors.New("failed to read from stdin")
var ErrEmptyStdin = errors.New("stdin must not be empty")
var ErrEmptyPath = errors.New("cannot read empty path")
var ErrNoFormatsProvided = errors.New("no formats provided")
