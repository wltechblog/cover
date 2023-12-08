package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func main() {
	var width int
	var height int
	var filename string
	var fit bool
	var output string

	flag.IntVar(&width, "width", 320, "Width of output image")
	flag.IntVar(&height, "height", 240, "Height of output image")
	flag.StringVar(&filename, "filename", "", "Input image filename")
	flag.StringVar(&output, "output", "", "Output filename")
	flag.BoolVar(&fit, "fit", false, "Resize image so the entire thing fits within the specified dimensions")
	flag.Parse()

	in, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	d, err := png.Decode(in)
	if err != nil {
		in.Close()

		in, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		d, err = jpeg.Decode(in)
		if err != nil {
			log.Fatal(err)
		}
	}

	var dst *image.NRGBA

	if fit {
		bounds := d.Bounds()

		imageaspect := float64(bounds.Dy()) / float64(bounds.Dx())
		fitaspect := float64(height) / float64(width)
		log.Printf("imageaspect is %f, fitaspect is %f", imageaspect, fitaspect)
		if imageaspect < fitaspect {
			d = resize.Resize(uint(width), 0, d, resize.Lanczos3)
		} else {
			d = resize.Resize(0, uint(height), d, resize.Lanczos3)

		}
		bounds = d.Bounds()
		gap := height - bounds.Dy()
		drop := int(math.Round(float64(gap) / 2))
		gap = width - bounds.Dx()
		slide := int(math.Round(float64(gap) / 2))

		blank := imaging.New(width, height, color.Black)

		d = imaging.Paste(blank, d, image.Pt(slide, drop))
	}

	dst = imaging.Fill(d, width, height, imaging.Center, imaging.Lanczos)
	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(out, dst)
	out.Close()

}
