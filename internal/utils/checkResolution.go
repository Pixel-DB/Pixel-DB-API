package utils

import (
	"image"
	_ "image/png"
	"io"
)

type Result struct {
	width  int
	height int
}

func checkResolution(i io.Reader) (Result, error) {
	im, _, err := image.DecodeConfig(i)
	if err != nil {
		return Result{}, err
	}

	return Result{width: im.Width, height: im.Height}, nil
}
