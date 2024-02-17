package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


func main() {
  fmt.Println("Hello World !")

  /*
  // This is to learn you about pointers
  // #1 Wrong initializing value
	  // decleare pointer variable
	  var name *string
	  // Initialization
	  name := "yash"
	  fmt.Println(name)
  */
  /*
  // #2 Correct initailization but Wrong declaration in case of pointer
  	// declaration pointer variable
	var name *string
	// Initialization
	*name := "yash"
	fmt.Println(name)
  */
  // #3 The Correct way
  	// declaration pointer variable
	var name *string = new(string)
	//Initialization
	*name = "yash"
	fmt.Printf("address: %v, value: %v\n", name, *name)

	// We can notice that if we change the value but the address doesn't change
	*name = "Raj"
	fmt.Printf("address: %v, value: %v\n", name, *name)

  dsn := "root:rootpass@/go_admin"
  db, err := gorm.Open(mysql.Open(dsn))
  if err != nil {
    panic("Couln not connect to the Database, Terminating")
  }
  fmt.Println(db)


  app := fiber.New()
  // An Api for url path "/"
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World !")
  })

  app.Listen(":8000")
}
