package image

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"golang.org/x/image/draw"
)

func GetPng(filename string) (image.Image, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v\n", err)
	}

	pic, err := toPng(data)
	if err != nil {
		return nil, fmt.Errorf("Error converting to png: %v\n", err)
	}

	r := bytes.NewReader(pic)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("Error decoding png: %v\n", err)
	}

	return img, nil
}

func toPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
		return imageBytes, nil
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to png", contentType)
}

func Resize(img image.Image) image.Image {
	resizeFactor := 6

	if img.Bounds().Max.X < 1600 {
		resizeFactor = 3
	}

	img2 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/resizeFactor, img.Bounds().Max.Y/resizeFactor))
	draw.NearestNeighbor.Scale(img2, img2.Rect, img, img.Bounds(), draw.Over, nil)
	return img2
}
