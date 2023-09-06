package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthRoute(router *gin.RouterGroup)  {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
}