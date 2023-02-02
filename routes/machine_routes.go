package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/repositories"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/labstack/echo/v4"
)

func InitMachineRoutes(e *echo.Echo, apiV1Prefix string) {

	machineRepo := repositories.NewMachineRepo(orm.Db)
	machineService := services.NewMachineService(machineRepo)
	machineController := controllers.NewMachineController(machineService)

	g := e.Group(apiV1Prefix + "/machine")
	g.GET("", machineController.FindMachineList)
	g.GET("/full", machineController.FindMachineFull)
	g.GET("/empty", machineController.FindMachineEmpty)
	g.POST("", machineController.CreateMachine)
	g.PUT("", machineController.UpdateMachine)
	g.DELETE("/:id", machineController.DeleteMachine)

}
