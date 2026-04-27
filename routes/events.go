package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
		})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Parse Event ID",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could Not Find the Event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Found the Event",
		"event":   event,
	})
}

func createEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Parse incoming Data",
			"error":   err.Error(),
		})
		return
	}

	userId := c.GetInt64(("userId"))
	event.UserID = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"data":    event})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Parse Event ID",
		})
		return
	}
	eventInDb, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could Not Find the Event",
		})
		return
	}

	userId := c.GetInt64("userId")

	if eventInDb.ID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You do Not Have Permisson To Update this Event!",
		})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Parse incoming Data",
			"error":   err.Error(),
		})
		return
	}
	event.UserID = userId
	event.ID = eventId
	err = event.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Update Event",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event Updated!",
		"event":   event,
	})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Parse Event ID",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Could Not Find the Event",
		})
		return
	}

	userId := c.GetInt64("userId")
	if event.ID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You do Not Have Permisson To Delete this Event!",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Delete Event",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event Deleted!",
	})

}
