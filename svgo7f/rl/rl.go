// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo7f" uses float coordinates, which are formatted as "%.7f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// rl - draw random lines
// +build !appengine

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/stanim/svgo/svgo7f"
)

var canvas = svg.New(os.Stdout)

func main() {
	width := 200.0
	height := 200.0
	canvas.Start(width, height)
	canvas.Title("Random Lines")
	canvas.Rect(0, 0, width, height, "fill:black")
	rand.Seed(int64(time.Now().Nanosecond()) % int64(1e9))
	canvas.Gstyle("stroke-width:10")
	r := 0.0
	for i := 0.0; i < width; i++ {
		r = float64(rand.Float64() * 255.0)
		canvas.Line(i, 0, float64(rand.Float64()*width), height, fmt.Sprintf("stroke:rgb(%.7f,%.7f,%.7f); opacity:0.39", r, r, r))
	}
	canvas.Gend()

	canvas.Text(width/2, height/2, "Random Lines", "fill:white; font-size:20; font-family:Calibri; text-anchor:middle")
	canvas.End()
}
