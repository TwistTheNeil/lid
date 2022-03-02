package router

import (
	"lid/interfaces"

	"github.com/gin-gonic/gin"
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

func (c controller) FindAllDevices(ctx *gin.Context) {
	result, _ := c.drs.FindAll()
	ctx.JSON(200, gin.H{
		"result": result,
	})
}
