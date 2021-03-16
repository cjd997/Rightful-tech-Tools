package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	*fiber.App
}

func NewServer() *Server {
	server := &Server{
		App: fiber.New(),
	}
	server.addRoutes()
	return server
}

func (s *Server) addRoutes() {
	s.Post("/generate", s.generateHandler)
}

func (s *Server) Serve(port string) error {
	return s.Listen(port)
}