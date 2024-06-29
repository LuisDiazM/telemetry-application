package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "online",
	})
}
