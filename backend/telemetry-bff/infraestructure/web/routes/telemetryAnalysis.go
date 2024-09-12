package routes

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers/telemetryanalysis"
	"github.com/gin-gonic/gin"
)

func RegisterTelemtryRouter(group *gin.RouterGroup, app *app.Application) {
	group.GET("/measures/:device_id", telemetryanalysis.GetLastWeekMeasuresByDevice(app))
}
