package models

import (
	"errors"
	"fmt"
)

type Base16Palette struct {
	b00 *Color
	b01 *Color
	b02 *Color
	b03 *Color
	b04 *Color
	b05 *Color
	b06 *Color
	b07 *Color
	b08 *Color
	b09 *Color
	b0a *Color
	b0b *Color
	b0c *Color
	b0d *Color
	b0e *Color
	b0f *Color
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

func (b *Base16Builder) Base00(value string) *Base16Builder {
	return b.setColor(&b.palette.b00, value)
}

func (b *Base16Builder) Base01(value string) *Base16Builder {
	return b.setColor(&b.palette.b01, value)
}

func (b *Base16Builder) Base02(value string) *Base16Builder {
	return b.setColor(&b.palette.b02, value)
}

func (b *Base16Builder) Base03(value string) *Base16Builder {
	return b.setColor(&b.palette.b03, value)
}

func (b *Base16Builder) Base04(value string) *Base16Builder {
	return b.setColor(&b.palette.b04, value)
}

func (b *Base16Builder) Base05(value string) *Base16Builder {
	return b.setColor(&b.palette.b05, value)
}

func (b *Base16Builder) Base06(value string) *Base16Builder {
	return b.setColor(&b.palette.b06, value)
}

func (b *Base16Builder) Base07(value string) *Base16Builder {
	return b.setColor(&b.palette.b07, value)
}

func (b *Base16Builder) Base08(value string) *Base16Builder {
	return b.setColor(&b.palette.b08, value)
}

func (b *Base16Builder) Base09(value string) *Base16Builder {
	return b.setColor(&b.palette.b09, value)
}

func (b *Base16Builder) Base0a(value string) *Base16Builder {
	return b.setColor(&b.palette.b0a, value)
}

func (b *Base16Builder) Base0b(value string) *Base16Builder {
	return b.setColor(&b.palette.b0b, value)
}

func (b *Base16Builder) Base0c(value string) *Base16Builder {
	return b.setColor(&b.palette.b0c, value)
}

func (b *Base16Builder) Base0d(value string) *Base16Builder {
	return b.setColor(&b.palette.b0d, value)
}

func (b *Base16Builder) Base0e(value string) *Base16Builder {
	return b.setColor(&b.palette.b0e, value)
}

func (b *Base16Builder) Base0f(value string) *Base16Builder {
	return b.setColor(&b.palette.b0f, value)
}

func newPalette() *Base16Palette {
	return &Base16Palette{
		b00: nil,
		b01: nil,
		b02: nil,
		b03: nil,
		b04: nil,
		b05: nil,
		b06: nil,
		b07: nil,
		b08: nil,
		b09: nil,
		b0a: nil,
		b0b: nil,
		b0c: nil,
		b0d: nil,
		b0e: nil,
		b0f: nil,
	}
}

func (b *Base16Builder) validate() error {
	if b.err != nil {
		return b.err
	}

	missingFields := make([]string, 0)

	if b.palette.b00 == nil {
		missingFields = append(missingFields, "b00")
	}

	if b.palette.b01 == nil {
		missingFields = append(missingFields, "b01")
	}

	if b.palette.b02 == nil {
		missingFields = append(missingFields, "b02")
	}

	if b.palette.b03 == nil {
		missingFields = append(missingFields, "b03")
	}

	if b.palette.b04 == nil {
		missingFields = append(missingFields, "b04")
	}

	if b.palette.b05 == nil {
		missingFields = append(missingFields, "b05")
	}

	if b.palette.b06 == nil {
		missingFields = append(missingFields, "b06")
	}

	if b.palette.b07 == nil {
		missingFields = append(missingFields, "b07")
	}

	if b.palette.b08 == nil {
		missingFields = append(missingFields, "b08")
	}

	if b.palette.b09 == nil {
		missingFields = append(missingFields, "b09")
	}

	if b.palette.b0a == nil {
		missingFields = append(missingFields, "b0a")
	}

	if b.palette.b0b == nil {
		missingFields = append(missingFields, "b0b")
	}

	if b.palette.b0c == nil {
		missingFields = append(missingFields, "b0c")
	}

	if b.palette.b0d == nil {
		missingFields = append(missingFields, "b0d")
	}

	if b.palette.b0e == nil {
		missingFields = append(missingFields, "b0e")
	}

	if b.palette.b0f == nil {
		missingFields = append(missingFields, "b0f")
	}

	if len(missingFields) > 0 {
		errMsg := fmt.Sprintf("missing fields: %v", missingFields)
		return errors.New(errMsg)
	}

	return nil
}
