package image

import (
	"sync"
)

// Filter with "R", "G" or "B" values. Increase the given r/g/b values
// of pixel by the given percentage.
func (img *Image) Filter(color string, percentage float32) *Image {
	var wg sync.WaitGroup

	for rowIndex := 0; rowIndex < img.Height; rowIndex++ {
		wg.Add(1)
		go (func(rowIndex int, img *Image) {
			for colIndex := 0; colIndex < img.Width; colIndex++ {
				pixel := img.Pixels[rowIndex][colIndex]
				enhanced := pixel.Get(color) * int(1+percentage)
				if enhanced > 255 {
					enhanced = 255
				}
				img.Pixels[rowIndex][colIndex] = pixel.Set(color, enhanced)
			}
			wg.Done()
		})(rowIndex, img)
	}

	wg.Wait()
	return img
}
