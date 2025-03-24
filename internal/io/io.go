package io

import (
	"io"
	"os"
)

func ReadStdin() ([]byte, error) {
	json, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, ErrEmptyStdin
	}

	return json, nil
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
