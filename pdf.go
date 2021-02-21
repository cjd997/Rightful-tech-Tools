package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func createAppealsPdf(d *Datas) {
	fileName := "appeals.pdf"
	title := "APPEALS: APPELLANT"

	pageSize := gopdf.PageSizeA4

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *pageSize})
	pdf.AddPage()

	// draw barrier
	drawBarrier(pdf, *pageSize)

	// write title
	writeTitle(pdf, *pageSize, title)

	// write sub title
	subTitleArr := [3]string{"CLIENT / SOLICITOR", "BARRISTER", "FIXED COST PER ACTIVITY"}
	writeSubTitle(pdf, *pageSize, subTitleArr)

	// write data to pdf
	writeData(pdf, *pageSize, d)

	//  write pdf file
	err := pdf.WritePdf(fileName)
	if err != nil {
		panic(err)
	}
}

func writeTitle(pdf *gopdf.GoPdf, pageSize gopdf.Rect, title string) {
	// setup family font
	err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
	if err != nil {
		log.Print(err.Error())
		return
	}

	// write title
	err = pdf.SetFont("Poppins", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX((pageSize.W / 5.0) * 2.0)
	_ = pdf.Cell(nil, title)
}

func writeSubTitle(pdf *gopdf.GoPdf, pageSize gopdf.Rect, title [3]string) {
	// setup family font
	err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
	if err != nil {
		log.Print(err.Error())
		return
	}

	// write sub title client/solicitor
	err = pdf.SetFont("Poppins", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}

	subTitleY := (pageSize.H / 20.0) * 1.5

	clientSolicitor := title[0]
	pdf.SetX((pageSize.W / 10.0) * 1.0)
	pdf.SetY(subTitleY)
	_ = pdf.Cell(nil, clientSolicitor)

	// write sub title BARRISTER
	barrister := title[1]
	pdf.SetX((pageSize.W / 10.0) * 6.0)
	pdf.SetY(subTitleY)
	_ = pdf.Cell(nil, barrister)

	// write sub title FIXED COST PER ACTIVITY
	cost := title[2]
	pdf.SetX((pageSize.W / 10.0) * 7.0)
	pdf.SetY(subTitleY)
	_ = pdf.Cell(nil, cost)
}

func writeData(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *Datas) {
	// setup family font
	err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
	if err != nil {
		log.Print(err.Error())
		return
	}

	// write title
	err = pdf.SetFont("Poppins", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// setup position for left data
	leftDataXPos := (pageSize.W / 10.0) * 1.0
	leftDataYPos := (pageSize.H / 20.0) * 3.0 // same as barrier data top position

	// setup position for right data
	rightDataXPos := (pageSize.W / 2.0) + 17.0 + 35.0 // 17.0 is barrier width
	rightDataYPos := (pageSize.H / 20.0) * 3.0        // same as barrier data top position

	// text color
	pdf.SetTextColor(0, 0, 0)

	for _, dItem := range *d {
		if dItem.Direction == Left {
			leftDataXPos, leftDataYPos = writeDataOnLeft(pdf, pageSize, &dItem, leftDataXPos, leftDataYPos)

		} else {
			rightDataXPos, rightDataYPos = writeDataOnRight(pdf, pageSize, &dItem, rightDataXPos, rightDataYPos)

		}
	}

}

func writeDataOnRight(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *DatasItem, x float64, y float64) (nextX float64, nextY float64) {

	// draw arrow
	drawArrow(pdf, x, y, -35, color{255, 0, 0})

	// write text data
	xText := x + 18.0
	yText := y - 18.0
	pdf.SetX(xText)
	texts, _ := pdf.SplitText(d.Text, 150)
	for i, text := range texts {
		pdf.SetX(xText)
		pdf.SetY(yText + float64(i+1)*10.0)
		_ = pdf.Cell(nil, text)
	}

	// draw cost rectangle
	nextX = x
	nextY = y + (float64(len(texts)) * 10.0) + 20.0
	return nextX, nextY

}

func writeDataOnLeft(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *DatasItem, x float64, y float64) (nextX float64, nextY float64) {
	// draw arrow
	drawArrow(pdf, x, y, -35, color{255, 0, 0})

	// write text data
	pdf.SetX(x)
	texts, _ := pdf.SplitText(d.Text, 150)
	for i, text := range texts {
		pdf.SetX(x)
		pdf.SetY(y + float64(i+1)*10.0)
		_ = pdf.Cell(nil, text)
	}

	// draw cost rectangle
	nextX = x
	nextY = y + (float64(len(texts)) * 10.0) + 20.0
	return nextX, nextY
}

// func createAppealsPdf(d *Datas) {
// 	fileName := "appeals.pdf"

// 	pageSize := gopdf.PageSizeA4

// 	pdf := &gopdf.GoPdf{}
// 	pdf.Start(gopdf.Config{PageSize: *pageSize})
// 	pdf.AddPage()

// 	// Draw Arrow
// 	drawArrow(pdf, 80, 30, -35, color{255, 0, 0})

// 	// Draw barrier (vertical rectangle)
// 	drawBarrier(pdf, *pageSize)

// 	// Draw text
// 	err := pdf.AddTTFFont(font("Poppins", "Poppins-Regular"))
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}

// 	err = pdf.SetFont("Poppins", "", 10)
// 	if err != nil {
// 		log.Print(err.Error())
// 		return
// 	}
// 	pdf.SetX(80 + 35 + 0)
// 	pdf.SetY(30)
// 	pdf.SetStrokeColor(0, 0, 0)
// 	pdf.SetFillColor(0, 0, 0)

// 	texts, _ := pdf.SplitText(`Within 21 days after date judgment
// pronounced or leave granted, or by the
// date fixed by the court appealed from, the
// appellant to file a noice of appearl and
// serve it on each person who was a party
// in the proceeding in the court appealed
// from or given leave to intervene.`, 150)
// 	for i, text := range texts {
// 		_ = pdf.Cell(nil, text)
// 		pdf.SetX(80 + 35 + 0)
// 		pdf.SetY(30.0 + float64(i+1)*12.0)
// 	}

// 	err = pdf.WritePdf(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// }
