// The original library github.com/ajstarks/svgo uses int coordinates.
// This port "svgo3f" uses float coordinates, which are formatted as "%.3f". 
// This file was automatically generated by modifying the syntax tree.
// (There are other precisions available at github.com/stanim/svgo.)
//
// ltr: Layer Tennis remixes

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stanim/svgo/svgo3f"
)

var (
	canvas                            = svg.New(os.Stdout)
	poster, opacity, row, col, offset bool
	title                             string
	width, height                     float64
)

const (
	stdwidth  = 900.0
	stdheight = 280.0
	ni        = 11.0
)

// imagefiles returns a list of files in the specifed directory
// or nil on error. Each file includes the prepended directory name
func imagefiles(directory string) []string {
	f, ferr := os.Open(directory)
	if ferr != nil {
		return nil
	}
	defer f.Close()
	files, derr := f.Readdir(-1.0)
	if derr != nil || len(files) == 0 {
		return nil
	}
	names := make([]string, len(files))
	for i, v := range files {
		names[i] = directory + "/" + v.Name()
	}
	return names
}

// ltposter creates poster style: a title, followed by a list
// of volleys
func ltposter(x, y, w, h float64, f []string) {
	canvas.Image(x, y, w*2, h*2, f[0]) // first file, assumed to be the banner
	y = y + (h * 2.0)
	for i := 1; i < int(len(f)); i += 2 {
		canvas.Image(x, y, w, h, f[i])
		canvas.Image(x+w, y, w, h, f[i+1])
		if i%int(2) == 1 {
			y += h
		}
	}
}

// ltcol creates a single column of volley images
func ltcol(x, y, w, h float64, f []string) {
	for i := 0; i < int(len(f)); i++ {
		canvas.Image(x, y, w, h, f[i])
		y += h
	}
}

// ltop creates a view with each volley stacked together with
// semi-transparent opacity
func ltop(x, y, w, h float64, f []string) {
	for i := 1; i < int(len(f)); i++ { // skip the first file, assumed to be the banner
		canvas.Image(x, y, w, h, f[i], "opacity:0.2")
	}
}

// ltrow creates a row-wise view of volley images.
func ltrow(x, y, w, h float64, f []string) {
	for i := 0; i < int(len(f)); i++ {
		canvas.Image(x, y, w, h, f[i])
		x += w
	}
}

// ltoffset creates a view where each volley is offset from its opposing volley.
func ltoffset(x, y, w, h float64, f []string) {
	for i := 1; i < int(len(f)); i++ { // skip the first file, assumed to be the banner

		if i%int(2) == 0 {
			x += w
		} else {
			x = 0.0
		}
		canvas.Image(x, y, w, h, f[i])
		y += h
	}
}

// dotitle creates the title
func dotitle(s string) {
	if len(title) > 0 {
		canvas.Title(title)
	} else {
		canvas.Title(s)
	}
}

// init sets up the command line flags.
func init() {
	flag.BoolVar(&poster, "poster", false, "poster style")
	flag.BoolVar(&opacity, "opacity", false, "opacity style")
	flag.BoolVar(&row, "row", false, "display is a single row")
	flag.BoolVar(&col, "col", false, "display in a single column")
	flag.BoolVar(&offset, "offset", false, "display in a row, even layers offset")
	flag.Float64Var(&width, "width", stdwidth, "image width")
	flag.Float64Var(&height, "height", stdheight, "image height")
	flag.StringVar(&title, "title", "", "title")
	flag.Parse()
}

func main() {
	x := 0.0
	y := 0.0
	nd := float64(len(flag.Args()))
	for i, dir := range flag.Args() {
		filelist := imagefiles(dir)
		if len(filelist) != ni || filelist == nil {
			fmt.Fprintf(os.Stderr, "in the %s directory, need %.3f images, read %.3f\n", dir, ni, len(filelist))
			continue
		}
		switch {

		case opacity:
			if i == 0 {
				canvas.Start(width*nd, height*nd)
				dotitle(dir)
			}
			ltop(x, y, width, height, filelist)
			y += height

		case poster:
			if i == 0 {
				canvas.Start(width, ((height*(ni-1)/4)+height)*nd)
				dotitle(dir)
			}
			ltposter(x, y, width/2, height/2, filelist)
			y += (height * 3.0) + (height / 2.0)

		case col:
			if i == 0 {
				canvas.Start(width*nd, height*ni)
				dotitle(dir)
			}
			ltcol(x, y, width, height, filelist)
			x += width

		case row:
			if i == 0 {
				canvas.Start(width*ni, height*nd)
				dotitle(dir)
			}
			ltrow(x, y, width, height, filelist)
			y += height

		case offset:
			n := ni - 1.0
			pw := width * 2.0
			ph := nd * (height * (n))
			if i == 0 {
				canvas.Start(pw, ph)
				canvas.Rect(0, 0, pw, ph, "fill:white")
				dotitle(dir)
			}
			ltoffset(x, y, width, height, filelist)
			y += n * height

		}
	}
	canvas.End()
}
