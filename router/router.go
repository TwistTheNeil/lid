package router

import (
	"lid/interfaces"

	"github.com/labstack/echo/v4"
)

// TODO: make config model
func Start(drs interfaces.DeviceRepository) {
	router := echo.New()

	controller := newController(drs)

	router.StaticFS("/", getFrontendAssets())
	v1API := router.Group("/api/v1")
	{
		v1API.GET("/devices", controller.FindAllDevices)
		// v1deviceAPI := v1API.Group("/device")
		// {
		// 	v1deviceAPI.GET("/name/:name")
		// 	v1deviceAPI.GET("/uuid/:uuid")
		// }
	}

	router.Logger.Info(router.Start(":8080"))
}