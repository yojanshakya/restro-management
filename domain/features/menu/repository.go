package menu

import (
	"Restro/models"
	"Restro/pkg/infrastructure"
	"fmt"
)

type MenuRepository struct{
	db infrastructure.Database
}

func NewMenuRepository(db infrastructure.Database) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) GetAllMenu(menuItems *[]models.MenuItem) (*[]models.MenuItem, error) {
	result := r.db.Preload("Category").Find(&menuItems);

	if result.Error != nil{
		fmt.Printf("%+v\n", result.Error)
	}
	

	return menuItems, result.Error
}

func (r *MenuRepository) CreateMenu(menuItem *models.MenuItem) (error){
	result := r.db.Table("menu_item").Create(&menuItem)

	fmt.Printf("%+v\n", result)
	return result.Error
}