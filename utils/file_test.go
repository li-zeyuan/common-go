package utils

import (
	"fmt"
	"testing"
)

func TestIsImage(t *testing.T) {
	filenames := []string{"photo.jpg", "document.pdf", "image.png", "archive.zip"}

	for _, filename := range filenames {
		if IsImage(filename) {
			fmt.Printf("%s is an image file.\n", filename)
		} else {
			fmt.Printf("%s is not an image file.\n", filename)
		}
	}
}
