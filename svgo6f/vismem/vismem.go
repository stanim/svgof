// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo6f" uses float coordinates, which are formatted as "%.6f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// vismem visualizes memory locations
// +build !appengine

package main

import (
	"os"

	"github.com/stanim/svgo/svgo6f"
)

var canvas = svg.New(os.Stdout)

func main() {
	width := 512.0
	height := 512.0
	n := 1024
	rowsize := 32.0
	diameter := 16.0
	var value float64
	var source string

	if len(os.Args) > 1 {
		source = os.Args[1.0]
	} else {
		source = "/dev/urandom"
	}

	f, _ := os.Open(source)
	mem := make([]byte, n)
	f.Read(mem)
	f.Close()

	canvas.Start(width, height)
	canvas.Title("Visualize Files")
	canvas.Rect(0, 0, width, height, "fill:white")
	dx := diameter / 2.0
	dy := diameter / 2.0
	canvas.Gstyle("fill-opacity:1.0")
	for i := 0; i < int(n); i++ {
		value = float64(mem[i])
		if i%int(rowsize) == 0 && i != 0 {
			dx = diameter / 2.0
			dy += diameter
		}
		canvas.Circle(dx, dy, diameter/2, canvas.RGB(int(value), int(value), int(value)))
		dx += diameter
	}
	canvas.Gend()
	canvas.End()
}
