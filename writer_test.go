package ppm_test

import (
	"bytes"
	"image"
	"strconv"
	"testing"

	"github.com/lmittmann/ppm"
)

func TestEncode(t *testing.T) {
	for i, test := range testImgs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			img := test.ImgFn()
			buf := bytes.NewBuffer(nil)
			ppm.Encode(buf, img)
			if wantEnc, gotEnc := test.Enc, buf.Bytes(); !bytes.Equal(wantEnc, gotEnc) {
				t.Fatalf("(-want +got)\n- %q\n+ %q", wantEnc, gotEnc)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	benchmarks := []struct {
		Img image.Image
	}{
		{image.NewRGBA(image.Rect(0, 0, 128, 128))},
		{&interfaceImage{image.NewRGBA(image.Rect(0, 0, 128, 128))}},
	}

	for i, bm := range benchmarks {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.ReportAllocs()

			buf := bytes.NewBuffer(nil)
			for i := 0; i < b.N; i++ {
				ppm.Encode(buf, bm.Img)
				buf.Reset()
			}
		})
	}
}
