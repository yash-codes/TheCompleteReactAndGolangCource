package controllers

import (
  "github.com/gofiber/fiber/v2"
  "go-admin/models"
  "go-admin/database"
)

func AllPermissions(c *fiber.Ctx) error {
  var permissions []models.Permission

  database.DB.Find(&permissions)

  return c.JSON(permissions)
}

func StorePerssions(){

  var permissions = [...]models.Permission{
    {Name:"view_users"},
    {Name:"edit_users"},
    {Name:"view_roles"},
    {Name:"edit_roles"},
    {Name:"view_product"},
    {Name:"edit_product"},
    {Name:"view_orders"},
    {Name:"edit_orders"},
}
    for i, _ := range permissions {
      database.StoreData(&permissions[i])
    }

}
