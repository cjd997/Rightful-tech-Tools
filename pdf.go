package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

var (
	pink  color = color{255, 187, 177}
	blue  color = color{178, 214, 239}
	grey  color = color{153, 153, 153}
	black color = color{0, 0, 0}
	white color = color{255, 255, 255}
)

func createAppealsPdf(d *Datas) {
	fileName := "appeals.pdf"
	title := "APPEALS: APPELLANT"

	pageSize := gopdf.PageSizeA4

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *pageSize})
	pdf.AddPage()

	// draw barrier
	b := Barrier{Thick: 17.0,
		Color:   grey,
		TopX:    pageSize.W / 2.0,
		TopY:    (pageSize.H / 20.0) * 3.0,
		BottomX: pageSize.W / 2.0,
		BottomY: (pageSize.H / 20.0) * 19.0}
	b.DrawBarrier(pdf)

	// write title
	writeTitle(pdf, *pageSize, title)

	// write sub title
	subTitleArr := [3]string{"CLIENT / SOLICITOR", "BARRISTER", "FIXED COST PER ACTIVITY"}
	writeSubTitle(pdf, *pageSize, subTitleArr)

	// write initial client consultation
	writeInitClientC(pdf, b)

	// write hearing of appeal
	writeHearingOfAppeal(pdf, b)

	// write data to pdf
	writeData(pdf, *pageSize, d, b)

	// write high level visual
	highLevelXPos := b.BottomX + (b.Thick / 2.0) //(pageSize.W / 2.0) + 17.0 + 35.0 // 17.0 is barrier width
	highLevelYPos := b.BottomY                   //(pageSize.H / 20.0) * 3.0        // same as barrier data top position
	writeHighLevelVisual(pdf, *pageSize, highLevelXPos, highLevelYPos)

	// write key block
	writeKeyBlock(pdf, b)

	//  write pdf file
	err := pdf.WritePdf(fileName)
	if err != nil {
		panic(err)
	}
}

func writeInitClientC(pdf *gopdf.GoPdf, b Barrier) {
	// draw left and right curved rectangle
	leftRect := Rectangle{fillColor: pink,
		strokeColor: pink,
		tlPoint:     gopdf.Point{X: b.TopX - 34, Y: b.TopY - 34},
		trPoint:     gopdf.Point{X: b.TopX, Y: b.TopY - 34},
		brPoint:     gopdf.Point{X: b.TopX, Y: b.TopY},
		blPoint:     gopdf.Point{X: b.TopX - 34, Y: b.TopY}}
	leftCurvedRectangle := CurvedRectangle{rectangle: leftRect, curveLength: 20.0}
	leftCurvedRectangle.DrawLeftCurved(pdf)

	rightRect := Rectangle{fillColor: blue,
		strokeColor: blue,
		tlPoint:     gopdf.Point{X: b.TopX + 34, Y: b.TopY - 34},
		trPoint:     gopdf.Point{X: b.TopX, Y: b.TopY - 34},
		brPoint:     gopdf.Point{X: b.TopX, Y: b.TopY},
		blPoint:     gopdf.Point{X: b.TopX + 34, Y: b.TopY}}
	rightCurvedRectangle := CurvedRectangle{rectangle: rightRect, curveLength: 20.0}
	rightCurvedRectangle.DrawRightCurved(pdf)

	// TODO: right text initial client consultation
}

func writeHearingOfAppeal(pdf *gopdf.GoPdf, b Barrier) {
	// draw left and right curved rectangle
	leftRect := Rectangle{fillColor: pink,
		strokeColor: pink,
		tlPoint:     gopdf.Point{X: b.BottomX - 34, Y: b.BottomY - 68},
		trPoint:     gopdf.Point{X: b.BottomX, Y: b.BottomY - 68},
		brPoint:     gopdf.Point{X: b.BottomX, Y: b.BottomY - 34},
		blPoint:     gopdf.Point{X: b.BottomX - 34, Y: b.BottomY - 34}}
	leftCurvedRectangle := CurvedRectangle{rectangle: leftRect, curveLength: 20.0}
	leftCurvedRectangle.DrawLeftCurved(pdf)

	rightRect := Rectangle{fillColor: blue,
		strokeColor: blue,
		tlPoint:     gopdf.Point{X: b.BottomX + 34, Y: b.BottomY - 68},
		trPoint:     gopdf.Point{X: b.BottomX, Y: b.BottomY - 68},
		brPoint:     gopdf.Point{X: b.BottomX, Y: b.BottomY - 34},
		blPoint:     gopdf.Point{X: b.BottomX + 34, Y: b.BottomY - 34}}
	rightCurvedRectangle := CurvedRectangle{rectangle: rightRect, curveLength: 20.0}
	rightCurvedRectangle.DrawRightCurved(pdf)

	// TODO: right text hearing of appeals
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
	pdf.SetX(((pageSize.W / 10.0) * 8.0) - 12.0)
	pdf.SetY(subTitleY)
	_ = pdf.Cell(nil, cost)

}

func writeData(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *Datas, b Barrier) {
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
	leftDataXPos := b.TopX - (b.Thick / 2.0) //(pageSize.W / 10.0) * 1.0
	//leftDataYPos := b.TopY                   //(pageSize.H / 20.0) * 3.0 // same as barrier data top position

	// setup position for right data
	rightDataXPos := b.TopX + (b.Thick / 2.0) //(pageSize.W / 2.0) + 17.0 + 35.0 // 17.0 is barrier width
	//rightDataYPos := b.TopY                   //(pageSize.H / 20.0) * 3.0        // same as barrier data top position

	dataYPos := b.TopY + 12.0

	// text color
	pdf.SetTextColor(0, 0, 0)

	for _, dItem := range *d {
		if dItem.Direction == Left {
			leftDataXPos, dataYPos = writeDataOnLeft(pdf, pageSize, &dItem, leftDataXPos, dataYPos)
		} else {
			rightDataXPos, dataYPos = writeDataOnRight(pdf, pageSize, &dItem, rightDataXPos, dataYPos)
		}
	}

}

func writeDataOnRight(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *DatasItem, x float64, y float64) (nextX float64, nextY float64) {
	// draw arrow
	drawToLeftArrow(pdf, x+35, y, -35, pink, 12)

	// write text data
	xText := x + 35.0 + 12.0 // 35.0 is length of arrow, and 12.0 is radian on circle arrow
	yText := y - 18.0
	pdf.SetX(xText)
	texts, _ := pdf.SplitText(d.Text, 150)
	for i, text := range texts {
		pdf.SetX(xText)
		pdf.SetY(yText + float64(i+1)*10.0)
		_ = pdf.Cell(nil, text)
	}

	// draw cost rectangle
	xCost := xText + 150.0 + 35.0 // 150.0 is length of text data, adn 35.0 is length of arrow
	yCost := yText + 12.0

	pdf.SetStrokeColor(black.r, black.g, black.b)
	pdf.SetFillColor(white.r, white.g, white.b)
	pdf.Polygon([]gopdf.Point{
		{X: xCost - 10, Y: yCost - 5.0},
		{X: xCost - 10 + 60, Y: yCost - 5.0},
		{X: xCost - 10 + 60, Y: yCost + 5.0 + 12.0},
		{X: xCost - 10, Y: yCost + 5.0 + 12.0}}, "DF")

	pdf.SetX(xCost)
	pdf.SetY(yCost)
	_ = pdf.Cell(nil, "$"+fmt.Sprintf("%.2f", d.Cost))

	nextX = x
	nextY = y + (float64(len(texts)) * 10.0) + 20.0
	return nextX, nextY

}

func writeDataOnLeft(pdf *gopdf.GoPdf, pageSize gopdf.Rect, d *DatasItem, x float64, y float64) (nextX float64, nextY float64) {
	// draw arrow
	drawToRightArrow(pdf, x, y, -35, pink, 12)

	// write text data
	xText := x - 35.0 - 12.0 - 150.0
	yText := y - 18.0
	pdf.SetX(xText) // 35 is length of arrow, and 12 is rad of circle arraow)
	texts, _ := pdf.SplitText(d.Text, 150)
	for i, text := range texts {
		pdf.SetX(xText)
		pdf.SetY(yText + float64(i+1)*10.0)
		_ = pdf.Cell(nil, text)
	}

	nextX = x
	nextY = y + (float64(len(texts)) * 10.0) + 20.0
	return nextX, nextY
}

func writeHighLevelVisual(pdf *gopdf.GoPdf, pageSize gopdf.Rect, x float64, y float64) {
	texts := [4]string{"* Highlevel visual ret. of", "process. Activities can", "change based on a", "case by case basis."}

	// draw arrow
	drawToLeftArrow(pdf, x+45, y, -45, grey, 36)

	// write text data
	xText := x + 45.0 + 36.0 + 12.0 // 45.0 is length of arrow, and 36.0 is radian on circle arrow
	yText := y - 18.0
	pdf.SetX(xText)
	//texts, _ := pdf.SplitText(d.Text, 150)
	for i, text := range texts {
		pdf.SetX(xText)
		pdf.SetY(yText + float64(i+1)*10.0)
		_ = pdf.Cell(nil, text)
	}

}

func writeKeyBlock(pdf *gopdf.GoPdf, b Barrier) {
	// draw outer rectangle
	rectUpperLeftX := b.BottomX - 250.0
	rectUpperLeftY := b.BottomY - 65.0
	rectWidth := 180.0
	rectHeight := 90.0

	pdf.SetLineWidth(1.0)
	pdf.SetStrokeColor(black.r, black.g, black.b)
	pdf.RectFromUpperLeft(rectUpperLeftX, rectUpperLeftY, rectWidth, rectHeight)

	// draw KEY text
	keyTextX := rectUpperLeftX + 17.0
	keyTextY := rectUpperLeftY + 6.0
	pdf.SetX(keyTextX)
	pdf.SetY(keyTextY)
	_ = pdf.Cell(nil, "KEY")

	// draw blue circle
	blueCircleX := keyTextX
	blueCircleY := keyTextY + 24.0
	drawCircle(pdf, blueCircleX, blueCircleY, 12.0, blue)

	// draw pink circle
	redCircleX := blueCircleX
	redCircleY := blueCircleY + 18.0
	drawCircle(pdf, redCircleX, redCircleY, 12.0, pink)

	// draw left and right curved rectangle
	leftRect := Rectangle{fillColor: pink,
		strokeColor: pink,
		tlPoint:     gopdf.Point{X: redCircleX, Y: redCircleY + 18.0},
		trPoint:     gopdf.Point{X: redCircleX + 6.0, Y: redCircleY + 18.0},
		brPoint:     gopdf.Point{X: redCircleX + 6.0, Y: redCircleY + 27.0},
		blPoint:     gopdf.Point{X: redCircleX, Y: redCircleY + 27.0}}
	leftCurvedRectangle := CurvedRectangle{rectangle: leftRect, curveLength: 5.0}
	leftCurvedRectangle.DrawLeftCurved(pdf)

	rightRect := Rectangle{fillColor: blue,
		strokeColor: blue,
		tlPoint:     gopdf.Point{X: redCircleX + 12.0, Y: redCircleY + 18.0},
		trPoint:     gopdf.Point{X: redCircleX + 6.0, Y: redCircleY + 18.0},
		brPoint:     gopdf.Point{X: redCircleX + 6.0, Y: redCircleY + 27.0},
		blPoint:     gopdf.Point{X: redCircleX + 12.0, Y: redCircleY + 27.0}}

	rightCurvedRectangle := CurvedRectangle{rectangle: rightRect, curveLength: 5.0}
	rightCurvedRectangle.DrawRightCurved(pdf)

	// write BARRISTER ACTIVITY text
	pdf.SetX(blueCircleX + 24.0)
	pdf.SetY(blueCircleY - 6.0)
	_ = pdf.Cell(nil, "BARRISTER ACTIVITY")

	// write CLIENT/SOLICITOR ACTIVITY text
	pdf.SetX(redCircleX + 24.0)
	pdf.SetY(redCircleY - 6.0)
	_ = pdf.Cell(nil, "CLIENT/SOLICITOR ACTIVITY")

	// write JOINT COLLABORATION ACTIVITY text
	pdf.SetX(redCircleX + 24.0)
	pdf.SetY(redCircleY + 12.0)
	_ = pdf.Cell(nil, "JOINT COLLABORATION")

	pdf.SetX(redCircleX + 24.0 + 32.0)
	pdf.SetY(redCircleY + 24.0)
	_ = pdf.Cell(nil, "ACTIVITY")

}
