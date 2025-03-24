package io_test

import (
	"errors"
	"os"
	"splash/internal/io"
	"testing"
)

func TestReadFile(t *testing.T) {
	path := getTempFile()
	res, err := io.ReadFile(path)

	if !errors.Is(err, nil) {
		t.Errorf("correct file should not yield an error, got: %v", err)
	}

	if string(res) != string("foo") {
		t.Errorf("result '%s' does not match expected '%s'", string(res), string("foo"))
	}
}

func TestReadEmptyPath(t *testing.T) {
	_, err := io.ReadFile("")
	if !errors.Is(err, io.ErrEmptyPath) {
		t.Errorf("resulting error '%v' does not match expected '%v", err, io.ErrEmptyPath)
	}
}

func TestReadNonExistentPath(t *testing.T) {
	_, err := io.ReadFile("bad")
	expected := io.ErrFailedToReadFile("bad")
	if err.Error() != expected.Error() {
		t.Errorf("resulting error '%v' does not match expected '%v", err, expected)
	}
}

func TestWriteToFile(t *testing.T) {
	dir := os.TempDir()
	contents := "foo"
	path := dir + "/splash-test"
	err := io.WriteToFile(path, []byte(contents))
	if !errors.Is(err, nil) {
		t.Errorf("writing to file should not fail, but got: %v", err)
	}

	got, _ := io.ReadFile(path)
	defer os.Remove(path)
	if string(got) != contents {
		t.Errorf("expected file contents '%s' does not match expected '%s'", string(got), contents)
	}
}

func TestReadEmptyStdin(t *testing.T) {
	path := getEmptyTempFile()
	stdin, _ := os.Open(path)
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = stdin

	_, err := io.ReadStdin()
	if err.Error() != io.ErrEmptyStdin.Error() {
		t.Errorf("resulting error '%v' does not match expected '%v'", err, io.ErrEmptyStdin)
	}
}

func TestBrokenStdin(t *testing.T) {
	stdin, err := os.Open("non-existent")
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	defer stdin.Close()
	os.Stdin = stdin

	_, err = io.ReadStdin()
	if !errors.Is(err, io.ErrIoFailure) {
		t.Errorf("resulting error '%v' does not match expected '%v", err, io.ErrIoFailure)
	}
}

func TestReadStdin(t *testing.T) {
	path := getTempFile()
	stdin, err := os.Open(path)
	if !errors.Is(err, nil) {
		t.Errorf("failed to open file to replace with stdin: %v", err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = stdin

	contents, err := io.ReadStdin()
	if !errors.Is(err, nil) {
		t.Errorf("failed to read stdin: %v", err)
	}

	if string(contents) != "foo" {
		t.Errorf("contents from stdin '%v' does not match expected '%v'", string(contents), "foo")
	}
}

func getTempFile() string {
	tmp, err := os.CreateTemp("", "tmp.json")
	defer tmp.Close()
	if err != nil {
		panic("could not create temporary file")
	}

	_, err = tmp.Write([]byte("foo"))
	if err != nil {
		panic("could not write to temporary file")
	}

	return tmp.Name()
}

func getEmptyTempFile() string {
	tmp, err := os.CreateTemp("", "tmp.json")
	defer tmp.Close()
	if err != nil {
		panic("could not create temporary file")
	}

	return tmp.Name()
}
