package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Base16Palette struct {
	B00 *Color `json:"b00"`
	B01 *Color `json:"b01"`
	B02 *Color `json:"b02"`
	B03 *Color `json:"b03"`
	B04 *Color `json:"b04"`
	B05 *Color `json:"b05"`
	B06 *Color `json:"b06"`
	B07 *Color `json:"b07"`
	B08 *Color `json:"b08"`
	B09 *Color `json:"b09"`
	B0a *Color `json:"b0a"`
	B0b *Color `json:"b0b"`
	B0c *Color `json:"b0c"`
	B0d *Color `json:"b0d"`
	B0e *Color `json:"b0e"`
	B0f *Color `json:"b0f"`
}

func PaletteFromJson(data []byte) (*Base16Palette, error) {
	palette := newPalette()
	err := json.Unmarshal(data, palette)
	if err != nil {
		return nil, err
	}

	return palette, nil
}

func (p *Base16Palette) UnmarshalJSON(data []byte) error {
	var palette any
	if err := json.Unmarshal(data, &palette); err != nil {
		return err
	}

	switch v := palette.(type) {
	case map[string]any:
		builder := NewPaletteBuilder()
		builder = builder.
			Base00(v["b00"]).
			Base01(v["b01"]).
			Base02(v["b02"]).
			Base03(v["b03"]).
			Base04(v["b04"]).
			Base05(v["b05"]).
			Base06(v["b06"]).
			Base07(v["b07"]).
			Base08(v["b08"]).
			Base09(v["b09"]).
			Base0a(v["b0a"]).
			Base0b(v["b0b"]).
			Base0c(v["b0c"]).
			Base0d(v["b0d"]).
			Base0e(v["b0e"]).
			Base0f(v["b0f"])

		pal, err := builder.Build()
		if err != nil {
			return err
		}

		*p = *pal
		return nil
	default:
		return errors.New("json input must be an object")
	}
}

type Base16Builder struct {
	palette Base16Palette
	convert func(string) (*Color, error)
	err     error
}

func NewPaletteBuilder() *Base16Builder {
	pal := newPalette()
	return &Base16Builder{
		palette: *pal,
		convert: ColorFromHexString,
	}
}

func (b *Base16Builder) Build() (*Base16Palette, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	return &b.palette, nil
}

func (b *Base16Builder) setColor(field **Color, value string) *Base16Builder {
	if b.err != nil {
		return b
	}

	color, err := b.convert(value)
	if err != nil {
		b.err = err
		return b
	}

	*field = color
	return b
}

func (b *Base16Builder) Base00(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B00, str)
	}

	return b
}

func (b *Base16Builder) Base01(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B01, str)
	}
	return b
}

func (b *Base16Builder) Base02(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B02, str)
	}
	return b
}

func (b *Base16Builder) Base03(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B03, str)
	}
	return b
}

func (b *Base16Builder) Base04(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B04, str)
	}
	return b
}

func (b *Base16Builder) Base05(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B05, str)
	}
	return b
}

func (b *Base16Builder) Base06(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B06, str)
	}
	return b
}

func (b *Base16Builder) Base07(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B07, str)
	}
	return b
}

func (b *Base16Builder) Base08(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B08, str)
	}
	return b
}

func (b *Base16Builder) Base09(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B09, str)
	}
	return b
}

func (b *Base16Builder) Base0a(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0a, str)
	}
	return b
}

func (b *Base16Builder) Base0b(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0b, str)
	}
	return b
}

func (b *Base16Builder) Base0c(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0c, str)
	}
	return b
}

func (b *Base16Builder) Base0d(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0d, str)
	}
	return b
}

func (b *Base16Builder) Base0e(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0e, str)
	}
	return b
}

func (b *Base16Builder) Base0f(value any) *Base16Builder {
	if str, ok := value.(string); ok {
		return b.setColor(&b.palette.B0f, str)
	}
	return b
}

func newPalette() *Base16Palette {
	return &Base16Palette{
		B00: nil,
		B01: nil,
		B02: nil,
		B03: nil,
		B04: nil,
		B05: nil,
		B06: nil,
		B07: nil,
		B08: nil,
		B09: nil,
		B0a: nil,
		B0b: nil,
		B0c: nil,
		B0d: nil,
		B0e: nil,
		B0f: nil,
	}
}

func (b *Base16Builder) validate() error {
	if b.err != nil {
		return b.err
	}

	missingFields := make([]string, 0)

	if b.palette.B00 == nil {
		missingFields = append(missingFields, "b00")
	}

	if b.palette.B01 == nil {
		missingFields = append(missingFields, "b01")
	}

	if b.palette.B02 == nil {
		missingFields = append(missingFields, "b02")
	}

	if b.palette.B03 == nil {
		missingFields = append(missingFields, "b03")
	}

	if b.palette.B04 == nil {
		missingFields = append(missingFields, "b04")
	}

	if b.palette.B05 == nil {
		missingFields = append(missingFields, "b05")
	}

	if b.palette.B06 == nil {
		missingFields = append(missingFields, "b06")
	}

	if b.palette.B07 == nil {
		missingFields = append(missingFields, "b07")
	}

	if b.palette.B08 == nil {
		missingFields = append(missingFields, "b08")
	}

	if b.palette.B09 == nil {
		missingFields = append(missingFields, "b09")
	}

	if b.palette.B0a == nil {
		missingFields = append(missingFields, "b0a")
	}

	if b.palette.B0b == nil {
		missingFields = append(missingFields, "b0b")
	}

	if b.palette.B0c == nil {
		missingFields = append(missingFields, "b0c")
	}

	if b.palette.B0d == nil {
		missingFields = append(missingFields, "b0d")
	}

	if b.palette.B0e == nil {
		missingFields = append(missingFields, "b0e")
	}

	if b.palette.B0f == nil {
		missingFields = append(missingFields, "b0f")
	}

	if len(missingFields) > 0 {
		errMsg := fmt.Sprintf("missing fields: %v", missingFields)
		return errors.New(errMsg)
	}

	return nil
}
