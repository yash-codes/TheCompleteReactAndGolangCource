package models

import "gorm.io/gorm"

type Order struct {
  // can be use and replace updatedAt createdAt, id
  //gorm.Model
  Id		uint	`json:"id"`
  FirstName	string	`json:"-"`                // by mentioning json:"-" this will not be added in the reponse
  LastName	string	`json:"-"`
  Name		string	`json:"name" gorm:"-"`    // by mentioning gorm:"-" this coloumn be not be created in database
  Total		float32	`json:"total" gorm:"-"`    // by mentioning gorm:"-" this coloumn be not be created in database
  Email		string	`json:"email"`
  UpdatedAt	string	`json:"updated_at"`
  CreatedAt	string	`json:"created_at"`
  OrderItems	[]OrderItem  `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id		uint	`json:"id"`
	OrderId		uint	`json:"order_id"`
	ProductTitle	string	`json:"product_title"`
	Price		float32	`json:"price"`
	Quantity	uint	`json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {

  var total int64
  db.Model(&Order{}).Count(&total)
  return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {

  var orders []Order
  db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

  for i, _ := range orders {
    orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
    var total float32 = 0

    for j, _ := range orders[i].OrderItems {
      total += orders[i].OrderItems[j].Price * float32(orders[i].OrderItems[j].Quantity)
    }
    orders[i].Total = total
  }

  return orders
}
