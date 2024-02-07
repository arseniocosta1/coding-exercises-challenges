// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		samplingSize           = 2
		f                      = 1.0 / float64(2*samplingSize)
	)

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

					y := suby/height*(ymax-ymin) + ymin
					x := subx/width*(xmax-xmin) + xmin
					z := complex(x, y)
					subpixelColors[sy*samplingSize+sx] = mandelbrot(z)
				}
			}

			// Average the colors of the 4 subpixel
			avgColor := averageColors(subpixelColors)
			img.Set(px, py, avgColor)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

func mandelbrot(z complex128) color.Color {
	const iterations = 255
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

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
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
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
