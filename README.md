# pnm #

pnm is a simple Go-library for writing one-dimensional arrays to binary Portable Anymap files, such as:

- Portable Bitmap (.pbm)
- Portable Graymap (.pgm)
- Portable Pixmap (.ppm)

## Installation ##
```
go get github.com/lmittmann/pnm
```

## Usage ##

```go
import "github.com/lmittmann/pnm"
```

## Example ##

```go
package main

import (
	"github.com/lmittmann/pnm"
	"math"
)

func main() {
	width, height := 256, 256
	pixel := make([]uint8, width*height*3)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := (y*width + x) * 3 // position of pixel at x, y in array
			pixel[i] = 255                                                                // red
			pixel[i+1] = uint8(math.Sin(float64(x)/16)*math.Sin(float64(y)/16)*128 + 128) // green
			pixel[i+2] = 0                                                                // blue
		}
	}
	err := pnm.WritePixelmap("test.ppm", &pixel, width, height)
	if err != nil {
		panic(err)
	}
}
```
