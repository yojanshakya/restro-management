package menu

import (
	"Restro/models"
)

type MenuService struct {
	menuRepo *MenuRepository
}

func NewMenuService(menuRepo *MenuRepository) *MenuService{
	return &MenuService{
		menuRepo: menuRepo,
	}
}	

func (s *MenuService) GetAllMenu(menuItems *[]models.MenuItem) ( *[]models.MenuItem) {
	s.menuRepo.GetAllMenu(menuItems)
	return menuItems
}

func (s *MenuService) CreateMenu(menuItem *models.MenuItem) (error){
	return s.menuRepo.CreateMenu(menuItem)
}
