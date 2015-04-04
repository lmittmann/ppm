package ppm

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
)

var errUnsupportedColorMode = errors.New("ppm: color mode not supported")

// Encode writes the Image img to Writer w in PPM format.
func Encode(w io.Writer, img image.Image) error {
	bw := bufio.NewWriter(w)

	switch img.ColorModel() {
	case color.RGBAModel:
		rec := img.Bounds()

		// write header
		fmt.Fprintf(bw, "P6\n%d %d\n255\n", rec.Dx(), rec.Dy())

		// write pixels
		pixel := make([]byte, 3)
		for y := rec.Min.Y; y < rec.Max.Y; y++ {
			for x := rec.Min.X; x < rec.Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				pixel[0], pixel[1], pixel[2] = byte(r), byte(g), byte(b)
				bw.Write(pixel)
			}
		}
		bw.Flush()
	default:
		return errUnsupportedColorMode
	}
	return nil
}
