package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  "strconv"
  //"golang.org/x/crypto/bcrypt"
)

func AllUsers(c *fiber.Ctx) error {

  page, _ := strconv.Atoi(c.Query("page","1"))
  // limit := 5
  // offset := (page - 1) * limit
  // var total int64

  // var users []models.User

  // database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
  // database.DB.Model(&models.User{}).Count(&total)
  // //return c.JSON(users)
  // return c.JSON(fiber.Map{
  //   "date": users,
  //   "meta": fiber.Map{
  //     "total": total,
  //     "page":page,
  //     "last_page": math.Ceil(float64(int(total)/limit)),
  //   },
  // })

  return c.JSON(models.Paginate(database.DB, &models.User{}, page))
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

  database.DB.Preload("Role").Find(&user)

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
