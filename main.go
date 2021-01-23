package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pageSize := gopdf.PageSizeA4

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *pageSize})
	pdf.AddPage()

	// Draw Arrow
	drawArrow(pdf, 80, 30, -35, color{255, 0, 0})

	// Draw barrier (vertical rectangle)
	drawBarrier(pdf, *pageSize)

	// Draw text
	err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Poppins", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(80 + 35 + 5)
	pdf.SetY(30)
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(0, 0, 0)

	texts, _ := pdf.SplitText(`Within 21 days after date judgment
pronounced or leave granted, or by the
date fixed by the court appealed from, the
appellant to file a noice of appearl and
serve it on each person who was a party
in the proceeding in the court appealed
from or given leave to intervene.`, 480)
	for i, text := range texts {
		_ = pdf.Cell(nil, text)
		pdf.SetY(30.0 + float64(i)*24.0)
	}

	err = pdf.WritePdf("hello.pdf")
	if err != nil {
		panic(err)
	}
}
