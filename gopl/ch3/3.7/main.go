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
					subpixelColors[sy*samplingSize+sx] = newton(z)
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

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 255
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
