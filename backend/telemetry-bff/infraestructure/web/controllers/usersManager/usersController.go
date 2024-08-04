package usersManager

import (
	"net/http"

	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers/usersManager/requests"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers/usersManager/utils"
	"github.com/gin-gonic/gin"
)

func CheckUserController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Request.Header.Get("sub")
		userData := app.UsersManagerUsecase.GetUserData(ctx, userId)

		ctx.JSON(http.StatusOK, userData)
	}
}

func SaveUserController(usecase *usersManager.UsersManagerUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequest requests.UserRequest
		err := ctx.BindJSON(&userRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		// //check permisions
		// userId := ctx.Request.Header.Get("sub")
		// userProfile := usecase.GetUserData(ctx.Request.Context(), userId)
		// if userProfile != nil {
		// 	if userProfile.Role == ROLE_ANALYST {
		// 		ctx.JSON(http.StatusForbidden, nil)
		// 		return
		// 	}
		// }

		userData, err := utils.DecodeIDToken(userRequest.UserData, "ANALYST", "TRIAL")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		userResponse, err := usecase.SaveUserData(ctx.Request.Context(), *userData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, userResponse)

	}
}
