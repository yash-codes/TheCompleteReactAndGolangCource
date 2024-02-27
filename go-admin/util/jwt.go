// We may have multiple routes/Api's, and while reaching to those api's we need to validate the user
// So it does not make sence to write authorization logic again and again in eacch of the Api/Handler functions
// For that perpose we would add middleware, it will help us to secure our API's as well.
// Middleware is a function that executes before reaching up to the Handler functions directly and reject if the user is unauthorize

package util

import (

  "github.com/dgrijalva/jwt-go"
  "time"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error){
  // now we got the user, so we want to store some info for this user, so have to create some claims
  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
    Issuer: issuer,
    ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // this will expires in 1 day
  })

  return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error){
  token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte("secret"), nil
  })

  if err != nil || !token.Valid {
    return "", err
  }

  claims := token.Claims.(*jwt.StandardClaims)

  return claims.Issuer, nil
}
