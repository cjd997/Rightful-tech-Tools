package main

import (
	"github.com/signintech/gopdf"
)

type Barrier struct {
	Thick   float64
	Color   color
	TopX    float64
	TopY    float64
	BottomX float64
	BottomY float64
}

func (b *Barrier) DrawBarrier(pdf *gopdf.GoPdf) {
	pdf.SetLineWidth(b.Thick)
	pdf.SetStrokeColor(b.Color.r, b.Color.g, b.Color.b)
	pdf.Line(b.TopX, b.TopY, b.BottomX, b.BottomY)
}
