package database

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect() {
  dsn := "root:rootpass@/go_admin"
  db, err := gorm.Open(mysql.Open(dsn))
  if err != nil {
    panic("Couln not connect to the Database, Terminating")
  }
  fmt.Println(db)

}

