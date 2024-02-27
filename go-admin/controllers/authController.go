package controllers

import (
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
  "time"
  "go-admin/models"
  "go-admin/database"
  "fmt"
)

func Hello(c *fiber.Ctx) error {
  return c.SendString("Hello world this is /")
}

func Register(c *fiber.Ctx) error {

  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  if data["password"] != data["password_confirm"] {
    c.Status(400)
    return c.JSON(fiber.Map{
      "Error Message": "password does not match",
    }) 
  }

  // password will be changed to hash & send as hash in response
  password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

  user := models.User{
    FirstName: data["first_name"],
    LastName: data["last_name"],
    Email: data["email"],
    Password: password,
  }

  // store user details in db
  database.StoreData(&user)

  return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  var user models.User

  database.DB.Where("email = ?", data["email"]).First(&user)

  fmt.Println("user=> ", user)
  if user.Id == 0 {
    c.Status(400)
    return c.JSON(fiber.Map{
	    "Error Message": "User not found",
    })
  }
  if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
    c.Status(400)
    return c.JSON(fiber.Map{
      "Error Message": "Incorrect Password",
    })
  }

  userIdStr := fmt.Sprintf("%v", user.Id)
  // now we got the user, so we want to store some info for this user, so have to create some claims
  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
    Issuer: userIdStr,
    ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // this will expires in 1 day
  })

  token, err := claims.SignedString([]byte("secret"))
  if err != nil {
    c.SendStatus(fiber.StatusInternalServerError)
  }

  cookie := fiber.Cookie{
    Name: "jwt",
    Value: token,
    Expires: time.Now().Add(time.Hour * 24),
    HTTPOnly: true,
  }

  c.Cookie(&cookie)

  //return c.JSON(token)
  return c.JSON(fiber.Map{
    "Message": "successfully generated the token and stored as cookies",
  })
}

type Claims struct {
  jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")

  token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte("secret"), nil
  })

  if err != nil || !token.Valid {
    c.Status(fiber.StatusUnauthorized)
    return c.JSON(fiber.Map{
      "Message": "Unauthenticated",
    })
  }

  //claims := token.Claims
  claims := token.Claims.(*Claims)

  var user models.User
  database.DB.Where("id = ?", claims.Issuer).First(&user)

  //return c.JSON(claims.Issuer)
  return c.JSON(user)
}
