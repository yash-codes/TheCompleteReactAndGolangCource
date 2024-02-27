package middlewares

import (
  "go-admin/util"
  "github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")

  if _, err := util.ParseJwt(cookie); err != nil {
    c.Status(fiber.StatusUnauthorized)
    return c.JSON(fiber.Map{
      "Message": "Unauthenticated by middleware",
    })
  }

  return c.Next()

}
