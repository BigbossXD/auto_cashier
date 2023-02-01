package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/labstack/echo/v4"
)

func InitConfigsRoutes(e *echo.Echo, apiV1Prefix string) {

	g := e.Group(apiV1Prefix + "/configs")
	g.GET("", controllers.FindConfig)
	g.POST("", controllers.CreateConfig)
	g.PUT("", controllers.UpdateConfig)
	g.DELETE("/:id", controllers.DeleteConfig)

}
