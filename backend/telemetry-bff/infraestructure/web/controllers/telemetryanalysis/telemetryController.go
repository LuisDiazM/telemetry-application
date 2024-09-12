package telemetryanalysis

import (
	"net/http"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/gin-gonic/gin"
)

func GetLastWeekMeasuresByDevice(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("device_id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		measures := app.TelemetryAnalysisUsecae.GetLastWeekMeasuresByDevice(ctx.Request.Context(), id)
		if measures == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		ctx.JSON(http.StatusOK, measures)
	}
}
