package image

import (
	"testing"
)

func TestFilter(t *testing.T) {
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

	previousPixelVal := img.Pixels[0][0].R
	percentage := float32(2.0)

	img.Filter("R", percentage)

	// Manually calculating the expected pixel.R value.
	expectedPixel := previousPixelVal * int(1+percentage)
	if expectedPixel > 255 {
		expectedPixel = 255
	}

	if img.Pixels[0][0].R != expectedPixel {
		t.Error("filter error")
	}

	// pixel.R at (0,0) is now 3. Testing the max -> 255
	img.Filter("R", 255)
	if img.Pixels[0][0].R != 255 {
		t.Error("filter error")
	}
}
