package ppm

import (
	"bytes"
	"image"
	"image/color"
	"testing"
)

func createImage(data []uint16) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 3, 2))

	for i := 0; i < len(data); i += 5 {
		imageColor := color.RGBA64{
			data[i+2],
			data[i+3],
			data[i+4],
			0xffff,
		}
		img.Set(int(data[i]), int(data[i+1]), imageColor)
	}

	return img
}

func TestEncode(t *testing.T) {
	pixelData := []uint16{
		0, 0, 0xffff, 0, 0,
		1, 0, 0, 0xffff, 0,
		2, 0, 0, 0, 0xffff,
		0, 1, 0xffff, 0xffff, 0,
		1, 1, 0xffff, 0xffff, 0xffff,
		2, 1, 0, 0, 0,
	}

	var w bytes.Buffer
	img := createImage(pixelData)
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
		}
		i++
	}

	for j := 0; j < len(pixelData); j += 5 {
		r := byte(pixelData[j+2])
		g := byte(pixelData[j+3])
		b := byte(pixelData[j+4])

		if wBytes[i] != r || wBytes[i+1] != g || wBytes[i+2] != b {
			t.Error("Error encoding body")
			t.FailNow()
		}
		i += 3
	}
}
