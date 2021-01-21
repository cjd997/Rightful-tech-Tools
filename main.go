package main

import (
	// "log"

	"github.com/signintech/gopdf"
)

func main() {

	pageSize := gopdf.PageSizeA4

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *pageSize})
	pdf.AddPage()

	// Draw Arrow
	drawArrow(pdf, gopdf.Point{X: 80, Y: 120}, gopdf.Point{X: 50, Y: 30}, color{255, 0, 0})

	// Draw barrier (vertical rectangle)
	drawBarrier(pdf, *pageSize)

	// err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }

	// err = pdf.SetFont("Poppins", "", 14)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }
	// pdf.Cell(nil, "Hello")
	err := pdf.WritePdf("hello.pdf")
	if err != nil {
		panic(err)
	}
}
