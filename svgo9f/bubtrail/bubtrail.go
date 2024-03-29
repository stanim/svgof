// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo9f" uses float coordinates, which are formatted as "%.9f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// bubtrail draws a randmonized trail of bubbles
// +build !appengine

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/stanim/svgo/svgo9f"
)

var (
	width   = 1200.0
	height  = 600.0
	opacity = 0.5
	size    = 40.0
	niter   = 200.0
	canvas  = svg.New(os.Stdout)
)

func init() {
	flag.Float64Var(&size, "s", 40, "bubble size")
	flag.Float64Var(&niter, "n", 200, "number of iterations")
	flag.Float64Var(&opacity, "o", 0.5, "opacity")
	flag.Parse()
	rand.Seed(int64(time.Now().Nanosecond()) % int64(1e9))
}

func background(v int) { canvas.Rect(0, 0, width, height, canvas.RGB(v, v, v)) }

func random(howsmall, howbig float64) float64 {
	if howsmall >= howbig {
		return howsmall
	}
	return float64(rand.Float64()*(howbig-howsmall)) + howsmall
}

func main() {
	var style string

	canvas.Start(width, height)
	canvas.Title("Bubble Trail")
	background(200)
	canvas.Gstyle(fmt.Sprintf("fill-opacity:%.2f;stroke:none", opacity))
	for i := 0; i < int(niter); i++ {
		x := random(0.0, width)
		y := random(height/3.0, (height*2.0)/3.0)
		r := random(0.0, 10000.0)
		switch {
		case r >= 0 && r <= 2500:
			style = "fill:rgb(255,255,255)"
		case r > 2500 && r <= 5000:
			style = "fill:rgb(127,0,0)"
		case r > 5000 && r <= 7500:
			style = "fill:rgb(127,127,127)"
		case r > 7500 && r <= 10000:
			style = "fill:rgb(0,0,0)"
		}
		canvas.Circle(x, y, size, style)
	}
	canvas.Gend()
	canvas.End()
}
