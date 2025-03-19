package models_test

import (
	"errors"
	"reflect"
	"splash/internal/models"
	"testing"
)

func TestPaletteFromJson(t *testing.T) {
	jsonPalette := `{
        "b00": "#000000",
        "b01": "#010101",
        "b02": "#020202",
        "b03": "#030303",
        "b04": "#040404",
        "b05": "#050505",
        "b06": "#060606",
        "b07": "#070707",
        "b08": "#080808",
        "b09": "#090909",
        "b0a": "#0a0a0a",
        "b0b": "#0b0b0b",
        "b0c": "#0c0c0c",
        "b0d": "#0d0d0d",
        "b0e": "#0e0e0e",
        "b0f": "#0f0f0f"
    }`

	palette, err := models.PaletteFromJson([]byte(jsonPalette))
	if !errors.Is(err, nil) {
		t.Errorf("Expected json unmarshalling to work, got error \n\n%v\n\n", err)
	}

	want, _ := getStandardPalette()
	if !reflect.DeepEqual(palette, want) {
		t.Errorf("Generated palette does not match the expected palette %v != %v", palette, want)
	}
}

func TestIncorrectJson(t *testing.T) {
	jsonPalette := `{
        "b00": "#000000",
        "b01": "#010101",
        "b02": "#020202",
        "b03": "#030303",
        "b04": "#040404",
        "b05": "#050505",
        "b06": "#060606",
        "b07": "#070707",
        "b08": "#080808",
        "b09": "#090909",
        "b0a": "#0a0a0a",
        "b0b": "#0b0b0b",
        "b0c": "#0c0c0c",
        "b0d": "#0d0d0d",
        "b0f": "#0f0f0f"
    }`

	_, err := models.PaletteFromJson([]byte(jsonPalette))
	if errors.Is(err, nil) {
		t.Error("Expected an error due to a missing field")
	}
}

func TestIncorrectJsonFormat(t *testing.T) {
	jsonPalette := `[
        "b00": "#000000",
        "b01": "#010101",
        "b02": "#020202",
        "b03": "#030303",
        "b04": "#040404",
        "b05": "#050505",
        "b06": "#060606",
        "b07": "#070707",
        "b08": "#080808",
        "b09": "#090909",
        "b0a": "#0a0a0a",
        "b0b": "#0b0b0b",
        "b0c": "#0c0c0c",
        "b0d": "#0d0d0d",
        "b0f": "#0f0f0f"
    ]`

	_, err := models.PaletteFromJson([]byte(jsonPalette))
	if errors.Is(err, nil) {
		t.Error("Expected an error due to a missing field")
	}

	jsonPalette = `"foo"`
	_, err = models.PaletteFromJson([]byte(jsonPalette))
	if errors.Is(err, nil) {
		t.Error("Expected an error due to non-JSON input")
	}

	jsonPalette = `{}`
	_, err = models.PaletteFromJson([]byte(jsonPalette))
	if errors.Is(err, nil) {
		t.Error("Expected an error due to non-JSON input")
	}

	_, err = models.PaletteFromJson(nil)
	if errors.Is(err, nil) {
		t.Error("Expected an error due to nil input")
	}
}

func getStandardPalette() (*models.Base16Palette, error) {
	builder := models.NewPaletteBuilder()
	builder = builder.
		Base00("#000000").
		Base01("#010101").
		Base02("#020202").
		Base03("#030303").
		Base04("#040404").
		Base05("#050505").
		Base06("#060606").
		Base07("#070707").
		Base08("#080808").
		Base09("#090909").
		Base0a("#0a0a0a").
		Base0b("#0b0b0b").
		Base0c("#0c0c0c").
		Base0d("#0d0d0d").
		Base0e("#0e0e0e").
		Base0f("#0f0f0f")

	palette, err := builder.Build()
	return palette, err
}

func TestCorrectPaletteBuilder(t *testing.T) {
	palette, err := getStandardPalette()
	if !errors.Is(err, nil) {
		t.Errorf("Expected builder to correctly build palette, got error %v", err)
	}

	if palette == nil {
		t.Errorf("Expected builder to correctly build palette")
	}
}

func TestOutOfRangeHex(t *testing.T) {
	_, err := models.ColorFromHexString("#gg0000")
	if errors.Is(err, nil) {
		t.Error("r component should throw an error")
	}

	_, err = models.ColorFromHexString("#00gg00")
	if errors.Is(err, nil) {
		t.Error("g component should throw an error")
	}

	_, err = models.ColorFromHexString("#0000gg")
	if errors.Is(err, nil) {
		t.Error("b component should throw an error")
	}
}

func TestIncorrectHexLength(t *testing.T) {
	_, err := models.ColorFromHexString("#00000000000")
	if errors.Is(err, nil) {
		t.Error("hex string should be too long")
	}

	_, err = models.ColorFromHexString("#0")
	if errors.Is(err, nil) {
		t.Error("hex string should be too short")
	}
}

func TestMissingFields(t *testing.T) {
	builder := models.NewPaletteBuilder()
	_, err := builder.Build()

	expected := errors.New("missing fields: [b00 b01 b02 b03 b04 b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base00("#000000")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b01 b02 b03 b04 b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base01("#010101")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b02 b03 b04 b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base02("#020202")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b03 b04 b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base03("#030303")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b04 b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base04("#040404")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b05 b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base05("#050505")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b06 b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base06("#060606")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b07 b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base07("#070707")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b08 b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base08("#080808")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b09 b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base09("#090909")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0a b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base0a("#0a0a0a")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0b b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base0b("#0b0b0b")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0c b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base0c("#0c0c0c")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0d b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base0d("#0d0d0d")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0e b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}

	builder = builder.Base0e("#0e0e0e")
	_, err = builder.Build()
	expected = errors.New("missing fields: [b0f]")
	if err.Error() != expected.Error() {
		t.Errorf("expected\n%v\ngot error\n%v", expected, err)
	}
}
