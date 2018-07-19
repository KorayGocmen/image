package image

import (
	"testing"
)

func TestGrayscale(t *testing.T) {
	height, width := 25, 25

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			pixel := Pixel{R: 1, G: 2, B: 3, A: 10}
			row = append(row, pixel)
		}
		pixels = append(pixels, row)
	}

	img := &Image{
		Pixels: pixels,
		Width:  width,
		Height: height,
	}

	grayscaleAverage := img.Grayscale(GrayscaleAverage)
	testPix1 := grayscaleAverage.Pixels[0][0]
	expectedPix := 2
	if testPix1.R != expectedPix || testPix1.G != expectedPix || testPix1.B != expectedPix {
		t.Error("grayscale average error")
	}

	grayscaleDesaturation := img.Grayscale(GrayscaleDesaturation)
	testPix2 := grayscaleDesaturation.Pixels[0][0]
	expectedPix2 := int((max(testPix2.R, testPix2.G, testPix2.B) + min(testPix2.R, testPix2.G, testPix2.B)) / 2)
	if testPix2.R != expectedPix2 || testPix2.G != expectedPix2 || testPix2.B != expectedPix2 {
		t.Error("grayscale desaturation error")
	}

	grayscaleLuma := img.Grayscale(GrayscaleLuma)
	testPix3 := grayscaleLuma.Pixels[0][0]
	expectedPix3 := int(float32(testPix3.R)*0.2126 + float32(testPix3.G)*0.7152 + float32(testPix3.B)*0.0722)
	if testPix3.R != expectedPix3 || testPix3.G != expectedPix3 || testPix3.B != expectedPix3 {
		t.Error("grayscale luma error")
	}
}
