// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo2f" uses float coordinates, which are formatted as "%.2f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// cube: draw cubes
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"

	"github.com/stanim/svgo/svgo2f"
)

var canvas = svg.New(os.Stdout)

// randcolor returns a random color
func randcolor() string {
	rgb := []byte{0, 0, 0} // read error returns black
	rand.Read(rgb)
	return fmt.Sprintf("fill:rgb(%d,%d,%d)", rgb[0], rgb[1], rgb[2])
}

// rcube makes a cube with three visible faces, each with a random color
func rcube(x, y, l float64) {
	tx := []float64{x, x + (l * 3.0), x, x - (l * 3.0), x}
	ty := []float64{y, y + (l * 2.0), y + (l * 4.0), y + (l * 2.0), y}

	lx := []float64{x - (l * 3.0), x, x, x - (l * 3.0), x - (l * 3.0)}
	ly := []float64{y + (l * 2.0), y + (l * 4.0), y + (l * 8.0), y + (l * 6.0), y + (l * 2.0)}

	rx := []float64{x + (l * 3.0), x + (l * 3.0), x, x, x + (l * 3.0)}
	ry := []float64{y + (l * 2.0), y + (l * 6.0), y + (l * 8.0), y + (l * 4.0), y + (l * 2.0)}

	canvas.Polygon(tx, ty, randcolor())
	canvas.Polygon(lx, ly, randcolor())
	canvas.Polygon(rx, ry, randcolor())
}

// lattice draws a grid of cubes, n rows deep.
// The grid begins at (xp, yp), with hspace between cubes in a row, and vspace between rows.
func lattice(xp, yp, w, h, size, hspace, vspace, n float64, bgcolor string) {
	if bgcolor == "" {
		canvas.Rect(0, 0, w, h, randcolor())
	} else {
		canvas.Rect(0, 0, w, h, "fill:"+bgcolor)
	}
	y := yp
	for r := 0; r < int(n); r++ {
		for x := xp; x < w; x += hspace {
			rcube(x, y, size)
		}
		y += vspace
	}
}

func main() {
	var (
		width  = flag.Float64("w", 600, "canvas width")
		height = flag.Float64("h", 600, "canvas height")
		x      = flag.Float64("x", 60, "begin x location")
		y      = flag.Float64("y", 60, "begin y location")
		size   = flag.Float64("size", 20, "cube size")
		rows   = flag.Float64("rows", 3, "rows")
		hs     = flag.Float64("hs", 120, "horizontal spacing")
		vs     = flag.Float64("vs", 160, "vertical spacing")
		bg     = flag.String("bg", "", "background")
	)
	flag.Parse()
	canvas.Start(*width, *height)
	lattice(*x, *y, *width, *height, *size, *hs, *vs, *rows, *bg)
	canvas.End()
}
