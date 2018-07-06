// Package image implements a simple library for image operations.
// The library can work with pngs or jpgs. Same functions can be
// used for both of those image types.

package image

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Pixel is a single pixel in 2d array
type Pixel struct {
	R int
	G int
	B int
	A int
}

// Image is the main object that holds information about the
// image file. Also is a wrapper around the decoded image
// from the standard image library.
type Image struct {
	Pixels [][]Pixel
	Width  int
	Height int
	_Rect  image.Rectangle
	_Image image.Image
}

const (
	GRAYSCALE_AVERAGE      = 0
	GRAYSCALE_LUMA         = 1
	GRAYSCALE_DESATURATION = 2
)

// Get pixel value with key name
func (pix *Pixel) Get(keyName string) int {
	switch keyName {
	case "R":
		return pix.R
	case "G":
		return pix.G
	case "B":
		return pix.B
	case "A":
		return pix.A
	default:
		return -1
	}
}

// Set pixel value with key name and new value
func (pix *Pixel) Set(keyName string, val int) Pixel {
	switch keyName {
	case "R":
		pix.R = val
	case "G":
		pix.G = val
	case "B":
		pix.B = val
	case "A":
		pix.A = val
	}
	return *pix
}

// rgbaToPixel alpha-premultiplied red, green, blue and alpha values
// to 8 bit red, green, blue and alpha values.
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257),
	}
}

// New reads an image from the given file path and return a
// new `Image` struct.
func New(filePath string) (*Image, error) {
	s := strings.Split(filePath, ".")
	imgType := s[len(s)-1]

	switch imgType {
	case "jpeg", "jpg":
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	case "png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	default:
		return nil, errors.New("unknown image type")
	}

	imgReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(imgReader)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			pixel := rgbaToPixel(img.At(x, y).RGBA())
			row = append(row, pixel)
		}
		pixels = append(pixels, row)
	}

	return &Image{
		Pixels: pixels,
		Width:  width,
		Height: height,
		_Rect:  img.Bounds(),
		_Image: img,
	}, nil
}

// WriteToFile writes iamges to the given filepath.
// Returns an error if it occurs.
func (img *Image) WriteToFile(outputPath string) error {
	cimg := image.NewRGBA(img._Rect)
	draw.Draw(cimg, img._Rect, img._Image, image.Point{}, draw.Over)

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			rowIndex, colIndex := y, x
			pixel := img.Pixels[rowIndex][colIndex]
			cimg.Set(x, y, color.RGBA{
				uint8(pixel.R),
				uint8(pixel.G),
				uint8(pixel.B),
				uint8(pixel.A),
			})
		}
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	s := strings.Split(outputPath, ".")
	imgType := s[len(s)-1]

	switch imgType {
	case "jpeg", "jpg":
		jpeg.Encode(fd, cimg, nil)
	case "png":
		png.Encode(fd, cimg)
	default:
		return errors.New("unknown image type")
	}

	return nil
}
