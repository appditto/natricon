package color

import (
	"math"
	"testing"
)

func TestHTMLToRGB(t *testing.T) {
	rgb, _ := HTMLToRGB("#884444")
	expectedR := 136.0
	expectedG := 68.0
	expectedB := 68.0
	if rgb.R != expectedR {
		t.Errorf("Expected R %f got %f", expectedR, rgb.R)
	}
	if rgb.G != expectedG {
		t.Errorf("Expected G %f got %f", expectedG, rgb.G)
	}
	if rgb.B != expectedB {
		t.Errorf("Expected B %f got %f", expectedB, rgb.B)
	}
	rgb, _ = HTMLToRGB("#FFFFFF")
	expectedR = 255.0
	expectedG = 255.0
	expectedB = 255.0
	if rgb.R != expectedR {
		t.Errorf("Expected R %f got %f", expectedR, rgb.R)
	}
	if rgb.G != expectedG {
		t.Errorf("Expected G %f got %f", expectedG, rgb.G)
	}
	if rgb.B != expectedB {
		t.Errorf("Expected B %f got %f", expectedB, rgb.B)
	}
	rgb, _ = HTMLToRGB("#000000")
	expectedR = 0.0
	expectedG = 0.0
	expectedB = 0.0
	if rgb.R != expectedR {
		t.Errorf("Expected R %f got %f", expectedR, rgb.R)
	}
	if rgb.G != expectedG {
		t.Errorf("Expected G %f got %f", expectedG, rgb.G)
	}
	if rgb.B != expectedB {
		t.Errorf("Expected B %f got %f", expectedB, rgb.B)
	}
}

func TestRGBtoHTML(t *testing.T) {
	expectedHTML := "#884444"
	rgb := RGB{
		R: 136.0,
		G: 68.0,
		B: 68.0,
	}
	if rgb.ToHTML(true) != expectedHTML {
		t.Errorf("Expected %s but got %s", expectedHTML, rgb.ToHTML(true))
	}
	expectedHTML = "#ffffff"
	rgb = RGB{
		R: 255.0,
		G: 255.0,
		B: 255.0,
	}
	if rgb.ToHTML(true) != expectedHTML {
		t.Errorf("Expected %s but got %s", expectedHTML, rgb.ToHTML(true))
	}
	expectedHTML = "#000000"
	rgb = RGB{
		R: 0.0,
		G: 0.0,
		B: 0.0,
	}
	if rgb.ToHTML(true) != expectedHTML {
		t.Errorf("Expected %s but got %s", expectedHTML, rgb.ToHTML(true))
	}
}

func TestRGBtoHSV(t *testing.T) {
	expectedHSV := HSV{
		H: 0.0,
		S: 0.5,
		V: 0.53,
	}
	rgb := RGB{
		R: 136.0,
		G: 68.0,
		B: 68.0,
	}
	convertedHSV := rgb.ToHSV()
	// Round to 2 decimals
	convertedHSV.H = math.Round(convertedHSV.H*100) / 100
	convertedHSV.S = math.Round(convertedHSV.S*100) / 100
	convertedHSV.V = math.Round(convertedHSV.V*100) / 100
	if convertedHSV.H != expectedHSV.H {
		t.Errorf("Expected H %f but got %f", expectedHSV.H, convertedHSV.H)
	}
	if convertedHSV.S != expectedHSV.S {
		t.Errorf("Expected S %f but got %f", expectedHSV.S, convertedHSV.S)
	}
	if convertedHSV.V != expectedHSV.V {
		t.Errorf("Expected V %f but got %f", expectedHSV.V, convertedHSV.V)
	}
	expectedHSV = HSV{
		H: 0.0,
		S: 0.0,
		V: 1.0,
	}
	rgb = RGB{
		R: 255.0,
		G: 255.0,
		B: 255.0,
	}
	convertedHSV = rgb.ToHSV()
	// Round to 2 decimals
	convertedHSV.H = math.Round(convertedHSV.H*100) / 100
	convertedHSV.S = math.Round(convertedHSV.S*100) / 100
	convertedHSV.V = math.Round(convertedHSV.V*100) / 100
	if convertedHSV.H != expectedHSV.H {
		t.Errorf("Expected H %f but got %f", expectedHSV.H, convertedHSV.H)
	}
	if convertedHSV.S != expectedHSV.S {
		t.Errorf("Expected S %f but got %f", expectedHSV.S, convertedHSV.S)
	}
	if convertedHSV.V != expectedHSV.V {
		t.Errorf("Expected V %f but got %f", expectedHSV.V, convertedHSV.V)
	}
	expectedHSV = HSV{
		H: 0.0,
		S: 0.0,
		V: 0.0,
	}
	rgb = RGB{
		R: 0.0,
		G: 0.0,
		B: 0.0,
	}
	convertedHSV = rgb.ToHSV()
	// Round to 2 decimals
	convertedHSV.H = math.Round(convertedHSV.H*100) / 100
	convertedHSV.S = math.Round(convertedHSV.S*100) / 100
	convertedHSV.V = math.Round(convertedHSV.V*100) / 100
	if convertedHSV.H != expectedHSV.H {
		t.Errorf("Expected H %f but got %f", expectedHSV.H, convertedHSV.H)
	}
	if convertedHSV.S != expectedHSV.S {
		t.Errorf("Expected S %f but got %f", expectedHSV.S, convertedHSV.S)
	}
	if convertedHSV.V != expectedHSV.V {
		t.Errorf("Expected V %f but got %f", expectedHSV.V, convertedHSV.V)
	}
}

func TestRGBtoHSL(t *testing.T) {
	expectedHSL := HSL{
		H: 50.0,
		S: 0.5,
		L: 0.5,
	}
	rgb := RGB{
		R: 191.0,
		G: 170.0,
		B: 64.0,
	}
	convertedHSL := rgb.ToHSL()
	// Round to 2 decimals
	convertedHSL.H = math.Round(convertedHSL.H*100) / 100
	convertedHSL.S = math.Round(convertedHSL.S*100) / 100
	convertedHSL.L = math.Round(convertedHSL.L*100) / 100
	if expectedHSL.H != expectedHSL.H {
		t.Errorf("Expected H %f but got %f", expectedHSL.H, convertedHSL.H)
	}
	if convertedHSL.S != expectedHSL.S {
		t.Errorf("Expected S %f but got %f", expectedHSL.S, convertedHSL.S)
	}
	if convertedHSL.L != expectedHSL.L {
		t.Errorf("Expected L %f but got %f", expectedHSL.L, convertedHSL.L)
	}
	// Convert back
	hslAsRGB := convertedHSL.ToRGB()
	if hslAsRGB.R != rgb.R {
		t.Errorf("Expected R %f but got %f", rgb.R, hslAsRGB.R)
	}
	if hslAsRGB.G != rgb.G {
		t.Errorf("Expected G %f but got %f", rgb.G, hslAsRGB.G)
	}
	if hslAsRGB.B != rgb.B {
		t.Errorf("Expected B %f but got %f", rgb.B, hslAsRGB.B)
	}
	// Another
	hsl := HSL{
		H: 248.0,
		S: 0.51,
		L: 0.36,
	}
	expectedRgb := RGB{
		R: 57.0,
		G: 45.0,
		B: 139.0,
	}
	convertedRgb := hsl.ToRGB()
	if expectedRgb.R != convertedRgb.R {
		t.Errorf("Expected R %f but got %f", expectedRgb.R, convertedRgb.R)
	}
	if expectedRgb.G != convertedRgb.G {
		t.Errorf("Expected G %f but got %f", expectedRgb.G, convertedRgb.G)
	}
	if expectedRgb.B != convertedRgb.B {
		t.Errorf("Expected B %f but got %f", expectedRgb.B, convertedRgb.B)
	}
}
