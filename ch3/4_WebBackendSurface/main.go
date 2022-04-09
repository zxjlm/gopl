// 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：
//
// w.Header().Set("Content-Type", "image/svg+xml")
//
//（这一步在Lissajous例子中不是必须的，因为服务器使用标准的PNG图像格式，可以根据前面的512个字节自动输出对应的头部。）允许客户端通过HTTP请求参数设置高度、宽度和颜色等参数。

package main

import (
	"fmt"
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
func main() {
	// example:  http://127.0.0.1:8001/drawer?color=red&height=400&weight=400
	http.HandleFunc("/drawer", drawer)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func drawer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	currentWeight, err := strconv.Atoi(r.URL.Query().Get("weight"))
	if err != nil {
		currentWeight = width
	}
	currentHeight, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		currentHeight = height
	}
	currentColor := r.URL.Query().Get("color")
	if currentColor == "" {
		currentColor = "black"
	}

	svgStr := ""
	svgStr += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, currentWeight, currentHeight)
			bx, by := corner(i, j, currentWeight, currentHeight)
			cx, cy := corner(i, j+1, currentWeight, currentHeight)
			dx, dy := corner(i+1, j+1, currentWeight, currentHeight)
			svgStr += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, currentColor)
		}
	}
	svgStr += fmt.Sprintf("</svg>")
	_, err = fmt.Fprintf(w, svgStr)
	if err != nil {
		return
	}
}

func corner(i, j, w, h int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(w/2) + (x-y)*cos30*xyscale
	sy := float64(h/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r

	if math.IsNaN(result) {
		return 0
	}

	return result
}
