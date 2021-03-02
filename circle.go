package main

import (
	"github.com/signintech/gopdf"
)

func drawCircle(pdf *gopdf.GoPdf, sX, sY float64, R int, c color) {
	pdf.SetLineWidth(1)
	pdf.SetStrokeColor(c.r, c.g, c.b)
	pdf.SetFillColor(c.r, c.g, c.b)

	// R := 12
	for r := 1; r < R; r++ {
		pdf.Oval(sX, sY-0.5*float64(r), sX+float64(r), sY+0.5*float64(r)) // Assume horizontal line for arrow
	}
}
