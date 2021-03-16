package server 

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/cjd997/Rightful-tech-Tools/chart"
)

func (s *Server) generateHandler(ctx *fiber.Ctx) error {
	var req []chart.Request
	err := ctx.BodyParser(&req)
	if err != nil {
		return fmt.Errorf("error parsing request data: %s", err.Error())
	}

	// TODO remove later
	fmt.Println(req)

	err = chart.Generate(req)
	if err != nil {
		return err
	}

	return nil
}