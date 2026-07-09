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

func openPrinter(device string) (*escpos.Escpos, *os.File, error) {
	f, err := os.OpenFile(device, os.O_RDWR, 0)
	if err != nil {
		return nil, nil, err
	}
	p := escpos.New(io.ReadWriter(f))
	p.SetConfig(escpos.ConfigEpsonTMT88II)
	return p, f, nil
}

func Print(device, text string) error {
	log.Println("Starting text print job...")

	p, f, err := openPrinter(device)
	if err != nil {
		return err
	}
	defer f.Close()

	p.Write(text)
	p.LineFeed()
	p.LineFeed()
	p.LineFeed()
	p.LineFeed()
	p.Print()
	p.Print()
	p.PrintAndCut()

	log.Println("Text print job completed.")
	return nil
}

// TODO: Add flag to make different types of lists (numbered, bulleted, checkbox)
func PrintList(device string, list []string) error {
	log.Println("Starting list print job...")

	p, f, err := openPrinter(device)
	if err != nil {
		return err
	}
	defer f.Close()

	p.Size(2, 2).Write(list[0])
	p.Print()
	p.LineFeed()

	for _, item := range list[1:] {
		p.Size(1, 1).Write(item)
		p.Print()
		p.LineFeed()
	}

	p.LineFeed()
	p.LineFeed()
	p.Print()
	p.Print()
	p.PrintAndCut()

	log.Println("List print job completed.")
	return nil
}

func PrintImage(device, filename string) error {
	log.Println("Starting image print job...")

	p, f, err := openPrinter(device)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := imaging.Open(filename)
	if err != nil {
		return fmt.Errorf("error retrieving image: %w", err)
	}

	dstImg := imaging.Resize(img, 710, 0, imaging.Lanczos)

	palette := []color.Color{color.Black, color.White}
	d := dither.NewDitherer(palette)
	d.Matrix = dither.ErrorDiffusionStrength(dither.FloydSteinberg, 0.75)
	outputImg := image.Image(d.Dither(dstImg))

	if _, err = p.PrintImage(outputImg); err != nil {
		return fmt.Errorf("failed to print image: %w", err)
	}
	p.Print()
	p.PrintAndCut()

	log.Println("Image print job completed.")
	return nil
}
