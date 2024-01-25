package menu

import "github.com/gin-gonic/gin"

type MenuController struct {

}

func NewMenuController() *MenuController{
	return &MenuController{}
}

func (menuController *MenuController) GetAllMenus(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "Hello World peace out",
	})
}