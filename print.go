package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hennedo/escpos"
)

func Print(text string) error {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	w := io.ReadWriter(f)

	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	return nil
}

func PrintImage(filename string) error {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := io.ReadWriter(f)

	p := escpos.New(w)
	p.SetConfig(escpos.ConfigEpsonTMT88II)

	img, err := GetPng(filename)
	if err != nil {
		return fmt.Errorf("Error getting png: %v\n", err)
	}

	img = Resize(img)

	_, err = p.PrintImage(img)
	if err != nil {
		return fmt.Errorf("Failed to print image: %v\n", err)
	}
	p.Print()

	// You need to use either p.Print() or p.PrintAndCut() at the end to send the data to the printer.
	p.PrintAndCut()

	return nil
}
