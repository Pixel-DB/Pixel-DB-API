package utils

import (
	"image"
	_ "image/png"
	"io"
)

type Result struct {
	Width  int
	Height int
}

func CheckResolution(i io.Reader) (Result, error) {
	im, _, err := image.DecodeConfig(i)
	if err != nil {
		return Result{}, err
	}

	return Result{Width: im.Width, Height: im.Height}, nil
}
