package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  "strconv"
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

func GetUser(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  database.DB.Find(&user)

  return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  if err := c.BodyParser(&user); err != nil {
    return err
  }

  database.DB.Model(&user).Updates(user)

  return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  database.DB.Delete(&user)

  return nil
}
