package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/labstack/echo/v4"
)

func InitMachineRoutes(e *echo.Echo, apiV1Prefix string) {

	g := e.Group(apiV1Prefix + "/machine")
	g.GET("", controllers.FindMachine)
	g.GET("/full", controllers.FindMachineFull)
	g.GET("/empty", controllers.FindMachineEmpty)
	g.POST("", controllers.CreateMachine)
	g.PUT("", controllers.UpdateMachine)
	g.DELETE("/:id", controllers.DeleteMachine)

}
