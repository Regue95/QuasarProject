package services

import (
	"testing"
)

// TestGetLocation tests GetLocation method
func TestGetLocation(t *testing.T) {

	data := []float32{627, 208, 450}

	var expectedX float32 = 49
	var expectedY float32 = 103

	calculatorService := NewCalculatorService()
	x, y := calculatorService.GetLocation(data...)

	if x != expectedX && y != expectedY {
		t.Fatal("(", x, ", ", y, ") es distinto a: ", "(", expectedX, ", ", expectedY, ")")
	}
}

// TestGetMessage tests GetMessage method
func TestGetMessage(t *testing.T) {

	data := [][]string{
		{" ", "es", "", " ", ""},
		{"este", "", "un", "", "secreto"},
		{"", "", "", "mensaje", ""},
	}

	expected := "este es un mensaje secreto"

	calculatorService := NewCalculatorService()
	result := calculatorService.GetMessage(data...)

	if result != expected {
		t.Fatal(result, " es distinto a: ", expected)
	}
}
