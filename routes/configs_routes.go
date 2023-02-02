package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/repositories"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/labstack/echo/v4"
)

func InitConfigsRoutes(e *echo.Echo, apiV1Prefix string) {

	configsRepo := repositories.NewConfigsRepo(orm.Db)
	configsService := services.NewConfigsService(configsRepo)
	configsController := controllers.NewConfigsController(configsService)

	g := e.Group(apiV1Prefix + "/configs")
	g.GET("", configsController.FindConfig)
	g.POST("", configsController.CreateConfig)
	g.PUT("", configsController.UpdateConfig)
	g.DELETE("/:id", configsController.DeleteConfig)

}
