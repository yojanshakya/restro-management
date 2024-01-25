package models

type MenuItem struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string `gorm:"column:description" json:"description"`
	Quantity int `gorm:"column:quantity" json:"quantity"`
	Price int `gorm:"column:price" json:"price"`
	Name int `gorm:"column:name" json:"name"`
	Category int `gorm:"foreignKey:category_id;references:id" json:"categoryId"`
}