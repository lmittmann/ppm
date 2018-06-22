# Package ppm [![GoDoc](https://godoc.org/github.com/lmittmann/ppm?status.svg)](https://godoc.org/github.com/lmittmann/ppm) [![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/ppm)](https://goreportcard.com/report/github.com/lmittmann/ppm)


```go
import "github.com/lmittmann/ppm"
```
Package ppm implements a Portable Pixel Map (PPM) image decoder and encoder.

The PPM specification is at http://netpbm.sourceforge.net/doc/ppm.html.


## func [Decode](reader.go#L27)
```go
func Decode(r io.Reader) (image.Image, error)
```
Decode reads a PPM image from Reader r and returns it as an image.Image.


## func [DecodeConfig](reader.go#L38)
```go
func DecodeConfig(r io.Reader) (image.Config, error)
```
DecodeConfig returns the color model and dimensions of a PPM image without decoding the entire image.


## func [Encode](writer.go#L15)
```go
func Encode(w io.Writer, img image.Image) error
```
Encode writes the Image img to Writer w in PPM format.
