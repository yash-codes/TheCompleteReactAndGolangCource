package main

import (
  "go-admin/database"
  "go-admin/routes"
  "github.com/gofiber/fiber/v2"
)


func main() {

    database.Connect()

    app := fiber.New()
    routes.Setup(app)

    //app.Get("/", Fun)
    app.Listen(":8000")
}
