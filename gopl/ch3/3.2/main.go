// Surface computes an SVG rendering of a 3-D surface function.
// we no longer need the debug flag bellow since we updated go beyond 1.22
// but im keeping it here for reference
//
//go:debug httpmuxgo121=0
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
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

type ShapeFunc func(x, y float64) float64

func main() {
	surface := func(w io.Writer, f ShapeFunc) {
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, ok1 := corner(i+1, j, f)
				bx, by, ok2 := corner(i, j, f)
				cx, cy, ok3 := corner(i, j+1, f)
				dx, dy, ok4 := corner(i+1, j+1, f)

				if !ok1 || !ok2 || !ok3 || !ok4 {
					continue
				}

				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)

			}
		}
		fmt.Fprintf(w, "</svg>")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/{shape}", func(writer http.ResponseWriter, req *http.Request) {
		shape := req.PathValue("shape")

		i, err := strconv.Atoi(shape)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "Invalid shape parameter: %s", shape)
			return
		}

		var fun ShapeFunc
		switch i {
		case 0:
			fun = f
		case 1:
			fun = feggbox
		case 2:
			fun = fmoguls
		case 3:
			fun = fsaddle
		case 4:
			fun = sinePattern

		default:
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, "Unknown shape: %d", i)
			return
		}

		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer, fun)
	})

	log.Fatal(http.ListenAndServe("localhost:8001", mux))
}

func corner(i, j int, f ShapeFunc) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func feggbox(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 5
}

func fmoguls(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 30
}

func sinePattern(x, y float64) float64 {
	return math.Sin(math.Hypot(x, y)) / 5
}

func fsaddle(x, y float64) float64 {
	return (math.Pow(x, 2) - math.Pow(y, 2)) / 500
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
