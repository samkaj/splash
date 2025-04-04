package generators

import (
	"fmt"
	"splash/internal/models"
)

type KittyGenerator struct{}

func (g *KittyGenerator) Generate(palette *models.Base16Palette) []byte {
	p := *palette
	output := fmt.Sprintf(`# Generated by splash
background %s
foreground %s
cursor %s
selection_background %s
selection_foreground %s
color0 %s
color1 %s
color2 %s
color3 %s
color4 %s
color5 %s
color6 %s
color7 %s
color8 %s
color9 %s
color10 %s
color11 %s
color12 %s
color13 %s
color14 %s
color15 %s
`,
		p.B00.ToHex(), // bg
		p.B05.ToHex(), // fg
		p.B05.ToHex(), // cursor
		p.B02.ToHex(), // sel-bg
		p.B00.ToHex(), // sel-fg
		p.B00.ToHex(),
		p.B01.ToHex(),
		p.B02.ToHex(),
		p.B03.ToHex(),
		p.B04.ToHex(),
		p.B05.ToHex(),
		p.B06.ToHex(),
		p.B07.ToHex(),
		p.B08.ToHex(),
		p.B09.ToHex(),
		p.B0a.ToHex(),
		p.B0b.ToHex(),
		p.B0c.ToHex(),
		p.B0d.ToHex(),
		p.B0e.ToHex(),
		p.B0f.ToHex(),
	)

	return []byte(output)
}
