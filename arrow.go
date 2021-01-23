package main

import (
	"math"

	"github.com/signintech/gopdf"
)

func drawArrow(pdf *gopdf.GoPdf, sX, sY, L float64, c color) {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(0, 0, 0)

	pdf.Line(sX, sY, sX+L, sY)
	pdf.Polygon(arrowHead(gopdf.Point{X: sX, Y: sY}, gopdf.Point{X: sX + L, Y: sY}), "DF")

	// Draw circle
	pdf.SetLineWidth(1)
	pdf.SetStrokeColor(c.r, c.g, c.b)
	pdf.SetFillColor(c.r, c.g, c.b)
	R := 12
	for r := 1; r < R; r++ {
		pdf.Oval(sX, sY-0.5*float64(r), sX+float64(r), sY+0.5*float64(r)) // Assume horizontal line for arrow
	}

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
