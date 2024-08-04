package routes

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers/usersManager"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(group *gin.RouterGroup, app *app.Application) {
	group.GET("/user/me", usersManager.CheckUserController(app))
	group.POST("/user", usersManager.SaveUserController(app.UsersManagerUsecase))
}
