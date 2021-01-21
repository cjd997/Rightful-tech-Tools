package main

import (
	"math"

	"github.com/signintech/gopdf"
)

func drawArrow(pdf *gopdf.GoPdf, s, e gopdf.Point, c color) {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(0, 0, 0)

	pdf.Line(s.X, s.Y, e.X, e.Y)
	pdf.Polygon(arrowHead(s, e), "DF")

	// Draw circle
	pdf.SetLineWidth(1)
	pdf.SetStrokeColor(c.r, c.g, c.b)
	pdf.SetFillColor(c.r, c.g, c.b)
	r := 5.0
	pdf.Oval(s.X, s.Y-0.5*r, s.X+r, s.Y+0.5*r) // Assume horizontal line for arrow
}

// https://math.stackexchange.com/questions/1314006/drawing-an-arrow
func arrowHead(s, e gopdf.Point) []gopdf.Point { //}   (gopdf.Point, gopdf.Point) {
	x1 := s.X
	x2 := e.X
	y1 := s.Y
	y2 := e.Y

	dx := x1 - x2
	dy := y1 - y2

	L1 := math.Sqrt(dx*dx + dy*dy)
	L2 := 5.0
	𝜃 := math.Pi / 6.0

	x3 := x2 + (L2/L1)*(dx*math.Cos(𝜃)+dy*math.Sin(𝜃))
	y3 := y2 + (L2/L1)*(dy*math.Cos(𝜃)-dx*math.Sin(𝜃))
	x4 := x2 + (L2/L1)*(dx*math.Cos(𝜃)-dy*math.Sin(𝜃))
	y4 := y2 + (L2/L1)*(dy*math.Cos(𝜃)+dx*math.Sin(𝜃))

	return []gopdf.Point{e, gopdf.Point{X: x3, Y: y3}, gopdf.Point{X: x4, Y: y4}}
}
