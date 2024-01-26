package menu

import (
	"Restro/models"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service *MenuService
}

func NewMenuController(service *MenuService) *MenuController{
	return &MenuController{
		service: service,
	}
}

func (menuController *MenuController) GetAllMenus(ctx *gin.Context){
	var menuItems []models.MenuItem
	menuController.service.GetAllMenu(&menuItems)

	ctx.JSON(200, gin.H{
		"data": menuItems,
	})
}