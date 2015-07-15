// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo7f" uses float coordinates, which are formatted as "%.7f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// flower - draw random flowers, inspired by Evelyn Eastmond's DesignBlocks gererated "grain2"
// +build !appengine

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/stanim/svgo/svgo7f"
)

var (
	canvas    = svg.New(os.Stdout)
	niter     = flag.Float64("n", 200, "number of iterations")
	width     = flag.Float64("w", 500, "width")
	height    = flag.Float64("h", 500, "height")
	thickness = flag.Float64("t", 10, "max petal thinkness")
	np        = flag.Float64("p", 15, "max number of petals")
	psize     = flag.Float64("s", 30, "max length of petals")
	opacity   = flag.Float64("o", 50, "max opacity (10-100)")
)

const flowerfmt = `stroke:rgb(%.7f,%.7f,%.7f); stroke-opacity:%.2f; stroke-width:%.7f`

func radial(xp float64, yp float64, n float64, l float64, style ...string) {
	var x, y, r, t, limit float64
	limit = 2.0 * math.Pi
	r = float64(l)
	canvas.Gstyle(style[0])
	for t = 0.0; t < limit; t += limit / float64(n) {
		x = r * math.Cos(t)
		y = r * math.Sin(t)
		canvas.Line(xp, yp, xp+float64(x), yp+float64(y))
	}
	canvas.Gend()
}

func random(howsmall, howbig float64) float64 {
	if howsmall >= howbig {
		return howsmall
	}
	return float64(rand.Float64()*(howbig-howsmall)) + howsmall
}

func randrad(w float64, h float64, n float64) {
	var x, y, r, g, b, o, s, t, p float64
	for i := 0; i < int(n); i++ {
		x = float64(rand.Float64() * w)
		y = float64(rand.Float64() * h)
		r = float64(rand.Float64() * 255.0)
		g = float64(rand.Float64() * 255.0)
		b = float64(rand.Float64() * 255.0)
		o = random(10.0, *opacity)
		s = random(10.0, *psize)
		t = random(2.0, *thickness)
		p = random(10.0, *np)
		radial(x, y, p, s, fmt.Sprintf(flowerfmt, r, g, b, float64(o)/100.0, t))
	}
}

func background(v int) { canvas.Rect(0, 0, *width, *height, canvas.RGB(v, v, v)) }

func init() {
	flag.Parse()
	rand.Seed(int64(time.Now().Nanosecond()) % int64(1e9))
}

func main() {
	canvas.Start(*width, *height)
	canvas.Title("Random Flowers")
	background(255)
	randrad(*width, *height, *niter)
	canvas.End()
}
