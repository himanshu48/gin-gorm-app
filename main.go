package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"

	"gin-gorm/configs"
	"gin-gorm/constants"
	"gin-gorm/middleware"
	"gin-gorm/models"
	"gin-gorm/routes"
)

func main() {
	// Load the env variables
    configs.Init()

	// connect to database
	models.ConnectDatabase()

	if configs.AppEnv() == constants.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimiter())
	router.Use(middleware.AccessLogger())
	router.Use(gin.Recovery()) // included in gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// router.Use(middleware.JwtMiddleware())

	// routes
	r := router.Group("api")
	routes.HealthRoute(r.Group("health")) // Public apis
	routes.AuthRoute(r.Group("auth")) // Public apis
	routes.UserRoute(r.Group("user", middleware.JwtMiddleware()))

	host := fmt.Sprintf(":%s", configs.AppPort())

    router.Run(host)
}