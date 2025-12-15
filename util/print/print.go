package print

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/hennedo/escpos"
	"github.com/jlee3227/simple-printer/util/image"
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

func PrintImage(filename string) error {
	log.Println("Starting image print job...")

	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := io.ReadWriter(f)
	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	img, err := image.GetPng(filename)
	if err != nil {
		return fmt.Errorf("Error getting png: %v\n", err)
	}

	img = image.Resize(img)

	_, err = p.PrintImage(img)
	if err != nil {
		return fmt.Errorf("Failed to print image: %v\n", err)
	}
	p.Print()
	p.PrintAndCut()

	log.Println("Image print job completed.")

	return nil
}

func PrintImage2(filename string) error {
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
	dstImg = imaging.Grayscale(dstImg)

	_, err = p.PrintImage(dstImg)
	if err != nil {
		return fmt.Errorf("Failed to print image: %v\n", err)
	}
	p.Print()
	p.PrintAndCut()

	log.Println("Image print job with new library completed.")

	return nil
}
