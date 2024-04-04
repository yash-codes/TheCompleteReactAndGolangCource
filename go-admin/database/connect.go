package database

import(
	"fmt"
	"go-admin/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
  dsn := "root:rootpass@/go_admin"
  db, err := gorm.Open(mysql.Open(dsn))
  if err != nil {
    panic("Couln not connect to the Database, Terminating")
  }
  fmt.Println(db)
  DB = db

  db.Migrator().DropTable(&models.User{})
  db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{}) //this will create the table automatically

}

func StoreData(data interface{}) {
	result := DB.Create(data)
	if result.Error != nil {
          fmt.Println("Error in storing data, ", result.Error)
	}
}
