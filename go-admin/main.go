package main

import (
	"go-admin/database"
	"go-admin/routes"
	"github.com/gofiber/fiber/v2"
)


func main() {
  fmt.Println("main() function called...")

  database.Connect()

  app := fiber.New()
  routes.Setup(app)

  app.Listen(":8000")
}
