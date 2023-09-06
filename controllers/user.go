package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm/models"
	"gin-gorm/responses"
)

func CreateUser(c *gin.Context) {
	var user models.User

	//validate the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	user.Prepare()
	
	err := user.Validate("")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	userCreated, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	
	c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"user": userCreated}})
}

func GetAllUsers(c *gin.Context) {
	user := models.User{}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	users, err := user.FindAllUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"users": users}},
	)
}

func GetAUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(uint32(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": userGotten}},
	)
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	var user models.User

	//validate the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	user.Prepare()
	err = user.Validate("update")

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	updatedUser, err := user.UpdateAUser(uint32(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": updatedUser}},
	)
}

func DeleteAUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	user := models.User{}
	_, err = user.DeleteAUser(uint32(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "User successfully deleted!"}},
	)
}
