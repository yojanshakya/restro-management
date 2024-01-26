package menu

import (
	"Restro/models"
	"Restro/pkg/utils"
)

type MenuService struct {
	menuRepo *MenuRepository
}

func NewMenuService(menuRepo *MenuRepository) *MenuService{
	return &MenuService{
		menuRepo: menuRepo,
	}
}	

func (s *MenuService) GetAllMenu(menuItems *[]models.MenuItem, count *int64, pagination *utils.Pagination) (error) {
	return s.menuRepo.GetAllMenu(menuItems, count, pagination)
}

func (s *MenuService) CreateMenu(menuItem *models.MenuItem) (error){
	return s.menuRepo.CreateMenu(menuItem)
}

func (s *MenuService) GetMenuById(id int, menuItem *models.MenuItem) (error){
	return s.menuRepo.GetMenuById(id, menuItem)
}