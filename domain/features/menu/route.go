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
	r.RegisterRoutes()
	r.groupRouter.GET("/all", r.menuController.GetAllMenus)
}