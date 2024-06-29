package web

import (
	"log"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	p := ginprom.New(
		ginprom.Engine(server),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	server.Use(p.Instrument())
	log.Println("start web server")
	return server
}
