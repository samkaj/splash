package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"splash/internal/generators"
	"splash/internal/io"
	"splash/internal/models"
)

func main() {
	inputFilePath := flag.String("i", "", "")
	flag.Usage = usage
	flag.Parse()

	var jsonPalette []byte
	var err error
	if *inputFilePath != "" {
		jsonPalette, err = io.ReadFile(*inputFilePath)
	} else {
		jsonPalette, err = io.ReadStdin()
	}
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
	filesCreated := make([]string, 0)
	for idx, generator := range generators {
		contents := generator.Generate(palette)
		ext, err := getFileExtension(outputFormats[idx])
		if err != nil {
			fail(err.Error())
		}

		name := fileName + ext
		err = io.WriteToFile(name, contents)
		if err != nil {
			fail(err.Error())
		}

		filesCreated = append(filesCreated, name)
	}

	fmt.Fprintf(os.Stderr, "splash generated %d files:\n", len(filesCreated))
	for _, name := range filesCreated {
		fmt.Fprintf(os.Stderr, "- %s\n", name)
	}
}

func getGenerator(format string) (*generators.Generator, error) {
	var generator generators.Generator
	var err error
	switch {
	case format == "nvim":
		generator = &generators.NvimGenerator{}
	case format == "ghostty":
		generator = &generators.GhosttyGenerator{}
	case format == "helix":
		generator = &generators.HelixGenerator{}
	case format == "alacritty":
		generator = &generators.AlacrittyGenerator{}
	case format == "kitty":
		generator = &generators.KittyGenerator{}
	case format == "wezterm":
		generator = &generators.WeztermGenerator{}
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
	case format == "helix":
		extension = "-helix.toml"
	case format == "alacritty":
		extension = "-alacritty.toml"
	case format == "kitty":
		extension = "-kitty.conf"
	case format == "wezterm":
		extension = "-wezterm.toml"
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
    <FORMAT>...  Output formats. Supported formats: [ nvim, ghostty, helix, alacritty, kitty, wezterm ]

OPTIONS:
    -i  JSON-file containing the palette. When omitted, stdin is used.
    -h  Display this message.
`)
}

func fail(message string) {
	fmt.Fprintf(os.Stderr, "splash: %s. run with -h for usage.\n", message)
	os.Exit(1)
}

var ErrNoFormatsProvided = errors.New("no formats provided")

func ErrUnsupportedFormat(format string) error {
	return fmt.Errorf("unsupported format '%s'", format)
}
