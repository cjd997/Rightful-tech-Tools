package main

import (
	"log"
	
	"github.com/cjd997/Rightful-tech-Tools/server"
)

const (
	port = ":3000"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.Serve(port))
}
