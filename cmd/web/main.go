package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"splash/internal/generators"
	"splash/internal/models"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", handlePost)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var generateReq GenerateRequest
	err := json.NewDecoder(r.Body).Decode(&generateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	generator, err := getGenerator(generateReq.Format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var gen generators.Generator = *generator
	colorscheme := string(gen.Generate(&generateReq.Palette))

	fmt.Fprintln(w, colorscheme)
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

type GenerateRequest struct {
	Format  string
	Palette models.Base16Palette
}

func ErrUnsupportedFormat(format string) error {
	return fmt.Errorf("unsupported format '%s'", format)
}
