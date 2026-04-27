package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Parse Event ID",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to Create Registration",
		})
		return
	}
	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to Create Registration",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event Registration complete"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Parse Event ID",
		})
		return
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not Cancel Registration",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Cancelled",
	})
}
