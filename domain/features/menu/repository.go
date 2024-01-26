package menu

import (
	"Restro/models"
	"Restro/pkg/framework"
	"Restro/pkg/infrastructure"
	"Restro/pkg/utils"
	"fmt"
)

type MenuRepository struct{
	infrastructure.Database
	framework.Logger
}

func NewMenuRepository(db infrastructure.Database, logger framework.Logger) *MenuRepository {
	return &MenuRepository{db, logger}
}

func (r *MenuRepository) GetAllMenu(menuItems *[]models.MenuItem, count *int64, pagination *utils.Pagination) (error) {
	result := r.Preload("Category").Offset(pagination.Offset).Limit(pagination.Limit).Find(&menuItems)
	r.Model(&models.MenuItem{}).Count(count)
	return result.Error
}

func (r *MenuRepository) CreateMenu(menuItem *models.MenuItem) (error){
	result := r.Create(&menuItem)
	fmt.Printf("%+v\n", result)
	return result.Error
}

func (r *MenuRepository) GetMenuById(id int, menuItem *models.MenuItem) (error){
	result := r.Preload("Category").First(&menuItem, id)
	return result.Error
}	

func (r *MenuRepository) DeleteMenuById(id int) (error){
	result := r.Delete(&models.MenuItem{}, id)
	return result.Error
}

