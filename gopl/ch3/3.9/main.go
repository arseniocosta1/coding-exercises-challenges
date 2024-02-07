package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/png")
		size, z := parseQuery(request.URL)

		x := 0.0
		y := 0.0

		img := newFractal(size, -2/z+x, -2/z+y, 2/z+x, 2/z+y)
		png.Encode(writer, img) // NOTE: ignoring errors
	})

	port := "8081"
	bindAddress := "localhost"
	log.Printf("Listening on http://%s:%s", bindAddress, port)
	log.Fatal(http.ListenAndServe(bindAddress+":"+port, nil))
}

func parseQuery(url *url.URL) (size int, zoom float64) {
	query := url.Query()
	if sizeStr := query.Get("size"); sizeStr != "" {
		size, _ = strconv.Atoi(sizeStr) // NOTE: ignoring errors
	} else {
		size = 1024
	}

	if zoomStr := query.Get("zoom"); zoomStr != "" {
		zoom, _ = strconv.ParseFloat(zoomStr, 64) // NOTE: ignoring errors
	} else {
		zoom = 1.0
	}

	return
}

func newFractal(size int, xmin, ymin, xmax, ymax float64) *image.RGBA {
	const (
		samplingSize = 2
		f            = 1.0 / float64(2*samplingSize)
	)

	width, height := size, size
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	subpixelColors := make([]color.Color, samplingSize*samplingSize)

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			//  super-sampling (4x)
			for sy := 0; sy < samplingSize; sy++ {
				for sx := 0; sx < samplingSize; sx++ {
					// Calculate the subpixel coordinates
					subx := float64(px) + (float64(sx)/float64(samplingSize) - 0.5 + f)
					suby := float64(py) + (float64(sy)/float64(samplingSize) - 0.5 + f)

					y := suby/float64(height)*(ymax-ymin) + ymin
					x := subx/float64(width)*(xmax-xmin) + xmin
					z := complex(x, y)
					subpixelColors[sy*samplingSize+sx] = mandelbrot(z)
				}
			}

			// Average the colors of the 4 subpixel
			avgColor := averageColors(subpixelColors)
			img.Set(px, py, avgColor)
		}
	}

	return img
}

// averageColor calculates the average of four color.Color values.
func averageColors(colors []color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r += cr
		g += cg
		b += cb
		a += ca
	}
	numColors := uint32(len(colors))
	// Convert back to 0-255 range
	return color.RGBA{
		R: uint8((r / numColors) >> 8),
		G: uint8((g / numColors) >> 8),
		B: uint8((b / numColors) >> 8),
		A: uint8((a / numColors) >> 8),
	}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		f := 1e-6
		if cmplx.Abs(z*z*z*z-1) < f {
			// Determine the root
			switch {
			case cmplx.Abs(z-1) < f:
				return color.RGBA{R: 255 - contrast*i, A: 255} // Red for root 1
			case cmplx.Abs(z+1) < f:
				return color.RGBA{G: 255 - contrast*i, A: 255} // Green for root -1
			case cmplx.Abs(z-complex(0, 1)) < f:
				return color.RGBA{B: 255 - contrast*i, A: 255} // Blue for root i
			case cmplx.Abs(z+complex(0, 1)) < f:
				return color.RGBA{R: 255 - contrast*i, G: 255 - contrast*i, A: 255} // Yellow for root -i
			}
		}
	}
	return color.Black
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	n := uint8(0)

	for ; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{G: 255 - contrast*n, A: 255}
		}
	}

	return color.Black
}
