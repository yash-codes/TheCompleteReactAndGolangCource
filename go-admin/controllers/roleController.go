package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  "strconv"
  //"golang.org/x/crypto/bcrypt"
)

func AllRoles(c *fiber.Ctx) error {
  var roles []models.Role

  database.DB.Find(&roles)

  return c.JSON(roles)

}

func CreateRole(c *fiber.Ctx) error {
  //var role models.Role

  var roleDto fiber.Map

  if err := c.BodyParser(&roleDto); err != nil {
    return err
  }

  list := roleDto["permissions"].([]interface{})

  permissions := make([]models.Permission, len(list))

  for i, permissionId := range list {
    id, _ := strconv.Atoi(permissionId.(string))

    permissions[i] = models.Permission{
	Id: uint(id),
    }
  }

  role := models.Role{
	Name:        roleDto["name"].(string),
	Permissions: permissions,
  }
  //user.Password = password
  database.DB.Create(&role)

  return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  role := models.Role{
    Id: uint(id),
  }

  database.DB.Find(&role)

  return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {

  id, _ := strconv.Atoi(c.Params("id"))

  var roleDto fiber.Map
  //role := models.Role{
    //Id: uint(id),
  //}

  if err := c.BodyParser(&roleDto); err != nil {
    return err
  }

  list := roleDto["permissions"].([]interface{})

  permissions := make([]models.Permission, len(list))

  for i, permissionId := range list {
    id, _ := strconv.Atoi(permissionId.(string))

    permissions[i] = models.Permission{
	Id: uint(id),
    }
  }

  var result interface{}

  database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

  role := models.Role{
	Id: uint(id),
	Name:        roleDto["name"].(string),
	Permissions: permissions,
  }

  database.DB.Model(&role).Updates(role)

  return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
  id, _ := strconv.Atoi(c.Params("id"))

  role := models.Role{
    Id: uint(id),
  }

  database.DB.Delete(&role)

  return nil
}
