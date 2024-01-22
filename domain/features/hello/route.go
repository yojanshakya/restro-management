package hello

import (
  "github.com/gin-gonic/gin"
  "Restro/pkg/infrastructure"
)

type HelloRoute struct {
    router *infrastructure.Router
    controller *HelloController
    groupRouter *gin.RouterGroup
}

func NewHelloRoute(router *infrastructure.Router, controller *HelloController) *HelloRoute {
	route := HelloRoute{router: router, controller: controller}
  route.groupRouter = route.router.Group("api/hello")
	route.RegisterHelloRoutes()
	return &route
}

func (r *HelloRoute) RegisterHelloRoutes() {
	r.groupRouter.GET("", r.controller.HandleRoot)
}
