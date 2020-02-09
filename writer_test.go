package ppm

import (
	"bytes"
	"image"
	"image/color"
	"testing"
)

func createImage(colors []color.RGBA64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 3, 2))

	pixels := map[image.Point]color.RGBA64{
		image.Pt(0, 0): colors[0],
		image.Pt(1, 0): colors[1],
		image.Pt(2, 0): colors[2],
		image.Pt(0, 1): colors[3],
		image.Pt(1, 1): colors[4],
		image.Pt(2, 1): colors[5],
	}

	for point, col := range pixels {
		img.Set(point.X, point.Y, col)
	}

	return img
}

func TestEncode(t *testing.T) {
	colors := []color.RGBA64{
		color.RGBA64{0xffff, 0, 0, 0xffff},
		color.RGBA64{0, 0xffff, 0, 0xffff},
		color.RGBA64{0, 0, 0xffff, 0xffff},
		color.RGBA64{0xffff, 0xffff, 0, 0xffff},
		color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
		color.RGBA64{0, 0, 0, 0xffff},
	}

	var w bytes.Buffer
	img := createImage(colors)
	err := Encode(&w, img)

	if err != nil {
		t.Error("Error encoding", err)
	}

	header := []byte("P6\n3 2\n255\n")
	wBytes := w.Bytes()

	i := 0
	for _, byte := range header {
		if wBytes[i] != byte {
			t.Error("Error encoding header")
			t.FailNow()
		}
		i++
	}

	for _, col := range colors {
		if wBytes[i] != byte(col.R) || wBytes[i+1] != byte(col.G) || wBytes[i+2] != byte(col.B) {
			t.Error("Error encoding body")
			t.FailNow()
		}

		i += 3
	}
}
