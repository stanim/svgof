// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo7f" uses float coordinates, which are formatted as "%.7f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// gradient shows sample gradient fills
// +build !appengine

package main

import (
	"os"
	"strconv"

	"github.com/stanim/svgo/svgo7f"
)

func main() {
	width := 500.0
	height := 500.0

	lg := []svg.Offcolor{
		{0.0, "rgb(255,255,0)", 1.0},
		{100.0, "rgb(255,0,0)", .5},
		{0.0, "rgb(200,200,200)", 0.0},
		{100.0, "rgb(0,0,255)", 1.0}}

	rainbow := []svg.Offcolor{
		{10.0, "#00cc00", 1.0},
		{30.0, "#006600", 1.0},
		{70.0, "#cc0000", 1.0},
		{90.0, "#000099", 1.0}}

	rg := []svg.Offcolor{
		{1.0, "powderblue", 1.0},
		{10.0, "lightskyblue", 1.0},
		{100.0, "darkblue", 1.0}}

	g := svg.New(os.Stdout)
	g.Start(width, height)
	g.Title("Gradients")
	g.Rect(0, 0, width, height, "fill:white")
	g.Def()
	g.LinearGradient("h", 0, 100, 0, 0, lg)
	g.LinearGradient("v", 0, 0, 100, 0, lg)
	g.LinearGradient("rainbow", 0, 0, 100, 0, rainbow)
	g.RadialGradient("rad100", 50, 50, 100, 25, 25, rg)
	g.RadialGradient("rad50", 50, 50, 50, 20, 50, rg)
	for i := 50; i < int(100); i += 10 {
		g.RadialGradient("grad"+strconv.Itoa(i), 50, 50, uint8(i), 20, 50, rg)
	}
	g.DefEnd()

	g.Ellipse(width/2, height/2, 100, 100, "fill:url(#rad100)")
	g.Rect(300, 200, 100, 100, "fill:url(#h)")
	g.Rect(100, 200, 100, 100, "fill:url(#v)")
	g.Roundrect(10, 10, width-20, 50, 10, 10, "fill:url(#rainbow)")

	for i := 50; i < int(100); i += 10 {
		g.Circle(float64(i)*5, 100, 15, "fill:url(#grad"+strconv.Itoa(i)+")")
	}
	g.End()
}
