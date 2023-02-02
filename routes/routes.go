package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/repositories"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, apiV1Prefix string) {

	configsRepo := repositories.NewConfigsRepo(orm.Db)
	configsService := services.NewConfigsService(configsRepo)
	configsController := controllers.NewConfigsController(configsService)

	g := e.Group(apiV1Prefix)
	g.GET("/maximum", configsController.GetMaximum)
	g.PUT("/deposit", configsController.Deposit)
	g.PUT("/withdraw", configsController.Withdraw)
	g.POST("/receive", configsController.Receive)

}
