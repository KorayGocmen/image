package image

import (
	"image/color"
	"os"
	"testing"
)

func TestPixelGet(t *testing.T) {
	pix := Pixel{R: 1, G: 1, B: 1, A: 1}

	if pix.Get("R") != pix.R {
		t.Error("pixel get error")
	}
	if pix.Get("G") != pix.G {
		t.Error("pixel get error")
	}
	if pix.Get("B") != pix.B {
		t.Error("pixel get error")
	}
	if pix.Get("A") != pix.A {
		t.Error("pixel get error")
	}
	if pix.Get("RANDOM") != -1 {
		t.Error("pixel get error")
	}
}

func TestPixelSet(t *testing.T) {
	pix := Pixel{R: 1, G: 1, B: 1, A: 1}
	newPix := Pixel{R: 1, G: 1, B: 1, A: 1}

	pix.Set("R", newPix.Get("R"))
	pix.Set("G", newPix.Get("G"))
	pix.Set("B", newPix.Get("B"))
	pix.Set("A", newPix.Get("A"))

	if pix.Get("R") != newPix.Get("R") {
		t.Error("pixel set error")
	}
	if pix.Get("G") != newPix.Get("G") {
		t.Error("pixel set error")
	}
	if pix.Get("B") != newPix.Get("B") {
		t.Error("pixel set error")
	}
	if pix.Get("A") != newPix.Get("A") {
		t.Error("pixel set error")
	}
}

func TestRGBAToPixel(t *testing.T) {
	expectedPix := Pixel{R: 128, G: 254, B: 200, A: 10}

	color := color.RGBA{
		uint8(expectedPix.R),
		uint8(expectedPix.G),
		uint8(expectedPix.B),
		uint8(expectedPix.A),
	}
	actualPix := rgbaToPixel(color.RGBA())

	if expectedPix.R != actualPix.R || expectedPix.G != actualPix.G || expectedPix.B != actualPix.B || expectedPix.A != actualPix.A {
		t.Error("rgba to pixel error", expectedPix, actualPix)
	}
}

func TestNew(t *testing.T) {
	imgPng, err := New("test/test.png")

	if err != nil {
		t.Error("error in new image", err)
	}

	if imgPng.Height != 10 || imgPng.Width != 10 {
		t.Error("error in new image width or height", imgPng.Height, imgPng.Width)
	}

	for y, row := range imgPng.Pixels {
		for x, pix := range row {
			if pix.R != 255 || pix.G != 255 || pix.B != 255 {
				t.Error("error in new image pixel", pix, y, x)
				return
			}
		}
	}

	imgJpg, err := New("test/test.jpg")
	if err != nil {
		t.Error("error in new image", err)
	}

	if imgJpg.Height != 10 || imgJpg.Width != 10 {
		t.Error("error in new image width or height", imgJpg.Height, imgJpg.Width)
	}

	for y, row := range imgJpg.Pixels {
		for x, pix := range row {
			if pix.R != 255 || pix.G != 255 || pix.B != 255 {
				t.Error("error in new image pixel", pix, y, x)
				return
			}
		}
	}

	_, err1 := New("test/test123")
	if err1 != nil {
		t.Error("error in new image path not found")
	}
}

func TestWriteToFile(t *testing.T) {
	imgPng, err := New("test/test.png")
	imgJpg, err := New("test/test.jpg")

	if err != nil {
		t.Error("error in new image", err)
	}

	imgPng.WriteToFile("test/modified_png.png")
	imgJpg.WriteToFile("test/modified_jpg.jpg")

	imgPngModified, err := New("test/modified_png.png")
	imgJpgModified, err := New("test/modified_jpg.jpg")

	if err != nil {
		t.Error("error in write to file", err)
	}

	if imgPngModified.Height != 10 || imgPngModified.Width != 10 {
		t.Error("error in write to file image width or height", imgPngModified.Height, imgPngModified.Width)
	}

	for y, row := range imgPngModified.Pixels {
		for x, pix := range row {
			if pix.R != 255 || pix.G != 255 || pix.B != 255 {
				t.Error("error in write to file pixel", pix, y, x)
				return
			}
		}
	}

	if imgJpgModified.Height != 10 || imgJpgModified.Width != 10 {
		t.Error("error in write to file width or height", imgJpgModified.Height, imgJpgModified.Width)
	}

	for y, row := range imgJpgModified.Pixels {
		for x, pix := range row {
			if pix.R != 255 || pix.G != 255 || pix.B != 255 {
				t.Error("error in write to file pixel", pix, y, x)
				return
			}
		}
	}

	os.Remove("test/modified_jpg.jpg")
	os.Remove("test/modified_png.png")
}
