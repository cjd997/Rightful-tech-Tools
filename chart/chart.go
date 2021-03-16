package chart

import (
	"fmt"

	"github.com/signintech/gopdf"
)

type Request struct {
	Text      string  `json:"text"`
	Direction string  `json:"direction"`
	Cost      float64 `json:"cost"`
}

func Generate(req []Request) error {
	return generate()
}

func generate() error {
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
		return fmt.Errorf("error adding font: %s", err.Error())
	}

	err = pdf.SetFont("Poppins", "", 10)
	if err != nil {
		return fmt.Errorf("error setting font: %s", err.Error())
	}
	pdf.SetX(80 + 35 + 0)
	pdf.SetY(30)
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetFillColor(0, 0, 0)

	texts, _ := pdf.SplitText(`Within 21 days after date judgment
pronounced or leave granted, or by the
date fixed by the court appealed from, the
appellant to file a noice of appearl and
serve it on each person who was a party
in the proceeding in the court appealed
from or given leave to intervene.`, 150)
	for i, text := range texts {
		_ = pdf.Cell(nil, text)
		pdf.SetX(80 + 35 + 0)
		pdf.SetY(30.0 + float64(i+1)*12.0)
	}

	err = pdf.WritePdf("hello.pdf")
	if err != nil {
		return fmt.Errorf("error writing chart to pdf file: %s", err.Error())
	}

	return nil
}
