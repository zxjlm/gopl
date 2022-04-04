// 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?cycles=20 这个URL，
//这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以调用strconv.Atoi函数。
//你可以在godoc里查看strconv.Atoi的详细说明。

package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", lissajousService)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousService(w http.ResponseWriter, r *http.Request) {
	const (
		res        = 0.001 // angular resolution
		size       = 100   // image canvas covers [-size..+size]
		nframes    = 64    // number of animation frames
		delay      = 8     // delay between frames in 10ms units
		blackIndex = 1     // next color in palette
	)

	var palette = []color.Color{color.White, color.Black}

	cycles, _ := strconv.Atoi(r.FormValue("cycles"))
	if cycles == 0 {
		cycles = 5
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(w, &anim)
	if err != nil {
		return
	} // NOTE: ignoring encoding errors
}
