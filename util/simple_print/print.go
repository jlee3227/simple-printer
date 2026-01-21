package simple_print

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/hennedo/escpos"
	"github.com/makeworld-the-better-one/dither/v2"
)

func Print(text string) error {
	log.Println("Starting text print job...")

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	w := io.ReadWriter(f)
	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	p.Write(text)

	p.LineFeed()
	p.LineFeed()
	p.LineFeed()
	p.LineFeed()

	// This is necessary to flush the buffer and force the printer to print
	// in case the supplied string is very long.
	p.Print()
	p.Print()
	p.PrintAndCut()

	log.Println("Text print job completed.")

	return nil
}

// TODO: Add flag to make different types of lists (numbered, bulleted, checkbox)
func PrintList(list []string) error {
	log.Println("Starting list print job...")

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	w := io.ReadWriter(f)
	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	// Print title of list
	p.Size(2, 2).Write(list[0])
	p.Print()
	p.LineFeed()

	for _, item := range list[1:] {
		p.Size(1, 1).Write("- " + item)
		p.Print()
		p.LineFeed()
	}

	p.LineFeed()
	p.LineFeed()

	// This is necessary to flush the buffer and force the printer to print
	// in case the supplied string is very long.
	p.Print()
	p.Print()
	p.PrintAndCut()

	log.Println("List print job completed.")
	return nil
}

func PrintImage(filename string) error {
	log.Println("Starting image print job with new library...")

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := io.ReadWriter(f)
	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	img, err := imaging.Open(filename)
	if err != nil {
		return fmt.Errorf("Error retrieving image: %v\n", err)
	}

	dstImg := imaging.Resize(img, 710, 0, imaging.Lanczos)
	outputImg := image.Image(dstImg)

	// Setting colors for dithering
	palette := []color.Color{
		color.Black,
		color.White,
	}

	// Create ditherer and dither image
	d := dither.NewDitherer(palette)
	d.Matrix = dither.ErrorDiffusionStrength(dither.FloydSteinberg, 0.75)
	outputImg = d.Dither(outputImg)

	_, err = p.PrintImage(outputImg)
	if err != nil {
		return fmt.Errorf("Failed to print image: %v\n", err)
	}
	p.Print()
	p.PrintAndCut()

	log.Println("Image print job completed.")

	return nil
}
