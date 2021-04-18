package main

import (
	"log"

	"github.com/cjd997/Rightful-tech-Tools/config"
	"github.com/cjd997/Rightful-tech-Tools/server"
)

const (
	configFile = "files/config.toml"
)

func mainCore() error {
	cfg, err := config.Load(configFile)
	if err != nil {
		return err
	}

	return server.NewServer(cfg).Serve()
}

func main() {
	if err := mainCore(); err != nil {
		log.Fatal(err)
	}
}
