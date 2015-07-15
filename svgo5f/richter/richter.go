// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo5f" uses float coordinates, which are formatted as "%.5f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// richter -- inspired by Gerhard Richter's 256 colors, 1974
// +build !appengine

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/stanim/svgo/svgo5f"
)

var canvas = svg.New(os.Stdout)

var width = 700.0
var height = 400.0

func main() {
	rand.Seed(int64(time.Now().Nanosecond()) % int64(1e9))
	canvas.Start(width, height)
	canvas.Title("Richter")
	canvas.Rect(0, 0, width, height, "fill:white")
	rw := 32.0
	rh := 18.0
	margin := 5.0
	for i, x := 0, 20.0; i < int(16); i++ {
		x += (rw + margin)
		for j, y := 0, 20.0; j < int(16); j++ {
			canvas.Rect(x, y, rw, rh, canvas.RGB(int(rand.Float64()*255), int(rand.Float64()*255), int(rand.Float64()*255)))
			y += (rh + margin)
		}
	}
	canvas.End()
}
