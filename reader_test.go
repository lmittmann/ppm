package ppm_test

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/lmittmann/ppm"
)

type interfaceImage struct{ image.Image }

var testImgs = []struct {
	ImgFn func() image.Image
	Enc   []byte
}{
	{
		ImgFn: func() image.Image {
			img := image.NewRGBA(image.Rect(0, 0, 1, 1))
			img.SetRGBA(0, 0, color.RGBA{0x11, 0x22, 0x33, 0xff})
			return img
		},
		Enc: []byte("P6\n1 1\n255\n\x11\x22\x33"),
	},
	{
		ImgFn: func() image.Image {
			img := image.NewRGBA(image.Rect(0, 0, 1, 1))
			img.SetRGBA(0, 0, color.RGBA{0x11, 0x22, 0x33, 0xff})
			return &interfaceImage{img}
		},
		Enc: []byte("P6\n1 1\n255\n\x11\x22\x33"),
	},
	{
		ImgFn: func() image.Image {
			img := image.NewRGBA(image.Rect(0, 0, 2, 1))
			img.SetRGBA(0, 0, color.RGBA{0x00, 0x00, 0x00, 0xff})
			img.SetRGBA(1, 0, color.RGBA{0xff, 0xee, 0xdd, 0xff})
			return img
		},
		Enc: []byte("P6\n2 1\n255\n\x00\x00\x00\xff\xee\xdd"),
	},
	{
		ImgFn: func() image.Image {
			img := image.NewRGBA(image.Rect(0, 0, 2, 1))
			img.SetRGBA(0, 0, color.RGBA{0x00, 0x00, 0x00, 0xff})
			img.SetRGBA(1, 0, color.RGBA{0xff, 0xee, 0xdd, 0xff})
			return &interfaceImage{img}
		},
		Enc: []byte("P6\n2 1\n255\n\x00\x00\x00\xff\xee\xdd"),
	},
}

func TestDecode(t *testing.T) {
	for i, test := range testImgs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			buf := bytes.NewBuffer(test.Enc)
			gotImg, err := ppm.Decode(buf)
			if err != nil {
				t.Fatalf("Failed to decode image: %v", err)
			}

			wantImg := test.ImgFn()
			if err := equal(wantImg, gotImg); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDecodeConfig(t *testing.T) {
	tests := []struct {
		Enc        io.Reader
		WantConfig image.Config
		WantErr    error
	}{
		{
			Enc:        bytes.NewBufferString("P6\n1 1\n255\n"),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 1},
		},
		{
			Enc:        bytes.NewBufferString("P6\n2 1\n255\n"),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 2, Height: 1},
		},
		{
			Enc:        bytes.NewBufferString("P6\n1 2\n255\n"),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 2},
		},
		{
			Enc:        bytes.NewBufferString("P6\n1 1\n255\n"),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 1},
		},
		{
			Enc:        bytes.NewBufferString("P6\n1\n1\n255\n"),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 1},
		},
		{
			Enc:        bytes.NewBufferString("P6 1 1 255 "),
			WantConfig: image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 1},
		},
		{
			Enc:     bytes.NewBufferString("P7\n1 1\n255\n"),
			WantErr: errors.New("ppm: invalid header"),
		},
		{
			Enc:     bytes.NewBufferString("P6\n1 1\n15\n"),
			WantErr: errors.New("ppm: unsupported format (maxVal != 255)"),
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotConfig, gotErr := ppm.DecodeConfig(test.Enc)
			if test.WantErr != nil && (gotErr == nil || test.WantErr.Error() != gotErr.Error()) {
				t.Fatalf("Err: want %q, got %q", test.WantErr, gotErr)
			}
			if test.WantConfig != gotConfig {
				t.Fatalf("Config: want %+v, got %+v", test.WantConfig, gotConfig)
			}
		})
	}
}

func equal(imgA, imgB image.Image) error {
	if imgA.Bounds() != imgB.Bounds() {
		return fmt.Errorf("bounds not equal")
	}

	rect := imgA.Bounds()
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			imgAR, imgAG, imgAB, imgAA := imgA.At(x, y).RGBA()
			imgBR, imgBG, imgBB, imgBA := imgB.At(x, y).RGBA()

			if imgAR != imgBR || imgAG != imgBG || imgAB != imgBB || imgAA != imgBA {
				return fmt.Errorf("%d %d: (%d %d %d %d) != (%d %d %d %d)\n", x, y, imgAR, imgAG, imgAB, imgAA, imgBR, imgBG, imgBB, imgBA)
			}
		}
	}
	return nil
}

func BenchmarkDecode(b *testing.B) {
	benchmarks := []struct {
		Enc []byte
	}{
		{[]byte("P6\n1 1\n255\n\x00\x00\x00")},
		{[]byte("P6\n10 10\n255\n" + strings.Repeat("\x00\x00\x00", 10*10))},
		{[]byte("P6\n100 100\n255\n" + strings.Repeat("\x00\x00\x00", 100*100))},
		{[]byte("P6\n1000 1000\n255\n" + strings.Repeat("\x00\x00\x00", 1000*1000))},
	}

	for i, bm := range benchmarks {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.ReportAllocs()

			buf := bytes.NewBuffer(nil)
			for i := 0; i < b.N; i++ {
				buf.Write(bm.Enc)
				ppm.Decode(buf)
				buf.Reset()
			}
		})
	}
}
