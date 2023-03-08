package main

import (
	"log"

	"github.com/antstalepresh/flightpath/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := server.New(fiber.New())

	if err := s.Start(); err != nil {
		log.Fatalf("failed to start the server: %s", err)
	}
}
