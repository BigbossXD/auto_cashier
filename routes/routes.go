package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, apiV2Prefix string) {

	g := e.Group(apiV2Prefix)
	g.GET("/maximum", controllers.GetMaximum)
	g.PUT("/deposit", controllers.Deposit)
	g.PUT("/withdraw", controllers.Withdraw)
	g.POST("/receive", controllers.Receive)

}
