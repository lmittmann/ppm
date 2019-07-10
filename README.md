# Package ppm [![GoDoc](https://godoc.org/github.com/lmittmann/ppm?status.svg)](https://godoc.org/github.com/lmittmann/ppm) [![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/ppm)](https://goreportcard.com/report/github.com/lmittmann/ppm)


```
import "github.com/lmittmann/ppm"
```
Package ppm implements a Portable Pixel Map (PPM) image decoder and encoder.

The PPM specification is at http://netpbm.sourceforge.net/doc/ppm.html.


## func [Decode](reader.go#L27)
<pre>
func Decode(r <a href="https://godoc.org/io">io</a>.<a href="https://godoc.org/io#Reader">Reader</a>) (<a href="https://godoc.org/image">image</a>.<a href="https://godoc.org/image#Image">Image</a>, <a href="https://godoc.org/builtin#error">error</a>)
</pre>
Decode reads a PPM image from Reader r and returns it as an image.Image.


## func [DecodeConfig](reader.go#L38)
<pre>
func DecodeConfig(r <a href="https://godoc.org/io">io</a>.<a href="https://godoc.org/io#Reader">Reader</a>) (<a href="https://godoc.org/image">image</a>.<a href="https://godoc.org/image#Config">Config</a>, <a href="https://godoc.org/builtin#error">error</a>)
</pre>
DecodeConfig returns the color model and dimensions of a PPM image without decoding the entire image.


## func [Encode](writer.go#L15)
<pre>
func Encode(w <a href="https://godoc.org/io">io</a>.<a href="https://godoc.org/io#Writer">Writer</a>, img <a href="https://godoc.org/image">image</a>.<a href="https://godoc.org/image#Image">Image</a>) <a href="https://godoc.org/builtin#error">error</a>
</pre>
Encode writes the Image img to Writer w in PPM format.
