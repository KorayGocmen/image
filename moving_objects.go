package image

import (
	"errors"
	"sort"
	"sync"
)

// medianPixel finds the median r, g, b values from the given
// pixel array and creates a new pixel from that median values
func medianPixel(pixels []Pixel) Pixel {
	var (
		rValues []int
		gValues []int
		bValues []int
	)

	for _, pix := range pixels {
		rValues = append(rValues, pix.R)
		gValues = append(gValues, pix.G)
		bValues = append(bValues, pix.B)
	}

	sort.Ints(rValues)
	sort.Ints(gValues)
	sort.Ints(bValues)

	rMedian := rValues[int(len(rValues)/2)]
	gMedian := gValues[int(len(gValues)/2)]
	bMedian := bValues[int(len(bValues)/2)]

	return Pixel{rMedian, gMedian, bMedian, 0}
}

// RemoveMovingObj iterates the given filepaths and generates new image
// objects. It then checks to see if all the heights and the widths
// of the images are matching. If they are, each pixel of every image is
// iterated and a median filter is applied to given images. Returns the
// output image object and an error if there is any.
func RemoveMovingObj(filePaths []string) (*Image, error) {
	var images []*Image
	for _, filePath := range filePaths {
		img, err := New(filePath)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	if len(images) < 5 {
		return nil, errors.New("not enough images to perform noise reduction")
	}

	outputImage := images[0]

	heigth := outputImage.Height
	width := outputImage.Width
	for _, img := range images {
		if heigth != img.Height || width != img.Width {
			return nil, errors.New("at least one image has a different width or height")
		}
	}

	var wg sync.WaitGroup

	for rowIndex := 0; rowIndex < heigth; rowIndex++ {
		wg.Add(1)

		go (func(rowIndex int) {
			for colIndex := 0; colIndex < width; colIndex++ {
				var pixels []Pixel
				for _, img := range images {
					pixels = append(pixels, img.Pixels[rowIndex][colIndex])
				}

				medPixel := medianPixel(pixels)
				outputImage.Pixels[rowIndex][colIndex].Set("R", medPixel.R)
				outputImage.Pixels[rowIndex][colIndex].Set("G", medPixel.G)
				outputImage.Pixels[rowIndex][colIndex].Set("B", medPixel.B)
			}
			wg.Done()
		})(rowIndex)
	}

	wg.Wait()
	return outputImage, nil
}
