package main

import (
	"github.com/signintech/gopdf"
)

// Draw barrier (vertical rectangle)
func drawBarrier(pdf *gopdf.GoPdf, pageSize gopdf.Rect) {
	thick := 17.0

	pdf.SetLineWidth(thick)
	pdf.SetStrokeColor(177, 178, 177)
	pdf.Line(pageSize.W/2.0, (pageSize.H/20.0)*3.0, pageSize.W/2.0, (pageSize.H/20.0)*19.0)
}
