package ppm

import (
  "os"
  "testing"
)

func TestDecode(t *testing.T) {
  file, err := os.Open("./test/p6.ppm")

  if err != nil {
    t.Error("Error opening ppm file", err)
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
