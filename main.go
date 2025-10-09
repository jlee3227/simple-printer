package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"net/http"
	"os"

	"github.com/hennedo/escpos"
	"golang.org/x/image/draw"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := io.ReadWriter(f)

	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	//	p.Bold(true).Size(2, 2).Write("ur a bitch")
	//	p.LineFeed()
	//	p.Print()
	//	p.Bold(false).Underline(2).Justify(escpos.JustifyCenter).Write("this is underlined")
	//	p.LineFeed()
	//	p.Print()
	//	p.QRCode("https://github.com/hennedo/escpos", true, 10, escpos.QRCodeErrorCorrectionLevelH)
	//	p.LineFeed()
	//	p.Print()

	data, err := os.ReadFile("test.png")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	pic, err := ToPng(data)
	if err != nil {
		fmt.Printf("Error converting to png: %v\n", err)
	}

	r := bytes.NewReader(pic)
	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Printf("Error decoding png: %v\n", err)
	}

	img = resize(img)

	_, err = p.PrintImage(img)
	if err != nil {
		fmt.Printf("Failed to print image: %v\n", err)
	}
	p.Print()

	// You need to use either p.Print() or p.PrintAndCut() at the end to send the data to the printer.
	p.PrintAndCut()
}

func ToPng(imageBytes []byte) ([]byte, error) {
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

func resize(img image.Image) image.Image {
	fmt.Println(img.Bounds().Max.X, img.Bounds().Max.Y)
	resizeFactor := 6

	if img.Bounds().Max.X < 1600 {
		resizeFactor = 3
	}

	img2 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/resizeFactor, img.Bounds().Max.Y/resizeFactor))
	draw.NearestNeighbor.Scale(img2, img2.Rect, img, img.Bounds(), draw.Over, nil)
	return img2
}
