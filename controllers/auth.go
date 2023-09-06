package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-gorm/configs"
	"gin-gorm/middleware"
	"gin-gorm/models"
	"gin-gorm/responses"
)

func LogInUser(c *gin.Context) {
	var user models.User

	//validate the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	err := user.Validate("login")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "Invalid credentials"}})
		return
	}
	te, err := user.FindUserByEmail(user.Email)

	err = models.VerifyPassword(te.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "Invalid credentials"}})
		return
	}
	token := middleware.GenerateToken(user.Email)
	c.SetCookie(configs.CookieConfig().Name, "Bearer "+token, configs.CookieConfig().MaxAge, configs.CookieConfig().Path, configs.CookieConfig().Domain, configs.CookieConfig().Secure, configs.CookieConfig().HttpOnly)

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{}})
}

func SignupUser(c *gin.Context) {
	CreateUser(c)
}

func LogOutUser(c *gin.Context) {

	c.SetCookie(configs.CookieConfig().Name, "", -1, configs.CookieConfig().Path, configs.CookieConfig().Domain, configs.CookieConfig().Secure, configs.CookieConfig().HttpOnly)

	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{}})
}