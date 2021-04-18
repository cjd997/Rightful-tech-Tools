package server

import (
	"github.com/cjd997/Rightful-tech-Tools/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	*fiber.App
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := &Server{
		App:    fiber.New(),
		config: cfg,
	}
	server.addRoutes()
	return server
}

func (s *Server) addRoutes() {
	s.Post("/generate", s.handler)
}

func (s *Server) Serve() error {
	return s.Listen(s.config.Server.Port)
}
