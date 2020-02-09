package ppm

import (
  "image"
  "image/color"
  "os"
  "testing"
)

func TestDecode(t *testing.T) {
  file, err := os.Open("./test/p6.ppm")

  if err != nil {
    t.Error("Error opening p6 ppm file", err)
  }

  img, err := Decode(file)

  if err != nil {
    t.Error("Error decoding", err)
  }

  bounds := img.Bounds()

  if bounds.Max.X != 3 {
    t.Error("Decoded incorrect width")
  }

  if bounds.Max.Y != 2 {
    t.Error("Decoded incorrect height")
  }

  pixels := map[image.Point]color.RGBA64{
    image.Pt(0, 0): color.RGBA64{0xffff, 0, 0, 0xffff},
    image.Pt(1, 0): color.RGBA64{0, 0xffff, 0, 0xffff},
    image.Pt(2, 0): color.RGBA64{0, 0, 0xffff, 0xffff},
    image.Pt(0, 1): color.RGBA64{0xffff, 0xffff, 0, 0xffff},
    image.Pt(1, 1): color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
    image.Pt(2, 1): color.RGBA64{0, 0, 0, 0xffff},
  }

  for point, col := range pixels {
    r, g, b, a := img.At(point.X, point.Y).RGBA()

    if uint16(r) != col.R || uint16(g) != col.G || uint16(b) != col.B || uint16(a) != col.A {
      t.Error("Decoded incorrect pixel values")
    }
  }
}

func TestDecodeP3(t *testing.T) {
  file, err := os.Open("./test/p3.ppm")

  if err != nil {
    t.Error("Error opening p3 ppm file", err)
  }

  _, err = Decode(file)

  if err != errBadHeader || err == nil {
    t.Error("Bad header error expected decoding p3", err)
  }

}

func TestDecodeConfig(t *testing.T) {
  file, err := os.Open("./test/p6.ppm")

  if err != nil {
    t.Error("Error opening p6 ppm file", err)
  }

  config, err := DecodeConfig(file)

  if err != nil {
    t.Error("Error decoding config", err)
  }

  if config.ColorModel != color.RGBAModel {
    t.Error("Decoded color model incorrectly")
  }

  if config.Width != 3 {
    t.Error("Decoded width incorrectly")
  }
  if config.Height != 2 {
    t.Error("Decoded height incorrectly")
  }
}
