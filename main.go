package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	var width int
	var height int
	var filename string
	var output string

	flag.IntVar(&width, "width", 320, "Width of output image")
	flag.IntVar(&height, "height", 240, "Height of output image")
	flag.StringVar(&filename, "filename", "", "Input image filename")
	flag.StringVar(&output, "output", "", "Output filename")
	flag.Parse()

	in, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	d, err := png.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	dstImageFill := imaging.Fill(d, width, height, imaging.Center, imaging.Lanczos)
	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(out, dstImageFill)
	out.Close()

}
