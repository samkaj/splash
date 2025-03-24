package io

import (
	"io"
	"os"
)

func ReadStdin() ([]byte, error) {
	contents, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, ErrIoFailure
	}

	if len(contents) <= 0 {
		return nil, ErrEmptyStdin
	}

	return contents, nil
}

func WriteToFile(path string, contents []byte) error {
	return os.WriteFile(path, contents, 0644)
}

func ReadFile(path string) ([]byte, error) {
	if path == "" {
		return nil, ErrEmptyPath
	}

	json, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrFailedToReadFile(path)
	}

	return json, nil
}
