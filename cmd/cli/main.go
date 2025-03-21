package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"splash/internal/generators"
	"splash/internal/models"
)

func main() {
	inputFilePath := flag.String("i", "", "")
	flag.Usage = usage
	flag.Parse()

	jsonPalette, err := readFile(*inputFilePath)
	if err != nil {
		fail(err.Error())
	}

	palette, err := models.PaletteFromJson(jsonPalette)
	if err != nil {
		fail(err.Error())
	}

	outputFormats := flag.Args()
	if len(outputFormats) < 1 {
		fail(ErrNoFormatsProvided.Error())
	}

	generators := make([]generators.Generator, 0)
	for _, format := range outputFormats {
		gen, err := getGenerator(format)
		if err != nil {
			fail(err.Error())
		}

		generators = append(generators, *gen)
	}

	fileName := "splash"
	for idx, generator := range generators {
		contents := generator.Generate(palette)
		ext, err := getFileExtension(outputFormats[idx])
		if err != nil {
			fail(err.Error())
		}

		err = writeToFile(fileName+ext, contents)
		if err != nil {
			fail(err.Error())
		}
	}
}

func readStdin() ([]byte, error) {
	json, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, ErrEmptyStdin
	}

	return json, nil
}

func readFile(path string) ([]byte, error) {
	if path == "" {
		return readStdin()
	}

	json, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrFailedToReadFile(path)
	}

	return json, nil
}

func writeToFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}

func getGenerator(format string) (*generators.Generator, error) {
	var generator generators.Generator
	var err error
	switch {
	case format == "nvim":
		generator = &generators.NvimGenerator{}
	case format == "ghostty":
		generator = &generators.GhosttyGenerator{}
	default:
		err = ErrUnsupportedFormat(format)
	}

	return &generator, err
}

func getFileExtension(format string) (string, error) {
	var extension string
	var err error
	switch {
	case format == "nvim":
		extension = "-nvim.lua"
	case format == "ghostty":
		extension = "-ghostty.conf"
	default:
		err = ErrUnsupportedFormat(format)
	}

	return extension, err
}

func usage() {
	fmt.Fprintf(os.Stderr, `splash exports base16 palettes to various formats

USAGE: 
    splash [OPTIONS] FORMAT [FORMAT ...]

POSITIONAL ARGUMENTS:
    <FORMAT>...  Output formats. Supported formats: [ nvim, ghostty, helix ]

OPTIONS:
    -i  JSON-file containing the palette. When omitted, stdin is used.
    -h  Display this message.
`)
}

func fail(message string) {
	fmt.Fprintf(os.Stderr, "splash: %s. run with -h for usage.\n", message)
	os.Exit(1)
}

func ErrFailedToReadFile(path string) error {
	return fmt.Errorf("failed to read file '%s'", path)
}

var ErrEmptyStdin = errors.New("failed to read from stdin")

var ErrNoFormatsProvided = errors.New("no formats provided")

func ErrUnsupportedFormat(format string) error {
	return fmt.Errorf("unsupported format '%s'", format)
}
