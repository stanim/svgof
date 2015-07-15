// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo6f" uses float coordinates, which are formatted as "%.6f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// funnel draws a funnel-like shape
// +build !appengine

package main

import (
	"os"

	"github.com/stanim/svgo/svgo6f"
)

var canvas = svg.New(os.Stdout)
var width = 320.0
var height = 480.0

func funnel(bg int, fg int, grid float64, dim float64) {
	h := dim / 2.0
	canvas.Rect(0, 0, width, height, canvas.RGB(bg, bg, bg))
	for size := grid; size < width; size += grid {
		canvas.Ellipse(h, size, size/2, size/2, canvas.RGBA(fg, fg, fg, 0.2))
	}
}

func main() {
	canvas.Start(width, height)
	canvas.Title("Funnel")
	funnel(0, 255, 25, width)
	canvas.End()
}