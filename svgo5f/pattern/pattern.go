// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo5f" uses float coordinates, which are formatted as "%.5f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// pattern: test the pattern function
package main

import (
	"fmt"
	"github.com/stanim/svgo/svgo5f"
	"os"
)

func main() {
	canvas := svg.New(os.Stdout)
	w, h := 500.0, 500.0
	pct := 5.0
	pw, ph := (w*pct)/100.0, (h*pct)/100.0
	canvas.Start(w, h)

	// define the pattern
	canvas.Def()
	canvas.Pattern("hatch", 0, 0, pw, ph, "user")
	canvas.Gstyle("fill:none;stroke-width:1")
	canvas.Path(fmt.Sprintf("M0,0 l%.5f,%.5f", pw, ph), "stroke:red")
	canvas.Path(fmt.Sprintf("M%.5f,0 l-%.5f,%.5f", pw, pw, ph), "stroke:blue")
	canvas.Gend()
	canvas.PatternEnd()
	canvas.DefEnd()

	// use the pattern
	canvas.Gstyle("stroke:black; stroke-width:2")
	canvas.Circle(w/2, h/2, h/8, "fill:url(#hatch)")
	canvas.CenterRect((w*4)/5, h/2, h/4, h/4, "fill:url(#hatch)")
	canvas.Gend()
	canvas.End()
}
