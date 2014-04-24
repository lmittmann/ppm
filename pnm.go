// Package pnm is a simple library for writing one-dimensional arrays to
// binary Portable Anymap files (such as Portable Bitmap, Portable Graymap
// and Portable Pixelmap).

package pnm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// writes array to Portable Bitmap file (.pbm)
func WriteBitmap(filename string, bitmap *[]bool, width, height int) error {
	if len(*bitmap) != width*height {
		return errors.New("invalid array size: len(array) != width * height")
	}

	file, err := os.Create(filename)
	if err == nil {
		w := bufio.NewWriter(file)

		// write pbm file header
		fmt.Fprintf(w, "P4\n%d %d\n", width, height)

		// write bits
		var b byte = 0
		for i, v := range *bitmap {
			b <<= 1
			if v {
				b |= 1
			}
			if bit := (i%width + 1) % 8; i%width == width-1 || bit == 0 {
				w.WriteByte(b << byte((8-bit)%8))
				b = 0
			}
		}

		w.Flush()
		err = file.Close()
	}
	return err
}

// writes array to Portable Graymap file (.pgm)
func WriteGraymap(filename string, graymap *[]uint8, width, height int) error {
	if len(*graymap) != width*height {
		return errors.New("invalid array size: len(array) != width * height")
	}

	file, err := os.Create(filename)
	if err == nil {
		w := bufio.NewWriter(file)

		// write pgm file header
		fmt.Fprintf(w, "P5\n%d %d\n255\n", width, height)

		// write bytes
		w.Write(*graymap)

		w.Flush()
		err = file.Close()
	}
	return err
}

// writes array to Portable Pixelmap file (.ppm)
func WritePixelmap(filename string, pixelmap *[]uint8, width, height int) error {
	if len(*pixelmap) != width*height*3 {
		return errors.New("invalid array size: len(array) != 3 * width * height")
	}

	file, err := os.Create(filename)
	if err == nil {
		w := bufio.NewWriter(file)

		// write ppm file header
		fmt.Fprintf(w, "P6\n%d %d\n255\n", width, height)

		// write bytes
		w.Write(*pixelmap)

		w.Flush()
		err = file.Close()
	}
	return err
}
