package controllers

import (
  "go-admin/models"
  "github.com/gofiber/fiber/v2"
  "go-admin/database"
  "strconv"
  "os"
  "encoding/csv"
  //"golang.org/x/crypto/bcrypt"
)

func AllOrders(c *fiber.Ctx) error {

  page, _ := strconv.Atoi(c.Query("page","1"))

 return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}

// export orders in CSV file
func Export(c *fiber.Ctx) error {
  filePath := "./csv/orders.csv"

  if err := CreateFile(filePath); err != nil {
    return err
  }

  return c.Download(filePath)
}

func CreateFile(filePath string) error {
  file, err := os.Create(filePath)

  if err != nil {
    return err
  }

  defer file.Close()

  writer := csv.NewWriter(file)
  defer writer.Flush()

  var orders []models.Order

  database.DB.Preload("OrderItems").Find(&orders)

  writer.Write([]string{
    "ID","Name", "Email", "Product Title" , "Price", "Quantity",
  })

  for _, order := range orders {
    data := []string{
      strconv.Itoa(int(order.Id)),
      order.FirstName + " " + order.LastName,
      order.Email,
      "",
      "",
      "",
    }
    if err := writer.Write(data); err != nil {
      return err
    }

    for _, orderItem := range order.OrderItems {
      data := []string{
        "",
        "",
        "",
        orderItem.ProductTitle,
        strconv.Itoa(int(orderItem.Price)),
        strconv.Itoa(int(orderItem.Quantity)),
      }
      if err := writer.Write(data); err != nil {
        return err
      }
    }
  }

  return nil
}

/*
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
*/
