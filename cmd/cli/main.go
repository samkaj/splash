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

	for _, generator := range generators {
		fmt.Println(string(generator.Generate(palette)))
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

func getGenerator(format string) (*generators.Generator, error) {
	var generator generators.Generator
	var err error
	switch {
	case format == "nvim":
		generator = &generators.NvimGenerator{}
	default:
		err = ErrUnsupportedFormat(format)
	}

	return &generator, err
}

func usage() {
	fmt.Fprintf(os.Stderr, `splash exports base16 palettes to various formats

USAGE: 
    splash [OPTIONS] FORMAT [FORMAT ...]

POSITIONAL ARGUMENTS:
    <FORMAT>...  Output formats. Supported formats: []

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
