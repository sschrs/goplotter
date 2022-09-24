package goplotter

import (
	"image"
	"image/png"
	"os"
)

func SavePNG(img *image.RGBA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return png.Encode(f, img)
}
