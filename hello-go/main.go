package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"math/cmplx"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		const (
			xmin, ymin, xmax, ymax = -2, -2, +2, +2
			width, height          = 4096, 4096
		)
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		jpeg.Encode(w, img, &jpeg.Options{Quality: 100}) // NOTE: ignoring errors
		// w.Header().Set("Content-Type", "text/plain")
		// fmt.Fprintln(w, "Hello Go!")
	})
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {}
