package web

import (
	"net/http"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/gin-gonic/gin"
)

func EnableCors(a *app.Application) *app.Application {
	a.WebServer.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Add("Access-Control-Max-Age", "3600")

		if ctx.Request.Method == http.MethodOptions {
			ctx.String(http.StatusOK, "")
		}
	})

	return a
}
