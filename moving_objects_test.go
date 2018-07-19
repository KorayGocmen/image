package image

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRemoveMovingObj(t *testing.T) {

	testFolder := "test/test-frames/"
	outPath := testFolder + "frames-out.png"

	var filePaths []string
	files, err := ioutil.ReadDir(testFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filePaths = append(filePaths, testFolder+f.Name())
	}

	img, err := RemoveMovingObj(filePaths)
	if err != nil {
		fmt.Println(err)
	}

	img.WriteToFile(outPath)
	os.Remove(outPath)
}
