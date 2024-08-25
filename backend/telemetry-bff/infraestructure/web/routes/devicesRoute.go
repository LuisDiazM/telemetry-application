package routes

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers/devicesmanager"
	"github.com/gin-gonic/gin"
)

func RegisterDevicesRouter(group *gin.RouterGroup, app *app.Application) {
	group.GET("/device/:id", devicesmanager.GetDeviceById(app))
	group.GET("/devices", devicesmanager.GetDevicesByLocation(app))
}
