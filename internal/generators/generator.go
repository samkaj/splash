package generators

import "splash/internal/models"

type Generator interface {
	Generate(palette *models.Base16Palette) []byte
}
