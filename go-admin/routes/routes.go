package routes

import(
  "go-admin/controllers"
  "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  app.Get("/", controllers.Hello)
  app.Get("/other", controllers.Other)
}
