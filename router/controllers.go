package router

import (
	"fmt"
	"lid/interfaces"

	"github.com/labstack/echo/v4"
)

type controller struct {
	drs interfaces.DeviceRepository
	frs interfaces.FileRepository
}

var c controller

func newController(drs interfaces.DeviceRepository, frs interfaces.FileRepository) controller {
	if (c == controller{}) {
		c = controller{drs: drs, frs: frs}
	}

	return c
}

func (c controller) GetAllDevices(ctx echo.Context) error {
	result, _ := c.drs.FindAllPreload()
	fmt.Printf("%#v", result)
	return ctx.JSON(200, result)
}

func (c controller) GetAllFiles(ctx echo.Context) error {
	result, _ := c.frs.FindAllPreload()
	return ctx.JSON(200, result)
}
