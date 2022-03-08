package router

import (
	"lid/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TODO: make config model
func Start(drs interfaces.DeviceRepository, nrs interfaces.NodeRepository) {
	router := echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	controller := newController(drs, nrs)

	router.StaticFS("/", getFrontendAssets())
	v1API := router.Group("/api/v1")
	{
		v1API.GET("/devices", controller.GetAllDevices)
		v1API.GET("/files", controller.GetAllFiles)
		// v1deviceAPI := v1API.Group("/device")
		// {
		// 	v1deviceAPI.GET("/name/:name")
		// 	v1deviceAPI.GET("/uuid/:uuid")
		// }
	}

	router.Logger.Info(router.Start(":8080"))
}
