package router

import (
	"lid/interfaces"

	"github.com/labstack/echo/v4"
)

type controller struct {
	drs interfaces.DeviceRepository
}

var c controller

func newController(drs interfaces.DeviceRepository) controller {
	if (c == controller{}) {
		c = controller{drs: drs}
	}

	return c
}

func (c controller) FindAllDevices(ctx echo.Context) error {
	result, _ := c.drs.FindAll()
	return ctx.JSON(200, result)
}
