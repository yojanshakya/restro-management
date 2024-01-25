package menu

import (
	"Restro/models"
	"fmt"
)


type MenuService struct {
	// db infrastructure.Database
}

func NewMenuService() *MenuService{
	return &MenuService{}
}	

func (s *MenuService) GetAllMenu() {
	menuItem := []models.MenuItem{}

	// s.db.Find(&menuItem)

	fmt.Println(menuItem)
}