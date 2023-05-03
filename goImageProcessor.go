package goimageprocessor

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

func ImageProcessor(data string, width, quality int) (*bytes.Buffer, error) {
	// Decode base64 string ke byte array
	imageBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// decode gambar
	img, format, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	// Resize image
	newHeight := img.Bounds().Dy() * width / img.Bounds().Dx()
	resized := image.NewRGBA(image.Rect(0, 0, width, newHeight))
	if err := resize(resized, img); err != nil {
		return nil, err
	}

	// create new varibale bytes to alocate new image
	buf := new(bytes.Buffer)

	// do compress image
	switch format {
	case "jpeg":
		// compress jpeg
		opts := jpeg.Options{
			Quality: quality,
		}
		err = jpeg.Encode(buf, resized, &opts)
		if err != nil {
			return nil, err
		}
	case "png":
		// compress png
		err = png.Encode(buf, resized)
		if err != nil {
			return nil, err
		}
	default:
		// format not valid
		return nil, fmt.Errorf("image format not supported")
	}

	return buf, nil
}

// resize resizes src to dst.
func resize(dst *image.RGBA, src image.Image) error {
	sw, sh := src.Bounds().Dx(), src.Bounds().Dy()
	dw, dh := dst.Bounds().Dx(), dst.Bounds().Dy()

	for dy := 0; dy < dh; dy++ {
		for dx := 0; dx < dw; dx++ {
			sx := dx * sw / dw
			sy := dy * sh / dh
			dst.Set(dx, dy, src.At(sx, sy))
		}
	}

	return nil
}
