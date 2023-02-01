package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, apiV1Prefix string) {

	g := e.Group(apiV1Prefix)
	g.GET("/maximum", controllers.GetMaximum)
	g.PUT("/deposit", controllers.Deposit)
	g.PUT("/withdraw", controllers.Withdraw)
	g.POST("/receive", controllers.Receive)
	g.GET("/transections", controllers.GetTransection)

}
