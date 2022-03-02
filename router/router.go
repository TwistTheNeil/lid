package router

import (
	"lid/interfaces"

	"github.com/gin-gonic/gin"
)

// TODO: make config model
func Start(drs interfaces.DeviceRepository) {
	router := gin.Default()

	controller := newController(drs)

	v1API := router.Group("/api/v1")
	{
		v1API.GET("/devices", controller.FindAllDevices)
		// v1deviceAPI := v1API.Group("/device")
		// {
		// 	v1deviceAPI.GET("/name/:name")
		// 	v1deviceAPI.GET("/uuid/:uuid")
		// }
	}

	router.Run(":8080")
}
