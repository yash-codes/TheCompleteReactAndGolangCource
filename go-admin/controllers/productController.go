package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  "strconv"
  //"golang.org/x/crypto/bcrypt"
)

func AllProducts(c *fiber.Ctx) error {

  page, _ := strconv.Atoi(c.Query("page","1"))

 // limit := 5
 // offset := (page - 1) * limit
 // var total int64

 // var products []models.Product

 // database.DB.Offset(offset).Limit(limit).Find(&products)
 // database.DB.Model(&models.Product{}).Count(&total)
 // //return c.JSON(products)
 // return c.JSON(fiber.Map{
 //   "date": products,
 //   "meta": fiber.Map{
 //     "total": total,
 //     "page":page,
 //     "last_page": math.Ceil(float64(int(total)/limit)),
 //   },
 // })


 //return c.JSON(models.Paginate(database.DB, page))
 // after add interface
 return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
  var product models.Product

  if err := c.BodyParser(&product); err != nil {
    return err
  }

  // Instead of using below commented statement we can use
  //password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
  //product.SetPassword("1234")

  //product.Password = password
  database.DB.Create(&product)

  return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  product := models.Product{
    Id: uint(id),
  }

  database.DB.Find(&product)

  return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {

  id, _ := strconv.Atoi(c.Params("id"))

  product := models.Product{
    Id: uint(id),
  }

  if err := c.BodyParser(&product); err != nil {
    return err
  }

  database.DB.Model(&product).Updates(product)

  return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  product := models.Product{
    Id: uint(id),
  }

  database.DB.Delete(&product)

  return nil
}
