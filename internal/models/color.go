package models

import (
	"strconv"
	"strings"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

// Create a new Color based on RGB components
func NewColor(red, green, blue int) *Color {
	return &Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

// Parse a hex string to a Color (e.g., '#ff00ff') -> (255, 0, 255))
// Supported prefixes are `0x` and `#`.
func ColorFromHexString(hexColor string) (*Color, error) {
	color := strings.Trim(hexColor, " \r\n\t")

	length := len(color)
	if strings.HasPrefix(color, "0x") {
		if length != 8 {
			return nil, ErrInvalidHexLength
		}

		color = strings.TrimPrefix(color, "0x")
	} else if strings.HasPrefix(color, "#") {
		if length != 7 {
			return nil, ErrInvalidHexLength
		}

		color = strings.TrimPrefix(color, "#")
	} else {
		return nil, ErrInvalidHexPrefix
	}

	r, err := hexToBase10(color[0:2])
	if err != nil {
		return nil, err
	}

	g, err := hexToBase10(color[2:4])
	if err != nil {
		return nil, err
	}

	b, err := hexToBase10(color[4:6])
	if err != nil {
		return nil, err
	}

	return NewColor(*r, *g, *b), nil
}

func hexToBase10(hexString string) (*int, error) {
	hex, err := strconv.ParseInt(hexString, 16, 0)
	if err != nil {
		return nil, ErrHexConversionFailure
	}

	i := int(hex)
	return &i, nil
}
