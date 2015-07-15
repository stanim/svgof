// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo4f" uses float coordinates, which are formatted as "%.4f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// html5logo draws the w3c HTML5 logo, with scripting added
// +build !appengine

package main

import (
	"github.com/stanim/svgo/svgo4f"
	"os"
)

func main() {
	// HTML5 logo data from
	// "Understanding and Optimizing Web Graphics", Session 508,
	// Dean Jackson, Apple WWDC 2011
	//
	// Draggable elements via Jeff Schiller's dragsvg Javascript library

	// shield
	var sx = []float64{71, 30, 481, 440, 255}
	var sy = []float64{460, 0, 0, 460, 512}
	// highlight
	var hx = []float64{256, 405, 440, 256}
	var hy = []float64{472, 431, 37, 37}
	// "five"
	var fx = []float64{181, 176, 392, 393, 396, 397, 114, 115, 129, 325, 318, 256, 192, 188, 132, 139, 256, 371, 372, 385, 387, 371}
	var fy = []float64{208, 150, 150, 138, 109, 94, 94, 109, 265, 265, 338, 355, 338, 293, 293, 382, 414, 382, 372, 223, 208, 208}

	canvas := svg.New(os.Stdout)
	width := 512.0
	height := 512.0

	// begin the document with the onload event, and namespace for dragging
	canvas.Start(width, height, `onload="initializeDraggableElements();"`, `xmlns:drag="http://www.codedread.com/dragsvg"`)
	canvas.Title("HTML5 Logo")
	canvas.Rect(0, 0, width, height)                                               // black background
	canvas.Script("application/javascript", "http://www.codedread.com/dragsvg.js") // reference the drag script
	canvas.Polygon(sx, sy, `drag:enable="true"`, canvas.RGB(227, 79, 38))          // draggable shield
	canvas.Polygon(hx, hy, `drag:enable="true"`, canvas.RGBA(255, 255, 255, 0.3))  // draggable highlight
	canvas.Polygon(fx, fy, `drag:enable="true"`, canvas.RGB(219, 219, 219))        // draggable five
	canvas.End()
}