package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  //"golang.org/x/crypto/bcrypt"
)

func AllUsers(c *fiber.Ctx) error {
  var users []models.User

  database.DB.Find(&users)

  return c.JSON(users)

}

func CreateUser(c *fiber.Ctx) error {
  var user models.User

  if err := c.BodyParser(&user); err != nil {
    return err
  }

  // Instead of using below commented statement we can use
  //password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
  user.SetPassword("1234")

  //user.Password = password

  database.DB.Create(&user)

  return c.JSON(user)

}
