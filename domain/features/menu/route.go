package menu

import (
	"Restro/pkg/infrastructure"

	"github.com/gin-gonic/gin"
)

type MenuRoute struct {
	router *infrastructure.Router
	groupRouter *gin.RouterGroup
	menuController *MenuController
}

func NewRoute(router *infrastructure.Router, menuController *MenuController) *MenuRoute{
	route := MenuRoute{router: router, menuController: menuController}
	route.RegisterRoutes()
	return &route
}

func (r *MenuRoute) RegisterRoutes() {
	r.groupRouter = r.router.Group("/menu")
	r.groupRouter.GET("/all", r.menuController.GetAllMenus)
	r.groupRouter.POST("/create", r.menuController.CreateMenu)
	r.groupRouter.GET("/:id", r.menuController.GetMenuById)
	r.groupRouter.DELETE("/:id", r.menuController.DeleteMenuById)
}