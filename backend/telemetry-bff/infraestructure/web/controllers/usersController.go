package controllers

import (
	"net/http"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/gin-gonic/gin"
)

func CheckUserController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "online",
			"sub":    ctx.Request.Header.Get("sub"),
		})
	}
}
