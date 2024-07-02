package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "pong",
		"message2": "Soy Crack",
	})
}

func Initial(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "Initial",
		"message2": "Mensaje de Bienvenida",
	})
}
