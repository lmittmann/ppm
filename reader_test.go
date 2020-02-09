package ppm

import (
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

  pixelData := []uint32{
    0, 0, 0xffff, 0, 0,
    1, 0, 0, 0xffff, 0,
    2, 0, 0, 0, 0xffff,
    0, 1, 0xffff, 0xffff, 0,
    1, 1, 0xffff, 0xffff, 0xffff,
    2, 1, 0, 0, 0,
  }

  for i := 0; i < len(pixelData); i += 5 {
    x := int(pixelData[i])
    y := int(pixelData[i+1])
    r, g, b, a := img.At(x, y).RGBA()

    if r != pixelData[i+2] || g != pixelData[i+3] || b != pixelData[i+4] || a != 0xffff {
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
