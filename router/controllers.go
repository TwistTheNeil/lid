package router

import (
	"lid/interfaces"

	"github.com/labstack/echo/v4"
)

type controller struct {
	drs interfaces.DeviceRepository
	nrs interfaces.NodeRepository
}

var c controller

func newController(drs interfaces.DeviceRepository, nrs interfaces.NodeRepository) controller {
	if (c == controller{}) {
		c = controller{drs: drs, nrs: nrs}
	}

	return c
}

func (c controller) GetAllDevices(ctx echo.Context) error {
	result, _ := c.drs.FindAll()
	return ctx.JSON(200, result)
}

func (c controller) GetAllFiles(ctx echo.Context) error {
	result, _ := c.nrs.FindAll()
	return ctx.JSON(200, result)
}
