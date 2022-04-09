// 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色(#0000ff)。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var maxZ = f(xyrange*0.5, xyrange*0.5)
var minZ = f(xyrange*-0.5, xyrange*-0.5)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, acolor := corner(i+1, j)
			bx, by, bcolor := corner(i, j)
			cx, cy, ccolor := corner(i, j+1)
			dx, dy, dcolor := corner(i+1, j+1)
			if acolor == "red" || bcolor == "red" || ccolor == "red" || dcolor == "red" {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: red'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else if acolor == "green" || bcolor == "green" || ccolor == "green" || dcolor == "green" {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: green'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if z > maxZ {
		return sx, sy, "red"
	} else if z < minZ {
		return sx, sy, "green"
	} else {
		return sx, sy, "black"
	}
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r

	if math.IsNaN(result) {
		return 0
	}

	return result
}
