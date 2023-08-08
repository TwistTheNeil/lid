package router

import (
	"context"
	"fmt"
	"lid/interfaces"
	"lid/services/logger"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type routerContext struct {
	router *echo.Echo
}

var r routerContext

// TODO: make config model
func Start(drs interfaces.DeviceRepository, frs interfaces.FileRepository) {
	log := logger.CreateLogger("router")
	log.Trace("router.Start called")

	router := echo.New()

	r = routerContext{router: router}

	r.router.Use(middleware.CORS())
	r.router.Use(middleware.Logger())
	controller := newController(drs, frs)

	r.router.StaticFS("/", getFrontendAssets())
	v1API := r.router.Group("/api/v1")
	{
		v1API.GET("/devices", controller.GetAllDevices)
		v1API.GET("/files", controller.GetAllFiles)
		// v1deviceAPI := v1API.Group("/device")
		// {
		// 	v1deviceAPI.GET("/name/:name")
		// 	v1deviceAPI.GET("/uuid/:uuid")
		// }
	}

	if err := r.router.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("something went wrong starting echo", err)
	}

	log.Info("echo has been gracefully shutdown")
}

func Stop() {
	log := logger.CreateLogger("router")
	log.Trace("router.Stop called")
	fmt.Println("what")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.router.Shutdown(ctx); err != nil {
		log.Fatal("something went wrong stopping echo", err)
	}
}
