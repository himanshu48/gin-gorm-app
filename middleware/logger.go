package middleware

import (
	"fmt"
	"gin-gorm/configs"
	"gin-gorm/constants"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AccessLogger() gin.HandlerFunc {
	currentTime := time.Now()
	yyyy, mm, dd := currentTime.Date()
	fileName := fmt.Sprintf("logs/access_log_%d-%d-%d.txt", yyyy, mm, dd)

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	
	if err != nil {
		fmt.Println("err", err)
	}
	err = os.Chmod(fileName, 0666)
	logger := log.New()
	logger.Out = src
	logger.SetLevel(getLogLevel(configs.LogLevel())) //  Trace, Debug, Info, Warning, Error, Fatal and Panic.
	logger.SetFormatter(&log.TextFormatter{}) // &log.TextFormatter{} | &log.JSONFormatter{}
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		// log format
		logger.Infof("%3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}

func getLogLevel(level constants.LogLevel) log.Level  {
	l, _ := log.ParseLevel(level.String())
	return l
}