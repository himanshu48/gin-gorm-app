package routes

import (
	"gin-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup)  {
    //All routes related to users comes here
	router.GET("/", controllers.GetAllUsers) 
	// router.POST("/", controllers.CreateUser)
	router.GET("/:userId", controllers.GetAUser)
    // router.PUT("/:userId", controllers.EditAUser)
    router.DELETE("/:userId", controllers.DeleteAUser)
}