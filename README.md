# `image`

 Package image implements a simple library for image operations. The library can work with pngs or jpgs. Same functions can be used for both of those image types.
 
Read more here:
http://www.koraygocmen.com/blog/writing-an-image-manipulation-library-in-go-part-1

98% Test coverage

---
#### Full Documentation:

https://godoc.org/github.com/KorayGocmen/image

```go
package main

import (
	"fmt"
	"log"

	"github.com/koraygocmen/image"
)

func main() {
	img, _ := image.New("test/test.jpg")

	err1 := img.Grayscale(image.GrayscaleAverage).WriteToFile("test/grayscale_average_method.jpg")
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := img.Grayscale(image.GrayscaleDesaturation).WriteToFile("test/grayscale_saturation_method.jpg")
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := img.Grayscale(image.GrayscaleLuma).WriteToFile("test/grayscale_luma_method.jpg")
	if err3 != nil {
		log.Fatal(err3)
	}

	err4 := img.Filter("R", 2).WriteToFile("test/red_color_filtered_200_percent.jpg")
	if err4 != nil {
		log.Fatal(err4)
	}

	// Works both with pngs and jpgs.
	imgPng, _ := image.New("test/test.png")
	fmt.Println(imgPng.Height, imgPng.Width)
}

```

---

### License

Released under the [MIT License](https://github.com/KorayGocmen/image/blob/master/LICENSE).
