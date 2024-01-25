package menu

import "github.com/gin-gonic/gin"

type MenuController struct {
	service *MenuService
}

func NewMenuController(service *MenuService) *MenuController{
	return &MenuController{
		service: service,
	}
}

func (menuController *MenuController) GetAllMenus(ctx *gin.Context){
	menuController.service.GetAllMenu()
	ctx.JSON(200, gin.H{
		"message": "Hello World peace out",
	})
}