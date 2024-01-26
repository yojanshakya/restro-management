package models

type MenuCategory struct {
	ID int `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:category_name" json:"name"`
}

func (MenuCategory) TableName() string {
    return "menu_category"
}