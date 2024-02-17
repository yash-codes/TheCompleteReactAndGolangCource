package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)


func main() {
  fmt.Println("Hello World !")

  app := fiber.New()
  // An Api for url path "/"
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World !")
  })

  app.Listen(":8000")
}
