// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
// をクライアントに書き出すWebサーバを作成する
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
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var (
	width, height = 600, 320              // canvas size in pixels
	xyscale       = width / 2 / xyrange   // pixels per x or y unit
	zscale        = float64(height) * 0.4 // pixels per z unit
	color         = "white"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if h := q.Get("height"); h != "" {
			height, _ = strconv.Atoi(h)
			zscale = float64(height) * 0.4 // pixels per z unit
		}
		if wi := q.Get("weight"); wi != "" {
			width, _ = strconv.Atoi(wi)
			xyscale = width / 2 / xyrange // pixels per x or y unit
		}
		if c := q.Get("color"); c != "" {
			color = c
			log.Printf("%s", c)
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		svgreder(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svgreder(out io.Writer) {
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	out.Write([]byte(str))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			str = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			out.Write([]byte(str))
		}
	}
	str = fmt.Sprintln("</svg>")
	out.Write([]byte(str))
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*float64(zscale)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
