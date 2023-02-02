package routes

import (
	"github.com/BigbossXD/auto_cashier/controllers"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/repositories"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/labstack/echo/v4"
)

func InitTransectionRoutes(e *echo.Echo, apiV1Prefix string) {

	transectionsRepo := repositories.NewTransectionsRepo(orm.Db)
	transectionsService := services.NewTransectionsService(transectionsRepo)
	transectionsController := controllers.NewTransectionsController(transectionsService)

	g := e.Group(apiV1Prefix)
	g.GET("/transections", transectionsController.GetTransection)

}
