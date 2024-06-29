package routes

import (
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRouter(group *gin.RouterGroup) {
	group.GET("/health", controllers.HealthController)
}
