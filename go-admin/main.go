package main

import (
  "go-admin/database"
  "go-admin/routes"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {

    database.Connect()

    app := fiber.New()
    app.Use(cors.New(cors.Config{
      AllowCredentials:false,
    }))
    routes.Setup(app)

    //app.Get("/", Fun)
    app.Listen(":8000")
}
