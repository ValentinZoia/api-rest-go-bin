package router

import (
	"github.com/ValentinZoia/gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.Initial)
	r.GET("/ping", handler.Ping)
	return r
}
