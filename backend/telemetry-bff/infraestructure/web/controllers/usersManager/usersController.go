package usersManager

import (
	"net/http"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/gin-gonic/gin"
)

func CheckUserController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Request.Header.Get("sub")
		userData := app.UsersManagerUsecase.GetUserData(ctx, userId)

		ctx.JSON(http.StatusOK, userData)
	}
}
