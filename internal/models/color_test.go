package models_test

import (
	"errors"
	"reflect"
	"splash/internal/models"
	"testing"
)

func TestHexConversionInverse(t *testing.T) {
	color := &models.Color{Red: 255, Green: 255, Blue: 255}
	colorString := "#ffffff"
	result := color.ToHexString()
	if !reflect.DeepEqual(colorString, result) {
		t.Errorf("color -> hex is not working, %v != %v", result, colorString)
	}

	c, _ := models.ColorFromHexString(colorString)
	if !reflect.DeepEqual(color, c) {
		t.Errorf("color <- hex is not working, %v != %v", color, colorString)
	}
}

func TestColorFromHexString(t *testing.T) {
	testCases := []struct {
		in            string
		expectedColor *models.Color
		expectedErr   error
	}{
		{
			in:            "#FFFFFF",
			expectedColor: &models.Color{Red: 255, Green: 255, Blue: 255},
			expectedErr:   nil,
		},
		{
			in:            "0xFFFFFF",
			expectedColor: &models.Color{Red: 255, Green: 255, Blue: 255},
			expectedErr:   nil,
		},
		{
			in:            "0x0f0f02",
			expectedColor: &models.Color{Red: 15, Green: 15, Blue: 2},
			expectedErr:   nil,
		},
		{
			in:            "0x000000",
			expectedColor: &models.Color{Red: 0, Green: 0, Blue: 0},
			expectedErr:   nil,
		},
		{
			in:            "#000000",
			expectedColor: &models.Color{Red: 0, Green: 0, Blue: 0},
			expectedErr:   nil,
		},
		{
			in:            "0x FFFFF",
			expectedColor: nil,
			expectedErr:   models.ErrHexConversionFailure,
		},
		{
			in:            "0xFXFFF",
			expectedColor: nil,
			expectedErr:   models.ErrInvalidHexLength,
		},
		{
			in:            "0xFFFff",
			expectedColor: nil,
			expectedErr:   models.ErrInvalidHexLength,
		},
		{
			in:            "##FFFff",
			expectedColor: nil,
			expectedErr:   models.ErrHexConversionFailure,
		},
		{
			in:            "0x#FFFff",
			expectedColor: nil,
			expectedErr:   models.ErrHexConversionFailure,
		},
		{
			in:            "invalid",
			expectedColor: nil,
			expectedErr:   models.ErrInvalidHexPrefix,
		},
		{
			in:            "",
			expectedColor: nil,
			expectedErr:   models.ErrInvalidHexPrefix,
		},
	}

	for _, tc := range testCases {
		color, err := models.ColorFromHexString(tc.in)

		if !reflect.DeepEqual(color, tc.expectedColor) {
			t.Errorf("For input %q, expected color %v, got %v", tc.in, tc.expectedColor, color)
		}

		if !errors.Is(err, tc.expectedErr) {
			t.Errorf("For input %q, expected error %v, got %v", tc.in, tc.expectedErr, err)
		}
	}
}
