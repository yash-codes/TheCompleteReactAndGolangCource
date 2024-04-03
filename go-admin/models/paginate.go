package models

import (
  "github.com/gofiber/fiber/v2"
  "gorm.io/gorm"
  "math"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
  limit := 5
  offset := (page - 1) * limit
  var total int64

  // If we specificaly use product moled this func will be only used for products not users
  //var products []Product

  data := entity.Take(db, limit, offset)
  total = entity.Count(db)
  //db.Offset(offset).Limit(limit).Find(&products)
  //db.Model(&Product{}).Count(&total)
  //return c.JSON(products)
  return fiber.Map{
    //"date": products,
    "date": data,
    "meta": fiber.Map{
      "total": total,
      "page":page,
      "last_page": math.Ceil(float64(int(total)/limit)),
    },
  }

}
