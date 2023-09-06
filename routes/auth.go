package routes

import (
	"gin-gorm/controllers"
	"gin-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.RouterGroup)  {
    //All routes related to users comes here
	router.POST("/login", controllers.LogInUser)
	router.POST("/signup", controllers.SignupUser)
	router.GET("/logout", middleware.JwtMiddleware(), controllers.LogOutUser)
}