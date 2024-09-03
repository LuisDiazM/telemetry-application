package devicesmanager

import (
	"net/http"
	"strconv"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/entities"
	"github.com/gin-gonic/gin"
)

func GetDeviceById(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		deviceData := app.DevicesManagerUsecase.GetDeviceById(ctx.Request.Context(), id)
		if deviceData == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		ctx.JSON(http.StatusOK, deviceData)
	}
}

func GetDevicesByLocation(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		latitude, latOk := ctx.GetQuery("latitude")
		longitude, longOk := ctx.GetQuery("longitude")

		if !latOk || !longOk {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "no latitude or longitude",
			})
			return
		}
		lat, errLat := strconv.ParseFloat(latitude, 64)
		long, errLong := strconv.ParseFloat(longitude, 64)

		if errLat != nil || errLong != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "lat or long not float",
			})
			return
		}
		location := entities.Location{Type: entities.Point, Coordinates: []float64{lat, long}}
		devices := app.DevicesManagerUsecase.GetDevicesByLocation(ctx.Request.Context(), location)
		if devices == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		ctx.JSON(http.StatusOK, devices)
	}
}

func GetStatusByLocation(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		latitude, latOk := ctx.GetQuery("latitude")
		longitude, longOk := ctx.GetQuery("longitude")

		if !latOk || !longOk {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "no latitude or longitude",
			})
			return
		}
		lat, errLat := strconv.ParseFloat(latitude, 64)
		long, errLong := strconv.ParseFloat(longitude, 64)

		if errLat != nil || errLong != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "lat or long not float",
			})
			return
		}
		location := entities.Location{Type: entities.Point, Coordinates: []float64{lat, long}}
		devices := app.DevicesManagerUsecase.GetStatusDevicesByLocation(ctx.Request.Context(), location)
		if devices == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		ctx.JSON(http.StatusOK, devices)
	}
}
