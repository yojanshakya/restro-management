package models

// todo json validation
type MenuItem struct {
	ID int `gorm:"primaryKey" json:"id"`
	Description string `gorm:"column:description" json:"description"`
	Quantity int `gorm:"column:quantity" json:"quantity"`
	Price int `gorm:"column:price" json:"price"`
	Name string `gorm:"column:name" json:"name"`
	CategoryID  int          `json:"categoryId"`
    Category    MenuCategory `gorm:"foreignKey:ID;references:CategoryID"` 
}

func (MenuItem) TableName() string {
    return "menu_item"
}