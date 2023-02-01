package main

import (
	"os"
	"time"

	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/routes"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {

	time.LoadLocation("Asia/Bangkok")
	utils.InitializeLogger()

	godotenv.Load(".env")
	appPort := os.Getenv("APP_PORT")

	orm.InitDB()

	e := echo.New()

	// /* CORS */
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// e.Use(middlewares.CheckXApiKey)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			utils.Logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	apiV1Prefix := "api/v1"

	routes.InitConfigsRoutes(e, apiV1Prefix)
	routes.InitRoutes(e, apiV1Prefix)
	routes.InitMachineRoutes(e, apiV1Prefix)

	e.Logger.Fatal(e.Start(":" + appPort))
}
