package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Parse incoming Data",
			"error":   err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Save User",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Saved Successfully!!",
	})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Parse incoming Data",
			"error":   err.Error(),
		})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something Went Wrong!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success!",
		"token":   token,
	})
}
