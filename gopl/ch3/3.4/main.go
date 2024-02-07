package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// colorToHex converts an RGBA color to a hex string.
func colorToHex(c color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func main() {
	surface := func(w io.Writer) {
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

		// uncomment to get min and max to use for interpolation
		// only required if xyrange, or cells are changed from the default
		// since func f arguments depend on those indirectly
		//mi := math.Inf(1)
		//ma := math.Inf(-1)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, z1, ok1 := corner(i+1, j)
				bx, by, z2, ok2 := corner(i, j)
				cx, cy, z3, ok3 := corner(i, j+1)
				dx, dy, z4, ok4 := corner(i+1, j+1)

				if !ok1 || !ok2 || !ok3 || !ok4 {
					continue
				}

				red := color.RGBA{R: 0xff, A: 0xff}
				blue := color.RGBA{B: 0xff, A: 0xff}

				a := -0.21722891503668823
				b := 0.9850673555377986
				avg := (z1 + z2 + z3 + z4) / 4.0
				t := (avg - a) / (b - a)

				// uncomment to get min and max to use for interpolation
				//mi = math.Min(mi, z1)
				//ma = math.Max(ma, z1)

				color := lerp(blue, red, t)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' stroke='%[9]s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, colorToHex(color))

			}
		}

		fmt.Fprintf(w, "</svg>")

		// uncomment to get min and max to use for interpolation
		//fmt.Printf("min: %g, max: %g\n", mi, ma)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer)
	})

	port := "8001"
	bindAddress := "localhost"
	log.Printf("Starting server at http://%s:%s\n", bindAddress, port)
	log.Fatal(http.ListenAndServe(bindAddress+":"+port, nil))
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func lerp(start, end color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(start.R) + t*(float64(end.R)-float64(start.R))),
		G: uint8(float64(start.G) + t*(float64(end.G)-float64(start.G))),
		B: uint8(float64(start.B) + t*(float64(end.B)-float64(start.B))),
		A: uint8(float64(start.A) + t*(float64(end.A)-float64(start.A))),
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
