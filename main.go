package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//	"github.com/signintech/gopdf"
)

func appealsCreate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var d Datas

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "Datas: %+v", d)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/appeals", appealsCreate)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// func main() {

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

// 	err = pdf.WritePdf("hello.pdf")
// 	if err != nil {
// 		panic(err)
// 	}
// }
