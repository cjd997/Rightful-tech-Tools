package chart

import (
	"fmt"
	"time"
	"math/rand"
	"path/filepath"

	"github.com/signintech/gopdf"
)

type Request struct {
	Text      string  `json:"text"`
	Direction string  `json:"direction"`
	Cost      float64 `json:"cost"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomString() string {
    b := make([]rune, 10)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func Generate(req []Request, filesFolder string) (string, error) {
	// generate filename
	filename := randomString() + ".pdf"
	filePath := filepath.Join(filesFolder, filename)
	return generate(filePath)
}

func generate(filePath string) (string, error) {
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
		return "", fmt.Errorf("error adding font: %s", err.Error())
	}

	err = pdf.SetFont("Poppins", "", 10)
	if err != nil {
		return "", fmt.Errorf("error setting font: %s", err.Error())
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

	err = pdf.WritePdf(filePath)
	if err != nil {
		return "", fmt.Errorf("error writing chart to pdf file: %s", err.Error())
	}

	return filePath, nil
}
