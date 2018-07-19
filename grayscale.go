package image

import (
	"sync"
)

// Grayscale turns the images to grayscale.
func (img *Image) Grayscale(algorithm int) *Image {
	var wg sync.WaitGroup

	for rowIndex := 0; rowIndex < img.Height; rowIndex++ {
		wg.Add(1)

		go (func(rowIndex int, img *Image) {
			for colIndex := 0; colIndex < img.Width; colIndex++ {
				pixel := img.Pixels[rowIndex][colIndex]

				var gray int
				if algorithm == GrayscaleLuma {
					gray = int(float32(pixel.R)*0.2126 + float32(pixel.G)*0.7152 + float32(pixel.B)*0.0722)
				} else if algorithm == GrayscaleDesaturation {
					gray = int((max(pixel.R, pixel.G, pixel.B) + min(pixel.R, pixel.G, pixel.B)) / 2)
				} else {
					gray = int((pixel.R + pixel.G + pixel.B) / 3)
				}
				pixel.Set("R", gray)
				pixel.Set("G", gray)
				pixel.Set("B", gray)

				img.Pixels[rowIndex][colIndex] = pixel
			}
			wg.Done()
		})(rowIndex, img)
	}

	wg.Wait()
	return img
}
