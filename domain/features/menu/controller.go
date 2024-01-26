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

func (menuController *MenuController) CreateMenu(ctx *gin.Context){
	var menuItem models.MenuItem
	ctx.BindJSON(&menuItem)
	
	err := menuController.service.CreateMenu(&menuItem)
	
	// todo send saved data back to FE
	if err == nil{
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"status": "success",
			},
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": gin.H{
			"status": "error",
			"message": err.Error(),
		},
	})
}