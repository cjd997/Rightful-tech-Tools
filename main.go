package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	// write json data to Datas struct
	//fmt.Fprintf(w, "Datas: %+v", d)

	pdfFileName := createAppealsPdf(&d)

	// read pdf generated file and serve to client as response
	serveAppealsPdfFile(pdfFileName, w)
}

func serveAppealsPdfFile(pdfFileName string, w http.ResponseWriter) {
	streamPDFbytes, err := ioutil.ReadFile(pdfFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := bytes.NewBuffer(streamPDFbytes)

	w.Header().Set("Content-type", "application/pdf")

	if _, err := b.WriteTo(w); err != nil {
		fmt.Fprintf(w, "%s", err)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/appeals", appealsCreate)
	mux.HandleFunc("/", appealsCreate)

	fmt.Println("web server running on port 4000.")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
