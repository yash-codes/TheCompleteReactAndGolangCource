package controllers

import (
  "github.com/gofiber/fiber/v2"
  "go-admin/models"
)

func Register(c *fiber.Ctx) error {

  user := models.User{
    FirstName: "first_name",
    LastName: "last_name",
    Email: "email",
  }

  return c.JSON(user)
}

